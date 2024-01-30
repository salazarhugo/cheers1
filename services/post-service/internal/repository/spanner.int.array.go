package repository

import (
	"cloud.google.com/go/spanner"
	"database/sql/driver"
	"errors"
	"fmt"
)

type ArrayInt []int64

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *ArrayInt) Scan(value interface{}) error {
	bytes, ok := value.([]spanner.NullInt64)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	*j = ToIntArray(bytes)

	return nil
}

// Value return json value, implement driver.Valuer interface
func (j ArrayInt) Value() (driver.Value, error) {
	v := make([]int64, 0)
	for _, s := range j {
		v = append(v, s)
	}
	return v, nil
}

func ToIntArray(s []spanner.NullInt64) []int64 {
	res := make([]int64, 0)
	for _, str := range s {
		res = append(res, str.Int64)
	}
	return res
}
