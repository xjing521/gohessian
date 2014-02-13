package gohessian

import (
  "bytes"
  "log"
  "reflect"
  "testing"
  "time"
  "runtime"
)
//GOPATH=/Users/weidewang/ownCloud/Workspace/go/gohessian go test -run=Test_*

var data *bytes.Buffer
var REPLY []byte = []byte{114, 1, 0}

func init() {
  _, filename, _, _ := runtime.Caller(1)
  log.SetPrefix(filename+"\n")
}

func Test_parse_binary(t *testing.T) {
  b := []byte{114, 1, 0, 66, 4, 0, 72, 31, 158, 110, 182, 222, 60, 211, 253, 178, 168, 141, 216, 138,
    38, 13, 108, 46, 61, 6, 3, 71, 155, 226, 59, 169, 181, 228, 144, 115, 89, 117, 33, 72, 200, 134,
    72, 52, 241, 151, 148, 184, 184, 203, 58, 202, 233, 190, 106, 158, 84, 212, 148, 170, 65, 128,
    159, 250, 0, 124, 130, 78, 104, 103, 168, 10, 217, 3, 112, 107, 63, 185, 237, 241, 12, 125, 29,
    72, 25, 217, 99, 59, 176, 104, 87, 138, 136, 207, 106, 247, 56, 62, 100, 242, 176, 32, 74, 122,
    233, 194, 100, 84, 162, 12, 9, 154, 28, 79, 38, 141, 50, 52, 120, 78, 209, 48, 77, 127, 159,
    158, 160, 212, 107, 112, 185, 38, 111, 223, 248, 126, 235, 108, 143, 211, 133, 99, 16, 62, 27,
    133, 82, 115, 222, 140, 139, 57, 44, 125, 234, 197, 39, 146, 168, 151, 212, 5, 150, 69, 93,
    34, 168, 97, 214, 80, 16, 180, 121, 252, 202, 93, 187, 225, 112, 124, 2, 222, 195, 75, 65, 44,
    152, 184, 177, 51, 207, 62, 87, 117, 37, 122, 244, 112, 210, 104, 40, 253, 56, 74, 252, 151,
    39, 156, 116, 93, 164, 120, 54, 194, 12, 166, 75, 91, 51, 26, 127, 243, 197, 121, 41, 182, 34,
    79, 6, 155, 243, 132, 79, 168, 16, 230, 6, 254, 185, 127, 47, 192, 230, 124, 216, 166, 148, 47,
    56, 83, 140, 48, 84, 83, 29, 242, 169, 198, 97, 141, 137, 99, 13, 13, 203, 23, 208, 21, 198,
    124, 53, 195, 120, 111, 129, 41, 57, 232, 80, 154, 214, 76, 230, 174, 97, 159, 210, 35, 168,
    190, 70, 106, 180, 107, 37, 139, 24, 30, 105, 184, 39, 204, 246, 53, 158, 4, 148, 64, 154, 163,
    117, 119, 118, 83, 124, 112, 145, 113, 189, 197, 158, 148, 75, 6, 190, 252, 196, 152, 189, 247,
    253, 223, 40, 205, 122, 44, 87, 136, 148, 140, 249, 40, 255, 91, 98, 129, 249, 2, 192, 186, 32,
    194, 162, 107, 126, 59, 225, 255, 119, 203, 0, 200, 97, 132, 13, 207, 227, 128, 218, 131, 200,
    22, 226, 123, 45, 21, 235, 13, 254, 0, 173, 52, 30, 191, 8, 118, 0, 188, 122, 13, 79, 177, 35,
    54, 162, 5, 236, 95, 118, 35, 132, 113, 204, 173, 189, 63, 2, 237, 244, 40, 159, 60, 172, 238,
    64, 55, 3, 55, 135, 128, 92, 253, 98, 221, 47, 180, 125, 126, 54, 49, 31, 59, 108, 121, 15, 188,
    96, 193, 198, 254, 56, 108, 230, 185, 97, 30, 241, 33, 187, 51, 220, 134, 176, 203, 72, 4, 169,
    72, 229, 230, 43, 68, 236, 213, 145, 208, 85, 110, 202, 140, 132, 155, 172, 190, 79, 121, 18,
    125, 3, 44, 165, 242, 94, 224, 4, 247, 158, 40, 182, 204, 128, 6, 91, 101, 142, 41, 123, 131,
    53, 42, 83, 35, 138, 190, 192, 226, 75, 18, 84, 65, 41, 128, 21, 126, 126, 16, 91, 181, 135,
    215, 147, 112, 43, 55, 161, 249, 219, 99, 109, 239, 155, 84, 195, 243, 235, 145, 107, 140, 189,
    29, 37, 52, 122, 31, 84, 125, 65, 187, 84, 127, 38, 117, 97, 0, 114, 247, 78, 66, 78, 11, 158,
    175, 72, 104, 12, 171, 0, 3, 46, 200, 127, 112, 184, 119, 245, 201, 179, 93, 193, 163, 215,
    121, 37, 153, 155, 211, 25, 77, 92, 137, 133, 239, 255, 77, 77, 92, 158, 81, 226, 102, 135,
    176, 222, 37, 125, 85, 247, 118, 109, 243, 172, 250, 194, 55, 49, 123, 213, 137, 150, 247, 214,
    25, 194, 169, 141, 153, 148, 90, 147, 110, 242, 46, 137, 235, 26, 122, 45, 208, 198, 93, 66,
    112, 195, 252, 143, 38, 123, 111, 131, 183, 147, 114, 244, 131, 107, 173, 51, 30, 87, 198, 43,
    151, 82, 173, 124, 46, 232, 116, 238, 197, 10, 142, 38, 140, 93, 225, 247, 185, 56, 115, 232,
    109, 113, 200, 113, 111, 139, 232, 187, 220, 167, 146, 84, 102, 106, 233, 204, 17, 240, 99, 202,
    37, 78, 24, 118, 66, 225, 24, 249, 19, 34, 201, 120, 25, 192, 52, 43, 42, 78, 244, 92, 11, 230,
    232, 49, 213, 99, 132, 166, 211, 78, 37, 16, 27, 132, 198, 118, 148, 79, 219, 213, 119, 246,
    165, 67, 138, 200, 150, 21, 52, 137, 213, 206, 89, 77, 31, 232, 82, 175, 173, 255, 67, 21,
    66, 97, 24, 126, 100, 40, 88, 184, 6, 230, 79, 244, 236, 53, 26, 4, 180, 3, 69, 47, 77, 7, 238,
    103, 255, 28, 199, 127, 39, 250, 248, 199, 216, 120, 91, 118, 152, 165, 56, 193, 77, 86, 239,
    98, 252, 71, 205, 162, 74, 225, 171, 221, 184, 255, 16, 148, 92, 235, 223, 75, 91, 209, 22, 222,
    87, 245, 217, 135, 75, 151, 179, 37, 182, 46, 130, 9, 138, 232, 247, 4, 244, 77, 114, 79, 131,
    38, 77, 135, 250, 66, 169, 78, 126, 198, 106, 71, 8, 175, 134, 49, 162, 124, 13, 48, 103, 58,
    207, 182, 143, 54, 112, 65, 215, 193, 98, 232, 36, 94, 146, 200, 195, 238, 112, 8, 198, 188, 84,
    226, 103, 101, 134, 217, 132, 201, 236, 68, 200, 170, 136, 186, 71, 220, 142, 164, 43, 166, 234,
    47, 180, 171, 205, 51, 80, 187, 103, 186, 172, 119, 255, 3, 84, 205, 57, 70, 184, 29, 138, 88,
    239, 143, 233, 107, 57, 157, 76, 44, 139, 11, 213, 76, 6, 247, 68, 120, 187, 119, 69, 237, 69,
    173, 179, 244, 3, 74, 250, 37, 134, 22, 42, 116, 95, 128, 71, 161, 85, 233, 153, 11, 21, 232,
    5, 143, 95, 126, 241, 153, 203, 202, 114, 236, 162, 94, 166, 113, 211, 218, 5, 135, 245,
    124, 227, 54, 23, 10, 17, 58, 35, 7, 166, 40, 96, 7, 177, 234, 179, 223, 5, 110, 103, 22,
    225, 69, 71, 40, 127, 42, 221, 7, 94, 94, 147, 210, 138, 80, 239, 198, 106}
  data = bytes.NewBuffer(b)
  h := NewHessian(bytes.NewReader(data.Bytes()))
  v, err := h.Parse()

  var binary_check []byte
  if reflect.TypeOf(v) != reflect.TypeOf(binary_check) {
    t.Fatalf("want []byte type,but got %v", reflect.TypeOf(v))
  }
  if err != nil {
    t.Fatalf("error: %v", err)
  }
}

func Test_parse_null(t *testing.T) {
  data = bytes.NewBuffer(append(REPLY, 'N'))
  h := NewHessian(bytes.NewReader(data.Bytes()))
  v, err := h.Parse()
  if err != nil {
    t.Fatalf("error: %v", err)
  }
  if v != nil {
    t.Fatalf("wnat nil,but got %v", v)
  }
}

func Test_parse_bool_true(t *testing.T) {
  data = bytes.NewBuffer(append(REPLY, 'T'))
  h := NewHessian(bytes.NewReader(data.Bytes()))
  v, err := h.Parse()
  if err != nil {
    t.Fatalf("error: %v", err)
  }

  if reflect.TypeOf(v) != reflect.TypeOf(true) {
    t.Fatalf("want bool type,but got %v", reflect.TypeOf(v))
  }

  if v != true {
    t.Fatalf("wnat true,but got %v", v)
  }
}

func Test_parse_bool_false(t *testing.T) {
  data = bytes.NewBuffer(append(REPLY, 'F'))
  h := NewHessian(bytes.NewReader(data.Bytes()))
  v, err := h.Parse()
  if err != nil {
    t.Fatalf("error: %v", err)
  }
  if reflect.TypeOf(v) != reflect.TypeOf(true) {
    t.Fatalf("want bool type,but got %v", reflect.TypeOf(v))
  }
  if v != false {
    t.Fatalf("wnat false,but got %v", v)
  }
}

func Test_parse_int32_89(t *testing.T) {
  data = bytes.NewBuffer(append(REPLY, []byte{'I', 0, 0, 0, 89}...))
  h := NewHessian(bytes.NewReader(data.Bytes()))
  v, err := h.Parse()
  if err != nil {
    t.Fatalf("error: %v", err)
  }
  if reflect.TypeOf(v) != reflect.TypeOf(int32(100)) {
    t.Fatalf("want int32 type,but got %v", reflect.TypeOf(v))
  }
  if v != int32(89) {
    t.Fatalf("wnat 89,but got %v", v)
  }
}

func Test_parse_long_19890604(t *testing.T) {
  data = bytes.NewBuffer(append(REPLY, []byte{'L', 0, 0, 0, 0, 1, 47, 129, 172}...))
  h := NewHessian(bytes.NewReader(data.Bytes()))
  v, err := h.Parse()
  if err != nil {
    t.Fatalf("error: %v", err)
  }
  if reflect.TypeOf(v) != reflect.TypeOf(int64(100)) {
    t.Fatalf("want int64 type,but got %v", reflect.TypeOf(v))
  }
  if v != int64(19890604) {
    t.Fatalf("wnat 19890604,but got %v", v)
  }
}

func Test_parse_double_1989_0604(t *testing.T) {
  data = bytes.NewBuffer(append(REPLY, []byte{'D', 64, 159, 20, 61, 217, 127, 98, 183}...))
  h := NewHessian(bytes.NewReader(data.Bytes()))
  v, err := h.Parse()
  if err != nil {
    t.Fatalf("error: %v", err)
  }
  if reflect.TypeOf(v) != reflect.TypeOf(float64(100)) {
    t.Fatalf("want float64 type,but got %v", reflect.TypeOf(v))
  }
  if v != 1989.0604 {
    t.Fatalf("wnat 1989.0604,but got %v", v)
  }
}

func Test_parse_date(t *testing.T) {
  data = bytes.NewBuffer(append(REPLY, []byte{'d', 0, 0, 1, 67, 206, 0, 226, 40}...))
  h := NewHessian(bytes.NewReader(data.Bytes()))
  v, err := h.Parse()
  if err != nil {
    t.Fatalf("error: %v", err)
  }
  if reflect.TypeOf(v) != reflect.TypeOf(time.Now()) {
    t.Fatalf("want time.Time type,but got %v", reflect.TypeOf(v))
  }
  if v != time.Unix(1390730601, 0) {
    t.Fatalf("wnat 2014-01-26 18:03:21 +0800 CST,but got %v", v)
  }
}

func Test_parse_string(t *testing.T) {
  b := []byte{
    83, 0, 34, 95, 95, 66, 69, 71,
    73, 78, 95, 95, 229, 133, 148,
    229, 133, 148, 229, 146, 140,
    229, 176, 143, 229, 167, 168,
    229, 173, 144, 49, 50, 51,
    52, 53, 54, 229, 133, 148,
    229, 133, 148, 231, 170, 129,
    231, 170, 129, 231, 170, 129,
    46, 95, 95, 69, 78, 68, 95, 95}
  data = bytes.NewBuffer(append(REPLY, b...))
  h := NewHessian(bytes.NewReader(data.Bytes()))
  v, err := h.Parse()
  if err != nil {
    t.Fatalf("error: %v", err)
  }
  if reflect.TypeOf(v) != reflect.TypeOf(string("")) {
    t.Fatalf("want string type,but got %v", reflect.TypeOf(v))
  }
  want := "__BEGIN__兔兔和小姨子123456兔兔突突突.__END__"
  if v != want {
    t.Fatalf("wnat %s,but got %v", want, v)
  }
}

func Test_parse_list_without_type(t *testing.T) {
  b := []byte{
    114, 1, 0, 86, 108, 0, 0, 0, 0, 83, 0, 5, 104, 101, 108, 108, 111,
    83, 0, 5, 119, 111, 114, 108, 100, 73, 0, 0, 0, 100, 84, 68, 64,
    195, 136, 0, 0, 0, 0, 0, 68, 64, 81, 128, 3, 70, 220, 93, 100,
    76, 0, 0, 0, 0, 0, 0, 3, 231, 100, 0, 0, 1, 67, 206, 150, 245, 230, 122}
  data = bytes.NewBuffer(b)
  h := NewHessian(bytes.NewReader(data.Bytes()))
  v, err := h.Parse()
  var list_check []Any
  if reflect.TypeOf(v) != reflect.TypeOf(list_check) {
    t.Fatalf("want []interface {} type,but got %v", reflect.TypeOf(v))
  }
  if err != nil {
    t.Fatalf("error: %v", err)
  }
}

func Test_parse_list_with_type(t *testing.T) {
  b := []byte{
    114, 1, 0, 86, 116, 0, 6, 115, 116, 114, 105, 110, 103, 108, 0, 0, 0, 10,
    83, 0, 5, 104, 101, 108, 108, 111, 83, 0, 5, 119, 111, 114, 108,
    100, 73, 0, 0, 0, 100, 84, 68, 64, 195, 136, 0, 0, 0, 0, 0, 68,
    64, 81, 128, 3, 70, 220, 93, 100, 76, 0, 0, 0, 0, 0, 0, 3, 231, 122}
  data = bytes.NewBuffer(b)
  h := NewHessian(bytes.NewReader(data.Bytes()))
  v, err := h.Parse()
  var list_check []Any
  if reflect.TypeOf(v) != reflect.TypeOf(list_check) {
    t.Fatalf("want []interface {} type,but got %v", reflect.TypeOf(v))
  }
  if err != nil {
    t.Fatalf("error: %v", err)
  }
}

func Test_parse_map(t *testing.T) {
  b := []byte{114, 1, 0, 77, 116, 0, 5, 102, 108, 111, 97, 116, 73,
    0, 0, 0, 10, 83, 0, 9, 109, 97, 112, 32, 118, 97, 108,
    117, 101, 83, 0, 3, 107, 101, 121, 84, 122}

  data = bytes.NewBuffer(b)
  h := NewHessian(bytes.NewReader(data.Bytes()))
  v, err := h.Parse()

  var map_check map[Any]Any
  if reflect.TypeOf(v) != reflect.TypeOf(map_check) {
    t.Fatalf("want map[Any]Any type,but got %v", reflect.TypeOf(v))
  }
  if err != nil {
    t.Fatalf("error: %v", err)
  }
}
