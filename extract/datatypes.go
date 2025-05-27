// library of types used in mapping fields
// for file consumption

package extract

import (
	"fmt"
	"strconv"
)

// main datatype interface
// declares that any DataType used in mapping of fields
// must include the following methods
type DataType interface{
	// Name returns the name of the datatype being used
	Name() string

	// Zero returns the zero value of the datatype
	Zero() any

	// Validate, validates that the assigned datatype works for the value
	// that was passed in
	// returns true/false if the value supplied is of the datatype used
	Validate(v any) bool

	// Format returns a formatted string of the value
	// corresponding to the datatype used
	Format(v any) string

	// Parse attempts to convert the string value provided into the
	// datattype assigned.
	// Returns the converted datatype value as a string, int, bool, etc
	// and error message if the string value could not be converted
	Parse(s string) (any, error)
}


// StringType implementation
type StringType struct{}

func (s StringType) Name() string { return "string" }
func (s StringType) Zero() any { return "" }
func (s StringType) Validate(v any) bool {
	_, ok := v.(string)
	return ok
}
func (s StringType) Format(v any) string {
	if s.Validate(v){
		return v.(string)
	}
	return fmt.Sprintf("%v", v)
}
func (s StringType) Parse(str string) (any, error) {
	return str, nil
}

// IntType implementation
type IntType struct{}

func (i IntType) Name() string { return "int" }
func (i IntType) Zero() any { return 0 }
func (i IntType) Validate(v any) bool {
	_, ok := v.(int)
	return ok
}
func (i IntType) Format(v any) string {
	if i.Validate(v) {
		return fmt.Sprintf("%d", v.(int))
	}
	return fmt.Sprintf("%v", v)
}
func (i IntType) Parse(str string) (any, error) {
	intVal, err := strconv.Atoi(str)
	if err != nil {
		return nil, fmt.Errorf("cannot parse '%v' as int: %w", str, err)
	}
	return intVal, nil
}

// BoolType implementation
type BoolType struct{}

func (b BoolType) Name() string { return "bool" }
func (b BoolType) Zero() any { return false }
func (b BoolType) Validate(v any) bool {
	_, ok := v.(bool)
	return ok
}
func (b BoolType) Format(v any) string {
	if b.Validate(v) {
		return fmt.Sprintf("%t", v.(bool))
	}
	return fmt.Sprintf("%v", v)
}
func (b BoolType) Parse(str string) (any, error) {
	boolVal, err := strconv.ParseBool(str)
	if err != nil {
		return nil, fmt.Errorf("cannot parse '%v' as bool: %w", str, err)
	}
	return boolVal, nil
}
