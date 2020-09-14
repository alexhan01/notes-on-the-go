package extractText

import (
	"encoding/json"
	"fmt"
	"github.com/alexhan01/notes-on-the-go/backend/app/entity"
	// entity/ExtractedText
	"github.com/alexhan01/notes-on-the-go/data/"
)

func extractText() {
	// Initializing local var
	var text ExtractedText
	Data := 

	// Error Checking
	err := json.Unmarshal(Data, &text)
	if err != nil {
		fmt.Println(err)
	}

	// Print details of decoded data
	fmt.Println("Struct is:", text)
	fmt.Printf("%s's capital is %s and it is in %s. \n", text.Name, text.Capital, text.Continent)
}
