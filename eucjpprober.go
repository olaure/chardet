package chardet

/*
######################## BEGIN LICENSE BLOCK ########################
# The Original Code is mozilla.org code.
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
	log "github.com/sirupsen/logrus"
)

// EUCJPProber prober
type EUCJPProber struct {
	MultiByteCharSetProber
	contextAnalyzer *EUCJPContextAnalysis
}

func newEUCJPProber() *EUCJPProber {
	mbcsp := newMultiByteCharSetProber(LFNone)
	mbcsp.distributionAnalyzer = newEUCJPDistributionAnalysis()
	mbcsp.codingSm = newCodingStateMachine(EUCJPSmModel)
	prober := EUCJPProber{*mbcsp, newEUCJPContextAnalysis()}
	prober.reset()

	return &prober
}

func (e *EUCJPProber) reset() {
	e.CharSetProber.reset()
	e.contextAnalyzer.reset()
}

func (e *EUCJPProber) charsetName() string {
	return "EUC-JP"
}

func (e *EUCJPProber) language() string {
	return "Japanese"
}

func (e *EUCJPProber) getConfidence() float64 {
	contextConf := e.contextAnalyzer.getConfidence()
	distribConf := e.distributionAnalyzer.getConfidence()
	if contextConf > distribConf {
		return contextConf
	}
	return distribConf
}

func (e *EUCJPProber) feed(data []byte) ProbingState {
	for i := 0; i < len(data); i++ {
		switch e.codingSm.nextState(data[i]) {
		case MSError:
			log.Debugf("EUCJP Prober SM machine said it's not her")
			e.state = PSNotMe
			break
		case MSItsMe:
			e.state = PSFound
			break
		case MSStart:
			charLen := e.codingSm.getCurrentCharLen()
			if i == 0 {
				e.lastChar[1] = data[0]
				e.contextAnalyzer.feed(e.lastChar, charLen)
				e.distributionAnalyzer.feed(e.lastChar, charLen)
			} else {
				e.contextAnalyzer.feed(data[i-1:i+1], charLen)
				e.distributionAnalyzer.feed(data[i-1:i+1], charLen)
			}
		}
	}
	e.lastChar[0] = data[len(data)-1]

	if e.state == PSDetecting {
		if e.contextAnalyzer.gotEnoughData() &&
			e.getConfidence() > CPShortcutThreshold {
			e.state = PSFound
		}
	}
	return e.state
}
