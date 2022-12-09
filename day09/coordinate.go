package day09

import (
	"math"
)

type coordinate struct {
	x int
	y int
}

func (c *coordinate) moveTowards(target coordinate) {
	distanceX := int(math.Abs(float64(target.x-c.x))) - 1
	distanceY := int(math.Abs(float64(target.y-c.y))) - 1

	offsetX := distanceX
	offsetY := distanceY

	if target.y-c.y < 0 {
		offsetY *= -1
	}
	if target.x-c.x < 0 {
		offsetX *= -1
	}

	if distanceX <= 0 && distanceY <= 0 {
		return
	}

	if distanceX > 0 {
		c.x += offsetX
		if c.y != target.y {
			if c.y-target.y > 0 {
				c.y--
			} else {
				c.y++
			}
		}
	} else {
		c.y += offsetY
		if c.x != target.x {
			if c.x-target.x > 0 {
				c.x--
			} else {
				c.x++
			}
		}
	}
}
