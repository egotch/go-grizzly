package extract

import (
	"fmt"
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

// testing the fieldmap gen
// 
// headers
// "UserLoginName","AccountNumber","AccountType","RoutingNumber","CurrentBalance","AccountDescription"
func TestDelimmitedFieldSetup(t *testing.T) {

	var fields []Field

	// UserName
	userName, _ := NewField("userName", StringType{}, WithIndex(0))
	// UserEmail
	email, _ := NewField("email", StringType{}, WithIndex(1))
	// UserPhone
	phone, _ := NewField("phone", StringType{}, WithIndex(2))
	// Status
	userStatus, _ := NewField("status", BoolType{}, WithIndex(3))

	// Assemble the fields
	fields = append(fields, userName, email, phone, userStatus)
	// Make field map
	fm := NewFieldMap(fields...)

	// get fields
	userNameField, _ := fm.GetFieldByName("userName")

	fmt.Println(userNameField.Type.Format("hello world"))

}

