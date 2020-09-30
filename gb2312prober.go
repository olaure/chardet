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

// GB2312Prober prober
type GB2312Prober struct {
	MultiByteCharSetProber
}

func newGB2312Prober() *GB2312Prober {
	mbcsp := newMultiByteCharSetProber(LFNone)
	mbcsp.distributionAnalyzer = newGB2312DistributionAnalysis()
	mbcsp.codingSm = newCodingStateMachine(GB2312SmModel)
	prober := GB2312Prober{*mbcsp}

	prober.reset()
	return &prober
}

func (b *GB2312Prober) charsetName() string {
	return "GB2312"
}

func (b *GB2312Prober) language() string {
	return "Chinese"
}
