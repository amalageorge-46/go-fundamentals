package main

import (
	"os"
	"time"

	"example.com/hello/math/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
