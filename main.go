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

	maxstat      line
	maxstatcache line

	lines []line

	returnlines []int
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
}

func trains() {
	var i, u int
	i, u = maxstatind, maxstatind

	for ; true; u++ {
		if lines[i].end < lines[u].start {
			break
		}
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
	sort(lines, maxstatind)
	returnlines = append(returnlines, lines[0].index)

	for {
		trains()
		returnlines = append(returnlines, maxstat.index)
		if maxstat.end == stationsc {
			fmt.Println(len(returnlines) - 1)
			fmt.Println(returnlines)
			break
		} else if maxstat == maxstatcache {
			fmt.Println(-1)
			break
		}
		maxstatcache = maxstat
	}
}
