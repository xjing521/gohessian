package gohessian

import (
  "bytes"
  "log"
  "runtime"
  "time"
  "unicode/utf8"
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
  CHUNK_SIZE = 0x8000
)

func init() {
  _, filename, _, _ := runtime.Caller(1)
  log.SetPrefix(filename + "\n")
}

func Encode(v interface{}) (b []byte, err error) {

  switch v.(type) {

  case []byte:
    b, err = encode_binary(v.([]byte))

  case bool:
    b, err = encode_bool(v.(bool))

  case time.Time:
    b, err = encode_time(v.(time.Time))

  case float64:
    b, err = encode_float64(v.(float64))

  case int32:
    b, err = encode_int32(v.(int32))

  case int64:
    b, err = encode_int64(v.(int64))

  case string:
    b, err = encode_string(v.(string))

  case nil:
    b, err = encode_null(v)

  case []Any:
    b, err = encode_list(v.([]Any))

  case map[Any]Any:
    b, err = encode_map(v.(map[Any]Any))

  default:
    panic("unknow type")
  }
  log.Printf(">>>>encode bytes: %v", b)
  return
}

//=====================================
//对各种数据类型的编码
//=====================================

// binary
func encode_binary(v []byte) (b []byte, err error) {
  var (
    tag   byte
    len_b []byte
    len_n int
  )

  if len(v) == 0 {
    len_b, err = pack_uint16(0)
    b = append(b, 'B')
    b = append(b, len_b...)
    return
  }

  r_buf := *bytes.NewBuffer(v)

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

// boolean
func encode_bool(v bool) (b []byte, err error) {
  if v == true {
    b = append(b, 'T')
  } else {
    b = append(b, 'F')
  }
  return
}

// date
func encode_time(v time.Time) (b []byte, err error) {
  b = append(b, 'd')
  tmp_v, err := pack_int64(v.UnixNano() / 1000000)
  b = append(b, tmp_v...)
  return
}

// double
func encode_float64(v float64) (b []byte, err error) {
  var tmp_v []byte
  b = append(b, 'D')
  tmp_v, err = pack_float64(v)
  b = append(b, tmp_v...)
  return
}

// int
func encode_int32(v int32) (b []byte, err error) {
  var tmp_v []byte
  b = append(b, 'I')
  tmp_v, err = pack_int32(v)
  b = append(b, tmp_v...)
  return
}

// long
func encode_int64(v int64) (b []byte, err error) {
  var tmp_v []byte
  b = append(b, 'L')
  tmp_v, err = pack_int64(v)
  b = append(b, tmp_v...)
  return

}

// null
func encode_null(v interface{}) (b []byte, err error) {
  b = append(b, 'N')
  return
}

// string
func encode_string(v string) (b []byte, err error) {
  var (
    len_b []byte
    s_buf = *bytes.NewBufferString(v)
    r_len = utf8.RuneCountInString(v)

    s_chunk = func(_len int) {
      for i := 0; i < _len; i++ {
        if r, s, err := s_buf.ReadRune(); s > 0 && err == nil {
          b = append(b, []byte(string(r))...)
        }
      }
    }
  )

  if v == "" {
    len_b, err = pack_uint16(uint16(r_len))
    b = append(b, 'S')
    b = append(b, len_b...)
    b = append(b, []byte{}...)
    return
  }

  for {
    r_len = utf8.RuneCount(s_buf.Bytes())
    if r_len == 0 {
      break
    }
    if r_len > CHUNK_SIZE {
      len_b, err = pack_uint16(uint16(CHUNK_SIZE))
      b = append(b, 's')
      b = append(b, len_b...)
      s_chunk(CHUNK_SIZE)
    } else {
      len_b, err = pack_uint16(uint16(r_len))
      b = append(b, 'S')
      b = append(b, len_b...)
      s_chunk(r_len)
    }
  }
  return
}

// list
func encode_list(v []Any) (b []byte, err error) {
  list_len := len(v)
  var (
    len_b []byte
    tmp_v []byte
  )

  b = append(b, 'V')
  if list_len > 0 {
    len_b, err = pack_int32(int32(list_len))
    b = append(b, 'l')
    b = append(b, len_b...)
  }
  for _, a := range v {
    tmp_v, err = Encode(a)
    b = append(b, tmp_v...)
  }
  b = append(b, 'z')
  return
}

// map
func encode_map(v map[Any]Any) (b []byte, err error) {
  // map_len := len(v)
  var (
    tmp_k []byte
    tmp_v []byte
  )
  b = append(b, 'M')
  for k, v := range v {
    tmp_k, err = Encode(k)
    tmp_v, err = Encode(v)
    b = append(b, tmp_k...)
    b = append(b, tmp_v...)
  }
  b = append(b, 'z')
  return
}
