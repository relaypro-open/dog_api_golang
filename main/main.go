package main

import (
	"dog_api_go/api"
	"fmt"
)

func main() {
	c := api.NewClient("my-key")

	HostsRes, HostsErr := c.GetHosts(nil)
	fmt.Printf("Hosts: %+v\n", HostsRes)
	fmt.Printf("Error: %+v\n", HostsErr)

	HostId := "eda000f6-0743-448f-874b-a7703ecddfaa"
	HostRes, HostErr := c.GetHost(HostId, nil)
	fmt.Printf("Host: %+v\n", HostRes)
	fmt.Printf("Error: %+v\n", HostErr)

}
