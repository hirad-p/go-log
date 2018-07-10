package util

import (
	"fmt"
	"html"
	"strconv"
	"time"

	"github.com/jb-hirad/go-log/model"
)

// Emoji returns the string version of emoji given its unicode
func Emoji(unicode int) string {
	return html.UnescapeString("&#" + strconv.Itoa(unicode) + ";")
}

// Stamp returns a custom formatted stamp
func Stamp(l model.Level) string {
	time := Timestamp()
	emoji := Emoji(l.Emoji)
	return fmt.Sprintf("[%s] - %s:", time, emoji)
}

// Timestamp returns a formatted timestamp
func Timestamp() string {
	return time.Now().String()
}
