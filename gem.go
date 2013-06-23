package rubygems

type Gem struct {
  api       *Api
  Name      string
  versions  []Version
}

func NewGem(name string) *Gem {
  api := NewApi(1, "json")
  return &Gem{ Name: name, api: api }
}

func (self *Gem) Versions() []Version {
  self.versions = self.api.versions(self.Name)
  return self.versions
}

func (self *Gem) LastVersion() Version {
  return self.Versions()[0]
}
