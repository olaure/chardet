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

// HZCls class
var HZCls = []int{
	1, 0, 0, 0, 0, 0, 0, 0, // 00 - 07
	0, 0, 0, 0, 0, 0, 0, 0, // 08 - 0f
	0, 0, 0, 0, 0, 0, 0, 0, // 10 - 17
	0, 0, 0, 1, 0, 0, 0, 0, // 18 - 1f
	0, 0, 0, 0, 0, 0, 0, 0, // 20 - 27
	0, 0, 0, 0, 0, 0, 0, 0, // 28 - 2f
	0, 0, 0, 0, 0, 0, 0, 0, // 30 - 37
	0, 0, 0, 0, 0, 0, 0, 0, // 38 - 3f
	0, 0, 0, 0, 0, 0, 0, 0, // 40 - 47
	0, 0, 0, 0, 0, 0, 0, 0, // 48 - 4f
	0, 0, 0, 0, 0, 0, 0, 0, // 50 - 57
	0, 0, 0, 0, 0, 0, 0, 0, // 58 - 5f
	0, 0, 0, 0, 0, 0, 0, 0, // 60 - 67
	0, 0, 0, 0, 0, 0, 0, 0, // 68 - 6f
	0, 0, 0, 0, 0, 0, 0, 0, // 70 - 77
	0, 0, 0, 4, 0, 5, 2, 0, // 78 - 7f
	1, 1, 1, 1, 1, 1, 1, 1, // 80 - 87
	1, 1, 1, 1, 1, 1, 1, 1, // 88 - 8f
	1, 1, 1, 1, 1, 1, 1, 1, // 90 - 97
	1, 1, 1, 1, 1, 1, 1, 1, // 98 - 9f
	1, 1, 1, 1, 1, 1, 1, 1, // a0 - a7
	1, 1, 1, 1, 1, 1, 1, 1, // a8 - af
	1, 1, 1, 1, 1, 1, 1, 1, // b0 - b7
	1, 1, 1, 1, 1, 1, 1, 1, // b8 - bf
	1, 1, 1, 1, 1, 1, 1, 1, // c0 - c7
	1, 1, 1, 1, 1, 1, 1, 1, // c8 - cf
	1, 1, 1, 1, 1, 1, 1, 1, // d0 - d7
	1, 1, 1, 1, 1, 1, 1, 1, // d8 - df
	1, 1, 1, 1, 1, 1, 1, 1, // e0 - e7
	1, 1, 1, 1, 1, 1, 1, 1, // e8 - ef
	1, 1, 1, 1, 1, 1, 1, 1, // f0 - f7
	1, 1, 1, 1, 1, 1, 1, 1, // f8 - ff
}

// HZStates states
var HZStates = []MachineState{
	MSStart, MSError, 3, MSStart, MSStart, MSStart, MSError, MSError, // 00-07
	MSError, MSError, MSError, MSError, MSItsMe, MSItsMe, MSItsMe, MSItsMe, // 08-0f
	MSItsMe, MSItsMe, MSError, MSError, MSStart, MSStart, 4, MSError, // 10-17
	5, MSError, 6, MSError, 5, 5, 4, MSError, // 18-1f
	4, MSError, 4, 4, 4, MSError, 4, MSError, // 20-27
	4, MSItsMe, MSStart, MSStart, MSStart, MSStart, MSStart, MSStart, // 28-2f
}

//HZCharLenTable HZ
var HZCharLenTable = []int{0, 0, 0, 0, 0, 0}

// HZSmModel HZ Model
var HZSmModel = Model{
	classTable:   HZCls,
	classFactor:  6,
	stateTable:   HZStates,
	charLenTable: HZCharLenTable,
	name:         "HZ-GB-2312",
	language:     "Chinese",
}

// ISO2022CNCls classes
var ISO2022CNCls = []int{
	2, 0, 0, 0, 0, 0, 0, 0, // 00 - 07
	0, 0, 0, 0, 0, 0, 0, 0, // 08 - 0f
	0, 0, 0, 0, 0, 0, 0, 0, // 10 - 17
	0, 0, 0, 1, 0, 0, 0, 0, // 18 - 1f
	0, 0, 0, 0, 0, 0, 0, 0, // 20 - 27
	0, 3, 0, 0, 0, 0, 0, 0, // 28 - 2f
	0, 0, 0, 0, 0, 0, 0, 0, // 30 - 37
	0, 0, 0, 0, 0, 0, 0, 0, // 38 - 3f
	0, 0, 0, 4, 0, 0, 0, 0, // 40 - 47
	0, 0, 0, 0, 0, 0, 0, 0, // 48 - 4f
	0, 0, 0, 0, 0, 0, 0, 0, // 50 - 57
	0, 0, 0, 0, 0, 0, 0, 0, // 58 - 5f
	0, 0, 0, 0, 0, 0, 0, 0, // 60 - 67
	0, 0, 0, 0, 0, 0, 0, 0, // 68 - 6f
	0, 0, 0, 0, 0, 0, 0, 0, // 70 - 77
	0, 0, 0, 0, 0, 0, 0, 0, // 78 - 7f
	2, 2, 2, 2, 2, 2, 2, 2, // 80 - 87
	2, 2, 2, 2, 2, 2, 2, 2, // 88 - 8f
	2, 2, 2, 2, 2, 2, 2, 2, // 90 - 97
	2, 2, 2, 2, 2, 2, 2, 2, // 98 - 9f
	2, 2, 2, 2, 2, 2, 2, 2, // a0 - a7
	2, 2, 2, 2, 2, 2, 2, 2, // a8 - af
	2, 2, 2, 2, 2, 2, 2, 2, // b0 - b7
	2, 2, 2, 2, 2, 2, 2, 2, // b8 - bf
	2, 2, 2, 2, 2, 2, 2, 2, // c0 - c7
	2, 2, 2, 2, 2, 2, 2, 2, // c8 - cf
	2, 2, 2, 2, 2, 2, 2, 2, // d0 - d7
	2, 2, 2, 2, 2, 2, 2, 2, // d8 - df
	2, 2, 2, 2, 2, 2, 2, 2, // e0 - e7
	2, 2, 2, 2, 2, 2, 2, 2, // e8 - ef
	2, 2, 2, 2, 2, 2, 2, 2, // f0 - f7
	2, 2, 2, 2, 2, 2, 2, 2, // f8 - ff
}

// ISO2022CNStates states
var ISO2022CNStates = []MachineState{
	MSStart, 3, MSError, MSStart, MSStart, MSStart, MSStart, MSStart, // 00-07
	MSStart, MSError, MSError, MSError, MSError, MSError, MSError, MSError, // 08-0f
	MSError, MSError, MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSItsMe, // 10-17
	MSItsMe, MSItsMe, MSItsMe, MSError, MSError, MSError, 4, MSError, // 18-1f
	MSError, MSError, MSError, MSItsMe, MSError, MSError, MSError, MSError, // 20-27
	5, 6, MSError, MSError, MSError, MSError, MSError, MSError, // 28-2f
	MSError, MSError, MSError, MSItsMe, MSError, MSError, MSError, MSError, // 30-37
	MSError, MSError, MSError, MSError, MSError, MSItsMe, MSError, MSStart, // 38-3f
}

// ISO2022CNCharLenTable ISO2022 CN
var ISO2022CNCharLenTable = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

// ISO2022CNSmModel model
var ISO2022CNSmModel = Model{
	classTable:   ISO2022CNCls,
	classFactor:  9,
	stateTable:   ISO2022CNStates,
	charLenTable: ISO2022CNCharLenTable,
	name:         "ISO-2022-CN",
	language:     "Chinese",
}

// ISO2022JPCls classes
var ISO2022JPCls = []int{
	2, 0, 0, 0, 0, 0, 0, 0, // 00 - 07
	0, 0, 0, 0, 0, 0, 2, 2, // 08 - 0f
	0, 0, 0, 0, 0, 0, 0, 0, // 10 - 17
	0, 0, 0, 1, 0, 0, 0, 0, // 18 - 1f
	0, 0, 0, 0, 7, 0, 0, 0, // 20 - 27
	3, 0, 0, 0, 0, 0, 0, 0, // 28 - 2f
	0, 0, 0, 0, 0, 0, 0, 0, // 30 - 37
	0, 0, 0, 0, 0, 0, 0, 0, // 38 - 3f
	6, 0, 4, 0, 8, 0, 0, 0, // 40 - 47
	0, 9, 5, 0, 0, 0, 0, 0, // 48 - 4f
	0, 0, 0, 0, 0, 0, 0, 0, // 50 - 57
	0, 0, 0, 0, 0, 0, 0, 0, // 58 - 5f
	0, 0, 0, 0, 0, 0, 0, 0, // 60 - 67
	0, 0, 0, 0, 0, 0, 0, 0, // 68 - 6f
	0, 0, 0, 0, 0, 0, 0, 0, // 70 - 77
	0, 0, 0, 0, 0, 0, 0, 0, // 78 - 7f
	2, 2, 2, 2, 2, 2, 2, 2, // 80 - 87
	2, 2, 2, 2, 2, 2, 2, 2, // 88 - 8f
	2, 2, 2, 2, 2, 2, 2, 2, // 90 - 97
	2, 2, 2, 2, 2, 2, 2, 2, // 98 - 9f
	2, 2, 2, 2, 2, 2, 2, 2, // a0 - a7
	2, 2, 2, 2, 2, 2, 2, 2, // a8 - af
	2, 2, 2, 2, 2, 2, 2, 2, // b0 - b7
	2, 2, 2, 2, 2, 2, 2, 2, // b8 - bf
	2, 2, 2, 2, 2, 2, 2, 2, // c0 - c7
	2, 2, 2, 2, 2, 2, 2, 2, // c8 - cf
	2, 2, 2, 2, 2, 2, 2, 2, // d0 - d7
	2, 2, 2, 2, 2, 2, 2, 2, // d8 - df
	2, 2, 2, 2, 2, 2, 2, 2, // e0 - e7
	2, 2, 2, 2, 2, 2, 2, 2, // e8 - ef
	2, 2, 2, 2, 2, 2, 2, 2, // f0 - f7
	2, 2, 2, 2, 2, 2, 2, 2, // f8 - ff
}

// ISO2022JPStates states
var ISO2022JPStates = []MachineState{
	MSStart, 3, MSError, MSStart, MSStart, MSStart, MSStart, MSStart, // 00-07
	MSStart, MSStart, MSError, MSError, MSError, MSError, MSError, MSError, // 08-0f
	MSError, MSError, MSError, MSError, MSItsMe, MSItsMe, MSItsMe, MSItsMe, // 10-17
	MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSError, MSError, // 18-1f
	MSError, 5, MSError, MSError, MSError, 4, MSError, MSError, // 20-27
	MSError, MSError, MSError, 6, MSItsMe, MSError, MSItsMe, MSError, // 28-2f
	MSError, MSError, MSError, MSError, MSError, MSError, MSItsMe, MSItsMe, // 30-37
	MSError, MSError, MSError, MSItsMe, MSError, MSError, MSError, MSError, // 38-3f
	MSError, MSError, MSError, MSError, MSItsMe, MSError, MSStart, MSStart, // 40-47
}

// ISO2022JPCharLenTable JP
var ISO2022JPCharLenTable = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// ISO2022JPSmModel JP model
var ISO2022JPSmModel = Model{
	classTable:   ISO2022JPCls,
	classFactor:  10,
	stateTable:   ISO2022JPStates,
	charLenTable: ISO2022JPCharLenTable,
	name:         "ISO-2022-JP",
	language:     "Japanese",
}

// ISO2022KRCls class
var ISO2022KRCls = []int{
	2, 0, 0, 0, 0, 0, 0, 0, // 00 - 07
	0, 0, 0, 0, 0, 0, 0, 0, // 08 - 0f
	0, 0, 0, 0, 0, 0, 0, 0, // 10 - 17
	0, 0, 0, 1, 0, 0, 0, 0, // 18 - 1f
	0, 0, 0, 0, 3, 0, 0, 0, // 20 - 27
	0, 4, 0, 0, 0, 0, 0, 0, // 28 - 2f
	0, 0, 0, 0, 0, 0, 0, 0, // 30 - 37
	0, 0, 0, 0, 0, 0, 0, 0, // 38 - 3f
	0, 0, 0, 5, 0, 0, 0, 0, // 40 - 47
	0, 0, 0, 0, 0, 0, 0, 0, // 48 - 4f
	0, 0, 0, 0, 0, 0, 0, 0, // 50 - 57
	0, 0, 0, 0, 0, 0, 0, 0, // 58 - 5f
	0, 0, 0, 0, 0, 0, 0, 0, // 60 - 67
	0, 0, 0, 0, 0, 0, 0, 0, // 68 - 6f
	0, 0, 0, 0, 0, 0, 0, 0, // 70 - 77
	0, 0, 0, 0, 0, 0, 0, 0, // 78 - 7f
	2, 2, 2, 2, 2, 2, 2, 2, // 80 - 87
	2, 2, 2, 2, 2, 2, 2, 2, // 88 - 8f
	2, 2, 2, 2, 2, 2, 2, 2, // 90 - 97
	2, 2, 2, 2, 2, 2, 2, 2, // 98 - 9f
	2, 2, 2, 2, 2, 2, 2, 2, // a0 - a7
	2, 2, 2, 2, 2, 2, 2, 2, // a8 - af
	2, 2, 2, 2, 2, 2, 2, 2, // b0 - b7
	2, 2, 2, 2, 2, 2, 2, 2, // b8 - bf
	2, 2, 2, 2, 2, 2, 2, 2, // c0 - c7
	2, 2, 2, 2, 2, 2, 2, 2, // c8 - cf
	2, 2, 2, 2, 2, 2, 2, 2, // d0 - d7
	2, 2, 2, 2, 2, 2, 2, 2, // d8 - df
	2, 2, 2, 2, 2, 2, 2, 2, // e0 - e7
	2, 2, 2, 2, 2, 2, 2, 2, // e8 - ef
	2, 2, 2, 2, 2, 2, 2, 2, // f0 - f7
	2, 2, 2, 2, 2, 2, 2, 2, // f8 - ff
}

// ISO2022KRStates prev ISO2022KR_ST
var ISO2022KRStates = []MachineState{
	MSStart, 3, MSError, MSStart, MSStart, MSStart, MSError, MSError, // 00-07
	MSError, MSError, MSError, MSError, MSItsMe, MSItsMe, MSItsMe, MSItsMe, // 08-0f
	MSItsMe, MSItsMe, MSError, MSError, MSError, 4, MSError, MSError, // 10-17
	MSError, MSError, MSError, MSError, 5, MSError, MSError, MSError, // 18-1f
	MSError, MSError, MSError, MSItsMe, MSStart, MSStart, MSStart, MSStart, // 20-27
}

// ISO2022KRCharLenTable KR
var ISO2022KRCharLenTable = []int{0, 0, 0, 0, 0, 0}

// ISO2022KRSmModel KR Model
var ISO2022KRSmModel = Model{
	classTable:   ISO2022KRCls,
	classFactor:  6,
	stateTable:   ISO2022KRStates,
	charLenTable: ISO2022KRCharLenTable,
	name:         "ISO-2022-KR",
	language:     "Korean",
}
