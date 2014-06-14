package api

import (
  "fmt"
  "io/ioutil"
  "net/http"
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

func (api *Api) Get(route string) (*http.Response, []byte, error) {
  // TODO: reuse the URL.Path
  endpoint := api.Url() + route
  fmt.Println("Trying endpoint", endpoint)

  resp, _ := http.Get(endpoint)
  defer resp.Body.Close()

  if body, err := ioutil.ReadAll(resp.Body); err != nil {
    fmt.Println("Could not contact endpoint")
    return nil, nil, err
  } else {
    fmt.Println("Got response code:", resp.StatusCode)
    fmt.Println("Got response body:", string(body))
    return resp, body, nil
  }
}

// Creates a NewApi with default params or using
// a user specified url. 9/10 times you will want
// to supply the url.
func NewApi(url string) (*Api, error) {
  if url == "" {
    return createDefaultApi()
  }

  return createApiWithUrl(url)
}
