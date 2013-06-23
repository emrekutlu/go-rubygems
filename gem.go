package rubygems

type Gem struct {
  Name      string
  versions  []Version
}

func NewGem(name string) *Gem {
  return &Gem{ Name: name }
}

func (self *Gem) Versions() []Version {
  self.versions = api.versions(self.Name)
  return self.versions
}

func (self *Gem) LastVersion() Version {
  return self.Versions()[0]
}
