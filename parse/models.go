package parse

type FieldMap struct{
	Fields []Field
}

type Field struct{
	Name string
	Index int
	Type interface{}
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
