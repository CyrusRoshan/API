package types

import (
	"encoding/json"
	"errors"

	"gopkg.in/gorp.v1"
)

type PhysicalLocation struct {
	Name    string     `json:"name"`
	City    string     `json:"city"`
	State   string     `json:"state"`
	Country string     `json:"country"`
	Coords  [2]float32 `json:"coords"`
}

type Prize struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Sponsor     string `json:"sponsor"`
}

type BusLocation struct {
	Name   string   `json:"name"`
	Time   int64    `json:"time"`
	Coords [2]int16 `json:"coords"`
}

type SocialLink struct {
	Name string `json:"name"`
	Link string `json:"link"`
	Logo string `json:"logo,omitempty"`
}

type HardwareItem struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
	Unit     string `json:"unit,omitempty"`
}

type Hackathon struct {
	Id             int64            `json:"id"`
	OwnerID        int64            `json:"ownerid"`
	Name           string           `json:"name"`
	Location       PhysicalLocation `json:"location"`
	StartDate      int64            `json:"startdate"`
	EndDate        int64            `json:"enddate"`
	CurrentState   int              `json:"currentstate"`
	Prizes         []Prize          `json:"prizes"`
	Reimbursements bool             `json:"reimbursements"`
	BusesOffered   bool             `json:"busesoffered"`
	BusLocations   []BusLocation    `json:"buslocations"`
	SocialLinks    []SocialLink     `json:"sociallinks"`
	Hardware       []HardwareItem   `json:"hardware"`
	Map            string           `json:"map"`
	Metadata       string           `json:"metadata"`
}

type HackathonTypeConverter struct{}

func (me HackathonTypeConverter) ToDb(val interface{}) (interface{}, error) {

	switch t := val.(type) {
	case PhysicalLocation, []Prize, []BusLocation, []SocialLink, []HardwareItem:
		b, err := json.Marshal(t)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

	return val, nil
}

func (me HackathonTypeConverter) FromDb(target interface{}) (gorp.CustomScanner, bool) {
	switch target.(type) {
	case *PhysicalLocation, *[]Prize, *[]BusLocation, *[]SocialLink, *[]HardwareItem:
		binder := func(holder, target interface{}) error {
			s, ok := holder.(*string)
			if !ok {
				return errors.New("FromDb: Unable to convert interface to *string")
			}
			b := []byte(*s)
			return json.Unmarshal(b, target)
		}
		return gorp.CustomScanner{
			Holder: new(string),
			Target: target,
			Binder: binder,
		}, true
	}

	return gorp.CustomScanner{}, false
}
