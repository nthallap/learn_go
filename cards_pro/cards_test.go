package main

import (
	"os"
	"testing"
)

func TestDesc(t *testing.T) {

	newdc := newDesk()

	if len(newdc) != 15 {
		t.Errorf("invalid length of variables %v", len(newdc))
	}

	if newdc[0] != "one heart" {
		t.Errorf("invalid card  %v", newdc[0])
	}
	if newdc[len(newdc)-1] != "queen diamond" {
		t.Errorf("invalid card  %v", newdc[len(newdc)-1])
	}
}

func TestSaveToDriveLoadFromDrive(t *testing.T) {
	os.Remove("_mydesctest")
	newdc := newDesk()
	newdc.saveToFile("_mydesctest")
	loadeddc := deckFromfile("_mydesctest")

	if len(newdc) == len(loadeddc) {
		t.Errorf("some thing went wrong!")
	}

}
