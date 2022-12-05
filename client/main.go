package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Second}
	req, err := http.NewRequest(`HEAD`, `https://api.github.com/`, nil)
	if err != nil {
		fmt.Printf("Error: %s\\n", err)
		return
	}
	req.Header.Add(`Accept`, `application/json`)
	req.Header.Add(`Authorization`, `Basic `+basicAuth(`myusername`, `mypassword`))
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Error: %s\\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Printf("Body: %s\n", body)
}

func basicAuth(u, p string) string {
	auth := u + `:` + p

	return base64.StdEncoding.EncodeToString([]byte(auth))
}
