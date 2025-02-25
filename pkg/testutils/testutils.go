package testutils

import (
	"encoding/json"
	"fmt"

	"os"
)

func OpenMock(path string, result interface{}) interface{} {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening json file")
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(result)
	if err != nil {
		fmt.Println("Error unmarshalling json file")
		return nil
	}

	return result
}
