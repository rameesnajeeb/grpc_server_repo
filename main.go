package main

import (
	"encoding/json"
	"fmt"
	"grpc_server_repo/models/files"
)

func main() {
	ietf := Interface{
		Name:                 "eth0",
		Description:          "Ethernet interface",
		IntfType:             "ethernet",
		Enabled:              true,
		LinkUpDownTrapEnable: 1,
	}

	f := files.File{FileName: "test.txt", FilePath: "/home/thinkpalm/grpc_server_repo", Contents: ietf}

	intfJSON, err := json.Marshal(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(intfJSON))
}
