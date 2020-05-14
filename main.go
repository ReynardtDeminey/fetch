package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := getURL()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	statusCode := resp.Status

	fmt.Printf("%s", data)
	fmt.Printf("%v\n", statusCode)

}

func getURL() string {
	url := os.Args[1]
	if strings.HasPrefix(url, "http://") == false {
		return "http://" + url
	}
	return url
}
