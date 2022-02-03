package main

import (
	"dog_api_go/api"
	"fmt"
)

func main() {
	c := api.NewClient("my-key")

	HostsRes, HostsErr := c.GetHosts(nil)
	fmt.Printf("res: %v\n", HostsRes)
	fmt.Printf("err: %v\n", HostsErr)

	HostId := "eda000f6-0743-448f-874b-a7703ecddfaa"
	res, err := c.GetHost(HostId, nil)
	fmt.Printf("res: %v\n", res)
	fmt.Printf("err: %v\n", err)

}
