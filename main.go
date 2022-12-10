package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func main() {
	baseUrl := "http://example.com"
	GetUsers(baseUrl)
	CreateUser(baseUrl, "Jack")
}

type errorResponse struct {
	Msg string `json:"msg"`
}

// GetUsers send http GET request to http://example.com/api/v1/users
// to get a list of users.
//
// Expect 200 OK and the following response format.
//
// [
//
//	"Jack",
//	"Marry",
//	"Sandy"
//
// ]
//
// If it is not 200 OK, it will return 4xx or 5xx with following message
// format.
//
//	{
//		"msg": "something wrong!"
//	}
func GetUsers(baseUrl string) ([]string, error) {
	url := baseUrl + "/api/v1/users"

	// Send the http request to the server.
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// Reading and parsing the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		// If the request is successful,
		// return the user information.
		var data []string
		err = json.Unmarshal(body, &data)
		if err != nil {
			return nil, err
		}
		return data, err
	} else {
		// If it fails, return the "msg" in the
		// response body.
		var data errorResponse
		err = json.Unmarshal(body, &data)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(data.Msg)
	}
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

type CreateUserResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CreateUser send http POST request to http://example.com/api/v1/user
// to create a user.
//
// Payload format:
//
//	{
//		"name": "Jack"
//	}
//
// Expect 201 Created and the following response format:
//
//	{
//		"id": "ABC-111",
//		"name": "Jack"
//	}
//
// If it is not 201 Created, it will return 4xx or 5xx with following message
// format:
//
//	{
//		"msg": "something wrong!"
//	}
func CreateUser(baseUrl, userName string) (*CreateUserResponse, error) {
	url := baseUrl + "/api/v1/user"

	// Create a payload that should be POSTed to the server.
	payload := CreateUserRequest{
		Name: userName,
	}

	// Encode the payload into json format.
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		return nil, err
	}

	// Create a new http POST request with the payload
	// and modify the Content-Type header.
	req, err := http.NewRequest(http.MethodPost, url, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	// Send the http request to the server.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Reading and parsing the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusCreated {
		// If the request is successful,
		// return the user information.
		var data CreateUserResponse
		err = json.Unmarshal(body, &data)
		if err != nil {
			return nil, err
		}
		return &data, nil
	} else {
		// If it fails, return the "msg" in the
		// response body.
		var data errorResponse
		err = json.Unmarshal(body, &data)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(data.Msg)
	}
}
