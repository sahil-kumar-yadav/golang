package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("URL in GoLang")
	myurl := "http://example.com:8080/api/resource?api_key=123&value"
	fmt.Printf("Type of URL: %T\n", myurl)

	parsedurl, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Printf("ParsedUrl type %T\n", parsedurl)
	fmt.Printf("Parsed URL: %+v\n", parsedurl)

	fmt.Println("--------------------------------")
	// Accessing different parts of the URL
	fmt.Println("Scheme:", parsedurl.Scheme)
	fmt.Println("Host:", parsedurl.Host)
	fmt.Println("Path:", parsedurl.Path)
	fmt.Println("Query:", parsedurl.Query())

	// Modifying parts of the URL
	parsedurl.Host = "new-example.com"
	parsedurl.Path = "/new-api/resource"
	query := url.Values{}
	query.Set("new_key", "new_value")
	parsedurl.RawQuery = query.Encode()
	fmt.Println("Modified URL:", parsedurl.String())

	fmt.Println("--------------------------------")
	// Encoding and Decoding URLs
	encodedURL := url.QueryEscape(parsedurl.String())
	fmt.Println("Encoded URL:", encodedURL)

	decodedURL, err := url.QueryUnescape(encodedURL)
	if err != nil {
		fmt.Println("Error decoding URL:", err)
		return
	}
	fmt.Printf("Decoded URL type %T\n", decodedURL)
	fmt.Println("Decoded URL:", decodedURL)

	fmt.Println("--------------------------------")
	// Parsing and validating URLs
	validURL := "https://www.example.com/valid-path?param=value#fragment"
    parsedValidURL, err := url.Parse(validURL)
    if err!= nil {
        fmt.Println("Error parsing valid URL:", err)
        return
    }
    fmt.Printf("Parsed valid URL type %T\n", parsedValidURL)
    fmt.Printf("Parsed valid URL: %+v\n", parsedValidURL)

    invalidURL := "invalid-url"
    parsedInvalidURL, err := url.Parse(invalidURL)
    if err!= nil {
        fmt.Println("Error parsing invalid URL:", err)
        return
    }
    fmt.Printf("Parsed invalid URL type %T\n", parsedInvalidURL)
    fmt.Printf("Parsed invalid URL: %+v\n", parsedInvalidURL)
	fmt.Println("--------------------------------")
	






}
