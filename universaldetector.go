package chardet

/*
######################## BEGIN LICENSE BLOCK ########################
# The Original Code is Mozilla Universal charset detector code.
#
# The Initial Developer of the Original Code is
# Netscape Communications Corporation.
# Portions created by the Initial Developer are Copyright (C) 2001
# the Initial Developer. All Rights Reserved.
#
# Contributor(s):
#   Mark Pilgrim - port to Python
#   Shy Shalom - original C code
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

/* Module containing the UniversalDetector detector struct, which is
   the primary struct a user of chardet should use.

   Authors :
	Oscar Laurent (port to go from python)
	Mark Pilgrim (initial port to python)
	Shy Shalom (original C code)
	Dan Blanchard (major refactoring for 3.0)
	Ian Cordasco
*/

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

// UDMinimumThreshold min threshhold for ???
var UDMinimumThreshold = 0.2

// UDHighByteDetector ???
var UDHighByteDetector = regexp.MustCompile(`[^[:ascii:]]`)

// UDEscDetector ???
var UDEscDetector = regexp.MustCompile(`(\x1B|~{)`)

// UDWinByteDetector ???
var UDWinByteDetector = regexp.MustCompile(`[\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8A\x8B\x8C\x8D\x8E\x8F\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9A\x9B\x9C\x9D\x9E\x9F]`)

// UDIsoWinMap correspondance ISO-Win names
var UDIsoWinMap = map[string]string{
	"iso-8859-1":  "Windows-1252",
	"iso-8859-2":  "Windows-1250",
	"iso-8859-5":  "Windows-1251",
	"iso-8859-6":  "Windows-1256",
	"iso-8859-7":  "Windows-1253",
	"iso-8859-8":  "Windows-1255",
	"iso-8859-9":  "Windows-1254",
	"iso-8859-13": "Windows-1257",
}

// Result contains the result or a detection
type Result struct {
	Encoding   string
	Confidence float64
	Language   string
}

// UniversalDetector is the detector of all the encodings.
type UniversalDetector struct {
	Result           *Result
	escCharSetProber Prober
	charsetProbers   []Prober
	lastChar         *byte
	langFilter       LanguageFilter
	Done             bool
	gotData          bool
	hasWinBytes      bool
	inputState       UniversalDetectorState
}

func newUniversalDetector(langFilter LanguageFilter) *UniversalDetector {
	u := UniversalDetector{
		&Result{},
		nil,
		[]Prober{},
		nil,
		langFilter,
		false,
		false,
		false,
		UDSPureASCII,
	}
	u.reset()
	return &u
}

/* Resets the UniversalDetector and all of its probers back to their
   initial states. This is called when creating a new detector so you only need
   to call this directly in between analyses of different documents
*/
func (u *UniversalDetector) reset() {
	u.Result = &Result{"", 0.0, ""}
	u.Done = false
	u.gotData = false
	u.hasWinBytes = false
	u.inputState = UDSPureASCII
	u.lastChar = nil
	if u.escCharSetProber != nil {
		u.escCharSetProber.reset()
	}
	for _, p := range u.charsetProbers {
		p.reset()
	}
}

// Feed takes a chunk of a document and feeds it through all of the relevant charset probers.
// After calling Feed, you can check the value of the Done attribute
// to see if you need to continue feeding the UniversalDetector mode data, or if it has
// made a prediction in the Result attribute
// Note: You should always call Close when you're done feeding in your document
//       if Done is not already true
func (u *UniversalDetector) Feed(data []byte) {
	if u.Done || len(data) == 0 {
		log.Debugf("Early return")
		return
	}

	// First check for known BOMs (byte order marks) since these are guaranteed to be correct
	if !u.gotData {
		bomResult := scanBOMs(data)
		u.gotData = true
		if bomResult != nil {
			u.Result = bomResult
			u.Done = true
			return
		}
	}

	// If none of those matched and we've only seen ASCII so far,
	// check for high bytes and escape sequences.
	highByteMark, escByteMark, winBytesMark := detectMarks(data, u.lastChar)
	if u.inputState == UDSPureASCII {
		//if UDHighByteDetector.Match(data) {
		if highByteMark {
			u.inputState = UDSHighByte
		} else if u.inputState == UDSPureASCII {
			//if UDEscDetector.Match(data) {
			if escByteMark {
				u.inputState = UDSEscASCII
				//} else if u.lastChar != nil {
				//	// This isn't pretty but we don't want to copy the full data, only the start
				//	maxLen := 5
				//	if len(data) < maxLen {
				//		maxLen = len(data)
				//	}
				//	//dataStart := append([]byte{*u.lastChar}, data[:maxLen]...)
				//	dataStart := append([]byte{*u.lastChar}, data...)
				//	if UDEscDetector.Match(dataStart) {
				//		u.inputState = UDSEscASCII
				//	}
			}
		}
	}
	u.lastChar = &data[len(data)-1]

	// If we've seen escape sequences, use the EscCharSetProber, which
	// Uses a simple state machine to check for known escape sequences in
	// HZ and ISO-2022 encodings, since those are the only encodings that
	// use such sequences.
	if u.inputState == UDSEscASCII {
		if u.escCharSetProber == nil {
			u.escCharSetProber = newEscCharSetProber(u.langFilter)
		}
		if u.escCharSetProber.feed(data) == PSFound {
			u.Result = &Result{
				Encoding:   u.escCharSetProber.charsetName(),
				Confidence: u.escCharSetProber.getConfidence(),
				Language:   u.escCharSetProber.language(),
			}
			u.Done = true
		}
	} else if u.inputState == UDSHighByte {
		// If we've seen high bytes (i.e. those with values greater than 127),
		// we need to do mode complicated checks using all our multi-byte and
		// single-byte probers that are left. The single-bite probers
		// use character bigram distributions to determine the encoding, whereas
		// the multi-byte probers use a combination of character unigram and
		// bigram distribitions.
		if len(u.charsetProbers) == 0 {
			u.charsetProbers = []Prober{newMBCSGroupProber(u.langFilter)}
			if (u.langFilter & LFNonCJK) > 0 {
				u.charsetProbers = append(u.charsetProbers, newSBCSGroupProber())
			}
			u.charsetProbers = append(u.charsetProbers, newLatin1Prober(LFNone))
		}
		for _, prober := range u.charsetProbers {
			tmp := prober.feed(data)
			if tmp == PSFound {
				u.Result = &Result{
					Encoding:   prober.charsetName(),
					Confidence: prober.getConfidence(),
					Language:   prober.language(),
				}
				u.Done = true
				break
			}
		}
		//u.hasWinBytes = detectWinBytes(data)
		u.hasWinBytes = winBytesMark
	}
}

func detectMarks(data []byte, lastChar *byte) (bool, bool, bool) {
	highByte := false
	escByte := false
	winBytes := false
	esc0 := false
	if lastChar != nil {
		esc0 = (*lastChar == '~')
	}
	for _, x := range data {
		highByte = highByte || (x >= 0x80)
		winBytes = winBytes || (x >= 0x80 && x <= 0x9f)
		escByte = escByte || (x == '\033') || (esc0 && x == '{')
		esc0 = (x == '~')
		if winBytes && escByte {
			return true, true, true
		}
	}
	return highByte, escByte, winBytes
}

func detectWinBytes(data []byte) bool {
	for _, x := range data {
		if x >= 0x80 && x <= 0x9f {
			return true
		}
	}
	return false
}

// Close stops analyzing the current document and come up with a final prediction.
// Returns : a Result struct
func (u *UniversalDetector) Close() *Result {
	// Don't bother with checks if we're already done.
	if u.Done {
		return u.Result
	}
	u.Done = true

	if !u.gotData {
		// log No data received!
		return nil
	}

	// Defaults to ASCII if it is all we've seen so far.
	if u.inputState == UDSPureASCII {
		u.Result = &Result{"ascii", 1.0, ""}
	} else if u.inputState == UDSHighByte {
		// If we have seen non ASCII, return the best that met UDMinimumThreshold
		proberConfidence := -1.0
		maxProberConfidence := -1.0
		var maxProber Prober = nil
		for _, prober := range u.charsetProbers {
			if prober == nil {
				continue
			}
			proberConfidence = prober.getConfidence()
			if proberConfidence > maxProberConfidence {
				maxProberConfidence = proberConfidence
				maxProber = prober
			}
		}
		if maxProber != nil && maxProberConfidence > UDMinimumThreshold {
			charsetName := maxProber.charsetName()
			lowerCharsetName := strings.ToLower(charsetName)
			confidence := maxProberConfidence
			// Use Windows encoding name instead of the ISO-8859 if we saw
			// any extra Windows-specific bytes
			if strings.HasPrefix(lowerCharsetName, "iso-8859") && u.hasWinBytes {
				charsetNameTmp, ok := UDIsoWinMap[lowerCharsetName]
				if ok {
					charsetName = charsetNameTmp
				}
			} else {
			}
			u.Result = &Result{
				Encoding:   charsetName,
				Confidence: confidence,
				Language:   maxProber.language(),
			}
		}

	}

	// Log all probers confidences if none met the UDMinimumThreshold
	if u.Result == nil {
		// Log No prober met minimum threshold
		u.printProbersDetails()
	}
	return u.Result
}

func (u *UniversalDetector) printProbersDetails() {
	for _, prober := range u.charsetProbers {
		if prober == nil {
			continue
		}
		// Distinguish between GroupProber and the rest before logging
		if mbcsprober, ok := prober.(*CharSetGroupProber); ok {
			for _, subprober := range mbcsprober.probers {
				log.Debugf(
					"Prober %v (lang %v) - %v",
					subprober.charsetName(),
					subprober.language(),
					subprober.getConfidence(),
				)
			}
		} else {
			log.Debugf(
				"Prober %v (lang %v) - %v",
				prober.charsetName(),
				prober.language(),
				prober.getConfidence(),
			)
		}
	}
}

func scanBOMs(data []byte) *Result {
	if bytes.HasPrefix(data, []byte{0xEF, 0xBB, 0xBF}) { // UTF8 BOM
		return &Result{"UTF-8-SIG", 1.0, ""}
	}
	if bytes.HasPrefix(data, []byte{0xff, 0xfe, 0x00, 0x00}) || // UTF32 LE
		bytes.HasPrefix(data, []byte{0x00, 0x00, 0xfe, 0xff}) { // UTF32 BE
		return &Result{"UTF-32", 1.0, ""}
	}
	if bytes.HasPrefix(data, []byte{0xFE, 0xFF, 0x00, 0x00}) { // UCS-4, Unusual octet BOM (3412)
		return &Result{"X-ISO-10646-UCS-4-3412", 1.0, ""}
	}
	if bytes.HasPrefix(data, []byte{0x00, 0x00, 0xFF, 0xFE}) { // UCS-4, Unusual octet BOM (2143)
		return &Result{"X-ISO-10646-UCS-4-2143", 1.0, ""}
	}
	if bytes.HasPrefix(data, []byte{0xff, 0xfe}) || // UTF16 LE
		bytes.HasPrefix(data, []byte{0xfe, 0xff}) { // UTF16 BE
		return &Result{"UTF-16", 1.0, ""}
	}
	return nil
}
