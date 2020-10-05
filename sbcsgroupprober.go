package chardet

/*
######################## BEGIN LICENSE BLOCK ########################
# The Original Code is Mozilla Universal charset detector code.
#
# The Initial Developer of the Original Code is
# Netscape Communications Corporation.
# Portions created by the Initial Developer are Copyright (C) 2001
# the Initial Developer. All Rights Reserved.
#
# Contributor(s):
#   Mark Pilgrim - port to Python
#   Shy Shalom - original C code
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

// SBCSGroupProber Group of Single Byte prober
type SBCSGroupProber struct {
	CharSetGroupProber
}

func newSBCSGroupProber() *SBCSGroupProber {
	hebrewProber := newHebrewProber()
	logicalHebrewProber := newSingleByteCharSetProber(&WINDOWS1255HebrewModel, false, hebrewProber)
	// TODO : See if ISO-8859-8 Hebrew model works better here
	//        since it's actually the visual one
	visualHebrewProber := newSingleByteCharSetProber(&WINDOWS1255HebrewModel, true, hebrewProber)
	hebrewProber.setModelProbers(logicalHebrewProber, visualHebrewProber)
	/* TODO:
	 *  ORDER MATTERS HERE. I changed the order vs what was in master
	 *  and several tests failed that did not before. Some thought
	 *  should be put into the ordering, and we should consider making
	 *  order not matter here, because that is very counter-intuitive.
	 *  TODO EDIT (go migration) : the order doesn't seem to matter anymore after migrating to go.
	 *      It seems the only thing changing is when equivalent model
	 *      (i.e. ISO8859_7 and WINDOWS1253) exchange order and they may have the same score.
	 */
	probers := []Prober{
		newSingleByteCharSetProber(&WINDOWS1251RussianModel, false, nil),
		newSingleByteCharSetProber(&KOI8RRussianModel, false, nil),
		newSingleByteCharSetProber(&ISO8859_5RussianModel, false, nil),
		newSingleByteCharSetProber(&MACCYRILLICRussianModel, false, nil),
		newSingleByteCharSetProber(&IBM866RussianModel, false, nil),
		newSingleByteCharSetProber(&IBM855RussianModel, false, nil),
		newSingleByteCharSetProber(&ISO8859_7GreekModel, false, nil),
		newSingleByteCharSetProber(&WINDOWS1253GreekModel, false, nil),
		newSingleByteCharSetProber(&ISO8859_5BulgarianModel, false, nil),
		// TODO : Restore Hungarian encodings (iso-8859-2 and windows-1250)
		//        After the model is retrained
		// newSingleByteCharSetProber(&ISO8859_2HungarianModel, false, nil),
		// newSingleByteCharSetProber(&WINDOWS1250HungarianModel, false, nil),
		newSingleByteCharSetProber(&TIS620ThaiModel, false, nil),
		newSingleByteCharSetProber(&ISO8859_9TurkishModel, false, nil),
		hebrewProber,
		logicalHebrewProber,
		visualHebrewProber,
	}
	csgp := newCharSetGroupProber(LFNone)
	csgp.probers = probers
	mgp := SBCSGroupProber{*csgp}
	mgp.reset()
	return &mgp
}
