// library of types used in mapping fields
// for file consumption

package parse

import (
)

type DataType interface{
	Name() string					// Returns the datatype name
	Zero() interface{}				// Provide the zero value of the data type
	Validate(v interface{}) bool	// Validate that datatype is mapped correctly
	Format(v interface{}) string	// Format the data value for display
	Parse(s string) (interface{}, error)	// Parse data value from read in string
}


// StringType implementation
type StringType struct{}

func (s StringType) Name() string { return "string" }
func (s StringType) Zero() interface{} { return "" }
func (s StringType) Validate(v interface{}) bool {
	_, ok := v.(string)
	return ok
}

// IntType implementation
type IntType struct{}

func (i IntType) Name() string { return "int" }
func (i IntType) Zero() interface{} { return 0 }
func (i IntType) Validate(v interface{}) bool {
	_, ok := v.(int)
	return ok
}

// BoolType implementation
type BoolType struct{}

func (b BoolType) Name() string { return "bool" }
func (b BoolType) Zero() interface{} { return false }
func (b BoolType) Validate(v interface{}) bool {
	_, ok := v.(bool)
	return ok
}
