package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/pallinder/go-randomdata"
	"os"
)

func main() {

	numRows := flag.Int("rows", 100, "Number of rows to generate")
	outputFilename := flag.String("output", "random_data.csv", "Output CSV filename")

	flag.Parse()

	fmt.Printf("Generating %d rows of random data into %s\n", *numRows, *outputFilename)

	file, err := os.Create(*outputFilename)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	// setup header for easy modification
	header := []string{
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
	if err := writer.Write(header); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing header: %v\n", err)
		return
	}

	for i := 0; i < *numRows; i++ {
		profile := randomdata.GenerateProfile(randomdata.RandomGender)
		row := []string{
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

		if err := writer.Write(row); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing row %d: %v\n", i+1, err)
			continue
		}
	}

	fmt.Printf("Successfully generated %d rows and saved to %s\n", *numRows, *outputFilename)
}
