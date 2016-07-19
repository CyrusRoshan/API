package cms

import (
	"encoding/json"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/jmoiron/sqlx"

	"github.com/cyrusroshan/API/types"
	"github.com/cyrusroshan/API/utils"
)

func GetHackathon(w http.ResponseWriter, r *http.Request, hackathonDB *sqlx.DB, params martini.Params) (int, string) {
	var existingHackathon types.Hackathon
	err := hackathonDB.Select(&existingHackathon, "SELECT * FROM hackathons WHERE id=$1", params["id"])
	utils.PanicIf(err)

	return 200, utils.MustMarshal(existingHackathon)
}

func NewHackathon(w http.ResponseWriter, r *http.Request, hackathonDB *sqlx.DB, params martini.Params) (int, string) {
	decoder := json.NewDecoder(r.Body)
	var newHackathon types.Hackathon

	err := decoder.Decode(&newHackathon)
	utils.PanicIf(err)

	// Change newHackathon.OwnerID here when auth has been implemented

	hackathonDB.MustExec("INSERT INTO hackathons (ownerID, name, location, startDate, endDate, currentState, prizes, reimbursements, busesOffered, busLocations, socialLinks, hardware, map, metadata) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)",
		newHackathon.OwnerID,
		newHackathon.Name,
		utils.MustMarshal(newHackathon.Location),
		newHackathon.StartDate,
		newHackathon.EndDate,
		newHackathon.CurrentState,
		utils.MustMarshal(newHackathon.Prizes),
		newHackathon.Reimbursements,
		newHackathon.BusesOffered,
		utils.MustMarshal(newHackathon.BusLocations),
		utils.MustMarshal(newHackathon.SocialLinks),
		utils.MustMarshal(newHackathon.Hardware),
		newHackathon.Map,
		newHackathon.Metadata,
	)

	return 200, utils.MustMarshal(newHackathon)
}

func EditHackathon(w http.ResponseWriter, r *http.Request, hackathonDB *sqlx.DB, params martini.Params) (int, string) {
	// var existingHackathon types.Hackathon
	// err := hackathonDB.SelectOne(&existingHackathon, "select * from hackathons where id=?", params["id"])
	// utils.PanicIf(err)
	//
	// // Make a check here to check existingHackathon.OwnerID here when auth has been implemented
	//
	// decoder := json.NewDecoder(r.Body)
	// var editedHackathon types.Hackathon
	//
	// err = decoder.Decode(&editedHackathon)
	// utils.PanicIf(err)
	//
	// editedHackathon.Id = existingHackathon.Id
	//
	// _, err = hackathonDB.Update(&editedHackathon)
	// utils.PanicIf(err)
	//
	// return 200, utils.MustMarshal(&editedHackathon)
	return 200, ""
}

func DeleteHackathon(w http.ResponseWriter, r *http.Request, hackathonDB *sqlx.DB, params martini.Params) (int, string) {
	// var existingHackathon types.Hackathon
	// err := hackathonDB.SelectOne(&existingHackathon, "select * from hackathons where id=?", params["id"])
	// utils.PanicIf(err)
	//
	// // Make a check here to check existingHackathon.OwnerID here when auth has been implemented
	//
	// _, err = hackathonDB.Delete(&existingHackathon)
	// utils.PanicIf(err)
	//
	// return 200, utils.MustMarshal(existingHackathon)
	return 200, ""
}
