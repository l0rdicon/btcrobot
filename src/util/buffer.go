package util

import (
	"bytes"
	"logger"
	"strconv"
)

// bytes.Buffer
type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func (this *Buffer) Append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			logger.Errorln("***** Not enough memory！******")
		}
	}()
	this.Buffer.WriteString(s)
	return this
}

func (this *Buffer) AppendInt(i int) *Buffer {
	return this.Append(strconv.Itoa(i))
}
