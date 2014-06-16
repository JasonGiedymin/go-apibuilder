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

func (mock *MockApiv10) ListContainers() (interface{}, error) {

  // Function to handle 200
  var status200 = func(body []byte) (interface{}, error) {
    var jsonResult MockData

    if err := json.Unmarshal(body, &jsonResult); err != nil {
      fmt.Println("Error trying to unmarshal data,", err)
      return MockData{}, err
    } else {
      return jsonResult, nil
    }
  }

  // Mapping status codes to functions
  handler := api.NewResponseHandler()
  handler.AddMethod(200, status200)

  // Get Route and handle response
  route := "/containers"
  if resp, body, err := mock.api.Get(route); err != nil {
    fmt.Println("Error getting route", route)
  } else {
    return handler.Handle(body, resp, &MockData{})
  }

  return MockData{}, nil
}
