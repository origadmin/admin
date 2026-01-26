package pubsub

import (
	"github.com/ThreeDotsLabs/watermill"

	"github.com/origadmin/runtime/log"
)

// watermillLoggerAdapter bridges watermill's logger to the application's runtime logger.
type watermillLoggerAdapter struct {
	logger log.Logger
	helper *log.Helper
}

// NewWatermillLogger creates a new watermill.LoggerAdapter that wraps the application's logger.
func NewWatermillLogger(logger log.Logger) watermill.LoggerAdapter {
	l := log.With(logger, "component", "watermill")
	return &watermillLoggerAdapter{
		logger: l,
		helper: log.NewHelper(l),
	}
}

func (w *watermillLoggerAdapter) fieldsToArgs(fields watermill.LogFields) []interface{} {
	args := make([]interface{}, 0, len(fields)*2)
	for k, v := range fields {
		args = append(args, k, v)
	}
	return args
}

func (w *watermillLoggerAdapter) Error(msg string, err error, fields watermill.LogFields) {
	args := w.fieldsToArgs(fields)
	args = append(args, "msg", msg)
	args = append(args, "err", err)
	w.helper.Errorw(args...)
}

func (w *watermillLoggerAdapter) Info(msg string, fields watermill.LogFields) {
	args := w.fieldsToArgs(fields)
	args = append(args, "msg", msg)
	w.helper.Infow(args...)
}

func (w *watermillLoggerAdapter) Debug(msg string, fields watermill.LogFields) {
	args := w.fieldsToArgs(fields)
	args = append(args, "msg", msg)
	w.helper.Debugw(args...)
}

func (w *watermillLoggerAdapter) Trace(msg string, fields watermill.LogFields) {
	// runtime/log doesn't have Trace, so we map it to Debug.
	args := w.fieldsToArgs(fields)
	args = append(args, "msg", msg)
	w.helper.Debugw(args...)
}

// With returns a new logger with the given fields.
func (w *watermillLoggerAdapter) With(fields watermill.LogFields) watermill.LoggerAdapter {
	// Create a new logger with the new fields using the top-level log.With function.
	newLogger := log.With(w.logger, w.fieldsToArgs(fields)...)
	return &watermillLoggerAdapter{
		logger: newLogger,
		helper: log.NewHelper(newLogger),
	}
}
