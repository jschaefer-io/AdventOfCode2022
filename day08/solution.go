package day08

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jschaefer-io/AdventOfCode2022/orchestration"
)

func Solver() orchestration.Solver {
	return &day08{}
}

type minHeight struct {
	left   int
	right  int
	top    int
	bottom int
}

type day08 struct {
	heightMap    [][]int
	visiblityMap [][]minHeight
	height       int
	width        int
}

func (d *day08) Setup(data string) error {
	// Build HeightMap
	rows := strings.Split(data, "\n")
	d.height = len(rows)
	d.width = len(rows[0])
	d.heightMap = make([][]int, d.height)
	d.visiblityMap = make([][]minHeight, d.height)
	for y, row := range rows {
		d.heightMap[y] = make([]int, d.width)
		d.visiblityMap[y] = make([]minHeight, d.height)
		for x, col := range row {
			val := int(col - '0')
			d.heightMap[y][x] = val
			d.visiblityMap[y][x] = minHeight{}
		}
	}

	if d.height != d.width {
		return fmt.Errorf("invalid treeMap dimensions %dx%d", d.width, d.height)
	}
	return nil
}

func (d *day08) A() (string, error) {
	// Build Visibility-Data
	dimensions := d.width - 1
	for a := 1; a < d.width; a++ {
		for b := 1; b < d.width; b++ {
			d.visiblityMap[b][a].left = int(math.Max(float64(d.visiblityMap[b][a-1].left), float64(d.heightMap[b][a-1])))
			d.visiblityMap[b][dimensions-a].right = int(math.Max(float64(d.visiblityMap[b][dimensions-a+1].right), float64(d.heightMap[b][dimensions-a+1])))
			d.visiblityMap[a][b].top = int(math.Max(float64(d.visiblityMap[a-1][b].top), float64(d.heightMap[a-1][b])))
			d.visiblityMap[dimensions-a][b].bottom = int(math.Max(float64(d.visiblityMap[dimensions-a+1][b].bottom), float64(d.heightMap[dimensions-a+1][b])))
		}
	}

	// count visible trees
	// -- unoptimized solution
	count := 0
	var height int
	var visbility minHeight
	for x := 1; x < d.width-1; x++ {
		for y := 1; y < d.height-1; y++ {
			height = d.heightMap[y][x]
			visbility = d.visiblityMap[y][x]
			if height > visbility.left ||
				height > visbility.right ||
				height > visbility.top ||
				height > visbility.bottom {
				count++
			}
		}
	}
	// add outer trees
	count += d.width*4 - 4
	return strconv.Itoa(count), nil
}

func (d *day08) getTreeRange(x, y int) int {
	v := minHeight{}
	currentHeight := d.heightMap[y][x]
	for offset := 1; offset < d.width; offset++ {
		if x+offset < 0 || x+offset >= d.width {
			break
		}
		v.left++
		if d.heightMap[y][x+offset] >= currentHeight {
			break
		}
	}
	for offset := 1; offset < d.width; offset++ {
		if x-offset < 0 || x-offset >= d.width {
			break
		}
		v.right++
		if d.heightMap[y][x-offset] >= currentHeight {
			break
		}
	}
	for offset := 1; offset < d.width; offset++ {
		if y+offset < 0 || y+offset >= d.height {
			break
		}
		v.bottom++
		if d.heightMap[y+offset][x] >= currentHeight {
			break
		}
	}
	for offset := 1; offset < d.width; offset++ {
		if y-offset < 0 || y-offset >= d.width {
			break
		}
		v.top++
		if d.heightMap[y-offset][x] >= currentHeight {
			break
		}
	}
	return v.top * v.left * v.right * v.bottom
}

func (d *day08) B() (string, error) {
	var current, max int
	for x := 0; x < d.width; x++ {
		for y := 0; y < d.height; y++ {
			current = d.getTreeRange(x, y)
			if current > max {
				max = current
			}
		}
	}
	return strconv.Itoa(max), nil
}
