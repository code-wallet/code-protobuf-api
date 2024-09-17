// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: chat/v2/model.proto

package chat

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on ChatId with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ChatId) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetValue()) != 32 {
		return ChatIdValidationError{
			field:  "Value",
			reason: "value length must be 32 bytes",
		}
	}

	return nil
}

// ChatIdValidationError is the validation error returned by ChatId.Validate if
// the designated constraints aren't met.
type ChatIdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatIdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatIdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatIdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatIdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatIdValidationError) ErrorName() string { return "ChatIdValidationError" }

// Error satisfies the builtin error interface
func (e ChatIdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatIdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatIdValidationError{}
