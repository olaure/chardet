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

// CP949Prober prober
type CP949Prober struct {
	MultiByteCharSetProber
}

func newCP949Prober() *CP949Prober {
	mbcsp := newMultiByteCharSetProber(LFNone)
	// Note: CP949 is a superset of EUC-KR so the distribution should
	// not be different
	mbcsp.distributionAnalyzer = newEUCKRDistributionAnalysis()
	mbcsp.codingSm = newCodingStateMachine(CP949SmModel)
	prober := CP949Prober{*mbcsp}

	prober.reset()
	return &prober
}

func (b *CP949Prober) charsetName() string {
	return "CP949"
}

func (b *CP949Prober) language() string {
	return "Korean"
}
