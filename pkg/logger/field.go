package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

// Field represents a key-value pair for structured logging
type Field struct {
	Key   string
	Value interface{}
}

// Fields represents a collection of fields with methods for manipulation
type Fields []Field

// Event creates an event field
func EventName(value string) Field {
	return Field{Key: "event_name", Value: value}
}

// String creates a string field
func String(key, value string) Field {
	return Field{Key: key, Value: value}
}

// Int creates an integer field
func Int(key string, value int) Field {
	return Field{Key: key, Value: value}
}

// Int64 creates an int64 field
func Int64(key string, value int64) Field {
	return Field{Key: key, Value: value}
}

// Float creates a float field
func Float(key string, value float64) Field {
	return Field{Key: key, Value: value}
}

// Bool creates a boolean field
func Bool(key string, value bool) Field {
	return Field{Key: key, Value: value}
}

// Struct creates a struct field
func Struct(key string, value interface{}) Field {
	return Field{Key: key, Value: value}
}

// Any creates a field with any type
func Any(key string, value interface{}) Field {
	return Field{Key: key, Value: value}
}

// NewFields creates a slice of fields from variadic Field arguments
func NewFields(fields ...Field) Fields {
	return Fields(fields)
}

// Add appends a field to the existing fields
func (f *Fields) Add(field Field) {
	*f = append(*f, field)
}

// fieldsToLogrus converts our Fields to logrus.Fields
func fieldsToLogrus(fields Fields) logrus.Fields {
	logrusFields := make(logrus.Fields)
	for _, field := range fields {
		logrusFields[field.Key] = field.Value
	}
	return logrusFields
}

// contextToLogrus converts context to logrus fields if it contains relevant information
func contextToLogrus(ctx context.Context) logrus.Fields {
	if ctx == nil {
		return logrus.Fields{}
	}

	// You can extend this to extract specific context values
	// For example, request ID, user ID, etc.
	return logrus.Fields{}
}
