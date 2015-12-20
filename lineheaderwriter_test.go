package lineheaderwriter

import (
	"bytes"
	"testing"
)

func TestLineHeaderWriter(t *testing.T) {
	var bb bytes.Buffer
	l := New(&bb, func() []byte {
		return []byte{'a', 'b', ' '}
	})
	l.Write([]byte{'c'})
	l.Write([]byte{'d'})
	l.Write([]byte{'\n', 'e', 'f'})
	expect := `ab cd
ab ef`
	if bb.String() != expect {
		t.Errorf("something went wrong")
	}
}
