package main

import (
	"bytes"
	"fmt"
)

func countdown(out *bytes.Buffer) {
	fmt.Fprint(out, "3")
}
