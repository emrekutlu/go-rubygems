package rubygems

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

func (self *Gem) Versions() []Version {
  return api.versions(self.Name)
}
