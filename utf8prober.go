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
	"math"
)

// UTF8OneCharProb probability of one character points
var UTF8OneCharProb = 0.5

// UTF8Prober prober
type UTF8Prober struct {
	CharSetProber
	codingSm   *codingStateMachine
	numMBChars int
	active     bool
}

func newUTF8Prober() Prober {
	return &UTF8Prober{
		*newCharSetProber(LFNone),
		newCodingStateMachine(UTF8SmModel),
		0,
		true,
	}
}

func (u *UTF8Prober) reset() {
	u.CharSetProber.reset()
	u.codingSm.reset()
	u.numMBChars = 0
	u.active = true
}

func (u *UTF8Prober) charsetName() string {
	return "utf-8"
}

func (u *UTF8Prober) language() string {
	return ""
}

func (u *UTF8Prober) getConfidence() float64 {
	unlike := 0.99
	if u.numMBChars < 6 {
		unlike *= math.Pow(UTF8OneCharProb, float64(u.numMBChars))
		return 1.0 - unlike
	}
	return unlike
}

func (u *UTF8Prober) feed(data []byte) ProbingState {
	for _, chr := range data {
		switch u.codingSm.nextState(chr) {
		case MSError:
			u.state = PSNotMe
			break
		case MSItsMe:
			u.state = PSFound
			break
		case MSStart:
			if u.codingSm.getCurrentCharLen() >= 2 {
				u.numMBChars++
			}
		}
	}
	if u.state == PSDetecting && u.getConfidence() > CPShortcutThreshold {
		u.state = PSFound
	}
	return u.state
}
