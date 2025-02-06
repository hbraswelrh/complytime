package log

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/hashicorp/go-hclog"
	"
)

var ErrMissingValue = fmt.Errorf("missing value")

// Need to initialize function

// Wrap the function
func Wrap(log *log.StandardLogOptions) hclog.Logger { return &CharmHclog{log} }

// CharmHclog will be a structure that accesses the attributes of charm log
type CharmHclog struct {
	logger *log.Logger
}

// LoggerOption is an option for a logger (from charm logger)
//type LoggerOption = func(*Logger)

// CharmHclog will implement the hclog.Logger
var _ hclog.Logger = &CharmHclog{}

var hclogCharmLevels = map[hclog.Level]log.Level{
	hclog.NoLevel: log.InfoLevel,  // There is no "NoLevel" equivalent in charm, use info
	hclog.Trace:   log.DebugLevel, // There is no "Trace" equivalent in charm, use debug
	hclog.Debug:   log.DebugLevel,
	hclog.Info:    log.InfoLevel,
	hclog.Warn:    log.WarnLevel,
	hclog.Error:   log.ErrorLevel,
	hclog.Off:     log.FatalLevel, // There is no "Off" level equivalent in charm
}

var charmHclogLevels = map[log.Level]hclog.Level{
	log.DebugLevel: hclog.Debug,
	log.InfoLevel:  hclog.Info,
	log.WarnLevel:  hclog.Warn,
	log.ErrorLevel: hclog.Error,
	log.FatalLevel: hclog.Error, // There is no "fatal" equivalent in hclog
}

func (c *CharmHclog) Log(level hclog.Level, msg string, args ...interface{}) {
	c.logger.Logf(hclogCharmLevels[level], fmt.Sprintf(msg, args...))
}
func (c *CharmHclog) Trace(msg string, args ...interface{}) {
	c.logger.Debugf(msg, args...)
}
func (c *CharmHclog) Debug(msg string, args ...interface{}) {
	c.logger.Debugf(msg, args...)
}
func (c *CharmHclog) Info(msg string, args ...interface{}) {
	c.logger.Infof(msg, args...)
}
func (c *CharmHclog) Warn(msg string, args ...interface{}) {
	c.logger.Warnf(msg, args...)
}
func (c *CharmHclog) Error(msg string, args ...interface{}) {
	c.logger.Errorf(msg, args...)
}

// Functions from go-hc-lo
// Functions from go-hc-log

func (c *CharmHclog) IsTrace() bool     { return false }
func (c *CharmHclog) IsDebug() bool     { return false }
func (c *CharmHclog) IsInfo() bool      { return false }
func (c *CharmHclog) IsWarn() bool      { return false }
func (c *CharmHclog) IsError() bool     { return false }
func (c *CharmHclog) ImpliedArgs() bool { return false }
func (c *CharmHclog) With(args ...interface{}) hclog.Logger {
	return &CharmHclog{c.logger.With(args...)}
}

// Need to configure a Name function
func (c *CharmHclog) Name() string { return c.Name() }

// Take input and then prepend name string
func (c *CharmHclog) Named(name string) hclog.Logger {
	return &CharmHclog{c.logger.With()}
}

// go-hclog logger resetnamed function to implement

//func (c *CharmHclog) ResetNamed(name string) hclog.Logger {
//	logger, err := log.NewWithOptions()
//	if err != nil { panic(err) }
//	return &CharmHclog{logger.WithAttrs(name)}}
//	return c.ResetNamed(name) }

// Enables setting log level
func (c *CharmHclog) SetLevel(level hclog.Level) {
	log.SetLevel(hclogCharmLevels[level])
}

func (c *CharmHclog) GetLevel() hclog.Level { return charmHclogLevels[log.Level(hclog.Level(level))] }

// Look at stdlog.go for forcing level
// The standard logger needs to be implemented to return a standard logger of the charm type
func (c *CharmHclog) StandardLogger(opts *hclog.StandardLoggerOptions) *log.Logger {
	return log.NewWithOptions(c.logger.StandardLog())
}

func (c *CharmHclog) StandardWriter(opts *hclog.StandardLoggerOptions) io.Writer { return os.Stdout }

func Logger() *log.Logger { return log.StandardLog() }
