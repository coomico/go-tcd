package tcd

import (
	"errors"
)

const (
	containNoput  = "noput~contains~"
	containAbsris = "absris~contains~"
	and           = "~and~"
)

func (q *Query) FilterByNoPut(noput string) error {
	if noput == "" {
		return errors.New("empty noput arguments")
	}

	if q.filter != "" {
		q.filter = ""
	}

	q.filter = containNoput + "'" + noput + "'"
	return nil
}

func (q *Query) FilterByAbsris(absris string) error {
	if absris == "" {
		return errors.New("empty absris arguments")
	}

	if q.filter != "" {
		q.filter = ""
	}

	q.filter = containAbsris + "'" + absris + "'"
	return nil
}

func (q *Query) FilterByNoPutAndAbsris(noput, absris string) error {
	if noput == "" || absris == "" {
		return errors.New("empty absris or noput arguments")
	}

	if q.filter != "" {
		q.filter = ""
	}

	noput = "'" + noput + "'"
	absris = "'" + absris + "'"
	q.filter = containNoput + noput + and + containAbsris + absris
	return nil
}
