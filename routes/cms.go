package cms

import (
	"encoding/json"
	"net/http"

	"github.com/go-martini/martini"
	"gopkg.in/gorp.v1"

	"github.com/cyrusroshan/API/types"
	"github.com/cyrusroshan/API/utils"
)

func GetHackathon(w http.ResponseWriter, r *http.Request, hackathonDB *gorp.DbMap, params martini.Params) (int, string) {
	var existingHackathon types.Hackathon
	err := hackathonDB.SelectOne(&existingHackathon, "select * from hackathons where id=?", params["id"])
	utils.PanicIf(err)

	return 200, utils.MustMarshal(existingHackathon)
}

func NewHackathon(w http.ResponseWriter, r *http.Request, hackathonDB *gorp.DbMap, params martini.Params) (int, string) {
	decoder := json.NewDecoder(r.Body)
	var newHackathon types.Hackathon

	err := decoder.Decode(&newHackathon)
	utils.PanicIf(err)

	// Change newHackathon.OwnerID here when auth has been implemented

	err = hackathonDB.Insert(&newHackathon)
	utils.PanicIf(err)

	return 200, utils.MustMarshal(newHackathon)
}

func EditHackathon(w http.ResponseWriter, r *http.Request, hackathonDB *gorp.DbMap, params martini.Params) (int, string) {
	var existingHackathon types.Hackathon
	err := hackathonDB.SelectOne(&existingHackathon, "select * from hackathons where id=?", params["id"])
	utils.PanicIf(err)

	// Make a check here to check existingHackathon.OwnerID here when auth has been implemented

	decoder := json.NewDecoder(r.Body)
	var editedHackathon types.Hackathon

	err = decoder.Decode(&editedHackathon)
	utils.PanicIf(err)

	editedHackathon.Id = existingHackathon.Id

	_, err = hackathonDB.Update(&editedHackathon)
	utils.PanicIf(err)

	return 200, utils.MustMarshal(&editedHackathon)
}

func DeleteHackathon(w http.ResponseWriter, r *http.Request, hackathonDB *gorp.DbMap, params martini.Params) (int, string) {
	var existingHackathon types.Hackathon
	err := hackathonDB.SelectOne(&existingHackathon, "select * from hackathons where id=?", params["id"])
	utils.PanicIf(err)

	// Make a check here to check existingHackathon.OwnerID here when auth has been implemented

	_, err = hackathonDB.Delete(&existingHackathon)
	utils.PanicIf(err)

	return 200, utils.MustMarshal(existingHackathon)
}
