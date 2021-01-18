package main

import (
	"encoding/json"
	"fmt"
	"json/common"

	"github.com/davecgh/go-spew/spew"
)

func main() {

	// SimpleTest()
	UnmarshalTest()
}

// SimpleTest - unmarshal
func SimpleTest() {
	data := []byte(`{ "name": "rich", "ages": [1,2,3,4]}`)
	var dat map[string]interface{}
	json.Unmarshal(data, &dat)
	spew.Dump(dat)
}

// UnmarshalTest - marshal and unmarshal test
func UnmarshalTest() {

	x := common.Account{
		Name: "Richard",
		Id:   1,
	}
	y, err := json.Marshal(x)
	if err != nil {
		panic(err)
	} else {
		spew.Dump(y)
	}

	dat := common.Account{}
	json.Unmarshal(y, &dat)

	fmt.Println(dat.Name)
	fmt.Println(dat.Id)
}
