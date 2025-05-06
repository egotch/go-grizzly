package test

import (
	"log"
	"os"
// 	"strings"

	"bufio"
	"github.com/egotch/go-grizzly/parse"
)

var (
	fm parse.FieldMap
)


// testing the fieldmap gen
// 
// headers
// "UserLoginName","AccountNumber","AccountType","RoutingNumber","CurrentBalance","AccountDescription"
func setupFieldmap() {

	field1, err := parse.NewField("UserLogin", 0, parse.StringType{})
	if err != nil {
		log.Panicf("Unable to generate field: %v", err)
	}

	log.Println(field1.Type.Name())

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
