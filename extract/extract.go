package extract

import (
	"bufio"
	"os"
	"encoding/csv"
)

// ParseDelimitted iterates over a delimitted file
// for each row it creates a rowData struct (rd) per the field mapping provided
//
// returns a slice of rowData structs in the order of the original file
func ParseDelimitted(fullPath string, delimiter string, fieldMap FieldMap) ([]RowData, error) {

	var rows []RowData

	// open file
	f_in, err := os.Open(fullPath)
	if err != nil {
		return rows, err
	}

	// set up buffer for reading
	csvReader := csv.NewReader(f_in)
	
	inScanner := bufio.NewScanner(f_in)

	for inScanner.Scan() {
		csvReader.ReadAll()

	}
	



	return rows, nil
}
