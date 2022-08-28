package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	FileNames := [5]string{"quadA", "quadB", "quadC", "quadD", "quadE"}
	var x int
	var y int
	val, _ := io.ReadAll(os.Stdin) //read from standart input
	for _, i := range val {
		if i == '\n' {
			y++
		} else {
			x++
		}
	}
	if x == 0 || y == 0 {
		fmt.Println("Not a Raid function")
		return
	}
	x /= y // size of one line
	twins := []string{}

	for _, i := range FileNames {
		res, _ := exec.Command("./"+i, strconv.Itoa(x), strconv.Itoa(y)).Output()
		// we look at output of each quad of particular x and y
		if string(val) == string(res) {
			twins = append(twins, i) //if output of cmd and standart input the same, twins + name of the quad
		}
	}

	if len(twins) == 0 {
		fmt.Println("Not a Raid function")
		return
	}

	for i, j := range twins {
		if i == 0 {
			fmt.Printf("[%s] [%d] [%d]", j, x, y) //%s the uninterpreted bytes of the string or slice
		} else {
			fmt.Printf(" || ") // next quad match
			fmt.Printf("[%s] [%d] [%d]", j, x, y) //%d integer base 10
		}
	}
	fmt.Println()
}
