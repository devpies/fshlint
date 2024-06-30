package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFshErrorListener tests the FshErrorListener's ability to capture syntax errors.
func TestFshErrorListener(t *testing.T) {
	// Given
	errorListener := NewFshErrorListener()

	// When
	errorListener.SyntaxError(nil, nil, 1, 1, "test error", nil)

	// Then
	expectedError := "line 1:1 test error"
	assert.Equal(t, 1, len(errorListener.errors))
	assert.Equal(t, expectedError, errorListener.errors[0])
}

// TestLintFile tests the lintFile function for identifying errors in a FSH file.
func TestLintFile(t *testing.T) {
	// Given
	content := `
	Alias: SCT = http://snomed.info/sct

	Profile: PatientProfile
	Parent: Patient
	Id: patient-profile

	* name 0..0l
	`
	filePath := "test.fsh"
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}
	defer os.Remove(filePath)

	// When
	errors, err := lintFile(filePath)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, 1, len(errors))
	assert.Contains(t, errors[0], "extraneous input '0..0l'")
}

// TestGetLineContent tests the getLineContent function for retrieving the correct line from a file.
func TestGetLineContent(t *testing.T) {
	// Given
	content := "line1\nline2\nline3\n"
	filePath := "test_lines.fsh"
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}
	defer os.Remove(filePath)

	// When
	lineContent, err := getLineContent(filePath, 2)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, "line2", lineContent)
}

// TestFormatError tests the formatError function for correct error formatting.
func TestFormatError(t *testing.T) {
	// Given
	filePath := "test.fsh"
	line := 3
	column := 5
	msg := "test error"

	content := "line1\nline2\nline3\n"
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}
	defer os.Remove(filePath)

	// When
	formattedError := formatError(filePath, line, column, msg)

	// Then
	expectedIndicator := "     ^"
	expectedFormattedError := fmt.Sprintf("%s%s%s:%d:%d: %s%s%s\n%s\n%s\n",
		colorBold, filePath, colorReset, line, column,
		colorLightRed, msg, colorReset,
		"line3",
		expectedIndicator,
	)
	assert.Equal(t, expectedFormattedError, formattedError)
}
