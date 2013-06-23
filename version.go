package rubygems

import "time"

type Version struct {
  Number          string
  BuiltAt         time.Time `json:"built_at"`
  Summary         string
  DownloadsCount  int `json:"downloads_count"`
  Platform        string
  Description     string
  prerelease      bool
}
