package errors

import (
	"bytes"
	"errors"
)

// Tail of the error chain: at the end
func Tail(err error) error {
	var (
		tail     errLinkedList
		causable causableError
		hasTail  hasInner
	)
	if errors.As(err, &tail) {
		return tail.Tail()
	}
	if errors.As(err, &causable) {
		return causable.Cause()
	}
	if errors.As(err, &hasTail) {
		i := hasTail.GetInner()
		if errors.Is(err, i) || i == nil {
			return err
		}
		return Tail(i)
	}
	return err
}

// Next error just below this one, or nil if there is no next error.  Note this may be an error created
// for you if you used annotations.  As a user, you probably don't want to use this.
func Next(err error) error {
	var (
		tail errLinkedList
		u    hasUnderline
		i    hasInner
	)
	if errors.As(err, &tail) {
		return tail.Next()
	}
	if errors.As(err, &u) {
		return u.Underlying()
	}
	if errors.As(err, &i) {
		return i.GetInner()
	}
	return nil
}

// Message is the error string at the Head of the linked list
func Message(err error) string {
	var (
		tail   errLinkedList
		hasMsg hasMessage
	)
	if errors.As(err, &tail) {
		return tail.Head().Error()
	}
	if errors.As(err, &hasMsg) {
		return hasMsg.Message()
	}
	return err.Error()
}

// Details are an easy to read concat of all the error strings in a chain
func Details(err error) string {
	b := bytes.Buffer{}
	PanicIfErr(b.WriteByte('['), "unexpected")
	first := true
	for ; err != nil; err = Next(err) {
		if !first {
			PanicIfErrWrite(b.WriteString(" | "))
		}
		PanicIfErrWrite(b.WriteString(Message(err)))
		first = false
	}
	PanicIfErr(b.WriteByte(']'), "unexpected")
	return b.String()
}
