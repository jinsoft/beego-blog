package util

import (
	"strings"
)

type String struct {
	value string
}

func NewString(s string) *String {
	var str String
	str.value = s
	return &str
}

func (str *String) ToString() string {
	return str.value
}

/*
  "123" => 3
*/
func (str *String) Len() int {
	return len(str.value)
}

func (str *String) Trim() *String {
	return NewString(strings.TrimSpace(str.value))
}

/**
  返回字符串的子串
*/
func (str *String) Substr(beginIndex, endIndex int) *String {
	return NewString(str.value[beginIndex:endIndex])
}

/*
  str := NewString("abcdef")
  str.SubstrBegin(2)
  return : "cdef"
*/
func (str *String) SubstrBegin(beginIndex int) *String {
	return str.Substr(beginIndex, str.Len())
}

/*
  str := NewString("abcdef")
  str.SubstrBegin(3)
  return : "abc"
*/
func (str *String) SubstrEnd(endIndex int) *String {
	return str.Substr(0, endIndex)
}
