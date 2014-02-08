package gohessian

import (
  "log"
  // "math"
  // "reflect"
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

func Test_001(t *testing.T) {
  log.Println("================")
  // Encode("新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好新年好")
  // Encode("")
  //log.Println(pack_int8(int8(13)))
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
  // log.Println(pack_int8(-1))
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
  // log.Println(pack_int32(int32(100)))
  // log.Println(byte(100>>24), byte(100>>16), byte(100>>8), byte(100))
  // log.Println(pack_int16(10000))
  //   log.Println(pack_int16(-32768))
  //   log.Println(pack_int16(32767))
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
  //list := []Any{"你好",int64(99999999), int32(100), float64(10.01), true, false, []byte{0, 1, 2, 3, 4, 5} }

  //Encode(list)

  var hmap = make(map[Any]Any)
  hmap["你好"] = "哈哈哈"
  hmap[int32(100)] = "嘿嘿"
  hmap[float64(100.1010)] = int32(101910)
  hmap[true] = true
  hmap[false] = true
  Encode(hmap)
}
