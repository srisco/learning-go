package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Example to show how to read and manipulate the minio config file in JSON
func main() {
	configFile := readFile("configtest.json")

	var content interface{}

	err := json.Unmarshal(configFile, &content)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		contentMap := content.(map[string]interface{})
		notifications := contentMap["notify"].(map[string]interface{})
		webhooks := notifications["webhook"].(map[string]interface{})
		//fmt.Println(webhooks["10"])

		newWebhook := make(map[string]interface{})
		newWebhook["enable"] = true
		newWebhook["endpoint"] = "http://localhost/test-endpoint"

		webhooks["1"] = newWebhook
		//fmt.Println(contentMap)

		result, _ := json.Marshal(contentMap)
		fmt.Println(string(result))
	}

}

func readFile(filename string) []byte {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return bs
}
