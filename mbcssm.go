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

// Model struct
type Model struct {
	language     string
	name         string
	classTable   []int
	charLenTable []int
	classFactor  int
	stateTable   []MachineState
}

// BIG5

// BIG5Cls previously BIG5_CLS
var BIG5Cls = []int{
	1, 1, 1, 1, 1, 1, 1, 1, // 00 - 07    #allow 0x00 as legal value
	1, 1, 1, 1, 1, 1, 0, 0, // 08 - 0f
	1, 1, 1, 1, 1, 1, 1, 1, // 10 - 17
	1, 1, 1, 0, 1, 1, 1, 1, // 18 - 1f
	1, 1, 1, 1, 1, 1, 1, 1, // 20 - 27
	1, 1, 1, 1, 1, 1, 1, 1, // 28 - 2f
	1, 1, 1, 1, 1, 1, 1, 1, // 30 - 37
	1, 1, 1, 1, 1, 1, 1, 1, // 38 - 3f
	2, 2, 2, 2, 2, 2, 2, 2, // 40 - 47
	2, 2, 2, 2, 2, 2, 2, 2, // 48 - 4f
	2, 2, 2, 2, 2, 2, 2, 2, // 50 - 57
	2, 2, 2, 2, 2, 2, 2, 2, // 58 - 5f
	2, 2, 2, 2, 2, 2, 2, 2, // 60 - 67
	2, 2, 2, 2, 2, 2, 2, 2, // 68 - 6f
	2, 2, 2, 2, 2, 2, 2, 2, // 70 - 77
	2, 2, 2, 2, 2, 2, 2, 1, // 78 - 7f
	4, 4, 4, 4, 4, 4, 4, 4, // 80 - 87
	4, 4, 4, 4, 4, 4, 4, 4, // 88 - 8f
	4, 4, 4, 4, 4, 4, 4, 4, // 90 - 97
	4, 4, 4, 4, 4, 4, 4, 4, // 98 - 9f
	4, 3, 3, 3, 3, 3, 3, 3, // a0 - a7
	3, 3, 3, 3, 3, 3, 3, 3, // a8 - af
	3, 3, 3, 3, 3, 3, 3, 3, // b0 - b7
	3, 3, 3, 3, 3, 3, 3, 3, // b8 - bf
	3, 3, 3, 3, 3, 3, 3, 3, // c0 - c7
	3, 3, 3, 3, 3, 3, 3, 3, // c8 - cf
	3, 3, 3, 3, 3, 3, 3, 3, // d0 - d7
	3, 3, 3, 3, 3, 3, 3, 3, // d8 - df
	3, 3, 3, 3, 3, 3, 3, 3, // e0 - e7
	3, 3, 3, 3, 3, 3, 3, 3, // e8 - ef
	3, 3, 3, 3, 3, 3, 3, 3, // f0 - f7
	3, 3, 3, 3, 3, 3, 3, 0, // f8 - ff
}

// BIG5States states, previously BIG5_ST
var BIG5States = []MachineState{
	MSError, MSStart, MSStart, 3, MSError, MSError, MSError, MSError, //00-07
	MSError, MSError, MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSError, //08-0f
	MSError, MSStart, MSStart, MSStart, MSStart, MSStart, MSStart, MSStart, //10-17
}

// BIG5CharLenTable Previously BIG5_CHAR_LEN_TABLE
var BIG5CharLenTable = []int{0, 1, 1, 2, 0}

// BIG5SmModel previously BIG5_SM_MODEL
var BIG5SmModel = Model{
	name:         "Big5",
	classTable:   BIG5Cls,
	charLenTable: BIG5CharLenTable,
	classFactor:  5,
	stateTable:   BIG5States,
}

// CP949

// CP949Cls prev CP949_CLS
var CP949Cls = []int{
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, // 00 - 0f
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, // 10 - 1f
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // 20 - 2f
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // 30 - 3f
	1, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, // 40 - 4f
	4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 1, 1, 1, // 50 - 5f
	1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, // 60 - 6f
	5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 1, 1, 1, // 70 - 7f
	0, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, // 80 - 8f
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, // 90 - 9f
	6, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 8, 8, 8, // a0 - af
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, // b0 - bf
	7, 7, 7, 7, 7, 7, 9, 2, 2, 3, 2, 2, 2, 2, 2, 2, // c0 - cf
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, // d0 - df
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, // e0 - ef
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0, // f0 - ff
}

// CP949States prev CP949_ST
var CP949States = []MachineState{
	//cls=    0      1      2      3      4      5      6      7      8      9  // previous state =
	MSError, MSStart, 3, MSError, MSStart, MSStart, 4, 5, MSError, 6, // MSStart
	MSError, MSError, MSError, MSError, MSError, MSError, MSError, MSError, MSError, MSError, // MSError
	MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSItsMe, // MSItsMe
	MSError, MSError, MSStart, MSStart, MSError, MSError, MSError, MSStart, MSStart, MSStart, // 3
	MSError, MSError, MSStart, MSStart, MSStart, MSStart, MSStart, MSStart, MSStart, MSStart, // 4
	MSError, MSStart, MSStart, MSStart, MSStart, MSStart, MSStart, MSStart, MSStart, MSStart, // 5
	MSError, MSStart, MSStart, MSStart, MSStart, MSError, MSError, MSStart, MSStart, MSStart, // 6
}

// CP949CharLenTable prev CP949_CHAR_LEN_TABLE
var CP949CharLenTable = []int{0, 1, 2, 0, 1, 1, 2, 2, 0, 2}

// CP949SmModel previously CP949_SM_MODEL
var CP949SmModel = Model{
	classTable:   CP949Cls,
	classFactor:  10,
	stateTable:   CP949States,
	charLenTable: CP949CharLenTable,
	name:         "CP949",
}

// EUC-JP

// EUCJPCls Prev EUCJP_CLS
var EUCJPCls = []int{
	4, 4, 4, 4, 4, 4, 4, 4, // 00 - 07
	4, 4, 4, 4, 4, 4, 5, 5, // 08 - 0f
	4, 4, 4, 4, 4, 4, 4, 4, // 10 - 17
	4, 4, 4, 5, 4, 4, 4, 4, // 18 - 1f
	4, 4, 4, 4, 4, 4, 4, 4, // 20 - 27
	4, 4, 4, 4, 4, 4, 4, 4, // 28 - 2f
	4, 4, 4, 4, 4, 4, 4, 4, // 30 - 37
	4, 4, 4, 4, 4, 4, 4, 4, // 38 - 3f
	4, 4, 4, 4, 4, 4, 4, 4, // 40 - 47
	4, 4, 4, 4, 4, 4, 4, 4, // 48 - 4f
	4, 4, 4, 4, 4, 4, 4, 4, // 50 - 57
	4, 4, 4, 4, 4, 4, 4, 4, // 58 - 5f
	4, 4, 4, 4, 4, 4, 4, 4, // 60 - 67
	4, 4, 4, 4, 4, 4, 4, 4, // 68 - 6f
	4, 4, 4, 4, 4, 4, 4, 4, // 70 - 77
	4, 4, 4, 4, 4, 4, 4, 4, // 78 - 7f
	5, 5, 5, 5, 5, 5, 5, 5, // 80 - 87
	5, 5, 5, 5, 5, 5, 1, 3, // 88 - 8f
	5, 5, 5, 5, 5, 5, 5, 5, // 90 - 97
	5, 5, 5, 5, 5, 5, 5, 5, // 98 - 9f
	5, 2, 2, 2, 2, 2, 2, 2, // a0 - a7
	2, 2, 2, 2, 2, 2, 2, 2, // a8 - af
	2, 2, 2, 2, 2, 2, 2, 2, // b0 - b7
	2, 2, 2, 2, 2, 2, 2, 2, // b8 - bf
	2, 2, 2, 2, 2, 2, 2, 2, // c0 - c7
	2, 2, 2, 2, 2, 2, 2, 2, // c8 - cf
	2, 2, 2, 2, 2, 2, 2, 2, // d0 - d7
	2, 2, 2, 2, 2, 2, 2, 2, // d8 - df
	0, 0, 0, 0, 0, 0, 0, 0, // e0 - e7
	0, 0, 0, 0, 0, 0, 0, 0, // e8 - ef
	0, 0, 0, 0, 0, 0, 0, 0, // f0 - f7
	0, 0, 0, 0, 0, 0, 0, 5, // f8 - ff
}

// EUCJPStates prev EUCJP_ST
var EUCJPStates = []MachineState{
	3, 4, 3, 5, MSStart, MSError, MSError, MSError, //00-07
	MSError, MSError, MSError, MSError, MSItsMe, MSItsMe, MSItsMe, MSItsMe, //08-0f
	MSItsMe, MSItsMe, MSStart, MSError, MSStart, MSError, MSError, MSError, //10-17
	MSError, MSError, MSStart, MSError, MSError, MSError, 3, MSError, //18-1f
	3, MSError, MSError, MSError, MSStart, MSStart, MSStart, MSStart, //20-27
}

// EUCJPCharLenTable prev EUCJP_CHAR_LEN_TABLE
var EUCJPCharLenTable = []int{2, 2, 2, 3, 1, 0}

// EUCJPSmModel prev EUCJP_SM_MODEL
var EUCJPSmModel = Model{
	classTable:   EUCJPCls,
	classFactor:  6,
	stateTable:   EUCJPStates,
	charLenTable: EUCJPCharLenTable,
	name:         "EUC-JP",
}

// EUC-KR

// EUCKRCls EUCKR Classes
var EUCKRCls = []int{
	1, 1, 1, 1, 1, 1, 1, 1, // 00 - 07
	1, 1, 1, 1, 1, 1, 0, 0, // 08 - 0f
	1, 1, 1, 1, 1, 1, 1, 1, // 10 - 17
	1, 1, 1, 0, 1, 1, 1, 1, // 18 - 1f
	1, 1, 1, 1, 1, 1, 1, 1, // 20 - 27
	1, 1, 1, 1, 1, 1, 1, 1, // 28 - 2f
	1, 1, 1, 1, 1, 1, 1, 1, // 30 - 37
	1, 1, 1, 1, 1, 1, 1, 1, // 38 - 3f
	1, 1, 1, 1, 1, 1, 1, 1, // 40 - 47
	1, 1, 1, 1, 1, 1, 1, 1, // 48 - 4f
	1, 1, 1, 1, 1, 1, 1, 1, // 50 - 57
	1, 1, 1, 1, 1, 1, 1, 1, // 58 - 5f
	1, 1, 1, 1, 1, 1, 1, 1, // 60 - 67
	1, 1, 1, 1, 1, 1, 1, 1, // 68 - 6f
	1, 1, 1, 1, 1, 1, 1, 1, // 70 - 77
	1, 1, 1, 1, 1, 1, 1, 1, // 78 - 7f
	0, 0, 0, 0, 0, 0, 0, 0, // 80 - 87
	0, 0, 0, 0, 0, 0, 0, 0, // 88 - 8f
	0, 0, 0, 0, 0, 0, 0, 0, // 90 - 97
	0, 0, 0, 0, 0, 0, 0, 0, // 98 - 9f
	0, 2, 2, 2, 2, 2, 2, 2, // a0 - a7
	2, 2, 2, 2, 2, 3, 3, 3, // a8 - af
	2, 2, 2, 2, 2, 2, 2, 2, // b0 - b7
	2, 2, 2, 2, 2, 2, 2, 2, // b8 - bf
	2, 2, 2, 2, 2, 2, 2, 2, // c0 - c7
	2, 3, 2, 2, 2, 2, 2, 2, // c8 - cf
	2, 2, 2, 2, 2, 2, 2, 2, // d0 - d7
	2, 2, 2, 2, 2, 2, 2, 2, // d8 - df
	2, 2, 2, 2, 2, 2, 2, 2, // e0 - e7
	2, 2, 2, 2, 2, 2, 2, 2, // e8 - ef
	2, 2, 2, 2, 2, 2, 2, 2, // f0 - f7
	2, 2, 2, 2, 2, 2, 2, 0, // f8 - ff
}

// EUCKRStates EUCKR States
var EUCKRStates = []MachineState{
	MSError, MSStart, 3, MSError, MSError, MSError, MSError, MSError, //00-07
	MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSError, MSError, MSStart, MSStart, //08-0f
}

// EUCKRCharLenTable table lengths
var EUCKRCharLenTable = []int{0, 1, 2, 0}

// EUCKRSmModel State Machine model
var EUCKRSmModel = Model{
	classTable:   EUCKRCls,
	classFactor:  4,
	stateTable:   EUCKRStates,
	charLenTable: EUCKRCharLenTable,
	name:         "EUC-KR"}

// EUC-TW

// EUCTWCls classes
var EUCTWCls = []int{
	2, 2, 2, 2, 2, 2, 2, 2, // 00 - 07
	2, 2, 2, 2, 2, 2, 0, 0, // 08 - 0f
	2, 2, 2, 2, 2, 2, 2, 2, // 10 - 17
	2, 2, 2, 0, 2, 2, 2, 2, // 18 - 1f
	2, 2, 2, 2, 2, 2, 2, 2, // 20 - 27
	2, 2, 2, 2, 2, 2, 2, 2, // 28 - 2f
	2, 2, 2, 2, 2, 2, 2, 2, // 30 - 37
	2, 2, 2, 2, 2, 2, 2, 2, // 38 - 3f
	2, 2, 2, 2, 2, 2, 2, 2, // 40 - 47
	2, 2, 2, 2, 2, 2, 2, 2, // 48 - 4f
	2, 2, 2, 2, 2, 2, 2, 2, // 50 - 57
	2, 2, 2, 2, 2, 2, 2, 2, // 58 - 5f
	2, 2, 2, 2, 2, 2, 2, 2, // 60 - 67
	2, 2, 2, 2, 2, 2, 2, 2, // 68 - 6f
	2, 2, 2, 2, 2, 2, 2, 2, // 70 - 77
	2, 2, 2, 2, 2, 2, 2, 2, // 78 - 7f
	0, 0, 0, 0, 0, 0, 0, 0, // 80 - 87
	0, 0, 0, 0, 0, 0, 6, 0, // 88 - 8f
	0, 0, 0, 0, 0, 0, 0, 0, // 90 - 97
	0, 0, 0, 0, 0, 0, 0, 0, // 98 - 9f
	0, 3, 4, 4, 4, 4, 4, 4, // a0 - a7
	5, 5, 1, 1, 1, 1, 1, 1, // a8 - af
	1, 1, 1, 1, 1, 1, 1, 1, // b0 - b7
	1, 1, 1, 1, 1, 1, 1, 1, // b8 - bf
	1, 1, 3, 1, 3, 3, 3, 3, // c0 - c7
	3, 3, 3, 3, 3, 3, 3, 3, // c8 - cf
	3, 3, 3, 3, 3, 3, 3, 3, // d0 - d7
	3, 3, 3, 3, 3, 3, 3, 3, // d8 - df
	3, 3, 3, 3, 3, 3, 3, 3, // e0 - e7
	3, 3, 3, 3, 3, 3, 3, 3, // e8 - ef
	3, 3, 3, 3, 3, 3, 3, 3, // f0 - f7
	3, 3, 3, 3, 3, 3, 3, 0, // f8 - ff
}

// EUCTWStates states
var EUCTWStates = []MachineState{
	MSError, MSError, MSStart, 3, 3, 3, 4, MSError, //00-07
	MSError, MSError, MSError, MSError, MSError, MSError, MSItsMe, MSItsMe, //08-0f
	MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSError, MSStart, MSError, //10-17
	MSStart, MSStart, MSStart, MSError, MSError, MSError, MSError, MSError, //18-1f
	5, MSError, MSError, MSError, MSStart, MSError, MSStart, MSStart, //20-27
	MSStart, MSError, MSStart, MSStart, MSStart, MSStart, MSStart, MSStart, //28-2f
}

// EUCTWCharLenTable charLenTable
var EUCTWCharLenTable = []int{0, 0, 1, 2, 2, 2, 3}

// EUCTWSmModel model
var EUCTWSmModel = Model{
	classTable:   EUCTWCls,
	classFactor:  7,
	stateTable:   EUCTWStates,
	charLenTable: EUCTWCharLenTable,
	name:         "x-euc-tw"}

// GB2312

//GB2312Cls classes
var GB2312Cls = []int{
	1, 1, 1, 1, 1, 1, 1, 1, // 00 - 07
	1, 1, 1, 1, 1, 1, 0, 0, // 08 - 0f
	1, 1, 1, 1, 1, 1, 1, 1, // 10 - 17
	1, 1, 1, 0, 1, 1, 1, 1, // 18 - 1f
	1, 1, 1, 1, 1, 1, 1, 1, // 20 - 27
	1, 1, 1, 1, 1, 1, 1, 1, // 28 - 2f
	3, 3, 3, 3, 3, 3, 3, 3, // 30 - 37
	3, 3, 1, 1, 1, 1, 1, 1, // 38 - 3f
	2, 2, 2, 2, 2, 2, 2, 2, // 40 - 47
	2, 2, 2, 2, 2, 2, 2, 2, // 48 - 4f
	2, 2, 2, 2, 2, 2, 2, 2, // 50 - 57
	2, 2, 2, 2, 2, 2, 2, 2, // 58 - 5f
	2, 2, 2, 2, 2, 2, 2, 2, // 60 - 67
	2, 2, 2, 2, 2, 2, 2, 2, // 68 - 6f
	2, 2, 2, 2, 2, 2, 2, 2, // 70 - 77
	2, 2, 2, 2, 2, 2, 2, 4, // 78 - 7f
	5, 6, 6, 6, 6, 6, 6, 6, // 80 - 87
	6, 6, 6, 6, 6, 6, 6, 6, // 88 - 8f
	6, 6, 6, 6, 6, 6, 6, 6, // 90 - 97
	6, 6, 6, 6, 6, 6, 6, 6, // 98 - 9f
	6, 6, 6, 6, 6, 6, 6, 6, // a0 - a7
	6, 6, 6, 6, 6, 6, 6, 6, // a8 - af
	6, 6, 6, 6, 6, 6, 6, 6, // b0 - b7
	6, 6, 6, 6, 6, 6, 6, 6, // b8 - bf
	6, 6, 6, 6, 6, 6, 6, 6, // c0 - c7
	6, 6, 6, 6, 6, 6, 6, 6, // c8 - cf
	6, 6, 6, 6, 6, 6, 6, 6, // d0 - d7
	6, 6, 6, 6, 6, 6, 6, 6, // d8 - df
	6, 6, 6, 6, 6, 6, 6, 6, // e0 - e7
	6, 6, 6, 6, 6, 6, 6, 6, // e8 - ef
	6, 6, 6, 6, 6, 6, 6, 6, // f0 - f7
	6, 6, 6, 6, 6, 6, 6, 0, // f8 - ff
}

// GB2312States states
var GB2312States = []MachineState{
	MSError, MSStart, MSStart, MSStart, MSStart, MSStart, 3, MSError, //00-07
	MSError, MSError, MSError, MSError, MSError, MSError, MSItsMe, MSItsMe, //08-0f
	MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSError, MSError, MSStart, //10-17
	4, MSError, MSStart, MSStart, MSError, MSError, MSError, MSError, //18-1f
	MSError, MSError, 5, MSError, MSError, MSError, MSItsMe, MSError, //20-27
	MSError, MSError, MSStart, MSStart, MSStart, MSStart, MSStart, MSStart, //28-2f
}

// GB2312CharLenTable charLenTable
// To be accurate, the length of class 6 can be either 2 or 4.
// But it is not necessary to discriminate between the two since
// it is used for frequency analysis only, and we are validating
// each code range there as well. So it is safe to set it to be
// 2 here.
var GB2312CharLenTable = []int{0, 1, 1, 1, 1, 1, 2}

// GB2312SmModel model
var GB2312SmModel = Model{
	classTable:   GB2312Cls,
	classFactor:  7,
	stateTable:   GB2312States,
	charLenTable: GB2312CharLenTable,
	name:         "GB2312"}

// Shift_JIS

// SJISCls classes
var SJISCls = []int{
	1, 1, 1, 1, 1, 1, 1, 1, // 00 - 07
	1, 1, 1, 1, 1, 1, 0, 0, // 08 - 0f
	1, 1, 1, 1, 1, 1, 1, 1, // 10 - 17
	1, 1, 1, 0, 1, 1, 1, 1, // 18 - 1f
	1, 1, 1, 1, 1, 1, 1, 1, // 20 - 27
	1, 1, 1, 1, 1, 1, 1, 1, // 28 - 2f
	1, 1, 1, 1, 1, 1, 1, 1, // 30 - 37
	1, 1, 1, 1, 1, 1, 1, 1, // 38 - 3f
	2, 2, 2, 2, 2, 2, 2, 2, // 40 - 47
	2, 2, 2, 2, 2, 2, 2, 2, // 48 - 4f
	2, 2, 2, 2, 2, 2, 2, 2, // 50 - 57
	2, 2, 2, 2, 2, 2, 2, 2, // 58 - 5f
	2, 2, 2, 2, 2, 2, 2, 2, // 60 - 67
	2, 2, 2, 2, 2, 2, 2, 2, // 68 - 6f
	2, 2, 2, 2, 2, 2, 2, 2, // 70 - 77
	2, 2, 2, 2, 2, 2, 2, 1, // 78 - 7f
	3, 3, 3, 3, 3, 2, 2, 3, // 80 - 87
	3, 3, 3, 3, 3, 3, 3, 3, // 88 - 8f
	3, 3, 3, 3, 3, 3, 3, 3, // 90 - 97
	3, 3, 3, 3, 3, 3, 3, 3, // 98 - 9f
	//0xa0 is illegal in sjis encoding, but some pages does
	//contain such byte. We need to be more error forgiven.
	2, 2, 2, 2, 2, 2, 2, 2, // a0 - a7
	2, 2, 2, 2, 2, 2, 2, 2, // a8 - af
	2, 2, 2, 2, 2, 2, 2, 2, // b0 - b7
	2, 2, 2, 2, 2, 2, 2, 2, // b8 - bf
	2, 2, 2, 2, 2, 2, 2, 2, // c0 - c7
	2, 2, 2, 2, 2, 2, 2, 2, // c8 - cf
	2, 2, 2, 2, 2, 2, 2, 2, // d0 - d7
	2, 2, 2, 2, 2, 2, 2, 2, // d8 - df
	3, 3, 3, 3, 3, 3, 3, 3, // e0 - e7
	3, 3, 3, 3, 3, 4, 4, 4, // e8 - ef
	3, 3, 3, 3, 3, 3, 3, 3, // f0 - f7
	3, 3, 3, 3, 3, 0, 0, 0, // f8 - ff
}

// SJISStates states
var SJISStates = []MachineState{
	MSError, MSStart, MSStart, 3, MSError, MSError, MSError, MSError, //00-07
	MSError, MSError, MSError, MSError, MSItsMe, MSItsMe, MSItsMe, MSItsMe, //08-0f
	MSItsMe, MSItsMe, MSError, MSError, MSStart, MSStart, MSStart, MSStart, //10-17
}

// SJISCharLenTable charLenTable
var SJISCharLenTable = []int{0, 1, 1, 2, 0, 0}

//SJISSmModel model
var SJISSmModel = Model{classTable: SJISCls,
	classFactor:  6,
	stateTable:   SJISStates,
	charLenTable: SJISCharLenTable,
	name:         "Shift_JIS"}

// UCS2-BE

// UCS2BECls classes
var UCS2BECls = []int{
	0, 0, 0, 0, 0, 0, 0, 0, // 00 - 07
	0, 0, 1, 0, 0, 2, 0, 0, // 08 - 0f
	0, 0, 0, 0, 0, 0, 0, 0, // 10 - 17
	0, 0, 0, 3, 0, 0, 0, 0, // 18 - 1f
	0, 0, 0, 0, 0, 0, 0, 0, // 20 - 27
	0, 3, 3, 3, 3, 3, 0, 0, // 28 - 2f
	0, 0, 0, 0, 0, 0, 0, 0, // 30 - 37
	0, 0, 0, 0, 0, 0, 0, 0, // 38 - 3f
	0, 0, 0, 0, 0, 0, 0, 0, // 40 - 47
	0, 0, 0, 0, 0, 0, 0, 0, // 48 - 4f
	0, 0, 0, 0, 0, 0, 0, 0, // 50 - 57
	0, 0, 0, 0, 0, 0, 0, 0, // 58 - 5f
	0, 0, 0, 0, 0, 0, 0, 0, // 60 - 67
	0, 0, 0, 0, 0, 0, 0, 0, // 68 - 6f
	0, 0, 0, 0, 0, 0, 0, 0, // 70 - 77
	0, 0, 0, 0, 0, 0, 0, 0, // 78 - 7f
	0, 0, 0, 0, 0, 0, 0, 0, // 80 - 87
	0, 0, 0, 0, 0, 0, 0, 0, // 88 - 8f
	0, 0, 0, 0, 0, 0, 0, 0, // 90 - 97
	0, 0, 0, 0, 0, 0, 0, 0, // 98 - 9f
	0, 0, 0, 0, 0, 0, 0, 0, // a0 - a7
	0, 0, 0, 0, 0, 0, 0, 0, // a8 - af
	0, 0, 0, 0, 0, 0, 0, 0, // b0 - b7
	0, 0, 0, 0, 0, 0, 0, 0, // b8 - bf
	0, 0, 0, 0, 0, 0, 0, 0, // c0 - c7
	0, 0, 0, 0, 0, 0, 0, 0, // c8 - cf
	0, 0, 0, 0, 0, 0, 0, 0, // d0 - d7
	0, 0, 0, 0, 0, 0, 0, 0, // d8 - df
	0, 0, 0, 0, 0, 0, 0, 0, // e0 - e7
	0, 0, 0, 0, 0, 0, 0, 0, // e8 - ef
	0, 0, 0, 0, 0, 0, 0, 0, // f0 - f7
	0, 0, 0, 0, 0, 0, 4, 5, // f8 - ff
}

// UCS2BEStates states
var UCS2BEStates = []MachineState{
	5, 7, 7, MSError, 4, 3, MSError, MSError, //00-07
	MSError, MSError, MSError, MSError, MSItsMe, MSItsMe, MSItsMe, MSItsMe, //08-0f
	MSItsMe, MSItsMe, 6, 6, 6, 6, MSError, MSError, //10-17
	6, 6, 6, 6, 6, MSItsMe, 6, 6, //18-1f
	6, 6, 6, 6, 5, 7, 7, MSError, //20-27
	5, 8, 6, 6, MSError, 6, 6, 6, //28-2f
	6, 6, 6, 6, MSError, MSError, MSStart, MSStart, //30-37
}

// UCS2BECharLenTable charLenTable
var UCS2BECharLenTable = []int{2, 2, 2, 0, 2, 2}

// UCS2BESmModel Model
var UCS2BESmModel = Model{classTable: UCS2BECls,
	classFactor:  6,
	stateTable:   UCS2BEStates,
	charLenTable: UCS2BECharLenTable,
	name:         "UTF-16BE"}

// UCS2-LE

// UCS2LECls classes
var UCS2LECls = []int{
	0, 0, 0, 0, 0, 0, 0, 0, // 00 - 07
	0, 0, 1, 0, 0, 2, 0, 0, // 08 - 0f
	0, 0, 0, 0, 0, 0, 0, 0, // 10 - 17
	0, 0, 0, 3, 0, 0, 0, 0, // 18 - 1f
	0, 0, 0, 0, 0, 0, 0, 0, // 20 - 27
	0, 3, 3, 3, 3, 3, 0, 0, // 28 - 2f
	0, 0, 0, 0, 0, 0, 0, 0, // 30 - 37
	0, 0, 0, 0, 0, 0, 0, 0, // 38 - 3f
	0, 0, 0, 0, 0, 0, 0, 0, // 40 - 47
	0, 0, 0, 0, 0, 0, 0, 0, // 48 - 4f
	0, 0, 0, 0, 0, 0, 0, 0, // 50 - 57
	0, 0, 0, 0, 0, 0, 0, 0, // 58 - 5f
	0, 0, 0, 0, 0, 0, 0, 0, // 60 - 67
	0, 0, 0, 0, 0, 0, 0, 0, // 68 - 6f
	0, 0, 0, 0, 0, 0, 0, 0, // 70 - 77
	0, 0, 0, 0, 0, 0, 0, 0, // 78 - 7f
	0, 0, 0, 0, 0, 0, 0, 0, // 80 - 87
	0, 0, 0, 0, 0, 0, 0, 0, // 88 - 8f
	0, 0, 0, 0, 0, 0, 0, 0, // 90 - 97
	0, 0, 0, 0, 0, 0, 0, 0, // 98 - 9f
	0, 0, 0, 0, 0, 0, 0, 0, // a0 - a7
	0, 0, 0, 0, 0, 0, 0, 0, // a8 - af
	0, 0, 0, 0, 0, 0, 0, 0, // b0 - b7
	0, 0, 0, 0, 0, 0, 0, 0, // b8 - bf
	0, 0, 0, 0, 0, 0, 0, 0, // c0 - c7
	0, 0, 0, 0, 0, 0, 0, 0, // c8 - cf
	0, 0, 0, 0, 0, 0, 0, 0, // d0 - d7
	0, 0, 0, 0, 0, 0, 0, 0, // d8 - df
	0, 0, 0, 0, 0, 0, 0, 0, // e0 - e7
	0, 0, 0, 0, 0, 0, 0, 0, // e8 - ef
	0, 0, 0, 0, 0, 0, 0, 0, // f0 - f7
	0, 0, 0, 0, 0, 0, 4, 5, // f8 - ff
}

// UCS2LEStates states
var UCS2LEStates = []MachineState{
	6, 6, 7, 6, 4, 3, MSError, MSError, //00-07
	MSError, MSError, MSError, MSError, MSItsMe, MSItsMe, MSItsMe, MSItsMe, //08-0f
	MSItsMe, MSItsMe, 5, 5, 5, MSError, MSItsMe, MSError, //10-17
	5, 5, 5, MSError, 5, MSError, 6, 6, //18-1f
	7, 6, 8, 8, 5, 5, 5, MSError, //20-27
	5, 5, 5, MSError, MSError, MSError, 5, 5, //28-2f
	5, 5, 5, MSError, 5, MSError, MSStart, MSStart, //30-37
}

// UCS2LECharLenTable charLenTable
var UCS2LECharLenTable = []int{2, 2, 2, 2, 2, 2}

// UCS2LESmModel model
var UCS2LESmModel = Model{
	classTable:   UCS2LECls,
	classFactor:  6,
	stateTable:   UCS2LEStates,
	charLenTable: UCS2LECharLenTable,
	name:         "UTF-16LE"}

// UTF-8

//UTF8Cls classes
var UTF8Cls = []int{
	1, 1, 1, 1, 1, 1, 1, 1, // 00 - 07  //allow 0x00 as a legal value
	1, 1, 1, 1, 1, 1, 0, 0, // 08 - 0f
	1, 1, 1, 1, 1, 1, 1, 1, // 10 - 17
	1, 1, 1, 0, 1, 1, 1, 1, // 18 - 1f
	1, 1, 1, 1, 1, 1, 1, 1, // 20 - 27
	1, 1, 1, 1, 1, 1, 1, 1, // 28 - 2f
	1, 1, 1, 1, 1, 1, 1, 1, // 30 - 37
	1, 1, 1, 1, 1, 1, 1, 1, // 38 - 3f
	1, 1, 1, 1, 1, 1, 1, 1, // 40 - 47
	1, 1, 1, 1, 1, 1, 1, 1, // 48 - 4f
	1, 1, 1, 1, 1, 1, 1, 1, // 50 - 57
	1, 1, 1, 1, 1, 1, 1, 1, // 58 - 5f
	1, 1, 1, 1, 1, 1, 1, 1, // 60 - 67
	1, 1, 1, 1, 1, 1, 1, 1, // 68 - 6f
	1, 1, 1, 1, 1, 1, 1, 1, // 70 - 77
	1, 1, 1, 1, 1, 1, 1, 1, // 78 - 7f
	2, 2, 2, 2, 3, 3, 3, 3, // 80 - 87
	4, 4, 4, 4, 4, 4, 4, 4, // 88 - 8f
	4, 4, 4, 4, 4, 4, 4, 4, // 90 - 97
	4, 4, 4, 4, 4, 4, 4, 4, // 98 - 9f
	5, 5, 5, 5, 5, 5, 5, 5, // a0 - a7
	5, 5, 5, 5, 5, 5, 5, 5, // a8 - af
	5, 5, 5, 5, 5, 5, 5, 5, // b0 - b7
	5, 5, 5, 5, 5, 5, 5, 5, // b8 - bf
	0, 0, 6, 6, 6, 6, 6, 6, // c0 - c7
	6, 6, 6, 6, 6, 6, 6, 6, // c8 - cf
	6, 6, 6, 6, 6, 6, 6, 6, // d0 - d7
	6, 6, 6, 6, 6, 6, 6, 6, // d8 - df
	7, 8, 8, 8, 8, 8, 8, 8, // e0 - e7
	8, 8, 8, 8, 8, 9, 8, 8, // e8 - ef
	10, 11, 11, 11, 11, 11, 11, 11, // f0 - f7
	12, 13, 13, 13, 14, 15, 0, 0, // f8 - ff
}

// UTF8States states
var UTF8States = []MachineState{
	MSError, MSStart, MSError, MSError, MSError, MSError, 12, 10, //00-07
	9, 11, 8, 7, 6, 5, 4, 3, //08-0f
	MSError, MSError, MSError, MSError, MSError, MSError, MSError, MSError, //10-17
	MSError, MSError, MSError, MSError, MSError, MSError, MSError, MSError, //18-1f
	MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSItsMe, //20-27
	MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSItsMe, MSItsMe, //28-2f
	MSError, MSError, 5, 5, 5, 5, MSError, MSError, //30-37
	MSError, MSError, MSError, MSError, MSError, MSError, MSError, MSError, //38-3f
	MSError, MSError, MSError, 5, 5, 5, MSError, MSError, //40-47
	MSError, MSError, MSError, MSError, MSError, MSError, MSError, MSError, //48-4f
	MSError, MSError, 7, 7, 7, 7, MSError, MSError, //50-57
	MSError, MSError, MSError, MSError, MSError, MSError, MSError, MSError, //58-5f
	MSError, MSError, MSError, MSError, 7, 7, MSError, MSError, //60-67
	MSError, MSError, MSError, MSError, MSError, MSError, MSError, MSError, //68-6f
	MSError, MSError, 9, 9, 9, 9, MSError, MSError, //70-77
	MSError, MSError, MSError, MSError, MSError, MSError, MSError, MSError, //78-7f
	MSError, MSError, MSError, MSError, MSError, 9, MSError, MSError, //80-87
	MSError, MSError, MSError, MSError, MSError, MSError, MSError, MSError, //88-8f
	MSError, MSError, 12, 12, 12, 12, MSError, MSError, //90-97
	MSError, MSError, MSError, MSError, MSError, MSError, MSError, MSError, //98-9f
	MSError, MSError, MSError, MSError, MSError, 12, MSError, MSError, //a0-a7
	MSError, MSError, MSError, MSError, MSError, MSError, MSError, MSError, //a8-af
	MSError, MSError, 12, 12, 12, MSError, MSError, MSError, //b0-b7
	MSError, MSError, MSError, MSError, MSError, MSError, MSError, MSError, //b8-bf
	MSError, MSError, MSStart, MSStart, MSStart, MSStart, MSError, MSError, //c0-c7
	MSError, MSError, MSError, MSError, MSError, MSError, MSError, MSError, //c8-cf
}

// UTF8CharLenTable charLenTable
var UTF8CharLenTable = []int{0, 1, 0, 0, 0, 0, 2, 3, 3, 3, 4, 4, 5, 5, 6, 6}

// UTF8SmModel model
var UTF8SmModel = Model{
	classTable:   UTF8Cls,
	classFactor:  16,
	stateTable:   UTF8States,
	charLenTable: UTF8CharLenTable,
	name:         "UTF-8"}
