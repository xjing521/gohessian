package gohessian

import (
  "bufio"
)

// interface{} 的别名
type Any interface{}

//hessian 数据结构定义
type Hessian struct {
  reader *bufio.Reader
  refs   []Any
}

