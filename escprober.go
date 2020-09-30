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

// EscCharSetProber Uses a "code scheme" approach for detecting encodings
// Whereby easily recognizable escape or shift sequences are relied on
// to identify these encodings
type EscCharSetProber struct {
	CharSetProber
	codingSms        []*codingStateMachine
	activeSmCount    int
	detectedCharset  string
	detectedLanguage string
}

func newEscCharSetProber(langFilter LanguageFilter) Prober {
	p := EscCharSetProber{
		*newCharSetProber(langFilter),
		[]*codingStateMachine{},
		0,
		"",
		"",
	}
	if langFilter != LFNone {
		if langFilter&LFChineseSimplified > 0 {
			p.codingSms = append(p.codingSms, newCodingStateMachine(HZSmModel))
			p.codingSms = append(p.codingSms, newCodingStateMachine(ISO2022CNSmModel))
		}
		if langFilter&LFJapanese > 0 {
			p.codingSms = append(p.codingSms, newCodingStateMachine(ISO2022JPSmModel))
		}
		if langFilter&LFKorean > 0 {
			p.codingSms = append(p.codingSms, newCodingStateMachine(ISO2022KRSmModel))
		}
	}
	p.reset()
	return &p
}

func (e *EscCharSetProber) reset() {
	e.CharSetProber.reset()
	for _, c := range e.codingSms {
		if c != nil {
			c.active = true
			c.reset()
		}
	}
	e.activeSmCount = len(e.codingSms)
	e.detectedCharset = ""
	e.detectedLanguage = ""
}

func (e *EscCharSetProber) charsetName() string {
	return e.detectedCharset
}

func (e *EscCharSetProber) language() string {
	return e.detectedLanguage
}

func (e *EscCharSetProber) getConfidence() float64 {
	if len(e.detectedCharset) > 0 {
		return 0.99
	}
	return 0.0
}

func (e *EscCharSetProber) feed(data []byte) ProbingState {
	for _, chr := range data {
		for _, sm := range e.codingSms {
			if sm == nil || !sm.active {
				continue
			}
			cState := sm.nextState(chr)
			if cState == MSError {
				sm.active = false
				e.activeSmCount--
				if e.activeSmCount <= 0 {
					e.state = PSNotMe
					return e.state
				}
			} else if cState == MSItsMe {
				e.state = PSFound
				e.detectedCharset = sm.getCodingStateMachine()
				e.detectedLanguage = sm.language()
				return e.state
			}
		}
	}
	return e.state
}
