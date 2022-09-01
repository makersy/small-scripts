package main

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := make(map[string]string)
	m["j"] = "l"
	m["a"] = "l"
	m["c"] = "l"
	fmt.Printf("%s\n", m)

}
