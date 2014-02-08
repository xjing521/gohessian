package gohessian

import (
  "bytes"
  "log"
  // "math"
  // "reflect"
  "runtime"
  "strings"
  "time"
  "unicode/utf8"
  // "encoding/binary"
  "strconv"
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

  default:
    panic("unknow type")
  }
  log.Println(">>>>encode bytes: ", b)
  return
}

//=====================================
//对各种数据类型的编码
//=====================================

// 二进制数据
// 二进制数据语法
// binary ::= b b1 b0 <binary-data> binary
//        ::= B b1 b0 <binary-data>
//        ::= [x20-x2f] <binary-data>
//
// 二进制数据被分割成chunk.
//  十六进制数x42('B')标识最后一个chunk
//  十六进制数x62('b')标识 普通的 chunk.
//
// 每个 chunk 有一个 16-bit 的长度值.
//   len = 256*b1 + b0
//
// 压缩格式:短二进制
//   长度小于 15 的二进制数据只需要用单个十六进制数字来表示[x20-x2f].
//   len = code - 0x20
//
// Binary 实例
//  x20 #零长度的二进制数据
//  x23x01x02x03 #长度为3的数据
//  B x10 x00 .... #4k 大小的 final chunk
//  b x04 x00 .... #1k 大小的 non-final chunk
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

// boolean
// Boolean 语法
//   boolean ::= T
//           ::= F
// 用 16 进制'F'来表示 false,用'T'表示 true.
// Boolean 实例
//  T # true
//  F # false
func encode_bool(v bool) (b []byte, err error) {
  if v == true {
    b = append(b, 'T')
  } else {
    b = append(b, 'F')
  }
  return
}

// date 语法
//   date ::=d b7 b6 b5 b4 b3 b2 b1 b0
//   Date采用64-bit来表示距1970 00:00H, UTC以来经过的milliseconds.
// Date实例
//   D x00 x00 xd0 x4b x92 x84 xb8  #2:51:31 May 8, 1998 UTC
func encode_time(v time.Time) (b []byte, err error) {
  b = append(b, 'd')
  tmp_v, err := pack_int64(v.UnixNano() / 1000000)
  b = append(b, tmp_v...)
  return
}

// double 语法
//   double  ::= D b7 b6 b5 b4 b3 b2 b1 b0
//       ::= x67
//       ::= x68
//       ::= x69 b0
//       ::= x6a b1 b0
//       ::= x6b b3 b2 b1 b0
//
// 符合IEEE标准的 64-bit浮点数.
//
// 压缩格式：double表示的 0
//   0.0 用十六进制x67来表示(对应ascii中字符g的ascii值)
//
// 压缩格式：double 表示的 1
//   1.0 用十六进制x68来表示
//
// 压缩格式：单字节double
//   介于-128.0和127.0之间的无小数位的double型可以用两个十六进制来表示
//     (如x3b表示的),也即相当于一个byte值转换成double:
//   value = (double)b0
//
// 压缩格式：short型double
//   介于 -32768.0 和 32767.0 (16位)之间的无小数位的double型可以用
//    三个十六进制数来表示，也即相当于一个short值转换成double:
//   value=(double)(256*b1 + b0)
//
// float型double
//   和32-bit float型等价的double能够用4个十六进制的 float来表示.
//
// Double实例
//   x67        # 0.0
//   x68        # 1.0
//
//   x69 x00      # 0.0
//   x69 x80      # -128.0
//   x69 xff      # 127.0
//
//   x70 x00 x00    # 0.0
//   x70 x80 x00    # -32768.0
//   x70 x7f xff      # 32767.0
//
//   D x40 x28 x80 x00 x00 x00 x00 x00    # 12.25
func encode_float64(v float64) (b []byte, err error) {
  if v == float64(0.0) { // 压缩格式标示 0.0
    b = append(b, 0x67)
    return
  }
  if v == float64(1.0) { // 压缩格式标示 1.0
    b = append(b, 0x68)
    return
  }
  // 判断传入的值是否有小数
  //有小数点返回       true
  //没有小数点否则返回 false
  hr := func(v float64) bool {
    if strings.IndexByte(strconv.FormatFloat(v, 'G', 100, 64), '.') == -1 { //没有小数点
      return false
    }
    return true
  }

  if !hr(v) { //没有小数位,使用压缩格式进行编码
    var bts []byte
    if v >= -128 && v <= 127 { // 0x69 单字节编码,0x00~0x79表示正数，0x80~0xFF表示负数
      bts, err = pack_int8(int8(v))
      b = append(b, 0x69)
    } else if v >= -32768 && v <= 32767 { // 0x70,short double  表示
      bts, err = pack_int16(int16(v))
      b = append(b, 0x70)
    }
    b = append(b, bts...)
    return
  }

  b = append(b, 'D')
  tmp_v, err := pack_float64(v)
  b = append(b, tmp_v...)
  return
}

// int 语法
//   int   ::=  'I' b3 b2 b1 b0
//         ::=  [x80-xbf]
//         ::=  [xc0-xcf]  b0
//         ::=  [xd0-xd7] b1 b0
//
//   32-bit的有符号整型.
//   一个整型由跟随在x49('I')之后的4个大端序(big-endian)排位的十六进制数来表示。
//   value = (b3<<24) + (b2<<16) + (b1<<8)  + b0;
//
// 单字节整型
//   介于-16和47之间的整型可以用单个字节来表示，用十六进制来表示范围为x80到xbf.
//   value = code – 0x90    # 这里是0x90, 如果code=x80，则value = x80 – x90 = -16
//
// 双字节整型
//   介于-2048和2047之间的整型可以用两个字节来表示，头字节的范围从xc0到xcf.
//   value = ((code – 0xc8)<<8) + b0;
//
// 三字节整型
//   介于- 262144和262143之间的整型可以用三个字节来表示，头字节的范围从xd0到xd7.
//
// 整型实例
//   x90      # 0
//   x80      # -16
//   xbf      # 47
//
//   xc8 x00    # 0
//   xc0 x00    # -2048
//   xc7 x00    # -256
//   xcf xff    # 2047
//
//   xd4 x00 x00  # 0
//   xd0 x00 x00  # -262144
//   xd7 xff xff    # 262143
//
//   I x00 x00 x00 x00  # 0
//   I x00 x00 x01 x2c  # 300
func encode_int32(v int32) (b []byte, err error) {
  var tmp_v []byte
  switch {

  case v >= -16 && v <= 47: //单字节整型
    b = append(b, byte(v+0x90))

  case v >= -2048 && v <= 2047: //双字节整型
    tmp_v, err = pack_int16(int16(v))
    tmp_v[0] += 0xc8
    b = append(b, tmp_v...)

  case v >= -262144 && v <= 262143: //三字节整型
    tmp_v, err = pack_int16(int16(v))
    b = append(b, byte(v>>16+0xd4)) // code,右移16位,相当于除以 65536
    b = append(b, tmp_v...)

  default:
    b = append(b, 'I')
    tmp_v, err = pack_int32(v)
    b = append(b, tmp_v...)
  }
  return
}

// long 语法
// long  ::=  L b7 b6 b5 b4 b3 b2 b1 b0
//       ::=  [xd8-xef]
//       ::=  [xf0-xff] b0
//       ::=  [x38-x3f] b1 b0
//       ::=  x77 b3 b2 b1 b0
//
// 一个64-bit的有符号整数. 起头由十六进制x4c('L')标识, 后面为8字节的大端（big-endian）序的整数。
//
// 压缩格式: 单字节long
//   介于-8和15之间的long由单个字节表示，范围为xd8到xef.
//   value = (code – 0xe0)
//
// 压缩格式: 双字节long
//   介于-2048和2047之间的long由两个字节标识, 起头字节的取值范围为xf0到xff.
//   value = ((code – 0xf8)<<8) + b0
//
// 压缩格式: 3字节long
//   介于-262144和262143之间的long由3个字节编码，起头字节的取值范围为x38到x3f.
//   value = ((code – 0x3c)<<16) + (b1<<8) + b0
//
// 压缩格式: 四字节long
//   可以用32-bit的整数来标识的long在这里需要用5个字节作编码,起头字节由x77标识.
//   value = (b3<<24) + (b2<<16) + (b1<<8) + b0
//
// long实例
//   xe0        # 0
//   xd8        # -8
//   xef        # 15
//
//   xf8 x00        # 0
//   xf0 x00        # -2048
//   xf7 x00        # -256
//   xff xff        # 2047
//
//   x3c x00 x00       # 0
//   x38 x00 x00      # -262144
//   x3f xff xff        # 262143
//
//   x77 x00 x00 x00 x00  # 0
//   x77 x00 x00 x01 x2c  # 300
//
//   L x00 x00 x00 x00 x00 x00 x01 x2c  # 300
func encode_int64(v int64) (b []byte, err error) {
  b = append(b, 'L')
  tmp_v, err := pack_int64(v)
  b = append(b, tmp_v...)
  return

}

// null 语法
//   null  ::= N
//   null表示一个空指针。字符 'N' 用来标识null值。
func encode_null(v interface{}) (b []byte, err error) {
  b = append(b, 'N')
  return
}

// string 语法
// string  ::= s b1 b0 <utf8-data> string
//         ::= s b1 b0 <utf8-data>
//         ::= [x00-x1f] <utf8-data>
//
// 16-bit字符，UTF-8编码的字符串。字符串分块(chunk)编码.
// x53('S')起头来标识最后一个chunk
// x73('s')起头标识非最终chunk
// 每个chunk有一个16-bit的长度值字段。
// length用来标识字符串的长度，而非字节数量.
// String chunks may not split surrogate pairs.
//
// 压缩格式: 短字符串
//   长度小于32的字符串可以用单字节长的length来标识，范围为[x00-x1f].
//   value = code
//
// 字符串实例
//   x00        # "", 空字符串
//   x05 hello      # "hello"
//   x01 xc3 x83    # "\u00c3" Ã
//
//   S   x00 x05 hello  # 长字符串格式的"hello"
//   s   x00 x07 hello,  # "hello, world"被分割成两个chunk
//       X05 world
func encode_string(v string) (b []byte, err error) {
  if v == "" {
    b = append(b, 0)
    return
  }
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

  if r_len < 32 {
    len_b, err = pack_int8(int8(r_len))
    b = append(b, len_b...)
    b = append(b, s_buf.Bytes()...)
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
// list  ::= V type? length? value* z
//       ::=V int int value*
//
// 一个有序列表.每个list都包含一个type字符串，长度length和一个值列表，
// 以十六进制x7a(‘z’)作为结尾。
//
// Type可以是任意的UTF-8编码的字符串。
//
// Length指定了list值列表的长度。
//
// list的每个值都被添加到一个引用列表中，
// 这样，所有list中的相同条目都共享同一份引用以节省空间。
//
// Any parser expecting a list must also accept a null or a shared ref.
//
// Type的有效取值在文档中并没有详细指定，这依赖于特定的应用.
// 比如， 在一个由静态类型语言实现的server所暴露的Hessian接口可以
//  使用类型信息来实例化特定的数组类型，
// 反之，在一个由动态类型语言(e.g.: python)实现的server中，将会忽略类型信息。
//
// 压缩格式: repeated list
//   Hessian2.0 制定了一个格式紧凑的list，其中list元素类型type和元素个数length都用整型来编码，
//   其中类型type是对先前定义的原始数据类型的引用。
//
// List实例
//   强类型int数组的序列化: int[] = {0, 1}
// V
//   t x00 x04 [int  # int[] 类型的编码
//   x6e x02    # length = 2
//   x90      # 整数 0
//   x91      # 整数 1
//   z
//
// 匿名变长list = {0, “foobar”}
// V
//   t x00 x04 [int    # int[] 类型编码
//   x6e x02      # length = 2
//   x90        # 整数0
//   x91        # 整数1
//   z
//
// Repeated list类型
// V
//   t x00 x04 [int    # int[]类型编码
//   x63 x02      # length=2
//   x90        # 整数 0
//   x91        # 整数 1
//   z
//
// V
//   x91        # int[]的类型引用  (integer #1)
//   x92        # length = 2
//   x92        # 整数2
//   x93        # 整数3
func encode_list(v []Any) (b []byte, err error) {
  return
}

// map
// map  ::= M type? (value value)* z
//
// 用来表示序列化map和对象.
// Type字段用来表示map类型，type可能为空（例如在length为0的情况下）。
// 如果类型未指定，则由解析器来负责选择类型。对于对象而言,不识别的key将会被忽略.
// 所有map元素也被存入一个引用列表. 在解析map时，可以同时支持空类型和引用类型。
// 类型由服务具体来进行选择。
//
// Map实例
//   Map map = new HashMap();
//   map.put(new Integer(1), “fee”);
//   map.put(new Integer(16), “fie”);
//   map.put(new Integer(256), “foe”);
//
// M
//   x91      # 1
//   x03 fee    # “fee”
//
//   xa0      # 16
//   x03 fie    # “fie”
//
//   xb9 x00    # 256
//   x03 foe    # “foe”
//
// z
//
// 由java对象表示的Map对象:
//
// Public class Car implements Serializable{
//   String color = “aquamarine”
//   String model = “Beetle”;
//   Int mileage = 65536;
// }
//
// M
//   t x00 x13 com.caucho.test.Car  #type
//   x05 color            # color field
//   x0a aquamarine
//
//   x05 model          # model field
//   x06 Beetle
//
//   x07 mileage          #mileage field
//   I x00 x01 x00 x00
func encode_map(v map[Any]Any) (b []byte, err error) {
  return
}
