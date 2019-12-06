package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var floatArrays []float64
	strArray := "[\"1243\",\"567\"]"
	if err := json.Unmarshal([]byte(strArray), &floatArrays); err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("%v", floatArrays)
}
