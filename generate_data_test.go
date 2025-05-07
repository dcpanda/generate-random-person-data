package main

import (
	"encoding/csv"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestParseFlags(t *testing.T) {

	numRows, outputFilename := parseFlags()

	// Default values should be used when no flags are provided
	if numRows != 100 {
		t.Errorf("Expected default numRows to be 100, got %d", numRows)
	}

	if outputFilename != "random_data.csv" {
		t.Errorf("Expected default outputFilename to be 'random_data.csv', got %s", outputFilename)
	}
}

func TestCreateCSVWriter(t *testing.T) {

	tempFile := "test_output.csv"
	defer os.Remove(tempFile) // Clean up after test

	file, writer, err := createCSVWriter(tempFile)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer file.Close()

	if file == nil {
		t.Error("Expected file to not be nil")
	}

	if writer == nil {
		t.Error("Expected writer to not be nil")
	}

}

func TestGetCSVHeader(t *testing.T) {
	expected := []string{
		"FirstName",
		"LastName",
		"Gender",
		"AddressLine1",
		"AddressLine2",
		"City",
		"State",
		"ZipCode",
		"Email",
		"PhoneNumber",
	}

	header := getCSVHeader()

	if !reflect.DeepEqual(header, expected) {
		t.Errorf("Expected header %v, got %v", expected, header)
	}
}

func TestWriteHeader(t *testing.T) {

	var sb strings.Builder
	writer := csv.NewWriter(&sb)

	err := writeHeader(writer)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	writer.Flush()

	expected := "FirstName,LastName,Gender,AddressLine1,AddressLine2,City,State,ZipCode,Email,PhoneNumber\n"
	if sb.String() != expected {
		t.Errorf("Expected header string %q, got %q", expected, sb.String())
	}
}

func TestGenerateRandomRow(t *testing.T) {
	row := generateRandomRow()

	expectedFields := 10
	if len(row) != expectedFields {
		t.Errorf("Expected row to have %d fields, got %d", expectedFields, len(row))
	}

	// Check that none of the fields are empty
	for i, field := range row {
		if field == "" {
			t.Errorf("Field %d is empty", i)
		}
	}
}
