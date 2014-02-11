package client

import (
  "bytes"
  // "gohessian"
  "log"
  "testing"
)

const (
  H_URL = "http://127.0.0.1:9000/"
)

func Test_http_post(t *testing.T) {
  log.Println("Test_request")
  data := bytes.NewBuffer([]byte{0, 1, 3, 4})
  rb, _ := http_post(H_URL, bytes.NewReader(data.Bytes()))
  log.Println(rb)
  log.Println(string(rb))
}

func Test_request(t *testing.T) {
  Request(H_URL, "method", 100, 200, 101.5, true, false, []byte{1, 2, 3, 5})
}
