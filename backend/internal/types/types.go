package types

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/lib/pq"
)

// NullTime is an alias for pq.NullTime data type
type NullTime pq.NullTime

// Scan implements the Scanner interface for NullTime
func (nt *NullTime) Scan(value interface{}) error {
	var t pq.NullTime
	if err := t.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nt = NullTime{t.Time, false}
	} else {
		*nt = NullTime{t.Time, true}
	}

	return nil
}

// MarshalJSON for NullTime
func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339))
	return []byte(val), nil
}

const dateFormat = "2006-01-02"

// UnmarshalJSON for NullTime
func (nt *NullTime) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(dateFormat, strings.Replace(
		string(b),
		"\"",
		"",
		-1,
	))

	if err != nil {
		return err
	}

	nt.Time = t
	nt.Valid = true

	return nil
}