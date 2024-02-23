package main

import (
	"fmt"

	"github.com/relaypro-open/dog_api_golang/api"
)

func main() {
	c := api.NewClient("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiIwZDY3NGRhNGFjNzYxMWVkYThmNzhmNzdmYTUyMWI2ZSJ9.dJAX9m8NJJ5ZzJLtDccEKUIGeIBnzhp8ze_zeq7Dei43LYC7JvzbPIG56LyA8vXPPCYicxp9Puk_U-mD5BhzacSA-v4_7ooj3olW4sed0EnpAXH-mEGojMWuBXsLfnls1GPyUjZZD1PNeF1sby4-FDU4P2pa-PylhPe7XMwvCjC169r3Ws59yiIVKvzQUQr2S1bR_EL1ongUn56okXzFrGPuT99xA8XlOAfNLTvTfMXD2JZffsknaeO0vk-IHvhQrJSzro8oqigj34S3FlqvSE5xAnMXNaJMv7NN1UvWWO_3n92k84CkfGfxpt1MuwkQsb0iBeqO4CyWe7e2gOknEA", "http://dog:8000/api/V2")

	HostsRes, HostsCode, HostsErr := c.GetHosts(nil)
	fmt.Printf("Hosts: %+v\n", HostsRes)
	fmt.Printf("HostsCode: %+v\n", HostsCode)
	fmt.Printf("Error: %+v\n", HostsErr)

	HostId := HostsRes[0].ID

	HostRes, HostCode, HostErr := c.GetHost(HostId, nil)
	fmt.Printf("Host: %+v\n", HostRes)
	fmt.Printf("HostCode: %+v\n", HostCode)
	fmt.Printf("Error: %+v\n", HostErr)

}
