package main
import (
"fmt"
"time"
)

func main() {
	go fmt.Printf("Printed in the background\n")
	i := 1
	go fmt.Printf("Currently, i is %d\n", i)
	go func() {
		fmt.Printf("i: %d\n", i)
	}()
	i++
	time.Sleep(100000000)
}
