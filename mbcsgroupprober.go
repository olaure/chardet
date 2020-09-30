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

// MBCSGroupProber regroups all group probers
type MBCSGroupProber struct {
	CharSetGroupProber
}

func newMBCSGroupProber(langFilter LanguageFilter) *MBCSGroupProber {
	csgp := newCharSetGroupProber(langFilter)
	csgp.probers = []Prober{
		newUTF8Prober(),
		newSJISProber(),
		newEUCJPProber(),
		newGB2312Prober(),
		newEUCKRProber(),
		newCP949Prober(),
		newBIG5Prober(),
		newEUCTWProber(),
	}

	mgp := MBCSGroupProber{
		*csgp,
	}
	mgp.reset()
	return &mgp
}
