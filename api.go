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

func (self *Api) baseUrl() string {
  return self.host + "v" + strconv.Itoa(self.version) + "/"
}

func (self *Api) versions(gemName string) ([]Version) {
  resp, err := self.client.Get(self.baseUrl() + "versions/" + gemName + ".json")

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
