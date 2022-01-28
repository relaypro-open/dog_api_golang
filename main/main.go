package main

import (
	"context"
	"dog_api_go/api"
	"fmt"
)

func main() {
	//c := NewClient(os.Getenv("HostST_INTEGRATION_API_KEY"))
	c := api.NewClient("DUMMY_API_KEY")
	fmt.Printf("c: %v\n", c)

	ctx := context.Background()
	res, err := c.GetHosts(ctx, nil)

	fmt.Printf("res: %v\n", res)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("res.Hosts: %v\n", res.Hosts)

}

//func main() {
//	url := "http://dog-ubuntu-server.lxd:7070/api/hosts"
//	res, getErr := http.Get(url)
//	if getErr != nil {
//		log.Fatal(getErr)
//	}
//	body, readErr := ioutil.ReadAll(res.Body)
//	if readErr != nil {
//		log.Fatal(readErr)
//	}
//	fmt.Println(body)
//	fullResponse := successResponse{
//		Data: v,
//	}
//	fmt.Println(json.NewDecoder(res.Body).Decode(&fullResponse)
//}
