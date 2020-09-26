package extracttext

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// ExtractedText represents the raw extracted json file in golang form
type ExtractedText struct {
	// Temporary struct variables
	Name      string `json:"Name"`
	Capital   string `json:"Capital"`
	Continent string `json:"Continent"`
}

func extractText(str string) string {
	// Initializing local var
	var text ExtractedText
	filePath := str

	// Open
	data, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened test.json")
	defer data.Close()

	byteValue, _ := ioutil.ReadAll(data)

	// Error Checking
	json.Unmarshal([]byte(byteValue), &text)

	// Print details of decoded data
	// var n string = text.Name
	// var cap string = text.Capital
	var cont string = text.Continent

	return cont
	// fmt.Printf("Struct is: %s", text.Name)
	// fmt.Printf("%s's capital is %s and it is in %s. \n", text.Name, text.Capital, text.Continent)
}
