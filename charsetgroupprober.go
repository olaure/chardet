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

// CharSetGroupProber is a group prober ???
type CharSetGroupProber struct {
	CharSetProber
	activeNum       int
	probers         []Prober
	bestGuessProber Prober
}

func newCharSetGroupProber(langFilter LanguageFilter) *CharSetGroupProber {
	return &CharSetGroupProber{
		*newCharSetProber(langFilter),
		0,
		[]Prober{},
		nil,
	}
}

func (c *CharSetGroupProber) reset() {
	c.CharSetProber.reset()
	c.activeNum = 0
	for _, prober := range c.probers {
		if prober != nil {
			prober.reset()
			prober.setActive(true)
			c.activeNum++
		}
	}
	c.bestGuessProber = nil
}

func (c *CharSetGroupProber) charsetName() string {
	if c.bestGuessProber == nil {
		c.getConfidence()
		if c.bestGuessProber == nil {
			return ""
		}
	}
	return c.bestGuessProber.charsetName()
}

func (c *CharSetGroupProber) language() string {
	if c.bestGuessProber == nil {
		c.getConfidence()
		if c.bestGuessProber == nil {
			return ""
		}
	}
	return c.bestGuessProber.language()
}

func (c *CharSetGroupProber) getConfidence() float64 {
	state := c.state
	if state == PSFound {
		return 0.99
	}
	if state == PSNotMe {
		return 0.01
	}
	var bestConf float64 = 0.0
	c.bestGuessProber = nil
	for _, prober := range c.probers {
		if prober == nil {
			continue
		}
		if !prober.getActive() {
			continue
		}
		conf := prober.getConfidence()
		if bestConf < conf {
			bestConf = conf
			c.bestGuessProber = prober
		}
	}
	if c.bestGuessProber == nil {
		return 0.0
	}
	return bestConf
}

func (c *CharSetGroupProber) feed(data []byte) ProbingState {
	for _, prober := range c.probers {
		if prober == nil || !prober.getActive() {
			continue
		}
		state := prober.feed(data)
		switch state {
		case PSDetecting:
			continue
		case PSFound:
			c.bestGuessProber = prober
			c.state = PSFound
			return c.state
		case PSNotMe:
			prober.setActive(false)
			c.activeNum--
			if c.activeNum <= 0 {
				c.state = PSNotMe
				return c.state
			}
		}
	}
	return c.state
}
