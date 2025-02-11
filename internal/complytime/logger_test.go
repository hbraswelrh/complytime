package complytime

import (
	"bytes"
	charmlogger "github.com/charmbracelet/log"
	"github.com/complytime/complytime/pkg/log"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
	//"strings"
	"testing"
)

func NewTestLogger(t *testing.T) log.CharmHclog {
	log.WrapLog(charmlogger.SetLevel())
}
func TestLogger(t *testing.T) log.CharmHclog {
	logger := NewTestLogger(t)
	logger.Debug("debug")
	logger.Warn("warn")
	logger.Error("error")
	logger.Info("info")
	logger.SetLevel(hclog.Info)
}

// TODO: WIP, general structure
// Decide on buffer vs string usage
// Info Level set initially
func TestNonExistentLevel(t *testing.T) {
	var buf bytes.Buffer
	l := charmlogger.New(&buf)
	l.SetLevel(charmlogger.InfoLevel)

	cases := []struct {
		name     string
		expected string
		level    charmlogger.Level
	}{
		{
			name:     " ",
			expected: "INFO prefix: info\n",
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
func TestStringPrefix(t *testing.T) {
	cases := []struct {
		name     string
		expected string
		prefix   string
		level    charmlogger.Level
	}{
		{
			name:     "named-prefix-info",
			expected: "INFO info\n",
			prefix:   "INFO prefix: info\n",
			level:    charmlogger.InfoLevel,
		},
		{
			name:     "named-prefix--debug",
			expected: "DEBUG debug\n",
			prefix:   "DEBUG prefix: debug\n",
			level:    charmlogger.DebugLevel,
		},
		{
			name:     "named-prefix-fatal",
			expected: "Fatal fatal\n",
			prefix:   "FATAL prefix: fatal\n",
			level:    charmlogger.FatalLevel,
		},
		{
			name:     "named-prefix-warn",
			expected: "WARN warn\n",
			prefix:   "WARN prefix: warn\n",
			level:    charmlogger.WarnLevel,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var buf bytes.Buffer
			l := charmlogger.New(&buf)
			l.SetPrefix(c.level.String())
			assert.Equal(t, c.expected, buf.String())
		})
	}
}
