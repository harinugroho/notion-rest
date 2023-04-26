package main

import (
	"encoding/json"
	"fmt"
)

func dd(vals ...interface{}) {
	for i, val := range vals {
		fmt.Printf("============== Start: %v | %T\n", i, val)
		if _, ok := val.([]byte); ok {
			fmt.Println(string(val.([]byte)))
		} else {
			raw, _ := json.MarshalIndent(val, "", "    ")
			fmt.Println(string(raw))
		}
		fmt.Printf("============== End    \n")
	}
}
