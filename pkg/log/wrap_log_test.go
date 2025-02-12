package log

import (
	"bytes"
	charmlogger "github.com/charmbracelet/log"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
	"time"
)

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
// Allocating a default options charm logger to l
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

// Testing different input types
func TestTypes(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		level    string
		msg      string
	}{
		{
			// Testing debug level names
			name:     "Debug",
			expected: "Debug",
			level:    "Debug",
			msg:      "The ComplyTime command has been executed.",
		},
		{
			// Testing info level names
			name:     "Info",
			expected: "Info",
			level:    "Info",
			msg:      "The ComplyTime command has been executed.",
		},
		{
			// Testing warn level names
			name:     "Warn",
			expected: "Warn",
			level:    "Warn",
			msg:      "The ComplyTime command has been executed.",
		},
		{
			// Testing error level names
			name:     "Error",
			expected: "Error",
			level:    "Error",
			msg:      "The ComplyTime command has been executed.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			charmlogger.New(&buf)
		})
	}

}
