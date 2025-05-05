package parse

import (
	"os"
)


// ParseDelimitted iterates over a delimitted file
// for each row it creates a rowData struct (rd) per the field mapping provided
//
// returns a slice of rowData structs in the order of the original file
func ParseDelimitted(fielPath string, delimiter string, fieldMap FieldMap) []RowData {

	// check validity 
}
