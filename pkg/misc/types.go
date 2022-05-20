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

type Date struct {
	time.Time
}

const dateLayout = "2006-01-02"

func (x *Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + x.Format(dateLayout) + `"`), nil

}
