package api

import (
  "testing"
)

func TestNewApi(t *testing.T) {
  if baseApi, err := NewApi(""); err != nil {
    t.Error("Could not generate new Api")
  } else {
    url := baseApi.Url()
    expected := "http://localhost/"

    if url != expected {
      t.Error("Default URL [" + url + "] was not correct, expected [" + expected + "].")
    }
  }
}

func TestNewApiWithUrl(t *testing.T) {
  userUrl := "http://localdomain"

  if baseApi, err := NewApi(userUrl); err != nil {
    t.Error("Could not generate new Api")
  } else {
    url := baseApi.Url()

    if url != userUrl {
      t.Error("Default URL [" + url + "] was not correct, expected [" + userUrl + "].")
    }
  }
}
