package helper

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(v interface{}) {
	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(bytes))
}
