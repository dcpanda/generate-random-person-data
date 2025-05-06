# generate-random-person-data
Generate File with Random Person Data for Testing

## Dependency 
1. Install the go-randomdata package (see go.mod). Run this command.
```bash
go get github.com/pallinder/go-randomdata
```
2. Some inbuild packages like `encoding/csv`, `fmt`, `flag`, `os`, `strconv` etc. are used.
## Running the program directly 
```bash
go run generate_data.go -rows 500 -output my_random_users.csv
```
1. `-rows` takes the number of rows to produce in the file.
2. `-output` is the name of the file with csv extension.

## Building and Running the executable (to share)

1. First build the executable 
```bash
go build generate_data.go
```
2. Running the executable

```bash
./generate_data -rows 1000 -output large_dataset.csv
```

## Default Values
if you run the program without any flags, it will use the default values: `100` rows and `random_data.csv` as the filename.
This code will work without any parameters.
```bash
./generate_data
```
or
```bash
go run generate_data.go
```

## Adding new columns
There are two sections that will require change to add or modify data columns. The `header` slice has the row reader and the `row` slice has the data rows.
Visit this pkg URL [go-randomdata](https://pkg.go.dev/github.com/grandper/go-randomdata) package for more functions.