package main

import (
	"bytes"
	"strconv"
	"testing"
)

//Test case for positive scenario
func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(`1
1
2
3
5
8
13
`)
	actual := strconv.Quote(buf.String())
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

//Test case for negative scenario
func TestMainOutputNegative(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(-7)

	expected := strconv.Quote(`1
-1
2
-3
5
-8
13
`)
	actual := strconv.Quote(buf.String())
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}
