package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	var (
		primString string
		primInt    int
		err        error
	)

	primString, _ = reader.ReadString('\n')
	primString = strings.TrimSuffix(primString, "\n")

	primInt, err = strconv.Atoi(primString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Please insert an integer!")
		os.Exit(1)
	}

	printPrime(primInt)

	os.Exit(0)
}

func printPrime(n int) {
	if n == 1 {
		fmt.Println("1")
		os.Exit(0)
	}
	a := make([]bool, n, n)

	for i := range a {
		a[i] = true
	}
	var k = 0
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if a[i] == true {
			for j := i * i; j < n; j = i*i + k*i {
				a[j] = false
				k++
			}
		}
	}

	for i := range a {
		if a[i] == true {
			fmt.Println(strconv.Itoa(i))
		}
	}

}
