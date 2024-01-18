package api

import (
	"fmt"
	"encoding/json"
)


func PrettyPrint(title string, incoming interface{}) {
	d, _ := json.MarshalIndent(incoming, "", "  ")
	fmt.Println("=", title)
	fmt.Println(string(d))
	fmt.Println("=end", title)
}
