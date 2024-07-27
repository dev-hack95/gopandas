package gopandas

import (
	"encoding/csv"
	"errors"
	"os"
)

type recods = [][]string

func ReadCsv(path string) (recods, error) {

	file, err := os.Open(path)

	if err != nil {
		return nil, errors.New("file path is not present")
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		return nil, errors.New("error occured at reading records")
	}

	return records, nil
}
