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

// EUCTWProber prober
type EUCTWProber struct {
	MultiByteCharSetProber
}

func newEUCTWProber() *EUCTWProber {
	mbcsp := newMultiByteCharSetProber(LFNone)
	mbcsp.distributionAnalyzer = newEUCTWDistributionAnalysis()
	mbcsp.codingSm = newCodingStateMachine(EUCTWSmModel)
	prober := EUCTWProber{*mbcsp}

	prober.reset()
	return &prober
}

func (b *EUCTWProber) charsetName() string {
	return "EUC-TW"
}

func (b *EUCTWProber) language() string {
	return "Taiwan"
}
