package gohessian

import (
  "bufio"
  "fmt"
  "io"
  "log"
  //  "strings"
  "runtime"
  "time"
  "util"
)

/*
对 Hessian 数据进行解码
*/

type Parser interface {
}

func init() {
  _, filename, _, _ := runtime.Caller(1)
  log.SetPrefix(filename+"\n")
}

func NewHessian(r io.Reader) (h *Hessian) {
  return &Hessian{reader: bufio.NewReader(r)}
}

//读取当前字节,指针不前移
func (h *Hessian) peek_byte() (b byte) {
  b = h.peek(1)[0]
  return
}

//添加引用
func (h *Hessian) append_refs(v interface{}) {
  h.refs = append(h.refs, v)
}

//获取缓冲长度
func (h *Hessian) len() (l int) {
  h.peek(1) //需要先读一下资源才能得到已缓冲的长度
  l = h.reader.Buffered()
  return
}

//读取 Hessian 结构中的一个字节,并后移一个字节
func (h *Hessian) read_byte() (c byte, err error) {
  c, err = h.reader.ReadByte()
  return
}

//读取指定长度的字节,并后移N个字节
func (h *Hessian) next(n int) (b []byte) {
  if n <= 0 {
    return
  }
  if n >= h.reader.Buffered() {
    n = h.reader.Buffered()
  }
  b = make([]byte, n)
  h.reader.Read(b)
  return
}

//读取指定长度字节,指针不后移
func (h *Hessian) peek(n int) (b []byte) {
  b, _ = h.reader.Peek(n)
  return
}

//读取指定长度的 utf8 字符
func (h *Hessian) next_rune(n int) (s []rune) {
  for i := 0; i < n; i++ {
    if r, ri, e := h.reader.ReadRune(); e == nil && ri > 0 {
      s = append(s, r)
    }
  }
  return
}

//读取数据类型描述,用于 list 和 map
func (h *Hessian) read_type() string {
  if h.peek_byte() != 't' {
    return ""
  }
  t_len, _ := util.UnpackAInt16(h.peek(5))
  t_name := h.next_rune(int(3 + t_len))
  return string(t_name)
}

//解析 hessian 数据包
func (h *Hessian) Parse() (v interface{}, err error) {
  t, err := h.read_byte()
  if err == io.EOF {
    return
  }
  switch t {
  case 'r': //reply
    h.next(2)
    return h.Parse()

  case 'f': //fault
    return h.Parse()
    code, _ := h.Parse()
    return h.Parse()
    message, _ := h.Parse()
    panic(fmt.Sprintln(code, message))

  case 'N': //null
    v = nil

  case 'T': //true
    v = true

  case 'F': //false
    v = false

  case 'I': //int
    if v, err = util.UnpackInt32(h.next(4)); err != nil {
      panic(err)
    }

  case 'L': //long
    if v, err = util.UnpackInt64(h.next(8)); err != nil {
      panic(err)
    }

  case 'D': //double
    if v, err = util.UnpackFloat64(h.next(8)); err != nil {
      panic(err)
    }

  case 'd': //date
    var ms int64
    if ms, err = util.UnpackInt64(h.next(8)); err != nil {
      panic(err)
    }
    v = time.Unix(ms/1000, ms%1000*10E5)

  case 'S', 's', 'X', 'x': //string,xml
    var str_chunks []rune
    for { //避免递归读取 Chunks
      len, _ := util.UnpackInt16(h.next(2))
      str_chunks = append(str_chunks, h.next_rune(int(len))...)
      if t == 'S' || t == 'X' {
        break
      }
      if t, err = h.read_byte(); err != nil {
        panic(err)
      }
    }
    v = string(str_chunks)

  case 'B', 'b': //binary
    var bin_chunks []byte //等同于 []uint8,在 反射判断类型的时候，会得到 []uint8
    for {                 //避免递归读取 Chunks
      len, _ := util.UnpackInt16(h.next(2))
      bin_chunks = append(bin_chunks, h.next(int(len))...)
      if t == 'B' {
        break
      }
      if t, err = h.read_byte(); err != nil {
        panic(err)
      }
    }
    v = bin_chunks

  case 'V': //list
    h.read_type() //TODO 类型怎么用?
    var list_chunks []Any
    if h.peek_byte() == 'l' {
      h.next(5)
    }
    for h.peek_byte() != 'z' {
      if _v, _e := h.Parse(); _e == nil {
        list_chunks = append(list_chunks, _v)
      }
    }
    h.read_byte()
    v = list_chunks
    h.append_refs(&list_chunks)

  case 'M': //map
    h.read_type() //TODO 类型怎么用?
    var map_chunks = make(map[Any]Any)
    for h.peek_byte() != 'z' {
      _kv, _ke := h.Parse()
      _vv, _ve := h.Parse()
      if _ke == nil && _ve == nil {
        map_chunks[_kv] = _vv
      }
    }
    h.read_byte()
    v = map_chunks
    h.append_refs(&map_chunks)

  case 'R': //ref
    if ref_idx, err := util.UnpackInt32(h.next(4)); err == nil {
      v = &h.refs[ref_idx]
    }

  default:
    panic(fmt.Sprintf("Invalid type: %v,>>%s<<<", string(t), h.peek(100)))
  } //switch

  return
} // Parse end
