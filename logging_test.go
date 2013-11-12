package logging

import (
	"os"
	"testing"
    "bytes"
)

var levelTests = []struct {
	lvl  Level
	want string
}{
	{STATS, "STATS"},
	{FATAL, "FATAL"},
	{ERROR, "ERROR"},
	{WARN, "WARN"},
	{INFO, "INFO"},
	{DEBUG, "DEBUG"},
	{TRACE, "TRACE"},
}

type mockWriter struct {
    buf bytes.Buffer
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
    return m.buf.Write(p)
}

func TestLevelStrings(t *testing.T) {
	for i, tt := range levelTests {
		s := tt.lvl.String()

		if s != tt.want {
			t.Errorf("[%d] Level strings: got: %q, want %q", i, s, tt.want)
		}
	}
}

func TestSetLevel(t *testing.T) {
	l := &Logger{}

	for i, tt := range levelTests {
		l.SetLevel(tt.lvl)
		lvl := l.Level

		if lvl != tt.lvl {
			t.Errorf("[%d] SetLevel(%s): got: %q, want %q", i, tt.lvl, lvl,
				tt.lvl)
		}
	}
}

func TestSetWriter(t *testing.T) {
	l := &Logger{}
	w := os.Stdout
	l.SetWriter(w)

	if l.Writer != w {
		t.Errorf("SetWriter(os.Stdout): got: %q, want %q (%s)", l.Writer, w,
			w.Name())
	}
}

func TestSetName(t *testing.T) {
	l := &Logger{}
	s := "TestLogger"
	l.SetName(s)

	if l.Name != s {
		t.Errorf("SetName(%q): got: %q, want %q (%s)", s, l.Name, s)
	}
}

func TestLog(t *testing.T) {
    msg := "test message"
    want := []byte("TRACE: test message")
    w := &mockWriter{}
	l := &Logger{Writer: w}
    l.Log(TRACE, msg)

	if bytes.HasSuffix(w.buf.Bytes(), []byte(want)) {
		t.Errorf("Log(%q): got: %q, want %q (%s)", msg, w.buf.String(), want)
	}
}
