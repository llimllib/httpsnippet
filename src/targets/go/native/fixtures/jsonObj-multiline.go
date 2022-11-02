package main

import (
	"fmt"
	"strings"
	"net/http"
	"io"
)

func main() {

	url := "http://mockbin.com/har"

	payload := strings.NewReader("{\n  \"foo\": \"bar\"\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}