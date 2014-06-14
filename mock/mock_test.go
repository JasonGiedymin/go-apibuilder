package mock

import (
  "../api"
  "testing"
)

// type MockApi interface {
//   GetData() string
// }

func TestGetSimpleData(t *testing.T) {
  if baseApi, err := api.NewApi(""); err != nil {
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

func TestListContainers(t *testing.T) {
  if baseApi, err := api.NewApi(""); err != nil {
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
