package extracttext

import (
	"testing"
)

func TestExtractText(t *testing.T) {
	var filePath string = "/home/alexh/Documents/Git/notesonthego/data/test.json"
	value := extractText(filePath)
	if value != "Asia" {
		t.Errorf("Data not imported properly")
	}
}
