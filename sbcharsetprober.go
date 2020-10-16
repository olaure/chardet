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

// SBCPSampleSize sample size
var SBCPSampleSize = 64

// SBEnoughRelThreshold ??? -> 0.25 * SBCPSampleSize^2
var SBEnoughRelThreshold = 1024

//SBCPPositiveShortcutThreshold ???
var SBCPPositiveShortcutThreshold = 0.95

//SBCPNegativeShortcutThreshold ???
var SBCPNegativeShortcutThreshold = 0.05

// SingleByteCharSetModel model for SBCSProber
type SingleByteCharSetModel struct {
	charsetName          string
	language             string
	charToOrderMap       map[byte]CharacterCategory
	charToOrderList      [256]CharacterCategory
	languageModel        map[CharacterCategory][]int
	typicalPositiveRatio float64
	keepASCIILetters     bool
	alphabet             string
}

// SingleByteCharSetProber is a single byte charset prober ???
type SingleByteCharSetProber struct {
	CharSetProber
	model       *SingleByteCharSetModel
	reversed    bool
	nameProber  Prober
	lastOrder   CharacterCategory
	seqCounters []int
	totalSeqs   int
	totalChar   int
	freqChar    int
}

func newSingleByteCharSetProber(model *SingleByteCharSetModel, rev bool, nameProber Prober) *SingleByteCharSetProber {
	return &SingleByteCharSetProber{
		*newCharSetProber(LFNone),
		model,
		// TRUE if we need to reverse every pair in the model lookup
		rev,
		// Optional auxiliary prober for name decision
		nameProber,
		255,
		[]int{},
		0,
		0,
		0,
	}
}

func (s *SingleByteCharSetProber) reset() {
	s.CharSetProber.reset()
	// Char order of the last character
	s.lastOrder = 255
	s.seqCounters = make([]int, SLNumCategories)
	s.totalSeqs = 0
	s.totalChar = 0
	// characters that fall in our sampling range
	s.freqChar = 0
}

func (s *SingleByteCharSetProber) charsetName() string {
	if s.nameProber != nil {
		return s.nameProber.charsetName()
	}
	return s.model.charsetName
}

func (s *SingleByteCharSetProber) language() string {
	if s.nameProber != nil {
		return s.nameProber.language()
	}
	return s.model.language
}

func (s *SingleByteCharSetProber) getConfidence() float64 {
	conf := 0.01
	if s.totalSeqs > 0 {
		// TODO POTENTIAL BUG
		conf = ((1.0 * float64(s.seqCounters[SLPositive])) /
			float64(s.totalSeqs) / float64(s.model.typicalPositiveRatio))
		conf = conf * float64(s.freqChar) / float64(s.totalChar)
		if conf >= 1.0 {
			conf = 0.99
		}
	}
	return conf
}

func (s *SingleByteCharSetProber) feed(data []byte) ProbingState {
	// TODO: Make filter_international_words keep things in self.alphabet
	newData := data
	if !s.model.keepASCIILetters {
		newData = filterInternationalWords(data)
	}
	if len(newData) == 0 {
		return s.state
	}
	//charToOrderMap := s.model.charToOrderMap
	charToOrderList := s.model.charToOrderList
	languageModel := s.model.languageModel
	for _, chr := range newData {
		var order CharacterCategory
		if chr > 255 {
			order = CCUndefined
		} else {
			order = charToOrderList[chr]
		}
		// order, ok := charToOrderMap[chr]
		// /*
		//  * XXX: This was SYMBOL_CAT_ORDER before, with a value of 250, but
		//  *      CharacterCategory.SYMBOL is actually 253, so we use CONTROL
		//  *      to make it closer to the original intent. The only difference
		//  *      is whether or not we count digits and control characters for
		//  *      _total_char purposes.
		//  */
		// if !ok {
		// 	order = CCUndefined
		// }
		if order < CCControl {
			s.totalChar++
		}
		/* TODO:
		 * Follow uchardet's lead and discount confidence for frequent
		 * control characters.
		 * See https://github.com/BYVoid/uchardet/commit/55b4f23971db61
		 */
		if int(order) < SBCPSampleSize {
			s.freqChar++
			if int(s.lastOrder) < SBCPSampleSize {
				s.totalSeqs++
				var lmCat int
				if !s.reversed {
					lmCat = languageModel[s.lastOrder][order]
				} else {
					lmCat = languageModel[order][s.lastOrder]
				}
				s.seqCounters[lmCat]++
			}
		}
		s.lastOrder = order
	}

	//charsetName := s.model.charsetName
	if s.state == PSDetecting {
		if s.totalSeqs > SBEnoughRelThreshold {
			confidence := s.getConfidence()
			if confidence > SBCPPositiveShortcutThreshold {
				// log
				s.state = PSFound
			} else if confidence < SBCPNegativeShortcutThreshold {
				// log
				s.state = PSNotMe

			}
		}
	}
	return s.state
}
