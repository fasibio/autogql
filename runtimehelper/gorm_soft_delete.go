package runtimehelper

import (
	"fmt"
	"io"
	"time"

	"gorm.io/gorm"
)

type SoftDelete struct {
	gorm.DeletedAt
}

func (y *SoftDelete) UnmarshalGQL(v interface{}) error {
	date, ok := v.(string)
	if !ok {
		return fmt.Errorf("SoftDelete must be a string")
	}
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return err
	}
	y.Time = t
	y.Valid = true

	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (y SoftDelete) MarshalGQL(w io.Writer) {
	w.Write([]byte(y.Time.Format(time.RFC3339)))
}
