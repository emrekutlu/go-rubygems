package rubygems

var api *Client = NewClient(1, "json")

type Gem struct {
  api       *Client
  Name      string
  versions  []Version
}

func NewGem(name string) *Gem {
  return &Gem{ Name: name, api: api }
}

func (self *Gem) Versions() []Version {
  self.versions = self.api.versions(self.Name)
  return self.versions
}

func (self *Gem) LastVersion() Version {
  return self.Versions()[0]
}
