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

// UniversalDetectorState is a possible state of the UniversalDetector
type UniversalDetectorState int

const (
	// UDSPureASCII state
	UDSPureASCII UniversalDetectorState = 0
	// UDSEscASCII state
	UDSEscASCII UniversalDetectorState = 1
	// UDSHighByte state
	UDSHighByte UniversalDetectorState = 2
)

// LanguageFilter represents the different language filters appliable to a UniversalDetector
type LanguageFilter int

const (
	// LFNone is no filter
	LFNone LanguageFilter = 0
	// LFChineseSimplified language filter
	LFChineseSimplified LanguageFilter = 1
	// LFChineseTraditional language filter
	LFChineseTraditional LanguageFilter = 2
	// LFJapanese language filterz
	LFJapanese LanguageFilter = 4
	// LFKorean language filter
	LFKorean LanguageFilter = 8
	// LFNonCJK language filter
	LFNonCJK LanguageFilter = 0x10
	// LFAll language filter
	LFAll LanguageFilter = 0x1f
	// LFChinese language filter
	LFChinese LanguageFilter = LFChineseSimplified | LFChineseTraditional
	// LFCJK language filter
	LFCJK LanguageFilter = LFChinese | LFJapanese | LFKorean
)

// ProbingState represents the different states a prober can be in
type ProbingState int

const (
	// PSDetecting state
	PSDetecting ProbingState = 0
	// PSFound state
	PSFound ProbingState = 1
	// PSNotMe state
	PSNotMe ProbingState = 2
)

// MachineState represents the different states a state machine can be in
type MachineState int

const (
	// MSStart state
	MSStart MachineState = 0
	// MSError state
	MSError MachineState = 1
	// MSItsMe state
	MSItsMe MachineState = 2
)

// SequenceLikelihood represents the likelihood of a character following the previous one
type SequenceLikelihood int

const (
	// SLNegative state
	SLNegative SequenceLikelihood = 0
	// SLUnlikely state
	SLUnlikely SequenceLikelihood = 1
	// SLLikely state
	SLLikely SequenceLikelihood = 2
	// SLPositive state
	SLPositive SequenceLikelihood = 3
)

// SLNumCategories the number of categories in SequenceLikelihood
const SLNumCategories = 4

// CharacterCategory represents the different categories language models
// for SBCP (SingleByteCharsetProber) put chars into
// Anything less than CCControl is considered a letter
type CharacterCategory int

const (
	// CCControl category
	CCControl CharacterCategory = 251
	// CCDigit category
	CCDigit CharacterCategory = 252
	// CCSymbol category
	CCSymbol CharacterCategory = 253
	// CCLineBreak category
	CCLineBreak CharacterCategory = 254
	// CCUndefined category
	CCUndefined CharacterCategory = 255
)
