package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// GetNimblFileList send a request to https://nimbl.blackbeartechhive.com/api/v1/list
func GetNimblFileList() ([]string, error) {
	// GET https://nimbl.blackbeartechhive.com/api/v1/list
	type response struct {
		Files []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"files"`
	}

	ret := make([]string, 0)

	var res response
	// send a http request
	resp, err := http.Get("https://nimbl.blackbeartechhive.com/api/v1/list")
	if err != nil {
		return ret, err
	}
	defer resp.Body.Close()
	// read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ret, err
	}

	fmt.Println("Response:", string(body))
	// parse response
	err = json.Unmarshal(body, &res)
	if err != nil {
		return ret, err
	}

	// Search all file start with bbnim and type is zip
	for _, file := range res.Files {
		if strings.HasPrefix(file.Name, "bbnim") && file.Type == ".zip" {
			ret = append(ret, file.Name)
		}
	}

	return ret, nil
}

// downloadFile download a file from https://nimbl.blackbeartechhive.com/api/v1/files/{filename}
func downloadFile(filename string, filepath string) error {
	// create a file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get("https://nimbl.blackbeartechhive.com/api/v1/files/" + filename)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// check response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad sataus %s", resp.Status)
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil

}
