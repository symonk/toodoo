package db

import "github.com/jmoiron/sqlx"

// rebindQuery rewrites the query string that utilises ? bindvars
// into something automatically applicable to the underlying database
// instance.
func RebindQuery(db *sqlx.DB, queryString string) string {
	return db.Rebind(queryString)
}
