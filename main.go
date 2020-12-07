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
  index := strings.LastIndex(x, y)
  return index != -1 && index == len(x) - len(y)
}
