package store

import (
	"database/sql"
	"fmt"
)

func nullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}
 
func locationToEWKT(l *Location) interface{} {
	if l == nil {
		return nil
	}
	return fmt.Sprintf("SRID=4326;POINT(%f %f)", l.Lng, l.Lat)
}
