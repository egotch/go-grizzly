package parse

import (
	"errors"
	"fmt"
)

type FieldMap struct{
	Fields []Field
}

type Field struct{
	Name string
	Index int
	Type DataType
}

type RowData struct{}

// Contstructor for new FieldMap
// returns a FieldMap
func NewFieldMap(firstField Field, additionalFields ...Field) *FieldMap {

	allFields := append([]Field{firstField}, additionalFields...)
	return &FieldMap{Fields: allFields}
}

// Getter method
func (fm *FieldMap) GetFields() []Field{
	return fm.Fields
}

// Add more fields
func (fm *FieldMap) AddField(f Field) error {
	// Check Field is valid
	if f.Name == "" || f.Type == nil {
		return fmt.Errorf("Unable to add invalid field.  Must include Name and Type")
	}

	// Check for duplicate names
	for _, field := range fm.Fields {
		if field.Name == f.Name {
			return fmt.Errorf("Unable to add Field, name already exists for '%v'", f.Name)
		}
	}

	// Add the field
	fm.Fields = append(fm.Fields, f)
	return nil
}

// GetFieldByName functions like a _,ok idiom
// if the name passed in matches one of the fields
// on FieldMap, then it is returned
// if not, an empty Field is returned with False
func (fm *FieldMap) GetFieldByName(inName string) (Field, bool) {
	for _, field := range fm.Fields{
		if field.Name == inName {
			return field, true
		}
	}
	return Field{}, false
}

// GetFieldByIndex functions like a _,ok idiom
// similar to GetFieldByName, but returns a field who's index
// matches the value passed into the function call
func (fm *FieldMap) GetFieldByIndex(inIdx int) (Field, bool) {

	for _, field := range fm.Fields{
		if field.Index == inIdx {
			return field, true
		}
	}
	return Field{}, false
}

// NewField creates a Field with validation
func NewField(name string, index int, dataType DataType) (Field, error) {

	if name == "" {
		return Field{}, errors.New("Error creating Field: Field name cannot be emptry")
	}
	if index < 0 {
		return Field{}, errors.New("Error creating Field: Field index cannot be less than zero")
	}
	return Field{Name: name, Index: index, Type: dataType}, nil
}
