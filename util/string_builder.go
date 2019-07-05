package util

import "bytes"

type StringBuilder struct {
	buffer bytes.Buffer
}

func NewStringBuilder() *StringBuilder {
	var builder StringBuilder
	return &builder
}

func NewStringBuilderString(str *String) *StringBuilder {
	var builder StringBuilder
	builder.buffer.WriteString(str.ToString())
	return &builder
}

func (builder *StringBuilder) Append(s string) *StringBuilder {
	builder.buffer.WriteString(s)
	return builder
}

func (builder *StringBuilder) ToString() string {
	return builder.buffer.String()
}
