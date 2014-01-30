package gohessian

import (
  "bufio"
)

//hessian 数据结构定义
type Hessian struct {
  reader *bufio.Reader
  refs   []interface{}
}


