package main

import (
	"hash/fnv"
	"io/ioutil"
	"testing"
)

func TestHighlightHomophones(t *testing.T) {
	t.Log("Testing that ansi highlights are added to known homophones...")
	expected_hash := uint32(1648753471)
	content, err := ioutil.ReadFile("test.txt")
	if err != nil {
		t.Errorf("Unable to open test file.")
	}

	text := HighlightHomphones(string(content))
	text_hash := hash(text)

	if text_hash != expected_hash {
		t.Errorf("Expected hash %d was %d instead.", expected_hash, text_hash)
	}
}

func TestHighlightApostrophes(t *testing.T) {
	t.Log("Testing that ansi highlights are added to apostrophes...")
	expected_hash := uint32(3485546507)
	content, err := ioutil.ReadFile("apostrophe_test.txt")
	if err != nil {
		t.Errorf("Unable to open test file.")
	}

	text := HighlightApostrophes(string(content))
	text_hash := hash(text)

	if text_hash != expected_hash {
		t.Errorf("Expected hash %d was %d instead.", expected_hash, text_hash)
	}
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
