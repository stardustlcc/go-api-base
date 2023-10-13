package stringutil

import "strconv"

type StrTo string

// 转成字符串类型
func (s StrTo) String() string {
	return string(s)
}

// 字符串转int
func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

// 字符串转int 忽略错误
func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

// 字符串转uint32
func (s StrTo) Uint32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

// 字符串转uint32 忽略错误
func (s StrTo) MustUint32() uint32 {
	v, _ := s.Uint32()
	return v
}
