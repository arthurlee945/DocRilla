package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Init push", os.Getenv("TEST_ENV"))
}
