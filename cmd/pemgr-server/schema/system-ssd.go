package schema

import (
	"errors"
	"strings"
)

type SsdCollection struct {
	// The type of a resource.
	OdataType string `json:"@odata.type"`

	// The name of the Resource or array member.
	Name string `json:"Name"`

	// The number of items in a collection.
	MemberOdataCount int64 `json:"Member@odata.count"`

	// The members of this collection.
	Members []CommonOid

	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`
}

type Ssd struct {
	// The type of a resource. [RO]
	OdataType string `json:"@odata.type"`

	// The identifier that uniquely identifies the Resource within
	// the collection of similar Resources. [RO]
	Id string `json:"Id"`

	// The name of the Resource or array member. [RO]
	Name string `json:"Name"`

	// Irene OEM
	Manufacturer   string  `json:"Manufacturer"`
	Model          string  `json:"Model"`
	SerialNumber   string  `json:"SerialNumber"`
	ReadingCelsius float64 `json:"ReadingCelsius,omitempty"`

	Status CommonStatus `json:"Status"`

	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`
}

type SsdOem struct {
	Custom SsdOemCustom `json:"Custom"`
}

type SsdOemCustom struct {
	// The action to control LED. [RW]
	// Valid values:
	// Reset:	Reset SSD.
	Action string `json:"Action"`
}

func (ssd SsdOemCustom) Validation() error {
	switch {
	case len(ssd.Action) == 0:
		return errors.New("The action is empty")
	case strings.ToLower(ssd.Action) == "reset":
		break
	default:
		return errors.New("The action is invalied")
	}
	return nil
}