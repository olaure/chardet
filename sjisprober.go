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

// SJISProber prober
type SJISProber struct {
	MultiByteCharSetProber
	contextAnalyzer *SJISContextAnalysis
}

func newSJISProber() Prober {
	mbcsp := newMultiByteCharSetProber(LFNone)
	mbcsp.distributionAnalyzer = newSJISDistributionAnalysis()
	mbcsp.codingSm = newCodingStateMachine(SJISSmModel)
	prober := SJISProber{*mbcsp, newSJISContextAnalysis()}

	prober.reset()
	return &prober
}

func (s *SJISProber) reset() {
	s.MultiByteCharSetProber.reset()
	s.contextAnalyzer.reset()
}

func (s *SJISProber) charsetName() string {
	return *s.contextAnalyzer.charSetName
}

func (s *SJISProber) language() string {
	return "Japanese"
}

func (s *SJISProber) getConfidence() float64 {
	contextConf := s.contextAnalyzer.getConfidence()
	distribConf := s.distributionAnalyzer.getConfidence()
	if contextConf > distribConf {
		return contextConf
	}
	return distribConf
}

func (s *SJISProber) feed(data []byte) ProbingState {
	for i, chr := range data {
		codingState := s.codingSm.nextState(chr)
		switch codingState {
		case MSError:
			s.state = PSNotMe
			break
		case MSItsMe:
			s.state = PSFound
			break
		case MSStart:
			charLen := s.codingSm.getCurrentCharLen()
			if i == 0 {
				s.lastChar[1] = chr
				s.contextAnalyzer.feed(s.lastChar[2-charLen:], charLen)
				s.distributionAnalyzer.feed(s.lastChar, charLen)
			} else {
				max := i + 3 - charLen
				if max > len(data) {
					max = len(data)
				}
				s.contextAnalyzer.feed(data[i+1-charLen:max], charLen)
				s.distributionAnalyzer.feed(data[i-1:i+1], charLen)
			}
		}
	}
	s.lastChar[0] = data[len(data)-1]

	if s.state == PSDetecting {
		if s.contextAnalyzer.gotEnoughData() &&
			s.getConfidence() > CPShortcutThreshold {
			s.state = PSFound
		}
	}
	return s.state
}
