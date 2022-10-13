package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var sites []string

	for fileScanner.Scan() {
		sites = append(sites, fileScanner.Text())
	}

	readFile.Close()

	for _, site := range sites {
		statusCode := VisitAndReturnStatusCode(site)
		fmt.Print(statusCode, " ")
		fmt.Println(site, StatusDesc(statusCode))

	}
}

func VisitAndReturnStatusCode(url string) int {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	client := new(http.Client)
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return response.StatusCode
}

func StatusDesc(statusCode int) string {
	switch statusCode {
	case 200:
		return "OK"
	case 301:
		return "Moved Permanently"
	case 302:
		return "Found"
	case 303:
		return "See Other"
	case 307:
		return "Temporary Redirect"
	case 400:
		return "Bad Request"
	case 401:
		return "Unauthorized"
	case 403:
		return "Forbidden"
	case 404:
		return "Not Found"
	case 500:
		return "Internal Server Error"
	case 502:
		return "Bad Gateway"
	case 503:
		return "Service Unavailable"
	default:
		return "Unknown"
	}
}
