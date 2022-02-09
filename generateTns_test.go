package main

import (
	"testing"
)

func TestGenerateTnsDeepSpeech(t *testing.T) {
	test_file := "en.wav"
	r, err := GenerateTnsDeepSpeech(test_file)
	if err != nil {
		t.Error(err)
	}
	if r != "i knocked at the door on the ancient site of the building\n" {
		t.Error("result should be ->i knocked at the door on the ancient site of the building, not ->", r)
	}
}
