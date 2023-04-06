package code

import (
	"fmt"
	"io"
)

var (
	Nil = NewCode(0, "", nil)
)

type Code interface {
	Code() int       // 错误码编号
	Message() string // 错误码消息
	Detail() any     // 错误码详情
	String() string  // 格式化错误码
}

type defaultCode struct {
	code    int
	message string
	detail  any
}

func NewCode(code int, message string, detail any) Code {
	return &defaultCode{
		code:    code,
		message: message,
		detail:  detail,
	}
}

// Code 错误码编号
func (c *defaultCode) Code() int {
	return c.code
}

// Message 错误码消息
func (c *defaultCode) Message() string {
	return c.message
}

// Detail 错误码详情
func (c *defaultCode) Detail() any {
	return c.detail
}

// String 格式化错误码
func (c *defaultCode) String() string {
	if c.message != "" {
		if c.detail != nil {
			return fmt.Sprintf("%d:%s:%v", c.code, c.message, c.detail)
		}
		return fmt.Sprintf("%d:%s", c.code, c.message)
	}
	return fmt.Sprintf("%d", c.code)
}

// Format 格式化输出
func (c *defaultCode) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		if c.message != "" {
			io.WriteString(s, fmt.Sprintf("%d:%s", c.code, c.message))
		} else {
			io.WriteString(s, fmt.Sprintf("%d", c.code))
		}
	case 'v':
		io.WriteString(s, c.String())
	}
}
