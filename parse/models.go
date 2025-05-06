package parse

import "errors"

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
func (fm *FieldMap) AddField(f Field) {
	fm.Fields = append(fm.Fields, f)
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
