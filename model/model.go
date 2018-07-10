package model

// Levels is a stuct that holds information of the levels in a logger
type Levels struct {
	Levels []Level `json:"Levels"`
}

// Level is a stuct that holds the priority level and the emoji accompanied by it
type Level struct {
	Emoji    int    `json:"Emoji"`
	Priority int    `json:"Priority"`
	Name     string `json:"Name"`
}

// Config is a struct that you use to configure a new logger. You can supply the following arguments:
// - config: the path to a yaml file that holds different logger priority levels
// - log: the path to a log file if you choose to persist the logs. If file exists, logger will append.
type Config struct {
	Config, Log string
}
