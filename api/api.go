package api

import (
  "net/url"
)

type ApiMethod func() interface{}

type Api struct {
  // scheme  string
  // domain  string
  // port    int
  // route   string
  url     url.URL
  Methods map[string]ApiMethod
}

func (api *Api) Url() string {
  return api.url.String()
}

func createDefaultApi() (*Api, error) {
  return &Api{
    url: url.URL{
      Scheme: "http",
      Host:   "localhost",
      Path:   "/",
    },
  }, nil
}

func createApiWithUrl(userUrl string) (*Api, error) {
  if parsedUrl, err := url.Parse(userUrl); err != nil {
    return nil, err
  } else {
    return &Api{
      url: *parsedUrl,
    }, nil
  }
}

func NewApi(url string) (*Api, error) {
  if url == "" {
    return createDefaultApi()
  }

  return createApiWithUrl(url)
}
