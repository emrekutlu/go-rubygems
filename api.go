package rubygems

import(
  "strconv"
  "net/http"
  "io/ioutil"
)

type Api struct {
  client  *http.Client
  version int
  host    string
  format  string
}



func (self *Api) baseUrl() string {
  return self.host + "v" + strconv.Itoa(self.version) + "/"
}

func (self *Api) url(s string) string {
  return self.baseUrl() + s + "." + self.format
}

func (self *Api) get(urlPart string) (*http.Response, error)  {
  resp, err := self.client.Get(self.url(urlPart))

  if err == nil {
    if resp.StatusCode >= 300 {

      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        return resp, err
      }

      apiError := ApiError{ Response: resp, StatusCode: resp.StatusCode, Message: string(body) }
      return resp, apiError
    }
  }

  return resp, err
}
