package chardet

/*
######################## BEGIN LICENSE BLOCK ########################
# The Original Code is Mozilla Communicator client code.
#
# The Initial Developer of the Original Code is
# Netscape Communications Corporation.
# Portions created by the Initial Developer are Copyright (C) 1998
# the Initial Developer. All Rights Reserved.
#
# Contributor(s):
#   Mark Pilgrim - port to Python
#
# This library is free software; you can redistribute it and/or
# modify it under the terms of the GNU Lesser General Public
# License as published by the Free Software Foundation; either
# version 2.1 of the License, or (at your option) any later version.
#
# This library is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
# Lesser General Public License for more details.
#
# You should have received a copy of the GNU Lesser General Public
# License along with this library; if not, write to the Free Software
# Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA
# 02110-1301  USA
######################### END LICENSE BLOCK #########################
*/

import (
	//log "github.com/sirupsen/logrus"
	"regexp"
)

// Prober interface shared by all probers XD
type Prober interface {
	reset()
	charsetName() string
	language() string
	feed([]byte) ProbingState
	getConfidence() float64
	setActive(bool)
	getActive() bool
}

// CPShortcutThreshold is the threshold after which certainty is absolute
var CPShortcutThreshold = 0.95

var reHighByteFilter = regexp.MustCompile(`([[:ascii:]])+`)

// [^[:ascii:]] == [\x80-\xFF]
/*
captures "words" around every non-ascii parts
	<= Takes every alpha char preceding
	=> Takes every alpha char following
	=> Takes an extra character if it is ASCII AND non alpha
If last character is ascii and non alpha
	change it for a space.
*/
var reInternationalFilter = regexp.MustCompile(
	`[[:alpha:]]*[^[:ascii:]]+[[:alpha:]]*([^[:alpha:][:^ascii:]])?`)

// var reIsAlpha = regexp.MustCompile(`[a-zA-Z]+`)

// CharSetProber probes a charset???
type CharSetProber struct {
	state      ProbingState
	langFilter LanguageFilter
	active     bool
}

func newCharSetProber(langFilter LanguageFilter) *CharSetProber {
	return &CharSetProber{
		PSDetecting,
		langFilter,
		true,
	}
}

func (c *CharSetProber) reset() {
	c.state = PSDetecting
}

func (c *CharSetProber) getActive() bool {
	return c.active
}

func (c *CharSetProber) setActive(act bool) {
	c.active = act
}

func (c *CharSetProber) getState() ProbingState {
	return c.state
}

func filterHighByteOnly(buf []byte) []byte {
	return reHighByteFilter.ReplaceAll(buf, []byte{' '})
}

/*  We define three types of bytes:
    - alphabet: english alphabets [a-zA-Z]
    - international: international characters [\x80-\xFF]
    - marker: everything else [^a-zA-Z\x80-\xFF]

	The input buffer can be thought to contain a series of words delimited
	by markers. This function works to filter all words that contain at
	least one internaitonal character. All contiguous sequences of markers
	are replaced by a single space ASCII character.

	This filter applies to all scripts which do not use english characters.
*/
func prevFilterInternationalWords(buf []byte) []byte {
	// This regexp filters out only words that have at least one
	// International character. The word may include one marker
	// character at the end.
	words := reInternationalFilter.FindAll(buf, -1)
	totalLength := 0
	for _, w := range words {
		totalLength += len(w)
	}
	filtered := make([]byte, 0, totalLength)
	for _, word := range words {
		filtered = append(filtered, word...)

		// If the last character in the word is a marker, replace it with a
		// space as markers shouldn't affect our analysis (they are used
		// similarly across all languages and may thus have similar frequencies).
		lastChar := word[len(word)-1]
		//if !reIsAlpha.Match([]byte{lastChar}) && lastChar < 0x80 {
		if lastChar < 0x80 && !(('a' <= lastChar && lastChar <= 'z') || ('A' <= lastChar && lastChar <= 'Z')) {
			filtered[len(filtered)-1] = ' '
			//lastChar = ' '
		}
		//filtered = append(filtered, lastChar)
	}
	return filtered
}

/*  We define three types of bytes:
    - alphabet: english alphabets [a-zA-Z]
    - international: international characters [\x80-\xFF]
    - marker: everything else [^a-zA-Z\x80-\xFF]

	The input buffer can be thought to contain a series of words delimited
	by markers. This function works to filter all words that contain at
	least one internaitonal character. All contiguous sequences of markers
	are replaced by a single space ASCII character.

	This filter applies to all scripts which do not use english characters.
*/
func filterInternationalWords(buf []byte) []byte {
	filtered := make([]byte, 0, len(buf))
	isAlpha := func(b byte) bool {
		return ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z')
	}
	start := 0
	hasHighByte := false
	inWord := false
	for idx, b := range buf {
		if !inWord {
			if isAlpha(b) || (b >= 0x80) { // Start of word
				inWord = true
				hasHighByte = (b >= 0x80)
				start = idx
			} else {
				continue
			}
		} else {
			if isAlpha(b) || b >= 0x80 { // In word
				hasHighByte = hasHighByte || (b >= 0x80)
			} else { // End of word
				inWord = false
				if hasHighByte { // End of interesting word
					filtered = append(filtered, buf[start:idx+1]...)
					// Is replacing the last marker character with a space useful?
					// It is the previous behavior though.
					filtered[len(filtered)-1] = ' '
				}
				// A word without any high bytes => Not interesting
			}
		}
	}
	if inWord && hasHighByte { // Last part of buffer is still interesting
		filtered = append(filtered, buf[start:]...)
	}
	// Let's release the unused memory
	return filtered[:len(filtered)]
}

/*  Returns a copy of buf that retains only the sequences of english alphabet
    and high byte characters that are not between <> characters.
    Also retains english alphabet and high byte characetrs immediately before
    occurences of >.
    This filter can be applied to all scripts which contain both english characters
    and extended ASCII characters, but is currently only used by Latin1Prober.
*/
func filterWithEnglishLetters(buf []byte) []byte {
	filtered := []byte{}
	inTag := false
	prev := 0

	for cur, chr := range buf {
		if chr == '<' {
			inTag = false
		} else if chr == '>' {
			inTag = true
		}

		// If current character is not extended-ASCII and not alphabetic...
		//if chr < 0x80 && !reIsAlpha.Match([]byte{chr}) {
		if chr < 0x80 && !(('a' <= chr && chr <= 'z') || ('A' <= chr && chr <= 'Z')) {
			// ... and we aren't in a tag
			if cur > prev && !inTag {
				// Keep everything after last non-extended-ASCII,
				// non-alphabetic character
				if cur < len(buf)-1 {
					filtered = append(filtered, buf[prev:cur]...)
					// Output a space to delimit stretch we keep
					filtered[len(filtered)-1] = ' '
				} else {
					filtered = append(filtered, buf[prev:cur]...)
				}
			}
			prev = cur + 1
		}
	}
	// If we aren't in a tag
	if !inTag && prev < len(buf) {
		// Keep everything after the last non-extended-ASCII, non-alphabetic character
		filtered = append(filtered, buf[prev:]...)
	}
	return filtered
}
