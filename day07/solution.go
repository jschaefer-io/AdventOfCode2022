package day07

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jschaefer-io/AdventOfCode2022/orchestration"
)

func Solver() orchestration.Solver {
	return &day07{}
}

type day07 struct {
	fileSystem *fileSystem
}

func (d *day07) Setup(data string) error {
	d.fileSystem = newFileSystem()
	currentDirectory := d.fileSystem.root

	lines := strings.Split(data, "\n")
	lineCount := len(lines)
	currentLine := 0
	for currentLine < lineCount {
		if lines[currentLine][0] != '$' {
			return fmt.Errorf("unexpected command response on line %d", currentLine)
		}
		command := strings.Split(lines[currentLine], " ")
		switch command[1] {
		case "cd":
			if command[2] == ".." {
				currentDirectory = currentDirectory.parent
			} else if command[2][0] == '/' {
				currentDirectory = d.fileSystem.getDirectory(
					fmt.Sprintf("%s", command[2]),
				)
			} else {
				currentDirectory = d.fileSystem.getDirectory(
					fmt.Sprintf("%s/%s", currentDirectory.String(), command[2]),
				)
			}
		case "ls":
			i := 1
			for currentLine+i < lineCount && lines[currentLine+i][0] != '$' {
				lsInfo := strings.Split(lines[currentLine+i], " ")
				if lsInfo[0] == "dir" {
					d.fileSystem.getDirectory(
						fmt.Sprintf("%s/%s", currentDirectory.String(), lsInfo[1]),
					)
				} else {
					size, err := strconv.Atoi(lsInfo[0])
					if err != nil {
						return fmt.Errorf("unable to parse file size on line %d: %w", currentLine+i, err)
					}
					currentDirectory.addFile(size, lsInfo[1])
				}
				i++
			}
			currentLine += i - 1
		}
		currentLine++
	}
	return nil
}

func (d *day07) A() (string, error) {
	sum := 0
	d.fileSystem.iterateWith(func(cd *directory) {
		size := cd.getSize(d.fileSystem)
		if size <= 100000 {
			sum += size
		}
	})
	return strconv.Itoa(sum), nil
}

func (d *day07) B() (string, error) {
	targetSpace := 30000000
	freeSpace := 70000000 - d.fileSystem.root.getSize(d.fileSystem)
	if targetSpace <= freeSpace {
		return "", errors.New("no deletions required")
	}
	minDeletion := targetSpace - freeSpace

	minSize := -1
	var minDirectory *directory
	d.fileSystem.iterateWith(func(cd *directory) {
		size := cd.getSize(d.fileSystem)
		if size < minDeletion {
			return
		}

		if size < minSize || minDirectory == nil {
			minSize = size
			minDirectory = cd
		}
	})
	if minDirectory == nil {
		return "", errors.New("no available directory for deletion found")
	}
	return strconv.Itoa(minSize), nil
}
