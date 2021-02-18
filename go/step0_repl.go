package main
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Malgo v0.0.1")
	fmt.Println("Last updated 2-17-2021")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("glo > ")
		s, _ := reader.ReadString('\n')
		fmt.Print(rep(s))
	}
}

func read(s string) string {
	return s
}

func eval(s string) string {
	return s
}

func print(s string) string {
	return s
}

func rep(s string) string {
	return print(eval(read(s)))
}