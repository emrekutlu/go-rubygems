package rubygems

import(
  "net/http"
  "log"
  "encoding/json"
  "io/ioutil"
  "strconv"
)

type Client struct {
  Version int
  Format  string
}

const host string = "https://rubygems.org/api/"

func NewClient(version int, format string) *Client {
  return &Client{ Version: version, Format: format }
}

func (self *Client) baseUrl() string {
  return host + "v" + strconv.Itoa(self.Version) + "/"
}

func (self *Client) versions(gemName string) ([]Version) {
  resp, err := http.Get(self.baseUrl() + "versions/" + gemName + "." + self.Format)
  defer resp.Body.Close()

  if err != nil {
    log.Fatal(err)
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
  log.Fatal("ioutil")
    log.Fatal(err)
  }

  var versions []Version
  json_err := json.Unmarshal(body, &versions)

  if json_err != nil {
  log.Fatal(string(self.Version))
    log.Fatal(json_err)
  }

  return versions
}
