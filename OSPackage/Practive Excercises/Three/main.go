package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Settings struct {
	Name     string `json:"name"`
	Theme    string `json:"theme"`
	FontSize int    `json:"fontSize"`
}

func main() {
	start := time.Now()
	fmt.Printf("Start Time: %v\n", start)

	err := os.MkdirAll("./workspace/data/", 0755)
	if err != nil {
		fmt.Printf("%q\n", err)
		return
	} else {
		fmt.Printf("%q\n", "Directory Created Successfully")
	}

	file, err := os.OpenFile("setting.json", os.O_RDWR|os.O_CREATE|os.O_EXCL|os.O_APPEND, 0755)
	if err != nil {
		fmt.Printf("%q\n", err)
	} else {
		fmt.Printf("%q\n", "file Is Created and Opened Successfully.")
	}

	err = file.Close()
	if err != nil {
		fmt.Printf("%q\n", err)
		return
	}

	info, err := os.Stat("./setting.json")
	if err != nil {
		fmt.Printf("%q\n", err)
		return
	} else {
		if info != nil {
			err := os.Rename("./setting.json", "./workspace/data/setting.json")
			if err != nil {
				fmt.Printf("%q\n", err)
				return
			} else {
				fmt.Printf("%q\n", "file Is Moved In Destination Successfully.")
			}
		}
	}

	var jsonSettingData Settings
	jsonSettingData.Name = "Kasra"
	jsonSettingData.Theme = "dark"
	jsonSettingData.FontSize = 16

	data, err := json.MarshalIndent(jsonSettingData, "", "\t")
	if err != nil {
		fmt.Printf("%q\n", err)
		return
	} else {
		fmt.Printf("%q\n", `"data" Has been inserted successfully
		DATA:`, data)
	}

	err = os.WriteFile("./workspace/data/setting.json", data, 0644)
	if err != nil {
		fmt.Printf("%q\n", err)
		return
	} else {
		fmt.Printf("%q\n", "JASON Content Has Wrote Successfully")
	}

	data, err = os.ReadFile("./workspace/data/setting.json")
	if err != nil {
		fmt.Printf("%q\n", err)
		return
	} else {
		fmt.Println(string(data))
	}

	defer file.Close()

	fmt.Printf("Elapsed: %v\n", time.Since(start))
}
