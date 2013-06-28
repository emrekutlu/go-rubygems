package rubygems

import(
  "net/http"
  "fmt"
)

type ApiError struct {
  Response    *http.Response
  StatusCode  int
  Message     string
}

func (self ApiError) Error() string {
  return fmt.Sprintf("%v", self.Message)
}
