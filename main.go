package main

import (
	"encoding/csv"
	"fmt"
	"github.com/hitchnsmile/gohunting"
	"os"
)

func main() {
	apiKey := os.Getenv("HUNTER_API")
	client := gohunting.Client(apiKey)

	fmt.Println(client)

	filePath := os.Getenv("FILE_PATH")
	// read the file
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	if err = f.Close(); err != nil {
		fmt.Println(err)
	}

	// add column
	l := len(lines)
	for i := 1; i < l; i++ {
		searchResponse, err := client.Search(lines[i][5])
		if err != nil {
			fmt.Println(err)
		}
		if len(searchResponse.Data.Emails) > 0 {
			lines[i] = append(lines[i], searchResponse.Data.Emails[0].Value)
		}
	}

	// write the file
	f, err = os.Create("new" + filePath)
	if err != nil {
		fmt.Println(err)
	}
	w := csv.NewWriter(f)
	if err = w.WriteAll(lines); err != nil {
		f.Close()
		fmt.Println(err)
	}

	f.Close()
}
