package test

import (
	"log"
	"os"

	// 	"strings"

	"bufio"

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
	userName, _ := extract.NewField("UserName", 0, extract.StringType{}, extract.AsCSV(","))
	// UserEmail
	email, _ := extract.NewField("Email", 1, extract.StringType{})
	// UserPhone
	phone, _ := extract.NewField("Phone", 2, extract.StringType{})
	// Status
	userStatus, _ := extract.NewField("IsActive", 3, extract.BoolType{})


	// Assemble the fields
	fields = append(fields, userName, email, phone, userStatus)
	// Make field map
	fm := extract.NewFieldMap(fields...)

}

func TestCSV() {
	log.Println("testing")

	setupFieldmap()

	f_in, err := os.Open("samples/LargeSample.csv")
	if err != nil {
		log.Panicf("Unable to open file: %v", err)
	}

	scanner := bufio.NewScanner(f_in)
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		// splitLine := strings.Split(line, ",")
		log.Printf("Line: %v | %v", lineNum, line)

	}
}
