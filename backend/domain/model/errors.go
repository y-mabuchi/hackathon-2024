package model

import "fmt"

type ErrNotFound struct {
	domain string
	id     string
}

func NewErrNotFound(domain, id string) error {
	return &ErrNotFound{
		domain,
		id,
	}
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf(
		"%s not found. id: %s",
		e.domain,
		e.id,
	)
}

type ErrInvalidArgument struct {
	domain string
	value  string
}

func NewErrInvalidArgument(domain, value string) error {
	return &ErrInvalidArgument{
		domain,
		value,
	}
}

func (e *ErrInvalidArgument) Error() string {
	return fmt.Sprintf(
		"%s is invalide in %s",
		e.value,
		e.domain,
	)
}
