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

type charDistributionAnalyzer interface {
	reset()
	feed([]byte, int)
	getConfidence() float64
	gotEnoughData() bool
	getOrder([]byte) int
}

// CDAEnoughDataThreshold threshold
var CDAEnoughDataThreshold = 1024

// CDASureYes threshold
var CDASureYes = 0.99

// CDASureNo threshold
var CDASureNo = 0.01

// CDAMinimumDataThreshold threshold
var CDAMinimumDataThreshold = 3

type charDistributionAnalysis struct {
	charToFreqOrder          []int
	tableSize                int
	typicalDistributionRatio float64
	done                     bool
	totalChars               int
	freqChars                int
	getOrder                 func([]byte) int
}

func (c *charDistributionAnalysis) reset() {
	// If this flag is set to true, detection is done an conclusion has been made
	c.done = false
	// Total characters encountered
	c.totalChars = 0
	// The number of characters whose frequency order is less than 512
	c.freqChars = 0
}

func defaultGetOrder(data []byte) int {
	return -1
}

// Return confidence based on existing data
func (c *charDistributionAnalysis) getConfidence() float64 {
	// If we didn't receive any character in our consideration range, return NO
	if c.totalChars <= 0 || c.freqChars <= CDAMinimumDataThreshold {
		return CDASureNo
	}

	if c.totalChars != c.freqChars {
		r := float64(c.freqChars) / (float64(c.totalChars-c.freqChars) * c.typicalDistributionRatio)
		if r < CDASureYes {
			return r
		}
	}
	return CDASureYes
}

func (c *charDistributionAnalysis) gotEnoughData() bool {
	return c.totalChars > CDAEnoughDataThreshold
}

func (c *charDistributionAnalysis) feed(data []byte, chrLen int) {
	var order int
	if chrLen == 2 {
		// We only care about 2-bytes characters in our distribution analysis
		order = c.getOrder(data)
	} else {
		order = -1
	}
	if order >= 0 {
		c.totalChars++
		// Order is valid
		if order < c.tableSize {
			if 512 > c.charToFreqOrder[order] {
				c.freqChars++
			}
		}
	}
}

// EUCTW

// EUCTWDistributionAnalysis EUCTW
type EUCTWDistributionAnalysis struct {
	charDistributionAnalysis
}

func getOrderEUCTW(data []byte) int {
	firstChr := data[0]
	if firstChr >= 0xc4 {
		return 94*int(firstChr-0xc4) + int(data[1]) - 0xa1
	}
	return -1
}

func (s *EUCTWDistributionAnalysis) getOrder(data []byte) int {
	return getOrderEUCTW(data)
}

func newEUCTWDistributionAnalysis() charDistributionAnalyzer {
	cda := EUCTWDistributionAnalysis{charDistributionAnalysis{
		charToFreqOrder:          EUCTWCharToFreqOrder,
		tableSize:                EUCTWTableSize,
		typicalDistributionRatio: EUCTWTypicalDistributionRatio,
		getOrder:                 getOrderEUCTW,
	}}
	cda.reset()
	return &cda
}

//EUCKR

// EUCKRDistributionAnalysis EUCKR
type EUCKRDistributionAnalysis struct {
	charDistributionAnalysis
}

// For euc-KR encoding we are interested
//  first byte range  : 0xb0 -- 0xfe
//  second bute range : 0xa1 -- 0xfe
// No validation needed here. State machine has done that.
func getOrderEUCKR(data []byte) int {
	firstChr := data[0]
	if firstChr >= 0xb0 {
		return 94*int(firstChr-0xb0) + int(data[1]) - 0xa1
	}
	return -1
}

func (s *EUCKRDistributionAnalysis) getOrder(data []byte) int {
	return getOrderEUCKR(data)
}

func newEUCKRDistributionAnalysis() charDistributionAnalyzer {
	cda := EUCKRDistributionAnalysis{charDistributionAnalysis{
		charToFreqOrder:          EUCKRCharToFreqOrder,
		tableSize:                EUCKRTableSize,
		typicalDistributionRatio: EUCKRTypicalDistributionRatio,
		getOrder:                 getOrderEUCKR,
	}}
	cda.reset()
	return &cda
}

//GB2312

// GB2312DistributionAnalysis GB2312
type GB2312DistributionAnalysis struct {
	charDistributionAnalysis
}

func getOrderGB2312(data []byte) int {
	firstChr, secondChr := data[0], data[1]
	if firstChr >= 0xb0 && secondChr >= 0xa1 {
		return 94*int(firstChr-0xb0) + int(secondChr-0xa1)
	}
	return -1
}

func (s *GB2312DistributionAnalysis) getOrder(data []byte) int {
	return getOrderGB2312(data)
}

func newGB2312DistributionAnalysis() charDistributionAnalyzer {
	cda := GB2312DistributionAnalysis{charDistributionAnalysis{
		charToFreqOrder:          GB2312CharToFreqOrder,
		tableSize:                GB2312TableSize,
		typicalDistributionRatio: GB2312TypicalDistributionRatio,
		getOrder:                 getOrderGB2312,
	}}
	cda.reset()
	return &cda
}

//BIG5

// BIG5DistributionAnalysis BIG5
type BIG5DistributionAnalysis struct {
	charDistributionAnalysis
}

func getOrderBIG5(data []byte) int {
	firstChr, secondChr := data[0], data[1]
	if firstChr >= 0xa4 {
		if secondChr >= 0xa1 {
			return 157*(int(firstChr)-0xa4) + int(secondChr) - 0xa1 + 63
		}
		return 157*int(firstChr-0xa4) + int(secondChr-0x40)
	}
	return -1
}

func (s *BIG5DistributionAnalysis) getOrder(data []byte) int {
	return getOrderBIG5(data)
}

func newBIG5DistributionAnalysis() charDistributionAnalyzer {
	cda := BIG5DistributionAnalysis{charDistributionAnalysis{
		charToFreqOrder:          BIG5CharToFreqOrder,
		tableSize:                BIG5TableSize,
		typicalDistributionRatio: BIG5TypicalDistributionRatio,
		getOrder:                 getOrderBIG5,
	}}
	cda.reset()
	return &cda
}

//SJIS

// SJISDistributionAnalysis SJIS
type SJISDistributionAnalysis struct {
	charDistributionAnalysis
}

func getOrderSJIS(data []byte) int {
	var order int
	firstChr, secondChr := data[0], data[1]
	if firstChr >= 0x81 && firstChr <= 0x9f {
		order = 188 * (int(firstChr) - 0x81)
	} else if firstChr >= 0xe0 && firstChr <= 0xef {
		order = 188 * (int(firstChr) - 0xe0 + 31)
	} else {
		return -1
	}
	order = order + int(secondChr-0x40)
	if secondChr > 0x7f {
		order = -1
	}
	return order
}

func (s *SJISDistributionAnalysis) getOrder(data []byte) int {
	return getOrderSJIS(data)
}

func newSJISDistributionAnalysis() charDistributionAnalyzer {
	cda := SJISDistributionAnalysis{charDistributionAnalysis{
		charToFreqOrder:          JISCharToFreqOrder,
		tableSize:                JISTableSize,
		typicalDistributionRatio: JISTypicalDistributionRatio,
		getOrder:                 getOrderSJIS,
	}}
	cda.reset()
	return &cda
}

//EUCJP

// EUCJPDistributionAnalysis EUCJP
type EUCJPDistributionAnalysis struct {
	charDistributionAnalysis
}

func getOrderEUCJP(data []byte) int {
	firstChr := data[0]
	if firstChr > 0xa0 {
		return 94*int(firstChr-0xa1) + int(data[1]) - 0xa1
	}
	return -1
}

func (s *EUCJPDistributionAnalysis) getOrder(data []byte) int {
	return getOrderEUCJP(data)
}

func newEUCJPDistributionAnalysis() charDistributionAnalyzer {
	cda := EUCJPDistributionAnalysis{charDistributionAnalysis{
		charToFreqOrder:          JISCharToFreqOrder,
		tableSize:                JISTableSize,
		typicalDistributionRatio: JISTypicalDistributionRatio,
		getOrder:                 getOrderEUCJP,
	}}
	cda.reset()
	return &cda
}
