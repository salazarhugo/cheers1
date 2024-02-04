package repository

import (
	"cloud.google.com/go/spanner"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type PostMediaArray []PostMedia

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *PostMediaArray) Scan(value interface{}) error {
	bytes, ok := value.([]spanner.NullJSON)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	*j = ToMedia(bytes)

	return nil
}

// Value return json value, implement driver.Valuer interface
func (j PostMediaArray) Value() (driver.Value, error) {
	v := make([]PostMedia, 0)
	for _, s := range j {
		v = append(v, s)
	}
	return v, nil
}

func ToMedia(s []spanner.NullJSON) []PostMedia {
	res := make([]PostMedia, 0)
	for _, j := range s {
		bytes, _ := j.MarshalJSON()
		var postMedia PostMedia
		json.Unmarshal(bytes, &postMedia)
		res = append(res, postMedia)
	}
	return res
}
