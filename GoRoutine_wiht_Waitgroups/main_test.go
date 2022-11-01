package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {

	stdout := os.Stdout

	r, w, _ := os.Pipe()
	t.Log(w)
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go PrintSomething("epsilon", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdout

	if !strings.Contains(output, "epsiln") {
		t.Errorf("Not Matched!! Errorr!!")
	}

}
