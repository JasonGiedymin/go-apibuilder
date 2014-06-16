package api

import (
  "errors"
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
)

// Type to describe an Api response (a la Tuple)
type Response struct {
  Data interface{} // Data encapsulates actual response
  Err error // The error if any
}

// Type which describes (atm) a basic Api response handler method
type RespFunc func(body []byte) Response

// Type which maps response codes to response handler methods
type ResponseMap map[int]RespFunc

// Struct holding
type ResponseHandler struct {
  responseMap ResponseMap
}

// For convenience
func NewResponseHandler() *ResponseHandler {
  return &ResponseHandler{ResponseMap{}}
}

// Adds func to map
func (handler *ResponseHandler) AddMethod(code int, respFunc RespFunc) {
  if currMethod := handler.responseMap; currMethod[code] == nil {
    currMethod[code] = respFunc
  } else {
    fmt.Println("Method for status code", code, "already added.")
  }
}

// Simple map lookup, if a func exists it is called else error
// Tested in integration form with a mock api
func (handler *ResponseHandler) Handle(body []byte, resp *http.Response, returnData interface{}) Response {
  if respMethod := handler.responseMap[resp.StatusCode]; respMethod == nil {
    msg := "Response code not mapped, no way to handle " +
      "this response code. Api library might be out " +
      "of date. Code: " + string(resp.StatusCode)

    fmt.Println(msg)
    return Response{ nil, errors.New(msg) }
  } else {
    return respMethod(body)
  }
}

type ApiMethod func() interface{}

type Api struct {
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

// func (api *Api) Post(...) (*http.Response, []byte, error) {
// func (api *Api) Put(...) (*http.Response, []byte, error) {

// Creates a NewApi with default params or using
// a user specified url. 9/10 times you will want
// to supply the url.
func NewApi(url string) (*Api, error) {
  if url == "" {
    return createDefaultApi()
  }

  return createApiWithUrl(url)
}
