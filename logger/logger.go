package logger

import (
	"fmt"

	"github.com/jb-hirad/go-log/model"
	"github.com/jb-hirad/go-log/util"
)

// Logger is a struct that holds information about the logger
type Logger struct {
	Path   string
	Levels []model.Level
}

// Log a message with a certain priority
func (l Logger) Log(priority int, message string) {
	level := l.Levels[priority]
	stamp := util.Stamp(level)
	fmt.Printf("%s %s\n", stamp, message)
}
