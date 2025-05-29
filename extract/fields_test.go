package extract

import (
	"log"
	"testing"
)

var (
	fm FieldMap
)

// TODO:
// - Write test for creating each data type
// - Write test for recalling field by name
// - Write test for recalling field by index
// - Write test for recalling field by offset
// - Write positive/negative test for format, parse, validate StringType
// - Write positive/negative test for format, parse, validate IntType
// - Write positive/negative test for format, parse, validate BoolType

// Tests the creation of delimited Field types and Field Map
// and their implicit functionality
func TestDelimitted(t *testing.T) {

	var fields []Field

	// string field test
	t.Run("TestStringField", func(t *testing.T){
		stringField, _ := NewField("stringField", StringType{}, WithIndex(0))
		fields = append(fields, stringField)

		// make the field map
		fm := NewFieldMap(fields...)

		// retrieve field by name
		gotByNameStringField, found := fm.GetFieldByName("stringField")
		if ! found {
			t.Error("Unable to retrieve stringField by name")
		}

		log.Println(gotByNameStringField.Type.Name())
		if gotByNameStringField.Type.Name() != "string" {
			t.Error("String field was not set up as a string")
		}

		log.Println(gotByNameStringField.Name)
		if gotByNameStringField.Name != "stringField" {
			t.Error("fields not set correctly on stringField")
		}


	})

}
