package lineheaderwriter

import (
	"bytes"
	"io"
)

type LineHeaderWriter struct {
	writer    io.Writer
	genHeader func() []byte
	midOfLine bool
}

func New(w io.Writer, f func() []byte) *LineHeaderWriter {
	return &LineHeaderWriter{
		writer:    w,
		genHeader: f,
	}
}

func (l *LineHeaderWriter) Write(buf []byte) (int, error) {
	var bb bytes.Buffer
	for _, chr := range buf {
		if !l.midOfLine {
			l.midOfLine = true
			bb.Write(l.genHeader())
		}
		if chr == '\n' {
			l.midOfLine = false
		}
		bb.Write([]byte{chr})
	}
	l.writer.Write(bb.Bytes())
	return len(buf), nil
}
