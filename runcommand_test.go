package main

import "testing"

func TestRunCommandValidation_InvalidArgs(t *testing.T) {
    invalidArgs := [3]string{"dummyScriptName", "run", "abc"}
    validFileName := "main.go"

    result := IsRunCommandValid(invalidArgs[:], validFileName)

    if result {
        t.Errorf("TestRunCommandValidation_InvalidArgs failed, expected %v, got %v", "false", result)
    }
}

func TestRunCommandValidation_InvalidFileName(t *testing.T) {
    validArgs := [3]string{"dummyScriptName", "run", "--file"}
    invalidFileName := "abc"

    result := IsRunCommandValid(validArgs[:], invalidFileName)

    if result {
        t.Errorf("TestRunCommandValidation_InvalidFileName failed, expected %v, got %v", "false", result)
    }
}

func TestRunCommandValidation_ValidParams(t *testing.T) {
    validArgs := [3]string{"dummyScriptName", "run", "--file"}
    validFileName := "main.go"

    result := IsRunCommandValid(validArgs[:], validFileName)

    if !result {
        t.Errorf("TestRunCommandValidation_ValidParams failed, expected %v, got %v", "true", result)
    }
}

func TestRunCommandValidation_InvalidParams(t *testing.T) {
    invalidArgs := [3]string{"dummyScriptName", "run", "abc"}
    invalidFileName := "abc"

    result := IsRunCommandValid(invalidArgs[:], invalidFileName)

    if result {
        t.Errorf("TestRunCommandValidation_InvalidParams failed, expected %v, got %v", "false", result)
    }
}