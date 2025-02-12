package log

import (
	"bytes"
	charmlogger "github.com/charmbracelet/log"
	"github.com/hashicorp/go-hclog"
	"github.com/jhump/protoreflect/dynamic/msgregistry"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
	"time"
)

//var (
//	buf        bytes.Buffer
//	TestLogger = charmlogger.New(&buf)
//)

// charmlogger = tlog.TestLogger

func TestWrapLog(t *testing.T) {
	wrap := CharmHclog{}
	charmlogger.StandardLog(wrap)

}
func TestRaceLog(t *testing.T) {
	t.Parallel()

	w := io.Discard
	l := charmlogger.New(w)
	//charmlogger := tlog.TestLogger
	for i := 0; i < 10; i++ {
		t.Run("race", func(t *testing.T) {
			t.Parallel()
			s := l.StandardLog()
			l.Info("test")
			l.GetLevel()
			l.Print("test")

			s.Print("test")
			s.Writer().Write([]byte("tester\n"))
			s.Output(1, "message alert")

			l.SetOutput(w)
			l.Debug("test")
			l.SetLevel(charmlogger.InfoLevel)
			l.GetPrefix()

			o := l.With("test", "tester\n")
			o.Printf("test %s", "tester\n")
			o.SetTimeFormat(time.Kitchen)
			o.Warn("test")
			o.SetOutput(w)
			o.Error("test")
			o.SetFormatter(charmlogger.LogfmtFormatter)
		})
	}
}

// Declaring the test logger to be a new charmlogger
func TestLogger(t *testing.T) {
	var buf bytes.Buffer
	lc := charmlogger.New(&buf) // declaring buffer as the charmlogger stored in lc
	cases := []struct {
		name     string
		expected string
		msg      string
	}{
		{
			name: "info",
			// expected: charmlogger.InfoLevel.String(),
			expected: "info",
			msg:      "info",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			lc.With(tc.name).Info(tc.msg)
			charmlogger.Print("info")
			assert.Equal(t, tc.expected, buf.String()) // Check if the expected string level == test val
		})
	}
}

// TODO: WIP, general structure
// Decide on buffer vs string usage
// Info Level set initially
// Allocating a default options charmlogger to l
// Setting level using the buffer
func TestNonExistentLevel(t *testing.T) {
	var buf bytes.Buffer
	lo := (*CharmHclog).StandardLogger()
	lo.SetOutput(&buf)

	l := charmlogger.New(&buf)
	l.SetLevel(charmlogger.InfoLevel)

	cases := []struct {
		name     string
		expected string
		level    charmlogger.Level
	}{
		{
			name:     " ",
			expected: "INFO: info\n",
			level:    charmlogger.InfoLevel,
		},
		{
			name:     "incorrect",
			expected: "INFO info\n",
			level:    charmlogger.InfoLevel,
		},
		{
			name:     "fake level",
			expected: "INFO info\n",
			level:    charmlogger.InfoLevel,
		},
		{
			name:     " ",
			expected: "INFO prefix: info\n",
			level:    charmlogger.InfoLevel,
		},
		// Print nonexistent level message
	}
	for _, c := range cases {
		buf.Reset()

		// Test switch statements for nonexistent level entry
		// Terminal output format, not logger formatting exactly
	}
}

// Constructing a string prefix tester
// Formatted to mimic https://github.com/charmbracelet/log/blob/main/logger_test.go
//func TestSPrefix(t *testing.T) {
//	var buf bytes.Buffer
//	cases := []struct {
//		name     string
//		expected string
//		prefix   string
//		msg      string
//	}{
//		{
//			name:     "include prefix",
//			expected: "INFO prefix: info\n",
//			prefix:   "prefix",
//			msg:      "info",
//		},
//	}
//	for _, c := range cases {
//		t.Run(c.name, func(t *testing.T) {
//			buf.Reset()
//			l := New(&buf)
//			l.SetPrefix(c.prefix)
//			l.Info(c.msg)
//			l.With(c.fields...).Info(c.msg, c.kvs...)
//			assert.Equal(t, c.expected, buf.String())
//		})
//	}
//}
