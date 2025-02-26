package libs

import "database/sql"

func ToNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: true}
}

func ToNullInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: true}
}
