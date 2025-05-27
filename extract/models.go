package extract

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

type Option func(*Field) error

type FieldMap struct{
	Fields []Field
}

type Field struct{
	Name string
	Index int
	Offset int
	OffsetLen int
	Type DataType
	Delimiter string
}

type RowData struct{}

// Contstructor for new FieldMap
// returns a FieldMap
func NewFieldMap(fields ...Field) *FieldMap {

	allFields := append([]Field{fields[0]}, fields[1:]...)
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
func NewField(name string, dataType DataType, options ...Option) (Field, error) {

	if name == "" {
		return Field{}, errors.New("Error creating Field: Field name cannot be emptry")
	}

	f := Field{Name: name, Type: dataType}

	for _, option := range options{
		err := option(&f)
		if err != nil {
			return f, fmt.Errorf("Failed to apply %v: %w", runtime.FuncForPC(reflect.ValueOf(option).Pointer()).Name(), err)
		}
	}
	
	return f, nil
}

// WithIndex specifies the index of the field (first is 0)
// this is to be used with DELIMITED files where the delimiter
// is later specified on the extraction engine.
//
// ex: someField = NewField("MyField", StringType{}, WithIndex(0))
func WithIndex(i int) Option {
	return func(f *Field) error {
		f.Index = i

		if f.Index < 0 {
			return errors.New("Invalid index, cannot be less than zero")
		}
		return nil
	}
}

// WithOffset specifies the start char (position 1 = 1) of the field
// followed by the length of the field.  This is to be used with
// FIXED WIDTH files
//
// ex: someField = NewField("MyField", StringType{}, WithOffset(1, 15))
func WithOffset(start int, length int) Option {
	return func(f *Field) error {
		f.Offset = start
		f.OffsetLen = length

		if f.Offset < 0 {
			return errors.New("Invalid offset, cannot be less than zero")
		}
		return nil
	}
}



