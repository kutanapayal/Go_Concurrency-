package main

import "testing"

func Test_updateMessage(t *testing.T) {

	msg = "gamma"

	wg.Add(2)
	go updateMessage("delta")
	go updateMessage("theta")
	wg.Wait()

	if msg != "theta" {
		t.Error("Error!!")
	}
}
