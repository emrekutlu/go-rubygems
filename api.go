package rubygems

import(
  "log"
  "encoding/json"
  "io/ioutil"
  "strconv"
  "net/http"
)

type Api struct {
  client  *http.Client
  version int
  host    string
}


func (self *Api) gem(name string) (Gem, error) {
  var gem_err error
  resp, err := self.get("gems/" + name)

  if err != nil {
    gem_err = err
  } else {
    defer resp.Body.Close()
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    gem_err = err
  }

  var gem Gem
  json_err := json.Unmarshal(body, &gem)

  if json_err != nil {
    gem_err = json_err
  }

  return gem, gem_err
}

func (self *Api) versions(name string) []Version {
  resp, err := self.get("versions/" + name)

  if err != nil {
    log.Fatal(err)
  } else {
    defer resp.Body.Close()
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }

  var versions []Version
  json_err := json.Unmarshal(body, &versions)

  if json_err != nil {
    log.Fatal(json_err)
  }

  return versions
}

func (self *Api) baseUrl() string {
  return self.host + "v" + strconv.Itoa(self.version) + "/"
}

func (self *Api) url(s string) string {
  return self.baseUrl() + s + ".json"
}

func (self *Api) get(urlPart string) (*http.Response, error)  {
  return self.client.Get(self.url(urlPart))
}
