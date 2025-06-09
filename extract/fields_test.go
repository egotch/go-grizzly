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
// - Write positive/negative test for format, parse, validate IntType
// - Write positive/negative test for format, parse, validate BoolType

// Tests the creation of a stringtype and
// validates each method
func TestStringField(t *testing.T) {

	var fields []Field

	// string field test
	stringField, _ := NewField("stringField", StringType{}, WithIndex(0))
	fields = append(fields, stringField)

	// make the field map
	fm := NewFieldMap(fields...)

	// retrieve field by name
	gotByNameStringField, found := fm.GetFieldByName("stringField")
	if ! found {
		t.Error("Unable to retrieve stringField by name")
	}

	// test returns "stringField" name
	log.Println("StringType - Testing field.Name")
	if gotByNameStringField.Name != "stringField" {
		t.Error("StringType - Unable to locate stringField by name")
	}
	// test name maches
	log.Println("StringType - Testing Type.Name()")
	if gotByNameStringField.Type.Name() != "string" {
		t.Error("String field was not set up as a string: Type.Name() return = ", gotByNameStringField.Type.Name())
	}
	// test zero val
	log.Println("StringType - Testing Type.Zero()")
	if gotByNameStringField.Type.Zero() != "" {
		t.Error("StringType - Unable to return correct zero value for string: Type.Zero() = ", gotByNameStringField.Type.Zero())
	}
	// test Validate()
	log.Println("StringType - Testing Type.Valdate()")
	if ! gotByNameStringField.Type.Validate(("I am a string")) {
		t.Error("StringType - Unable to confirm value IS A string: Type.Validate(I am a string) = ", gotByNameStringField.Type.Validate(("I am a string")))
	}
	if gotByNameStringField.Type.Validate(1234) {
		t.Error("StringType - unable to confirm value is NOT a string: Type.Validate(1234) = ", gotByNameStringField.Type.Validate(("I am a string")))
	}
	// test Format
	log.Println("StringType - Testing Type.Format()")
	if gotByNameStringField.Type.Format("I am a string") != "I am a string" {
		t.Error("StringType - failed to format string value: Type.Format(I am a string) = ", gotByNameStringField.Type.Format("I am a string"))
	}
	// test Parse
	log.Println("StringType - Testing Type.Parse()")
	if _, err := gotByNameStringField.Type.Parse("I am a string"); err != nil {
		t.Error("StringType - unable to successfully parse and return string value: gotByNameStringField.Type.Parse(I am a string) = ", err)
	}
}
