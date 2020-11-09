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
	"bytes"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"strings"
	"unicode/utf8"
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

// DetectShortestUTF8 detects the best encoding for a byte slice
// Given that the right encoding may not have the highest confidence.
// Heuristic is made with the fewer count of exotic UTF8 chars
// (neither letters, numbers, punctuation or spaces)
func DetectShortestUTF8(data []byte) *Result {
	detector := newUniversalDetector(LFAll)
	detector.Feed(data)
	highConfidenceResult := detector.Close()
	lowestConfidence := 2.0*highConfidenceResult.Confidence - 1.0 // conf - (1 - conf)
	if highConfidenceResult.Confidence >= 0.99 {
		return highConfidenceResult
	}

	currentBestResult := highConfidenceResult
	currentBestScore := scoreFromResult(data, currentBestResult, -1)

	if detector.inputState == UDSHighByte {
		for _, prober := range detector.charsetProbers {
			conf := prober.getConfidence()
			if conf <= lowestConfidence || prober.getConfidence() <= UDMinimumThreshold {
				break
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
			newResult := &Result{
				Encoding:   charsetName,
				Confidence: prober.getConfidence(),
			}
			newScore := scoreFromResult(data, newResult, currentBestScore)
			if newScore < currentBestScore {
				currentBestScore = newScore
				currentBestResult = newResult
			}
		}
	}
	if highConfidenceResult.Encoding != currentBestResult.Encoding {
		log.Debugf("Detected differing encoding from highest confidence : %v vs %v", highConfidenceResult, currentBestResult)
	}
	return currentBestResult
}

func intoCharset(b []byte, encoding string) ([]byte, error) {
	reader, err := charset.NewReaderLabel(encoding, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	decoded, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

func scoreFromResult(data []byte, result *Result, currentBestScore int) int {
	bestDecode, err := intoCharset(data, result.Encoding)
	if err != nil {
		return utf8.RuneCountInString(string(data))
	}
	return utf8.RuneCountInString(string(bestDecode))
}
