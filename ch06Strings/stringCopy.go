package main
import "fmt"

func main() {
	str := "A string"
	bytes := make([]byte, len(str))
	copy(bytes, str)
	strCopy := string(bytes)
	fmt.Printf("'%s'", strCopy)
}
