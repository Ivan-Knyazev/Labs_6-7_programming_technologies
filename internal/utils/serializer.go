package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

func SerializeAndPrint(object any) {
	objectJSON, err := json.MarshalIndent(object, "", "	")
	if err != nil {
		log.Fatal("Error marshaling to JSON:", err)
	}
	fmt.Println(string(objectJSON))
}
