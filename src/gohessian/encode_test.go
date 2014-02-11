package gohessian

import (
  "log"
  // "math"
  // "reflect"
  "bytes"
  "runtime"
  "testing"
  "time"
  // "strconv"
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

// func Test_encode_map(t *testing.T) {
//   log.Println("map")
//   var hmap = make(map[Any]Any)
//   hmap["你好"] = "浮渣"
//   hmap[true] = false
//   hmap[1989.0604] = 8964
//   hmap[false] = 19890604
//   b, err := Encode(hmap)
//   log.Println(b,"===============",err)
//   if err != nil || b == nil {
//     t.Fail()
//   }
// }

func Test_001(t *testing.T) {
  log.Println("================")
  // Encode("新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好")
  // Encode("")
  //log.Println(util.PackInt8(int8(13)))
  //Encode("")
  // Encode(time.Now())
  // Encode(true))
  // Encode(false)
  // Encode(int16(10))
  // Encode(true)
  // Encode(int32(10))
  // Encode(int64(10))
  // Encode(float64(10))
  // Encode(time.Now())
  // Encode([]byte{})
  // Encode(nil)
  // Encode([]byte{1, 2, 3, 4, 5})
  // Encode(Binary{1, 2, 3, 4, 5})
  // var large_b []byte
  //   for i := 0; i <= 65535*2; i++ {
  //     large_b = append(large_b,byte(i))
  //   }
  //   //log.Println(large_b)
  //   Encode(large_b)

  //Encode(time.Now())

  // Encode(float64(10))
  // Encode(float64(0))
  // Encode(float64(-10))
  // Encode(float64(-1))
  // Encode(float64(-2))
  // Encode(float64(-128))
  // Encode(float64(127))
  // log.Println(util.PackInt8(-1))
  // Encode(float64(32767))
  // Encode(float64(32767))
  // Encode(float64(32767.0000000000000))
  // Encode(float64(32767.0001))
  // Encode(float64(32767.111111111111111111111111111132))
  //Encode(false)
  //Encode(true)
  //Encode(nil)
  // Encode(Integer(100))

  //  d := float64(0.0)
  //   log.Println(d)
  //   if d == float64(0.0){
  //     log.Println("float eq")
  //   }
  //==================
  // l := 12345678
  //   log.Println(byte(l>>56))
  //   log.Println(byte(l>>48))
  //   log.Println(byte(l>>40))
  //   log.Println(byte(l>>32))
  //   log.Println(byte(l>>24))
  //   log.Println(byte(l>>16))
  //   log.Println(byte(l>>8))
  //   log.Println(byte(l))
  //%w(78 97 188 0 0 0 0 0)
  // log.Println(util.PackInt32(int32(100)))
  // log.Println(byte(100>>24), byte(100>>16), byte(100>>8), byte(100))
  // log.Println(util.PackInt16(10000))
  //   log.Println(util.PackInt16(-32768))
  //   log.Println(util.PackInt16(32767))
  // Encode(int32(-16))
  // Encode(int32(0))
  // Encode(int32(47))

  // Encode(int32(0))
  //   Encode(int32(-2048))
  //   Encode(int32(-256))
  //   Encode(int32(2047))

  // Encode(int32(0))
  // Encode(int32(-262144))
  // Encode(int32(262143))

  // Encode(int64(0))
  //   Encode(int64(-8))
  //   Encode(int64(15))
  // Encode(int64(0))
  //   Encode(int64(-2048))
  //   Encode(int64(-256))
  //   Encode(int64(2047))

  // Encode(int64(0))
  //   Encode(int64(-262144))
  //   Encode(int64(262143))
  //
  // Encode(int64(0))
  // Encode(int64(300))
  // Encode(int64(-300))
  // list := []Any{"你", 99999999, 100, 10.01, true, false, []byte{0, 1, 2, 3, 4, 5}}

  // log.Println(Encode(list))
  //
  // var hmap = make(map[Any]Any)
  // hmap["你好"] = "哈哈哈"
  // hmap[int32(100)] = "嘿嘿"
  // hmap[float64(100.1010)] = int32(101910)
  // hmap[true] = true
  // hmap[false] = true
  // Encode(hmap)
}
