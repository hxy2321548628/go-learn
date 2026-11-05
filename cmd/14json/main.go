package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	hxy := &User{
		Name:     "hxy",
		Age:      23,
		Password: "1233",
	}

	vo, err := json.MarshalIndent(hxy, "", "  ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("vo: %v\n", string(vo))
	}

	var hxy2 User
	if err := json.Unmarshal(vo, &hxy2); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("hxy2: %v\n", hxy2)

}
