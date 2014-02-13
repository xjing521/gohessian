package gohessian

import (
  "bytes"
  "log"
  "runtime"
  "testing"
  "time"
)

func init() {
  _, filename, _, _ := runtime.Caller(1)
  log.SetPrefix(filename + "\n")
  time.Now() //do nothing
}

var check_result = func(want, got []byte, t *testing.T) {
  if !bytes.Equal(want, got) {
    t.Fatalf("want %v , got %v", want, got)
  }
}

func Test_encode_binary(t *testing.T) {
  b, err := Encode([]byte{})
  if err != nil || b == nil {
    t.Fail()
  }

  want := []byte{0x42, 0x00, 0x00}
  check_result(b, want, t)

  raw := []byte{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 'a', 'b', 'c', 'd'}
  b, err = Encode(raw)
  want = []byte{0x42, 0x00, 0x0e, 0x0a, 0x09, 0x08, 0x07, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01, 0x61, 0x62, 0x63, 0x64}

  if err != nil || b == nil {
    t.Fail()
  }

  check_result(b, want, t)

}

func Test_encode_boolean(t *testing.T) {
  b, err := Encode(true)
  if err != nil || b[0] != 'T' {
    t.Fail()
  }
  want := []byte{0x54}
  check_result(want, b, t)

  b, err = Encode(false)
  if err != nil || b[0] != 'F' {
    t.Fail()
  }

  want = []byte{0x46}
  check_result(want, b, t)
}

func Test_encode_time(t *testing.T) {
  tz, _ := time.Parse("2006-01-02 15:04:05", "2014-02-09 06:15:23")
  b, err := Encode(tz)
  if err != nil || b == nil {
    t.Fail()
  }
  want := []byte{0x64, 0x00, 0x00, 0x01, 0x44, 0x15, 0x49, 0x34, 0x78}
  check_result(want, b, t)
}

func Test_encode_double(t *testing.T) {
  b, err := Encode(1989.0604)
  if err != nil || b == nil {
    t.Fail()
  }
}

func Test_encode_int_long(t *testing.T) {
  b, err := Encode(19890604)
  if err != nil || b == nil {
    t.Fail()
  }
  b, err = Encode(int32(19890604))
  if err != nil || b == nil {
    t.Fail()
  }
  b, err = Encode(int64(19890604))
  if err != nil || b == nil {
    t.Fail()
  }
}

func Test_encode_string(t *testing.T) {
  b, err := Encode("亡命之徒")
  if err != nil || b == nil {
    t.Fail()
  }
}

func Test_encode_nil(t *testing.T) {
  b, err := Encode(nil)
  if err != nil || b == nil {
    t.Fail()
  }
}

func Test_encode_list(t *testing.T) {
  list := []Any{100, 10.001, "不厌其烦", []byte{0, 2, 4, 6, 8, 10}, true, nil, false}

  b, err := Encode(list)
  if err != nil || b == nil {
    t.Fail()
  }
}

func Test_encode_map(t *testing.T) {
  var hmap = make(map[Any]Any)
  hmap["你好"] = "哈哈哈"
  hmap[100] = "嘿嘿"
  hmap[100.1010] = 101910
  hmap[true] = true
  hmap[false] = true
  b, err := Encode(hmap)
  if err != nil || b == nil {
    t.Fail()
  }

}
