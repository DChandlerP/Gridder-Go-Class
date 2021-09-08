package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	if fileExists("_decktesting") {
		os.Remove("_decktesting")
	}	
	d := newDeck()
	d.saveToFile("_decktesting")
	d2 := newDeckFromFile("_decktesting")
	if len(d2) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d2))
	}
}

func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}