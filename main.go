package main
import (
	"fmt"
	"strings"
)
func main() {
	fmt.Println("Hello, playground")
	fmt.Println(endsWith("go gopher", "go"))
}
func endsWith(x, y string) bool {
  return strings.LastIndex(x, y) == len(x) - len(y)
}
