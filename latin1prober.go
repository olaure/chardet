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

import (
//log "github.com/sirupsen/logrus"
)

var lat1FreqCatNum = 4

type lat1Class int

const (
	// UDF is undefined
	UDF lat1Class = 0
	// OTH is other
	OTH lat1Class = 1
	// ASC is ascii capital letter
	ASC lat1Class = 2
	// ASS is ascii small letter
	ASS lat1Class = 3
	// ACV is accent capital vowel
	ACV lat1Class = 4
	// ACO is accent capital other
	ACO lat1Class = 5
	// ASV is accent small vowel
	ASV lat1Class = 6
	// ASO is accent small other
	ASO lat1Class = 7
	// ClassNum is total classes
	ClassNum lat1Class = 8
)

var latin1CharToClass = []lat1Class{
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 00 - 07
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 08 - 0F
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 10 - 17
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 18 - 1F
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 20 - 27
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 28 - 2F
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 30 - 37
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 38 - 3F
	OTH, ASC, ASC, ASC, ASC, ASC, ASC, ASC, // 40 - 47
	ASC, ASC, ASC, ASC, ASC, ASC, ASC, ASC, // 48 - 4F
	ASC, ASC, ASC, ASC, ASC, ASC, ASC, ASC, // 50 - 57
	ASC, ASC, ASC, OTH, OTH, OTH, OTH, OTH, // 58 - 5F
	OTH, ASS, ASS, ASS, ASS, ASS, ASS, ASS, // 60 - 67
	ASS, ASS, ASS, ASS, ASS, ASS, ASS, ASS, // 68 - 6F
	ASS, ASS, ASS, ASS, ASS, ASS, ASS, ASS, // 70 - 77
	ASS, ASS, ASS, OTH, OTH, OTH, OTH, OTH, // 78 - 7F
	OTH, UDF, OTH, ASO, OTH, OTH, OTH, OTH, // 80 - 87
	OTH, OTH, ACO, OTH, ACO, UDF, ACO, UDF, // 88 - 8F
	UDF, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 90 - 97
	OTH, OTH, ASO, OTH, ASO, UDF, ASO, ACO, // 98 - 9F
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // A0 - A7
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // A8 - AF
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // B0 - B7
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // B8 - BF
	ACV, ACV, ACV, ACV, ACV, ACV, ACO, ACO, // C0 - C7
	ACV, ACV, ACV, ACV, ACV, ACV, ACV, ACV, // C8 - CF
	ACO, ACO, ACV, ACV, ACV, ACV, ACV, OTH, // D0 - D7
	ACV, ACV, ACV, ACV, ACV, ACO, ACO, ACO, // D8 - DF
	ASV, ASV, ASV, ASV, ASV, ASV, ASO, ASO, // E0 - E7
	ASV, ASV, ASV, ASV, ASV, ASV, ASV, ASV, // E8 - EF
	ASO, ASO, ASV, ASV, ASV, ASV, ASV, OTH, // F0 - F7
	ASV, ASV, ASV, ASV, ASV, ASO, ASO, ASO, // F8 - FF
}

// 0 : illegal
// 1 : very unlikely
// 2 : normal
// 3 : very likely
var latin1ClassModel = []int{
	// UDF OTH ASC ASS ACV ACO ASV ASO
	0, 0, 0, 0, 0, 0, 0, 0, // UDF
	0, 3, 3, 3, 3, 3, 3, 3, // OTH
	0, 3, 3, 3, 3, 3, 3, 3, // ASC
	0, 3, 3, 3, 1, 1, 3, 3, // ASS
	0, 3, 3, 3, 1, 2, 1, 2, // ACV
	0, 3, 3, 3, 3, 3, 3, 3, // ACO
	0, 3, 1, 3, 1, 1, 1, 3, // ASV
	0, 3, 1, 3, 1, 1, 3, 3, // ASO
}

// Latin1Prober prober
type Latin1Prober struct {
	CharSetProber
	lastCharClass lat1Class
	freqCounter   []int
}

func newLatin1Prober(langFilter LanguageFilter) *Latin1Prober {
	l := Latin1Prober{
		*newCharSetProber(langFilter),
		OTH,
		make([]int, lat1FreqCatNum), // Default init is 0
	}
	l.reset()
	return &l
}

func (l *Latin1Prober) reset() {
	l.CharSetProber.reset()
	l.lastCharClass = OTH
	l.freqCounter = make([]int, lat1FreqCatNum)
}

func (l *Latin1Prober) charsetName() string {
	return "ISO-8859-1"
}

func (l *Latin1Prober) language() string {
	return ""
}

func (l *Latin1Prober) getConfidence() float64 {
	if l.state == PSNotMe {
		return 0.01
	}
	total := 0
	for _, v := range l.freqCounter {
		total += v
	}
	var confidence float64
	if total <= 0 {
		confidence = 0.01
	} else {
		confidence = (float64(l.freqCounter[3]) - float64(l.freqCounter[1])*20.0) / float64(total)
	}
	if confidence < 0.0 {
		confidence = 0.0
	}
	// lower the confidence of latin1 so that other more accurate
	// detector can take priority.
	confidence = confidence * 0.73
	return confidence
}

func (l *Latin1Prober) feed(data []byte) ProbingState {
	newData := filterWithEnglishLetters(data)
	for _, chr := range newData {
		charClass := latin1CharToClass[chr]
		freq := latin1ClassModel[l.lastCharClass*ClassNum+charClass]
		if freq == 0 {
			l.state = PSNotMe
			break
		}
		l.freqCounter[freq]++
		l.lastCharClass = charClass
	}
	return l.state
}
