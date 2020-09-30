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

/*  A state machine to verify a byte sequence for a particular encoding.
    For each byte the detector receives, it will feed that byte to every active
    state machine available, one byte at a time. The state machine changes its
    state based on its previous state and the byte it receives. There are 3
    states in a state machine that are of interest to an auto-detector:
    - MSStart : This is the state to sstart with, or a legal byte sequence
			    (i.e. a valid code point) for each character has been identified.
	- MSItsMe : This indicates that the state machine identified a byte sequence
				that is specific to the charset it is designed for and that
				there is no other possible encoding which can contain this byte
				sequence. This will lead to an immediate positive answer for the
				detector.
	- MSError : This indicates the state machine identified an illegal byte sequence
				for the encoding. This will lead to an immediate negative answer for
				this encoding. Detector will exclude this encoding from consideration
				from here on.
*/
type codingStateMachine struct {
	model       Model
	currBytePos int
	currCharLen int
	currState   MachineState
	active      bool
}

func newCodingStateMachine(model Model) *codingStateMachine {
	return &codingStateMachine{
		model,
		0,
		0,
		MSStart,
		true,
	}
}

func (c *codingStateMachine) reset() {
	c.currState = MSStart
	c.active = true
}

func (c *codingStateMachine) nextState(chr byte) MachineState {
	// For each byte we get its class
	// If it is first byte we also get byte length
	byteClass := c.model.classTable[int(chr)]
	if c.currState == MSStart {
		c.currBytePos = 0
		c.currCharLen = c.model.charLenTable[byteClass]
	}
	// From byte's class and state_table, we get its next state
	currState := int(c.currState)*c.model.classFactor + byteClass
	c.currState = c.model.stateTable[currState]
	c.currBytePos++
	return c.currState

}

func (c *codingStateMachine) getCurrentCharLen() int {
	return c.currCharLen
}

func (c *codingStateMachine) getCodingStateMachine() string {
	return c.model.name
}

func (c *codingStateMachine) language() string {
	return c.model.language
}
