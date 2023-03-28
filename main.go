package main

import "fmt"

type line struct {
	start int
	end   int
	index int
}

var (
	stationsc  int
	linesc     int
	start      int
	end        int
	maxstatind int

	maxstat line

	lines []line
)

func input() {
	fmt.Scanln(&linesc, &stationsc)
	for i := 0; i < linesc; i++ {
		fmt.Scanln(&start, &end)
		lines = append(lines, newLine(start, end, i))
	}
}

func newLine(start, end, index int) line {
	return line{
		start: start,
		end:   end,
		index: index,
	}
}

func sort(line []line, i int) {
	for ; i < len(line)-1; i++ {
		if lines[i].start == line[i+1].start {
			if lines[i].end < lines[i+1].end {
				lines[i], lines[i+1] = lines[i+1], lines[i]
			}
		} else {
			break
		}

	}

	if lines[len(lines)-1].start == lines[len(lines)].start {
		if lines[len(lines)-1].end < lines[len(lines)].end {
			lines[len(lines)-1], lines[len(lines)] = lines[len(lines)], lines[len(lines)-1]
		}
	}
}

func trains() {

	var i, u int

	i, u = maxstatind, maxstatind

	for ; true; u++ {
		if lines[i].end < lines[u].start {
			break
		}
		fmt.Println(maxstat)
		if lines[i].start == lines[u].start {

		} else {
			if lines[u].end > maxstat.end {
				maxstat = lines[u]
				maxstatind = u
			}
		}
	}
}

func main() {
	input()

	//for true {

	for {

		if maxstat.end == stationsc {
			break
		}
		sort(lines, maxstatind)
		trains()
	}

	fmt.Println(lines)
	fmt.Println(maxstat)

	//}
}
