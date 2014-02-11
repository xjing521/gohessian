package client

import (
  "gohessian"
  "io"
  "io/ioutil"
  "log"
  "net/http"
  "util"
)

//向hessian服务发请求,并将解析结果返回
//url string hessian 服务地址
//method string hessian 公开的方法
//params ...Any 请求参数
func Request(url string, method string, params ...gohessian.Any) (interface{}, error) {
  for k, v := range params {
    log.Println(k, v)
  }

  return nil, nil
}

//http post 请求,返回body字节数组
func http_post(url string, body io.Reader) (rb []byte, err error) {
  var resp *http.Response
  if resp, err = http.Post(url, "application/binary", body); err != nil {
    return nil, err
  }
  defer resp.Body.Close()
  rb, err = ioutil.ReadAll(resp.Body)
  return
}

//返回  hessian 请求头
func pack_head(method string) (b []byte) {
  b = append(b, []byte{99, 0, 1, 109}...)
  tmp_b, _ := util.PackUint16(len(method))
  return
}
