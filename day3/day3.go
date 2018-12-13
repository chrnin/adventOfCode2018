package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type rectangle struct {
	id     int
	x      int
	y      int
	width  int
	height int
	plots  []*plot
}

type plot struct {
	x        int
	y        int
	numrects int
	rects    []*rectangle
}

func plots(rect *rectangle, mapping map[[2]int]plot) {

}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Panic()
	}

	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile(`(\d*) @ (\d*),(\d*): (\d*)x(\d*)`)
	f := make(map[int]*rectangle)
	mapping := make(map[[2]int]*plot)

	for scanner.Scan() {
		v := r.FindStringSubmatch(scanner.Text())
		rect := rectangle{}
		rect.id, err = strconv.Atoi(v[1])
		rect.x, err = strconv.Atoi(v[2])
		rect.y, err = strconv.Atoi(v[3])
		rect.width, err = strconv.Atoi(v[4])
		rect.height, err = strconv.Atoi(v[5])

		for i := rect.x; i < rect.x+rect.width; i++ {
			for j := rect.y; j < rect.y+rect.height; j++ {
				p, ok := mapping[[2]int{i, j}]
				if !ok {
					p = &plot{
						x:        i,
						y:        j,
						numrects: 0,
						rects:    []*rectangle{},
					}
				}
				rect.plots = append(rect.plots, p)
				p.rects = append(p.rects, &rect)
				p.numrects++
				mapping[[2]int{i, j}] = p
			}
		}
		f[rect.id] = &rect
	}
	star1 := 0
	for _, p := range mapping {
		if p.numrects > 1 {
			star1++
		}
	}
	star2 := 0
	for _, v := range f {
		winner := true
		for _, p := range v.plots {
			if p.numrects > 1 {
				winner = false
				break
			}
		}
		if winner == true {
			star2 = v.id
			break
		}
	}
	fmt.Println(star1)
	fmt.Println(star2)
}
