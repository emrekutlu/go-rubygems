package rubygems

import(
  "log"
  "encoding/json"
  "io/ioutil"
  "strconv"
)

type Api struct {
  Version int
  Format  string
}

const host string = "https://rubygems.org/api/"

func NewApi(version int, format string) *Api {
  return &Api{ Version: version, Format: format }
}

func (self *Api) baseUrl() string {
  return host + "v" + strconv.Itoa(self.Version) + "/"
}

func (self *Api) versions(gemName string) ([]Version) {
  resp, err := Client.Get(self.baseUrl() + "versions/" + gemName + "." + self.Format)

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
