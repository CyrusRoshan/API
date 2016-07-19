package db

import (
	"github.com/cyrusroshan/API/utils"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const (
	PendingState = iota
	OpenState
	ClosedState
)

func InitHackathons(url string) *sqlx.DB {
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	db, err := sqlx.Connect("postgres", connection)
	utils.PanicIf(err)

	db.MustExec("CREATE TABLE IF NOT EXISTS hackathons (id SERIAL PRIMARY KEY, ownerID INTEGER, name TEXT, location JSONB, startDate BIGINT, endDate BIGINT, currentState INTEGER, prizes JSONB, reimbursements BOOLEAN, busesOffered BOOLEAN, busLocations JSONB, socialLinks JSONB, hardware JSONB, map TEXT, metadata JSONB)")

	return db
}
