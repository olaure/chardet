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
)

// MultiByteCharSetProber is a multi byte charset prober ???
type MultiByteCharSetProber struct {
	CharSetProber
	distributionAnalyzer charDistributionAnalyzer
	codingSm             *codingStateMachine
	lastChar             []byte
	active               bool
}

func newMultiByteCharSetProber(langFilter LanguageFilter) *MultiByteCharSetProber {
	return &MultiByteCharSetProber{
		*newCharSetProber(langFilter),
		nil,
		nil,
		[]byte{0, 0},
		true,
	}
}

func (m *MultiByteCharSetProber) reset() {
	m.CharSetProber.reset()
	if m.codingSm != nil {
		(*m.codingSm).reset()
	}
	if m.distributionAnalyzer != nil {
		m.distributionAnalyzer.reset()
	}
	m.lastChar = []byte{0, 0}
}

/*
func (m *MultiByteCharSetProber) charsetName() string {
	panic("Not implemented")
}

func (m *MultiByteCharSetProber) language() string {
	panic("Not implemented")
}
*/

func (m *MultiByteCharSetProber) getConfidence() float64 {
	return m.distributionAnalyzer.getConfidence()
}

func (m *MultiByteCharSetProber) feed(data []byte) ProbingState {
	for i, chr := range data {
		codingState := m.codingSm.nextState(chr)
		switch codingState {
		case MSError:
			m.state = PSNotMe
			m.active = false
			break
		case MSItsMe:
			m.state = PSFound
			break
		case MSStart:
			charLen := m.codingSm.getCurrentCharLen()
			if i == 0 {
				m.lastChar[1] = chr
				m.distributionAnalyzer.feed(m.lastChar, charLen)
			} else {
				m.distributionAnalyzer.feed(data[i-1:i+1], charLen)
			}
		}
	}
	m.lastChar[0] = data[len(data)-1]
	if m.state == PSDetecting {
		if m.distributionAnalyzer.gotEnoughData() {
			conf := m.getConfidence()
			if conf > CPShortcutThreshold {
				m.state = PSFound
			}
		}
	}
	return m.state
}
