package misc

import (
	"database/sql"
	"encoding/json"
	"time"
)

type NullString sql.NullString

func (x *NullString) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(x.String)
}

type BirthDate struct {
	time.Time
}

const birthDateLayout = "2006-01-02"

func (x *BirthDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + x.Format(birthDateLayout) + `"`), nil

}
