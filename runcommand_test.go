package main

import "testing"

func TestIsRunCommandValid(t *testing.T) {
	invalidArgs := [3]string{"dummyScriptName", "run", "abc"}
	validFileName := "main.go"

	result := IsRunCommandValid(invalidArgs[:], validFileName)

	if result {
		t.Errorf("IsRunCommandValid failed, expected %v, got %v", "false", result)
	}

	validArgs := [3]string{"dummyScriptName", "run", "--file"}
	invalidFileName := "abc"
	result = IsRunCommandValid(validArgs[:], invalidFileName)
	if result {
		t.Errorf("IsRunCommandValid failed, expected %v, got %v", "false", result)
	}

	result = IsRunCommandValid(validArgs[:], validFileName)

	if !result {
		t.Errorf("IsRunCommandValid failed, expected %v, got %v", "true", result)
	}
}