package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// DivisionData holds the name and ID of a division
type DivisionData struct {
	Name string `json:"divisionName"`
	ID   string `json:"divisionId"`
}

func main() {

	// divisionsFile, err := os.Open("data.json")
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println("Successfully Opened users.json")
	// defer divisionsFile.Close()

	// log.Println(divisionsFile)

	// byteValue, _ := ioutil.ReadAll(divisionsFile)

	// var result map[string][]map[string]string
	// json.Unmarshal([]byte(byteValue), &result)
	// log.Println(result["production"][0]["divisionName"])

	// var conversationConfig ConversationConfig
	// err = json.Unmarshal([]byte(byteValue), &conversationConfig.test)
	// log.Println(conversationConfig)

	// if err != nil {
	// 	log.Println(err)
	// }

	f, err := os.Open("data.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	divisionConfig := make(map[string][]DivisionData)
	dec := json.NewDecoder(f)
	err = dec.Decode(&divisionConfig)
	if err != nil {
		panic(err)
	}
	for _, d := range divisionConfig["development"] {
		log.Println(d.ID)
	}
	fmt.Printf("Development: %+v\n", divisionConfig["development"][0].ID)
	fmt.Printf("Staging: %+v\n", divisionConfig["staging"])
	fmt.Printf("Production: %+v\n", divisionConfig["production"])
}
