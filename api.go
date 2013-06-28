package rubygems

import(
  "strconv"
  "net/http"
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
  return self.client.Get(self.url(urlPart))
}
