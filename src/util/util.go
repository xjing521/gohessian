package util

import (
  "bytes"
  "encoding/binary"
  "fmt"
  "strings"
)

func PackInt8(v int8) (r []byte, err error) {
  buf := new(bytes.Buffer)
  err = binary.Write(buf, binary.BigEndian, v)
  if err != nil {
    return
  }
  r = buf.Bytes()
  return
}

//[10].pack('n').bytes => [0, 10]
func PackInt16(v int16) (r []byte, err error) {
  buf := new(bytes.Buffer)
  err = binary.Write(buf, binary.BigEndian, v)
  if err != nil {
    return
  }
  r = buf.Bytes()
  return
}

//[10].pack('n').bytes => [0, 10]
func PackUint16(v uint16) (r []byte, err error) {
  buf := new(bytes.Buffer)
  err = binary.Write(buf, binary.BigEndian, v)
  if err != nil {
    return
  }
  r = buf.Bytes()
  return
}

//[10].pack('N').bytes => [0, 0, 0, 10]
func PackInt32(v int32) (r []byte, err error) {
  buf := new(bytes.Buffer)
  err = binary.Write(buf, binary.BigEndian, v)
  if err != nil {
    return
  }
  r = buf.Bytes()
  return
}

//[10].pack('q>').bytes => [0, 0, 0, 0, 0, 0, 0, 10]
func PackInt64(v int64) (r []byte, err error) {
  buf := new(bytes.Buffer)
  err = binary.Write(buf, binary.BigEndian, v)
  if err != nil {
    return
  }
  r = buf.Bytes()
  return
}

//[10].pack('G').bytes => [64, 36, 0, 0, 0, 0, 0, 0]
func PackFloat64(v float64) (r []byte, err error) {
  buf := new(bytes.Buffer)
  err = binary.Write(buf, binary.BigEndian, v)
  if err != nil {
    return
  }
  r = buf.Bytes()
  return
}

//(0,2).unpack('n')
func UnpackInt16(b []byte) (pi int16, err error) {
  err = binary.Read(bytes.NewReader(b), binary.BigEndian, &pi)
  if err != nil {
    return
  }
  return
}



//(0,4).unpack('N')
func UnpackInt32(b []byte) (pi int32, err error) {
  err = binary.Read(bytes.NewReader(b), binary.BigEndian, &pi)
  if err != nil {
    return
  }
  return
}

//long (0,8).unpack('q>')
func UnpackInt64(b []byte) (pi int64, err error) {
  err = binary.Read(bytes.NewReader(b), binary.BigEndian, &pi)
  if err != nil {
    return
  }
  return
}

//Double (0,8).unpack('G)
func UnpackFloat64(b []byte) (rs float64, err error) {
  err = binary.Read(bytes.NewReader(b), binary.BigEndian, &rs)
  if err != nil {
    return
  }
  return
}

//将字节数组格式化成 hex
func SprintHex(b []byte) (rs string) {
  rs = fmt.Sprintf("[]byte{")
  for _, v := range b {
    rs += fmt.Sprintf("0x%02x,", v)
  }
  rs = strings.TrimSpace(rs)
  rs += fmt.Sprintf("}\n")
  return
}


