package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution5(line string) string {
	strArray := strings.Split(line, ",")
	length := (len(strArray)) / 2
	var index int
	var value int
	for i, v := range strArray {
		if i == 0 {
			value, _ = strconv.Atoi(v)
			continue
		}
		value2, _ := strconv.Atoi(v)
		if value > value2 {
			index = i
			break
		}
		value, _ = strconv.Atoi(v)
	}
	if index <= length {
		return strArray[index+length]
	} else {
		return strArray[index-length-1]
	}
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution5(string(line)))
	}
}
