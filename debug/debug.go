package debug

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// PrintMarshal pretty prints a struct for use in debugging
func PrintMarshal(msg string, v interface{}) {
	log.Println(msg)
	pp, _ := json.MarshalIndent(v, "", "    ")
	fmt.Println(string(pp))
}

// PrintJSONBytes pretty prints a JSON byte array
func PrintJSONBytes(msg string, v []byte) {
	var out bytes.Buffer
	err := json.Indent(&out, v, "", "    ")
	if err != nil {
		log.Printf("error indenting JSON []byte: %v", err)
		return
	}
	log.Printf("%s %s\n", msg, out.String())
}

//Print allows for generic debug printing
func Print(msg string, v interface{}) {
	log.Printf("%s%v", msg, v)
}
