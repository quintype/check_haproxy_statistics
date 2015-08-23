package haproxy

import (
	"io"
	"encoding/csv"
	"errors"
)

type CSVReader struct {
	csvReader *csv.Reader
	headers []string
}

func NewCSVReader(reader io.Reader) (*CSVReader) {
	return &CSVReader{csvReader: csv.NewReader(reader)}
}

func (reader *CSVReader) Read() (map[string]string, error) {
	row, err := reader.csvReader.Read()
	if (err != nil) {
		return nil, err
	}
	if(reader.headers == nil) {
		reader.headers = row;
		return reader.Read()
	}

	if(len(reader.headers) != len(row)) {
		return nil, errors.New("Incorrect number of csv entries")
	}

	result := make(map[string]string)
	for i := 0; i < len(row); i++ {
		result[reader.headers[i]] = row[i]
	}

	return result, nil
}

func (reader *CSVReader) ReadAll() ([] map[string]string, error) {
	result := []map[string]string{}

	row, err := reader.Read()
	for err == nil {
		result = append(result, row)
		row, err = reader.Read()
	}

	if (err != io.EOF) {
		return nil, err
	}

	return result, nil
}
