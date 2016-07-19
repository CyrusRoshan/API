package db

import (
	"database/sql"
	"log"

	"gopkg.in/gorp.v1"

	"github.com/cyrusroshan/API/types"
	"github.com/cyrusroshan/API/utils"
	"github.com/lib/pq"
)

const (
	PendingState = iota
	OpenState
	ClosedState
)

func InitHackathons(url string) *gorp.DbMap {
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Println(err)
	}

	dbmap := &gorp.DbMap{
		Db:            db,
		Dialect:       gorp.PostgresDialect{},
		TypeConverter: types.HackathonTypeConverter{},
	}

	dbmap.AddTableWithName(types.Hackathon{}, "hackathons").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	utils.PanicIf(err)

	return dbmap
}
