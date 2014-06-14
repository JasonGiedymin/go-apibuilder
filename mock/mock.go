package mock

import (
  "../api"
  "encoding/json"
  "fmt"
  // "io/ioutil"
  // "net/http"
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

func (mock *MockApiv10) ListContainers() (MockData, error) {
  type respFunc func() interface{}
  var testFunc = func() interface{} {
    return nil
  }
  respPattern := map[int]respFunc{
    200: testFunc,
  }

  fmt.Println("Test:", respPattern)

  route := "/containers"
  if resp, body, err := mock.api.Get(route); err != nil {
    fmt.Println("Error getting route", route)
  } else {
    switch resp.StatusCode {
    case 200:
      var jsonResult MockData

      if err := json.Unmarshal(body, &jsonResult); err != nil {
        fmt.Println("Error trying to unmarshal data,", err)
        return MockData{}, err
      } else {
        return jsonResult, nil
      }
    case 404:
      fmt.Println("Cannot find containers")
    case 500:
      fmt.Println("Error while trying to communicate to endpoint")
    }
  }

  return MockData{}, nil
}
