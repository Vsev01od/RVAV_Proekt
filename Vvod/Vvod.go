package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	var vvod string
	fmt.Scanln(&vvod)
	data := map[string]interface{}{
		"message": vvod,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:8080/", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		println("Request sent successfully")
	} else {
		println("Error:", resp.Status)
	}
}
