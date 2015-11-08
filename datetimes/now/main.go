package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("%d seconds since the Epoc\n", time.Now().Unix())
	fmt.Printf("%d nanoseconds since the Epoc\n", time.Now().UnixNano())
}
