package test

import (
	"fmt"
	"log"

	"github.com/egotch/go-grizzly/extract"
)

var (
	fm extract.FieldMap
)


// testing the fieldmap gen
// 
// headers
// "UserLoginName","AccountNumber","AccountType","RoutingNumber","CurrentBalance","AccountDescription"
func setupFieldmap() {

	var fields []extract.Field

	// UserName
	userName, _ := extract.NewField("userName", extract.StringType{}, extract.WithIndex(0))
	// UserEmail
	email, _ := extract.NewField("email", extract.StringType{}, extract.WithIndex(1))
	// UserPhone
	phone, _ := extract.NewField("phone", extract.StringType{}, extract.WithIndex(2))
	// Status
	userStatus, _ := extract.NewField("status", extract.BoolType{}, extract.WithIndex(3))

	// Assemble the fields
	fields = append(fields, userName, email, phone, userStatus)
	// Make field map
	fm := extract.NewFieldMap(fields...)

	// get fields
	userNameField, _ := fm.GetFieldByName("userName")
	fmt.Println(userNameField.Type.Format("hello world"))

}

func TestCSV() {
	log.Println("testing")

	setupFieldmap()

}
