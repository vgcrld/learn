package main

import (
	"fmt"
	"rest/nutanix"
)

func main() {

	var api nutanix.API

	api = nutanix.NewRequest("galileosvc@disapoc.com", "vion123@")
	body := api.Request("clusters/list")
	fmt.Println(body)

}
