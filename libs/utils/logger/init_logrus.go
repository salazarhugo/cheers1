package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func InitLogrus() *logrus.Logger {
	log := logrus.New()
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             false,
		TimestampFormat:           "",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	}
	log.Out = os.Stdout

	return log
}
