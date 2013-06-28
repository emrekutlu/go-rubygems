package rubygems

import(
  "log"
  "encoding/json"
  "io/ioutil"
)

type Gem struct {
  Name              string
  Downloads         int
  Version           string
  VersionDownloads  int
  Info              string
  Authors           string
  ProjectUri        string `json:"project_uri"`
  GemUri            string `json:"gem_uri"`
  HomepageUri       string `json:"homepage_uri"`
  WikiUri           string `json:"wiki_uri"`
  DocumentationUri  string `json:"documentarion_uri"`
  MailingListUri    string `json:"mailing_list_uri"`
  SourceCodeUri     string `json:"source_code_uri"`
  BugTrackerUri     string `json:"bug_tracker_uri"`
  Dependencies struct {
    Development  []struct {
      Name          string
      Requirements  string
    }
    Runtime  []struct {
      Name          string
      Requirements  string
    }
  }
}

func (self *Gem) Get() (*Gem, error) {
  resp, err := api.get("gems/" + self.Name)

  if err != nil {
    return self, err
  } else {
    defer resp.Body.Close()
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return self, err
  }

  err = json.Unmarshal(body, &self)

  if err != nil {
    return self, err
  }

  return self, nil
}

func (self *Gem) Versions() []Version {
  resp, err := api.get("versions/" + self.Name)

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

func (self *Gem) Owners() ([]Owner, error) {
  var owners []Owner

  resp, err := api.get("gems/" + self.Name + "/owners")

  if err != nil {
    return owners, err
  } else {
    defer resp.Body.Close()
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return owners, err
  }

  json_err := json.Unmarshal(body, &owners)

  if json_err != nil {
    return owners, err
  }

  return owners, nil
}

func (self *Gem) TotalDownloadsForVersion(number string) (*TotalDownloadsForVersion, error) {
  var totalDownloads TotalDownloadsForVersion

  resp, err := api.get("downloads/" + self.Name + "-" + number)

  if err != nil {
    return &totalDownloads, err
  } else {
    defer resp.Body.Close()
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return &totalDownloads, err
  }

  json_err := json.Unmarshal(body, &totalDownloads)

  if json_err != nil {
    return &totalDownloads, err
  }

  return &totalDownloads, nil
}
