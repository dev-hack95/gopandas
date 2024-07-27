package gopandas_test

import (
	"fmt"
	"testing"

	pd "github.com/dev-hack95/gopandas"
)

func TestReadCsv(t *testing.T) {
	df, err := pd.ReadCsv("csv/train.csv")

	if err != nil {
		t.Error("Error: ", err.Error())
		return
	}

	fmt.Println(df)
}
