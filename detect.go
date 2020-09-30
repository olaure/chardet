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

import (
	"strings"
)

// Detect detects the encoding of a byte slice
func Detect(data []byte) *Result {
	detector := newUniversalDetector(LFAll)
	detector.Feed(data)
	return detector.Close()
}

// DetectAll detects all the possible encoding of a byte slice
// Outputs possible encodings only if their confidence is above a threshold
func DetectAll(data []byte) []*Result {
	detector := newUniversalDetector(LFAll)
	detector.Feed(data)
	detector.Close()

	if detector.inputState == UDSHighByte {
		results := []*Result{}
		for _, prober := range detector.charsetProbers {
			if prober.getConfidence() <= UDMinimumThreshold {
				continue
			}
			charsetName := prober.charsetName()
			lowerCharsetName := strings.ToLower(charsetName)
			// Use windows encoding name instead of ISO-8859
			// if we saw any extra windows-specific bytes
			if strings.HasPrefix(lowerCharsetName, "iso-8859") && detector.hasWinBytes {
				charsetNameTmp, ok := UDIsoWinMap[lowerCharsetName]
				if ok {
					charsetName = charsetNameTmp
				}
			}
			results = append(results, &Result{
				Encoding:   charsetName,
				Confidence: prober.getConfidence(),
			})
		}
	}
	return []*Result{detector.Result}
}
