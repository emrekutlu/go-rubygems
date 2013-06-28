package rubygems

import "net/http"

var api *Api

func Initialize(client *http.Client) {
  api = &Api{ client: client, version: 1, host: "https://rubygems.org/api/", format: "json" }
}

func NewGem(name string) *Gem {
  return &Gem{ Name: name }
}
