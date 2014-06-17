package main

import (
  "fmt"
  // "github.com/JasonGiedymin/go-apibuilder"
  "../"
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestGetSimpleData(t *testing.T) {
  if baseApi, err := apibuilder.NewApi(""); err != nil {
    t.Error("Could not generate new Api")
  } else {
    mockApi := MockApiv10{api: baseApi}
    expected := "200/OK"
    response := mockApi.GetSimpleData()

    if response != expected {
      t.Error("Response: [" + response + "] not as expected: [" + expected + "].")
    }
  }
}

func TestListContainers200(t *testing.T) {
  testToken := MockData{Data: "Hello world"}

  testServer := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
    resp.Header().Set("Content-Type", "application/json")
    resp.Write(testToken.Bytes())
  }))
  defer testServer.Close()

  fmt.Println("Trying URL:", testServer.URL)

  if baseApi, err := apibuilder.NewApi(testServer.URL); err != nil {
    t.Error("Could not generate new Api,", err)
  } else {
    mockApi := MockApiv10{api: baseApi}

    if response := mockApi.ListContainers(); response.Err != nil {
      t.Error("ListContainers() failed to connect...", err)
    } else {
      if response.Data != testToken {
        t.Error("Response: [", response, "] not as expected: [", testToken, "].")
      }
    }
  }
}

func TestListContainers404(t *testing.T) {
  testToken := "404 page not found\n"

  testServer := httptest.NewServer(http.NotFoundHandler())
  defer testServer.Close()

  fmt.Println("Trying URL:", testServer.URL)

  if baseApi, err := apibuilder.NewApi(testServer.URL); err != nil {
    t.Error("Could not generate new Api,", err)
  } else {
    mockApi := MockApiv10{api: baseApi}

    if response := mockApi.ListContainers(); response.Err != nil {
      t.Error("ListContainers() failed to connect...", err)
    } else {
      fmt.Println("Test Response:", response)
      if response.Data != testToken {
        t.Error("Response: [", response, "] not as expected: [", testToken, "].")
      }
    }
  }
}

func TestListContainersNoMethod(t *testing.T) {
  testServer := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
    resp.WriteHeader(http.StatusConflict)
  }))
  defer testServer.Close()

  fmt.Println("Trying URL:", testServer.URL)

  if baseApi, err := apibuilder.NewApi(testServer.URL); err != nil {
    t.Error("Could not generate new Api,", err)
  } else {
    mockApi := MockApiv10{api: baseApi}

    if response := mockApi.ListContainers(); response.Err == nil {
      t.Error("Expected error to be returned when no default method exists for handling.")
    }
  }
}

// func TestListContainersDefaultMethod(t *testing.T) {
//   testToken := MockData{string(http.StatusConflict)}
//
//   testServer := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
//     resp.WriteHeader(http.StatusConflict)
//   }))
//   defer testServer.Close()
//
//   fmt.Println("Trying URL:", testServer.URL)
//
//   if baseApi, err := api.NewApi(testServer.URL); err != nil {
//     t.Error("Could not generate new Api,", err)
//   } else {
//     mockApi := MockApiv10{api: baseApi}
//
//     if response := mockApi.ListContainers(); response.Err != nil {
//       t.Error("Api client error found while trying to ListContainers() ...", err)
//     } else {
//       fmt.Println("Test Response:", response)
//       if response.Data != testToken {
//         t.Error("Response: [", response, "] not as expected: [", testToken, "].")
//       }
//     }
//   }
// }
