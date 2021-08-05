package interface_test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
)

type GitHubResponse []struct {
	FullName string `json:"full_name"`
}

type customWriter struct{}

func (w customWriter) Write(p []byte) (n int, err error) {
	var resp GitHubResponse
	json.Unmarshal(p, &resp)
	for _, r := range resp {
		fmt.Println(r.FullName)
	}
	return len(p), nil
}

func TestWriter(t *testing.T) {
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	writer := customWriter{}
	io.Copy(writer, resp.Body)
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		_, _ = fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func TestHttp(t *testing.T) {
	db := database{"Go T-Shirt": 25, "Go Jacket": 55}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
