package mock

import (
  "../api"
  "encoding/json"
  "fmt"
)

type MockData struct {
  Data string `json:"data"`
}

func (rd *MockData) Bytes() []byte {
  bytes, err := json.Marshal(rd)
  if err != nil {
    fmt.Println("Could not marshal data for test server in unit test,", err)
    return nil
  }
  return bytes
}

type MockApiv10 struct {
  api *api.Api
}

func (mock *MockApiv10) GetSimpleData() string {
  return "200/OK"
}

func (mock *MockApiv10) ListContainers() api.Response {

  // Function to handle 200
  var status200 = func(body []byte) api.Response {
    var jsonResult MockData

    if err := json.Unmarshal(body, &jsonResult); err != nil {
      fmt.Println("Error trying to unmarshal data,", err)
      return api.Response{MockData{}, err}
    } else {
      return api.Response{jsonResult, nil}
    }
  }

  var status404 = func(body []byte) api.Response {
    fmt.Println("Container not found")
    return api.Response{string(body), nil}
  }

  var defaultHandler = func(body []byte) api.Response {
    msg := "Api response not expected. Server sent back: " + string(body)
    fmt.Println(msg)
    return api.Response{msg, nil}
  }

  // Mapping status codes to functions
  handler := api.NewResponseHandler()
  handler.AddMethod(200, status200)
  handler.AddMethod(404, status404)
  handler.AddDefault(defaultHandler)

  // Get Route and handle response
  route := "/containers"
  if resp, body, err := mock.api.Get(route); err != nil {
    fmt.Println("Error getting route", route)
  } else {
    return handler.Handle(body, resp, &MockData{})
  }

  return api.Response{MockData{}, nil}
}
