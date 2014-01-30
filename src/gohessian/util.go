package gohessian


import (
    "bytes"
    "encoding/binary"
)



func pack_int8(v int8) (r []byte, err error) {
  buf := new(bytes.Buffer)
  err = binary.Write(buf, binary.BigEndian, v)
  if err != nil {
    return
  }
  r = buf.Bytes()
  return
}

//[10].pack('n').bytes => [0, 10]
func pack_int16(v int16) (r []byte, err error) {
  buf := new(bytes.Buffer)
  err = binary.Write(buf, binary.BigEndian, v)
  if err != nil {
    return
  }
  r = buf.Bytes()
  return
}

//[10].pack('n').bytes => [0, 10]
func pack_uint16(v uint16) (r []byte, err error) {
  buf := new(bytes.Buffer)
  err = binary.Write(buf, binary.BigEndian, v)
  if err != nil {
    return
  }
  r = buf.Bytes()
  return
}

//[10].pack('N').bytes => [0, 0, 0, 10]
func pack_int32(v int32) (r []byte, err error) {
  buf := new(bytes.Buffer)
  err = binary.Write(buf, binary.BigEndian, v)
  if err != nil {
    return
  }
  r = buf.Bytes()
  return
}

//[10].pack('q>').bytes => [0, 0, 0, 0, 0, 0, 0, 10]
func pack_int64(v int64) (r []byte, err error) {
  buf := new(bytes.Buffer)
  err = binary.Write(buf, binary.BigEndian, v)
  if err != nil {
    return
  }
  r = buf.Bytes()
  return
}

//[10].pack('G').bytes => [64, 36, 0, 0, 0, 0, 0, 0]
func pack_float64(v float64) (r []byte, err error) {
  buf := new(bytes.Buffer)
  err = binary.Write(buf, binary.BigEndian, v)
  if err != nil {
    return
  }
  r = buf.Bytes()
  return
}



//(0,2).unpack('n')
func unpack_int16(b []byte) (pi int16, err error) {
    err = binary.Read(bytes.NewReader(b), binary.BigEndian, &pi)
    if err != nil {
        return
    }
    return
}

//unpack('an')
func unpack_aint16(b []byte) (pi int16, err error) {
    err = binary.Read(bytes.NewReader(b[1:3]), binary.BigEndian, &pi)
    if err != nil {
        return
    }
    return
}

//(0,4).unpack('N')
func unpack_int32(b []byte) (pi int32, err error) {
    err = binary.Read(bytes.NewReader(b), binary.BigEndian, &pi)
    if err != nil {
        return
    }
    return
}

//long (0,8).unpack('q>')
func unpack_int64(b []byte) (pi int64, err error) {
    err = binary.Read(bytes.NewReader(b), binary.BigEndian, &pi)
    if err != nil {
        return
    }
    return
}

//Double (0,8).unpack('G)
func unpack_float64(b []byte) (rs float64, err error) {
    err = binary.Read(bytes.NewReader(b), binary.BigEndian, &rs)
    if err != nil {
        return
    }
    return
}




