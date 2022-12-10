package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	t.Run("happy path, we can get users info", func(t *testing.T) {
		// Create a router that routes http requests to specific handlers.
		router := http.NewServeMux()

		// We expect to have the mock http server process /api/v1/users
		// while faking its response as we expect it to look.
		router.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
			/*
				This is where we can test the format of our request!
				Let's say assert that the request was made with the correct header.

				For example, check if the request comes with an authorization header
				and a value of "bearer xxx".

				assert.Equal(t, "bearer xxx", r.Header.Get("Authorization"))
			*/

			// We expect the http method is GET.
			assert.Equal(t, http.MethodGet, r.Method)

			/*
				This is where we mock the response of the the request. That is,
				what we expect the API Server to send back.
			*/

			// return 200 OK and users info.
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[
				"Jack",
				"Marry",
				"Sandy"
			]`))
		})

		// Create an http server and register the router with a
		// predefined mock handler.
		fakeServer := httptest.NewServer(router)

		// We should always close the http server at the end of the test
		// to release related resources.
		defer fakeServer.Close()

		// The URL from the mock http server is in the format of http://127.0.0.1
		// and the port is random.
		baseUrl := fakeServer.URL

		// Calling a function to be tested.
		users, err := GetUsers(baseUrl)

		// Test the results of the function as we expect.
		assert.NoError(t, err)
		assert.Len(t, users, 3)
		assert.Equal(t, []string{"Jack", "Marry", "Sandy"}, users)
	})

	t.Run("unhappy path, API server has some problems", func(t *testing.T) {
		// Create a router that routes http requests to specific handlers.
		router := http.NewServeMux()

		// We expect to have the mock http server process /api/v1/users
		// while faking its response as we expect it to look.
		router.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
			/*
				This is where we can test the format of our request!
				Let's say assert that the request was made with the correct header.

				For example, check if the request comes with an authorization header
				and a value of "bearer xxx".

				assert.Equal(t, "bearer xxx", r.Header.Get("Authorization"))
			*/

			// We expect the http method is GET.
			assert.Equal(t, http.MethodGet, r.Method)

			/*
				This is where we mock the response of the the request. That is,
				what we expect the API server to send back.
			*/

			// return 500 Internal Server Error.
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{
				"msg": "get error"
			}`))
		})

		// Create an http server and register the router with a
		// predefined mock handler.
		fakeServer := httptest.NewServer(router)

		// We should always close the http server at the end of the test
		// to release related resources.
		defer fakeServer.Close()

		// The URL from the mock http server is in the format of http://127.0.0.1
		// and the port is random.
		baseUrl := fakeServer.URL

		// Calling a function to be tested.
		_, err := GetUsers(baseUrl)

		// Test the results of the function as we expect.
		assert.Error(t, err)
		assert.EqualError(t, err, "get error")
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("happy path, API server has some problems", func(t *testing.T) {
		// Create a router that routes http requests to specific handlers.
		router := http.NewServeMux()

		// We expect to have the mock http server process /api/v1/user
		// while faking its response as we expect it to look.
		router.HandleFunc("/api/v1/user", func(w http.ResponseWriter, r *http.Request) {
			/*
				This is where we can test the format of our request!
				Let's say assert that the request was made with the correct header.

				For example, check if the request comes with an authorization header
				and a value of "bearer xxx".

				assert.Equal(t, "bearer xxx", r.Header.Get("Authorization"))
			*/

			// We expect the http method is POST.
			assert.Equal(t, http.MethodPost, r.Method)

			// Check if the Content-Type header is application/json.
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

			// Check the payload format of the request.
			body, err := io.ReadAll(r.Body)
			assert.NoError(t, err)
			assert.JSONEq(t, `{"name": "Jack"}`, string(body))

			/*
				This is where we mock the response of the the request. That is,
				what we expect the API Server to send back.
			*/

			// return 201 Created and user info.
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{
				"id": "id_foo",
				"name": "name_foo"
			}`))
		})

		// Create an http server and register the router with a
		// predefined mock handler.
		fakeServer := httptest.NewServer(router)

		// We should always close the http server at the end of the test
		// to release related resources.
		defer fakeServer.Close()

		// The URL from the mock http server is in the format of http://127.0.0.1
		// and the port is random.
		baseUrl := fakeServer.URL

		// Calling a function to be tested.
		user, err := CreateUser(baseUrl, "Jack")

		// Test the results of the function as we expect.
		assert.NoError(t, err)
		assert.Equal(t, "id_foo", user.ID)
		assert.Equal(t, "name_foo", user.Name)
	})
	t.Run("unhappy path, some error occur", func(t *testing.T) {
		// Create a router that routes http requests to specific handlers.
		router := http.NewServeMux()

		// We expect to have the mock http server process /api/v1/user
		// while faking its response as we expect it to look.
		router.HandleFunc("/api/v1/user", func(w http.ResponseWriter, r *http.Request) {
			/*
				This is where we can test the format of our request!
				Let's say assert that the request was made with the correct header.

				For example, check if the request comes with an authorization header
				and a value of "bearer xxx".

				assert.Equal(t, "bearer xxx", r.Header.Get("Authorization"))
			*/

			// We expect the http method is POST.
			assert.Equal(t, http.MethodPost, r.Method)

			// Check if the Content-Type header is application/json.
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

			// Check the payload format of the request.
			body, err := io.ReadAll(r.Body)
			assert.NoError(t, err)
			assert.JSONEq(t, `{"name": "Jack"}`, string(body))

			/*
				This is where we mock the response of the the request. That is,
				what we expect the API Server to send back.
			*/

			// return 500 Internal Server Error.
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{
				"msg": "get error"
			}`))
		})

		// Create an http server and register the router with a
		// predefined mock handler.
		fakeServer := httptest.NewServer(router)

		// We should always close the http server at the end of the test
		// to release related resources.
		defer fakeServer.Close()

		// The URL from the mock http server is in the format of http://127.0.0.1
		// and the port is random.
		baseUrl := fakeServer.URL

		// Calling a function to be tested.
		_, err := CreateUser(baseUrl, "Jack")

		// Test the results of the function as we expect.
		assert.Error(t, err)
		assert.EqualError(t, err, "get error")
	})
}
