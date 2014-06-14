package mock

import (
  "../api"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
)

type ReturnData struct {
  data []string
}

type MockApiv10 struct {
  api *api.Api
}

func (mock *MockApiv10) GetSimpleData() string {
  return "200/OK"
}

func (mock *MockApiv10) ListContainers() (ReturnData, error) {
  resp, _ := http.Get(mock.api.Url())
  defer resp.Body.Close()

  if body, err := ioutil.ReadAll(resp.Body); err != nil {
    fmt.Println("Could not contact endpoint")
    return ReturnData{}, err
  } else {
    switch resp.StatusCode {
    case 200:
      var jsonResult ReturnData
      if err := json.Unmarshal(body, &jsonResult); err != nil {
        fmt.Println("Error", err)
        return ReturnData{}, err
      } else {
        // Logger.Trace("Containers found for", imageName, ids)
        return jsonResult, nil
      }
    case 404:
      fmt.Println("Cannot find containers")
    case 500:
      fmt.Println("Error while trying to communicate to endpoint")
    }
  }

  return ReturnData{}, nil
}
