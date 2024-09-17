// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/v1/model.proto

package common

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

// Validate checks the field values on SolanaAccountId with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *SolanaAccountId) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetValue()) != 32 {
		return SolanaAccountIdValidationError{
			field:  "Value",
			reason: "value length must be 32 bytes",
		}
	}

	return nil
}

// SolanaAccountIdValidationError is the validation error returned by
// SolanaAccountId.Validate if the designated constraints aren't met.
type SolanaAccountIdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SolanaAccountIdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SolanaAccountIdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SolanaAccountIdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SolanaAccountIdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SolanaAccountIdValidationError) ErrorName() string { return "SolanaAccountIdValidationError" }

// Error satisfies the builtin error interface
func (e SolanaAccountIdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSolanaAccountId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SolanaAccountIdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SolanaAccountIdValidationError{}

// Validate checks the field values on InstructionAccount with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *InstructionAccount) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetAccount() == nil {
		return InstructionAccountValidationError{
			field:  "Account",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetAccount()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InstructionAccountValidationError{
				field:  "Account",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for IsSigner

	// no validation rules for IsWritable

	return nil
}

// InstructionAccountValidationError is the validation error returned by
// InstructionAccount.Validate if the designated constraints aren't met.
type InstructionAccountValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InstructionAccountValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InstructionAccountValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InstructionAccountValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InstructionAccountValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InstructionAccountValidationError) ErrorName() string {
	return "InstructionAccountValidationError"
}

// Error satisfies the builtin error interface
func (e InstructionAccountValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInstructionAccount.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InstructionAccountValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InstructionAccountValidationError{}

// Validate checks the field values on Transaction with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Transaction) Validate() error {
	if m == nil {
		return nil
	}

	if l := len(m.GetValue()); l < 1 || l > 1232 {
		return TransactionValidationError{
			field:  "Value",
			reason: "value length must be between 1 and 1232 bytes, inclusive",
		}
	}

	return nil
}

// TransactionValidationError is the validation error returned by
// Transaction.Validate if the designated constraints aren't met.
type TransactionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TransactionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TransactionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TransactionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TransactionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TransactionValidationError) ErrorName() string { return "TransactionValidationError" }

// Error satisfies the builtin error interface
func (e TransactionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTransaction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TransactionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TransactionValidationError{}

// Validate checks the field values on Blockhash with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Blockhash) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetValue()) != 32 {
		return BlockhashValidationError{
			field:  "Value",
			reason: "value length must be 32 bytes",
		}
	}

	return nil
}

// BlockhashValidationError is the validation error returned by
// Blockhash.Validate if the designated constraints aren't met.
type BlockhashValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BlockhashValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BlockhashValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BlockhashValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BlockhashValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BlockhashValidationError) ErrorName() string { return "BlockhashValidationError" }

// Error satisfies the builtin error interface
func (e BlockhashValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBlockhash.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BlockhashValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BlockhashValidationError{}

// Validate checks the field values on Signature with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Signature) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetValue()) != 64 {
		return SignatureValidationError{
			field:  "Value",
			reason: "value length must be 64 bytes",
		}
	}

	return nil
}

// SignatureValidationError is the validation error returned by
// Signature.Validate if the designated constraints aren't met.
type SignatureValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignatureValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignatureValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignatureValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignatureValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignatureValidationError) ErrorName() string { return "SignatureValidationError" }

// Error satisfies the builtin error interface
func (e SignatureValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignature.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignatureValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignatureValidationError{}

// Validate checks the field values on IntentId with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *IntentId) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetValue()) != 32 {
		return IntentIdValidationError{
			field:  "Value",
			reason: "value length must be 32 bytes",
		}
	}

	return nil
}

// IntentIdValidationError is the validation error returned by
// IntentId.Validate if the designated constraints aren't met.
type IntentIdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntentIdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntentIdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntentIdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntentIdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntentIdValidationError) ErrorName() string { return "IntentIdValidationError" }

// Error satisfies the builtin error interface
func (e IntentIdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntentId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntentIdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntentIdValidationError{}

// Validate checks the field values on UserId with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *UserId) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetValue()) != 16 {
		return UserIdValidationError{
			field:  "Value",
			reason: "value length must be 16 bytes",
		}
	}

	return nil
}

// UserIdValidationError is the validation error returned by UserId.Validate if
// the designated constraints aren't met.
type UserIdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserIdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserIdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserIdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserIdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserIdValidationError) ErrorName() string { return "UserIdValidationError" }

// Error satisfies the builtin error interface
func (e UserIdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserIdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserIdValidationError{}

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

// Validate checks the field values on DataContainerId with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DataContainerId) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetValue()) != 16 {
		return DataContainerIdValidationError{
			field:  "Value",
			reason: "value length must be 16 bytes",
		}
	}

	return nil
}

// DataContainerIdValidationError is the validation error returned by
// DataContainerId.Validate if the designated constraints aren't met.
type DataContainerIdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataContainerIdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataContainerIdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataContainerIdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DataContainerIdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataContainerIdValidationError) ErrorName() string { return "DataContainerIdValidationError" }

// Error satisfies the builtin error interface
func (e DataContainerIdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDataContainerId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataContainerIdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DataContainerIdValidationError{}

// Validate checks the field values on DeviceToken with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DeviceToken) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetValue()); l < 1 || l > 16348 {
		return DeviceTokenValidationError{
			field:  "Value",
			reason: "value length must be between 1 and 16348 runes, inclusive",
		}
	}

	return nil
}

// DeviceTokenValidationError is the validation error returned by
// DeviceToken.Validate if the designated constraints aren't met.
type DeviceTokenValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeviceTokenValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeviceTokenValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeviceTokenValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeviceTokenValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeviceTokenValidationError) ErrorName() string { return "DeviceTokenValidationError" }

// Error satisfies the builtin error interface
func (e DeviceTokenValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeviceToken.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeviceTokenValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeviceTokenValidationError{}

// Validate checks the field values on AppInstallId with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AppInstallId) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetValue()); l < 1 || l > 256 {
		return AppInstallIdValidationError{
			field:  "Value",
			reason: "value length must be between 1 and 256 runes, inclusive",
		}
	}

	return nil
}

// AppInstallIdValidationError is the validation error returned by
// AppInstallId.Validate if the designated constraints aren't met.
type AppInstallIdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppInstallIdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppInstallIdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppInstallIdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppInstallIdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppInstallIdValidationError) ErrorName() string { return "AppInstallIdValidationError" }

// Error satisfies the builtin error interface
func (e AppInstallIdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppInstallId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppInstallIdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppInstallIdValidationError{}

// Validate checks the field values on PhoneNumber with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *PhoneNumber) Validate() error {
	if m == nil {
		return nil
	}

	if !_PhoneNumber_Value_Pattern.MatchString(m.GetValue()) {
		return PhoneNumberValidationError{
			field:  "Value",
			reason: "value does not match regex pattern \"^\\\\+[1-9]\\\\d{1,14}$\"",
		}
	}

	return nil
}

// PhoneNumberValidationError is the validation error returned by
// PhoneNumber.Validate if the designated constraints aren't met.
type PhoneNumberValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PhoneNumberValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PhoneNumberValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PhoneNumberValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PhoneNumberValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PhoneNumberValidationError) ErrorName() string { return "PhoneNumberValidationError" }

// Error satisfies the builtin error interface
func (e PhoneNumberValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPhoneNumber.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PhoneNumberValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PhoneNumberValidationError{}

var _PhoneNumber_Value_Pattern = regexp.MustCompile("^\\+[1-9]\\d{1,14}$")

// Validate checks the field values on Domain with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Domain) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetValue()); l < 1 || l > 253 {
		return DomainValidationError{
			field:  "Value",
			reason: "value length must be between 1 and 253 runes, inclusive",
		}
	}

	if err := m._validateHostname(m.GetValue()); err != nil {
		return DomainValidationError{
			field:  "Value",
			reason: "value must be a valid hostname",
			cause:  err,
		}
	}

	return nil
}

func (m *Domain) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

// DomainValidationError is the validation error returned by Domain.Validate if
// the designated constraints aren't met.
type DomainValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DomainValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DomainValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DomainValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DomainValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DomainValidationError) ErrorName() string { return "DomainValidationError" }

// Error satisfies the builtin error interface
func (e DomainValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDomain.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DomainValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DomainValidationError{}

// Validate checks the field values on Relationship with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Relationship) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Type.(type) {

	case *Relationship_Domain:

		if m.GetDomain() == nil {
			return RelationshipValidationError{
				field:  "Domain",
				reason: "value is required",
			}
		}

		if v, ok := interface{}(m.GetDomain()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RelationshipValidationError{
					field:  "Domain",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RelationshipValidationError is the validation error returned by
// Relationship.Validate if the designated constraints aren't met.
type RelationshipValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RelationshipValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RelationshipValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RelationshipValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RelationshipValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RelationshipValidationError) ErrorName() string { return "RelationshipValidationError" }

// Error satisfies the builtin error interface
func (e RelationshipValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRelationship.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RelationshipValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RelationshipValidationError{}

// Validate checks the field values on Hash with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Hash) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetValue()) != 32 {
		return HashValidationError{
			field:  "Value",
			reason: "value length must be 32 bytes",
		}
	}

	return nil
}

// HashValidationError is the validation error returned by Hash.Validate if the
// designated constraints aren't met.
type HashValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HashValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HashValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HashValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HashValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HashValidationError) ErrorName() string { return "HashValidationError" }

// Error satisfies the builtin error interface
func (e HashValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHash.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HashValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HashValidationError{}

// Validate checks the field values on Locale with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Locale) Validate() error {
	if m == nil {
		return nil
	}

	if !_Locale_Value_Pattern.MatchString(m.GetValue()) {
		return LocaleValidationError{
			field:  "Value",
			reason: "value does not match regex pattern \"^[A-Za-z]{2,4}([_-][A-Za-z]{4})?([_-]([A-Za-z]{2}|[0-9]{3}))?$\"",
		}
	}

	return nil
}

// LocaleValidationError is the validation error returned by Locale.Validate if
// the designated constraints aren't met.
type LocaleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocaleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocaleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocaleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocaleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocaleValidationError) ErrorName() string { return "LocaleValidationError" }

// Error satisfies the builtin error interface
func (e LocaleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocale.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocaleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocaleValidationError{}

var _Locale_Value_Pattern = regexp.MustCompile("^[A-Za-z]{2,4}([_-][A-Za-z]{4})?([_-]([A-Za-z]{2}|[0-9]{3}))?$")

// Validate checks the field values on UUID with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *UUID) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetValue()) != 16 {
		return UUIDValidationError{
			field:  "Value",
			reason: "value length must be 16 bytes",
		}
	}

	return nil
}

// UUIDValidationError is the validation error returned by UUID.Validate if the
// designated constraints aren't met.
type UUIDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UUIDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UUIDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UUIDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UUIDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UUIDValidationError) ErrorName() string { return "UUIDValidationError" }

// Error satisfies the builtin error interface
func (e UUIDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUUID.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UUIDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UUIDValidationError{}

// Validate checks the field values on Request with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Version

	// no validation rules for Service

	// no validation rules for Method

	// no validation rules for Body

	return nil
}

// RequestValidationError is the validation error returned by Request.Validate
// if the designated constraints aren't met.
type RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestValidationError) ErrorName() string { return "RequestValidationError" }

// Error satisfies the builtin error interface
func (e RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestValidationError{}

// Validate checks the field values on Response with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	// no validation rules for Body

	// no validation rules for Message

	return nil
}

// ResponseValidationError is the validation error returned by
// Response.Validate if the designated constraints aren't met.
type ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResponseValidationError) ErrorName() string { return "ResponseValidationError" }

// Error satisfies the builtin error interface
func (e ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResponseValidationError{}

// Validate checks the field values on ServerPing with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ServerPing) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetTimestamp() == nil {
		return ServerPingValidationError{
			field:  "Timestamp",
			reason: "value is required",
		}
	}

	if m.GetPingDelay() == nil {
		return ServerPingValidationError{
			field:  "PingDelay",
			reason: "value is required",
		}
	}

	return nil
}

// ServerPingValidationError is the validation error returned by
// ServerPing.Validate if the designated constraints aren't met.
type ServerPingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerPingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerPingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerPingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerPingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerPingValidationError) ErrorName() string { return "ServerPingValidationError" }

// Error satisfies the builtin error interface
func (e ServerPingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerPing.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerPingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerPingValidationError{}

// Validate checks the field values on ClientPong with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ClientPong) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetTimestamp() == nil {
		return ClientPongValidationError{
			field:  "Timestamp",
			reason: "value is required",
		}
	}

	return nil
}

// ClientPongValidationError is the validation error returned by
// ClientPong.Validate if the designated constraints aren't met.
type ClientPongValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClientPongValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClientPongValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClientPongValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClientPongValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClientPongValidationError) ErrorName() string { return "ClientPongValidationError" }

// Error satisfies the builtin error interface
func (e ClientPongValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClientPong.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClientPongValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClientPongValidationError{}
