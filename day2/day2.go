package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type file []string

func twothree(l string) (int, int) {

	chars := make(map[rune]int)
	for _, c := range l {
		chars[c]++
	}

	two, three := 0, 0
	for _, v := range chars {
		if v == 2 {
			two = 1
		} else if v == 3 {
			three = 1
		}
	}
	return two, three
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Panic("No input file")
	}

	scanner := bufio.NewScanner(f)
	m := file{}

	i := 0
	twos, threes := 0, 0
	for scanner.Scan() {
		l := scanner.Text()
		m = append(m, l)
		two, three := twothree(m[i])
		i++
		twos = twos + two
		threes = threes + three
	}

	fmt.Println("Star 1: " + strconv.Itoa(twos*threes))
	mapDifference := make(map[[2]int]bool)

	for i := range m {
		for j := range m {
			if i < j {
				mapDifference[[2]int{i, j}] = false
			}
		}
	}

	for i := 0; i < len(m[0])-2; i++ {
		for j := range mapDifference {
			if m[j[0]][i] != m[j[1]][i] {
				if mapDifference[j] {
					delete(mapDifference, j)
				} else {
					mapDifference[j] = true
				}
			}
		}
	}
	for i, v := range mapDifference {
		if v {
			string1 := m[i[0]]
			string2 := m[i[1]]
			var res string
			for i := range string1 {
				if string1[i] == string2[i] {
					res = res + string(string1[i])
				}
			}
			fmt.Println("Star 2: " + res)
		}
	}

}
