// Simple example for recieving a standard input as a line of strings
package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')

	fmt.Println("Hello, World.");
	fmt.Println(line)

}