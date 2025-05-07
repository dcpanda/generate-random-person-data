package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/pallinder/go-randomdata"
	"os"
)

// parseFlags parses command-line flags and returns the number of rows to generate and the output filename
func parseFlags() (int, string) {
	numRows := flag.Int("rows", 100, "Number of rows to generate")
	outputFilename := flag.String("output", "random_data.csv", "Output CSV filename")

	flag.Parse()

	fmt.Printf("Generating %d rows of random data into %s\n", *numRows, *outputFilename)

	return *numRows, *outputFilename
}

// createCSVWriter creates a CSV writer for the specified filename and returns the file and writer
func createCSVWriter(filename string) (*os.File, *csv.Writer, error) {
	file, err := os.Create(filename)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		return nil, nil, err
	}

	writer := csv.NewWriter(file)

	return file, writer, nil
}

// getCSVHeader returns the header row for the CSV file
func getCSVHeader() []string {
	return []string{
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
}

// writeHeader writes the header row to the CSV file
func writeHeader(writer *csv.Writer) error {
	header := getCSVHeader()
	if err := writer.Write(header); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing header: %v\n", err)
		return err
	}
	return nil
}

// generateRandomRow generates a single row of random person data
func generateRandomRow() []string {
	profile := randomdata.GenerateProfile(randomdata.RandomGender)
	return []string{
		profile.Name.First,
		profile.Name.Last,
		profile.Gender,
		randomdata.StreetForCountry("US"),
		randomdata.Address(),
		randomdata.City(),
		randomdata.State(randomdata.Small),
		randomdata.PostalCode("US"),
		randomdata.Email(),
		randomdata.StringNumberExt(2, "-", 3) + "-" + randomdata.StringNumber(2, ""), //US phone number
	}
}

// writeRows generates and writes the specified number of random data rows to the CSV file
func writeRows(writer *csv.Writer, numRows int) {
	for i := 0; i < numRows; i++ {
		row := generateRandomRow()

		if err := writer.Write(row); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing row %d: %v\n", i+1, err)
			continue
		}
	}
}

func main() {
	numRows, outputFilename := parseFlags()

	file, writer, err := createCSVWriter(outputFilename)
	if err != nil {
		os.Exit(1)
	}

	defer file.Close()
	defer writer.Flush()

	if err := writeHeader(writer); err != nil {
		return
	}

	writeRows(writer, numRows)

	fmt.Printf("Successfully generated %d rows and saved to %s\n", numRows, outputFilename)
}
