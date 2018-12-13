package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

type frequency struct {
	freq      int
	variation int
	modulo    *[]int
}

type modulos map[int]*[]int

func main() {
	day1, err := os.Open("input")
	if err != nil {
		log.Panic("No file")
	}

	var cycle []*frequency

	oldFreq := frequency{
		freq: 0,
	}

	file := bufio.NewReader(day1)
	for {
		variation, _, err := file.ReadLine()
		if err == io.EOF {
			break
		}

		freq := frequency{}
		freq.variation, err = strconv.Atoi(string(variation))
		freq.freq = freq.variation + oldFreq.freq

		oldFreq = freq
		cycle = append(cycle, &freq)

	}

	shift := oldFreq.freq

	fmt.Println("first star: " + strconv.Itoa(shift))
	mods := modulos{}

	for _, f := range cycle {
		modulo := f.freq % shift
		if modulo < 0 {
			modulo = modulo + shift
		}
		if _, ok := mods[modulo]; !ok {
			mods[modulo] = new([]int)
		}
		*mods[modulo] = append(*mods[modulo], f.freq)
		f.modulo = mods[modulo]
	}

	for _, v := range mods {
		sort.Slice(*v, func(i, j int) bool { return i > j })
	}
	winner := struct {
		delta int
		freq  *frequency
		win   int
	}{
		delta: 1000000000000,
	}

	for _, f := range cycle {
		m := *f.modulo
		for i := 0; i < len(m)-1; i++ {
			delta := math.Abs(float64(m[i+1] - m[i]))
			if int(delta) < winner.delta && f.freq == m[i] {
				winner.delta = int(delta) / shift
				winner.win = int(math.Max(float64(m[i]), float64(m[i+1])))
			}

		}
	}

	fmt.Println("second star: " + strconv.Itoa(winner.win))
}
