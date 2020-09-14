package extracttext

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func extractText(str string) {
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
	fmt.Println("Struct is:", text)
	fmt.Printf("%s's capital is %s and it is in %s. \n", text.Name, text.Capital, text.Continent)
}
