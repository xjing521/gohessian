package gohessian

import (
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

func Test_001(t *testing.T) {
  log.Println("================")
  // Encode(int16(10))
  //   Encode(int32(10))
  //   Encode(int64(10))
  //   Encode(float64(10))
  //Encode(time.Now())
  //Encode([]byte{})
  //Encode([]byte{1, 2, 3, 4, 5})
  // var large_b []byte
  //   for i := 0; i <= 65535*2; i++ {
  //     large_b = append(large_b,byte(i))
  //   }
  //   //log.Println(large_b)
  //   Encode(large_b)

  //Encode(time.Now())

  //Encode(float64(10))
  //Encode(false)
  //Encode(true)
  //Encode(nil)

  //  d := float64(0.0)
  //   log.Println(d)
  //   if d == float64(0.0){
  //     log.Println("float eq")
  //   }
  //==================
  

}
