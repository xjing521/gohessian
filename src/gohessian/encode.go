package gohessian

import (
  "bytes"
  "log"
  "reflect"
  "runtime"
  "time"
)

/*
对 基本数据进行 Hessian 编码
支持:
int8 int16 int32 int64
float64
time.Time
[]byte
[]interface{}
map[interface{}]interface{}
nil
bool
*/

type Encoder struct {
}

const (
  CHUNK_SIZE = 0xffff
)

func init() {
  _, filename, _, _ := runtime.Caller(1)
  log.SetPrefix(filename + "\n")
}

func Encode(v interface{}) (b []byte, err error) {
  //TODO nil 值需要特殊处理,建一个叫 null 的数据类型处理,避免其他值是 nil 的时候造成困扰
  if v == nil {
    b, err = encode_null(v)
    return
  }
  
  t := reflect.ValueOf(v)
  log.Println("detect type", t.Type())
  switch t.Type() {

  case reflect.TypeOf(true):
    b, err = encode_bool(v.(bool))

  case reflect.TypeOf([]byte{0}):
    b, err = encode_binary(v.([]byte))

  case reflect.TypeOf(float64(0)):
    b, err = encode_float64(v.(float64))

  case reflect.TypeOf(int32(0)):
    b, err = encode_int32(v.(int32))

  case reflect.TypeOf(int64(0)):
    b, err = encode_int64(v.(int64))

  case reflect.TypeOf(time.Now()):
    b, err = encode_time(v.(time.Time))
  }
  log.Println(b)
  return
}

//=====================================
//对各种数据类型的编码
//=====================================

func encode_binary(v []byte) (b []byte, err error) {
  v_len := len(v)
  if v_len == 0 {
    b = []byte{0x20}
    return
  }
  if v_len < 15 {
    b = append(b, byte(0x20+v_len))
    b = append(b, v...)
    return
  }
  r_buf := *bytes.NewBuffer(v)
  var (
    tag   byte
    len_b []byte
    len_n int
  )
  for r_buf.Len() > 0 {
    if r_buf.Len() > CHUNK_SIZE {
      tag = 'b'
      len_b, err = pack_uint16(uint16(CHUNK_SIZE))
      len_n = CHUNK_SIZE
    } else {
      tag = 'B'
      len_b, err = pack_uint16(uint16(r_buf.Len()))
      len_n = r_buf.Len()
    }
    b = append(b, tag)
    b = append(b, len_b...)
    b = append(b, r_buf.Next(len_n)...)
  }
  return
}

func encode_bool(v bool) (b []byte, err error) {
  if v == true {
    b = append(b, 'T')
  } else {
    b = append(b, 'F')
  }
  return
}

func encode_time(v time.Time) (b []byte, err error) {
  b = append(b, 'd')
  tmp_v, err := pack_int64(v.UnixNano() / 1000000)
  b = append(b, tmp_v...)
  return
}

func encode_float64(v float64) (b []byte, err error) {
  b = append(b, 'D')
  tmp_v, err := pack_float64(v)
  b = append(b, tmp_v...)
  return
}

func encode_int32(v int32) (b []byte, err error) {
  b = append(b, 'I')
  tmp_v, err := pack_int32(v)
  b = append(b, tmp_v...)
  return
}

func encode_int64(v int64) (b []byte, err error) {
  b = append(b, 'L')
  tmp_v, err := pack_int64(v)
  b = append(b, tmp_v...)
  return

}

func encode_null(v interface{}) (b []byte, err error) {
  b = append(b, 'N')
  return
}

func encode_string(v interface{}) (b []byte, err error) {
  return
}

func encode_list(v interface{}) (b []byte, err error) {
  return
}

func encode_map(v interface{}) (b []byte, err error) {
  return
}
