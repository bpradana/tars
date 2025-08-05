package logger

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

// Logger wraps logrus.Logger to provide a simpler interface
type Logger struct {
	*logrus.Logger
}

// New creates a new logger instance
func New() *Logger {
	return &Logger{
		Logger: logrus.New(),
	}
}

// SetLevel sets the log level
func (l *Logger) SetLevel(level logrus.Level) {
	l.Logger.SetLevel(level)
}

// SetFormatter sets the log formatter
func (l *Logger) SetFormatter(formatter logrus.Formatter) {
	l.Logger.SetFormatter(formatter)
}

// Error logs an error message with optional fields
func (l *Logger) Error(err error, args ...interface{}) {
	if len(args) == 0 {
		l.Logger.WithError(err).Error(err.Error())
		return
	}

	// Check if the last argument is a slice of fields
	if fields, ok := args[len(args)-1].(Fields); ok {
		// Remove the fields from args
		messageArgs := args[:len(args)-1]
		message := fmt.Sprint(messageArgs...)

		logrusFields := fieldsToLogrus(fields)
		l.Logger.WithError(err).WithFields(logrusFields).Error(message)
		return
	}

	// No fields, just message
	message := fmt.Sprint(args...)
	l.Logger.WithError(err).Error(message)
}

// ErrorWithContext logs an error message with context and optional fields
func (l *Logger) ErrorWithContext(ctx context.Context, err error, args ...interface{}) {
	contextFields := contextToLogrus(ctx)

	if len(args) == 0 {
		l.Logger.WithError(err).WithFields(contextFields).Error(err.Error())
		return
	}

	// Check if the last argument is a slice of fields
	if fields, ok := args[len(args)-1].(Fields); ok {
		// Remove the fields from args
		messageArgs := args[:len(args)-1]
		message := fmt.Sprint(messageArgs...)

		logrusFields := fieldsToLogrus(fields)
		// Merge context fields with provided fields
		for k, v := range contextFields {
			logrusFields[k] = v
		}

		l.Logger.WithError(err).WithFields(logrusFields).Error(message)
		return
	}

	// No fields, just message
	message := fmt.Sprint(args...)
	l.Logger.WithError(err).WithFields(contextFields).Error(message)
}

// Warn logs a warning message with optional fields
func (l *Logger) Warn(args ...interface{}) {
	if len(args) == 0 {
		l.Logger.Warn("warning")
		return
	}

	// Check if the last argument is a slice of fields
	if fields, ok := args[len(args)-1].(Fields); ok {
		// Remove the fields from args
		messageArgs := args[:len(args)-1]
		message := fmt.Sprint(messageArgs...)

		logrusFields := fieldsToLogrus(fields)
		l.Logger.WithFields(logrusFields).Warn(message)
		return
	}

	// No fields, just message
	message := fmt.Sprint(args...)
	l.Logger.Warn(message)
}

// WarnWithContext logs a warning message with context and optional fields
func (l *Logger) WarnWithContext(ctx context.Context, args ...interface{}) {
	contextFields := contextToLogrus(ctx)

	if len(args) == 0 {
		l.Logger.WithFields(contextFields).Warn("warning")
		return
	}

	// Check if the last argument is a slice of fields
	if fields, ok := args[len(args)-1].(Fields); ok {
		// Remove the fields from args
		messageArgs := args[:len(args)-1]
		message := fmt.Sprint(messageArgs...)

		logrusFields := fieldsToLogrus(fields)
		// Merge context fields with provided fields
		for k, v := range contextFields {
			logrusFields[k] = v
		}

		l.Logger.WithFields(logrusFields).Warn(message)
		return
	}

	// No fields, just message
	message := fmt.Sprint(args...)
	l.Logger.WithFields(contextFields).Warn(message)
}

// Info logs an info message with optional fields
func (l *Logger) Info(args ...interface{}) {
	if len(args) == 0 {
		l.Logger.Info("info")
		return
	}

	// Check if the last argument is a slice of fields
	if fields, ok := args[len(args)-1].(Fields); ok {
		// Remove the fields from args
		messageArgs := args[:len(args)-1]
		message := fmt.Sprint(messageArgs...)

		logrusFields := fieldsToLogrus(fields)
		l.Logger.WithFields(logrusFields).Info(message)
		return
	}

	// No fields, just message
	message := fmt.Sprint(args...)
	l.Logger.Info(message)
}

// InfoWithContext logs an info message with context and optional fields
func (l *Logger) InfoWithContext(ctx context.Context, args ...interface{}) {
	contextFields := contextToLogrus(ctx)

	if len(args) == 0 {
		l.Logger.WithFields(contextFields).Info("info")
		return
	}

	// Check if the last argument is a slice of fields
	if fields, ok := args[len(args)-1].(Fields); ok {
		// Remove the fields from args
		messageArgs := args[:len(args)-1]
		message := fmt.Sprint(messageArgs...)

		logrusFields := fieldsToLogrus(fields)
		// Merge context fields with provided fields
		for k, v := range contextFields {
			logrusFields[k] = v
		}

		l.Logger.WithFields(logrusFields).Info(message)
		return
	}

	// No fields, just message
	message := fmt.Sprint(args...)
	l.Logger.WithFields(contextFields).Info(message)
}

// Debug logs a debug message with optional fields
func (l *Logger) Debug(args ...interface{}) {
	if len(args) == 0 {
		l.Logger.Debug("debug")
		return
	}

	// Check if the last argument is a slice of fields
	if fields, ok := args[len(args)-1].(Fields); ok {
		// Remove the fields from args
		messageArgs := args[:len(args)-1]
		message := fmt.Sprint(messageArgs...)

		logrusFields := fieldsToLogrus(fields)
		l.Logger.WithFields(logrusFields).Debug(message)
		return
	}

	// No fields, just message
	message := fmt.Sprint(args...)
	l.Logger.Debug(message)
}

// DebugWithContext logs a debug message with context and optional fields
func (l *Logger) DebugWithContext(ctx context.Context, args ...interface{}) {
	contextFields := contextToLogrus(ctx)

	if len(args) == 0 {
		l.Logger.WithFields(contextFields).Debug("debug")
		return
	}

	// Check if the last argument is a slice of fields
	if fields, ok := args[len(args)-1].(Fields); ok {
		// Remove the fields from args
		messageArgs := args[:len(args)-1]
		message := fmt.Sprint(messageArgs...)

		logrusFields := fieldsToLogrus(fields)
		// Merge context fields with provided fields
		for k, v := range contextFields {
			logrusFields[k] = v
		}

		l.Logger.WithFields(logrusFields).Debug(message)
		return
	}

	// No fields, just message
	message := fmt.Sprint(args...)
	l.Logger.WithFields(contextFields).Debug(message)
}

// Global logger instance
var defaultLogger = New()

// SetDefault sets the default logger instance
func SetDefault(logger *Logger) {
	defaultLogger = logger
}

// GetDefault returns the default logger instance
func GetDefault() *Logger {
	return defaultLogger
}

// Convenience functions that use the default logger

// Error logs an error message with optional fields using the default logger
func Error(err error, args ...interface{}) {
	defaultLogger.Error(err, args...)
}

// ErrorWithContext logs an error message with context and optional fields using the default logger
func ErrorWithContext(ctx context.Context, err error, args ...interface{}) {
	defaultLogger.ErrorWithContext(ctx, err, args...)
}

// Warn logs a warning message with optional fields using the default logger
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// WarnWithContext logs a warning message with context and optional fields using the default logger
func WarnWithContext(ctx context.Context, args ...interface{}) {
	defaultLogger.WarnWithContext(ctx, args...)
}

// Info logs an info message with optional fields using the default logger
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// InfoWithContext logs an info message with context and optional fields using the default logger
func InfoWithContext(ctx context.Context, args ...interface{}) {
	defaultLogger.InfoWithContext(ctx, args...)
}

// Debug logs a debug message with optional fields using the default logger
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// DebugWithContext logs a debug message with context and optional fields using the default logger
func DebugWithContext(ctx context.Context, args ...interface{}) {
	defaultLogger.DebugWithContext(ctx, args...)
}
