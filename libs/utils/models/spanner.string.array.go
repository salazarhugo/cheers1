package models

import (
	"cloud.google.com/go/spanner"
	"database/sql/driver"
	"errors"
	"fmt"
)

type ArrayString []string

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *ArrayString) Scan(value interface{}) error {
	bytes, ok := value.([]spanner.NullString)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	*j = ToStringArray(bytes)

	return nil
}

// Value return json value, implement driver.Valuer interface
func (j ArrayString) Value() (driver.Value, error) {
	str := make([]string, 0)
	for _, s := range j {
		str = append(str, s)
	}
	return str, nil
}
