package schema

type ChassisCollection struct {
	// The type of a resource. [RO]
	OdataType string `json:"@odata.type"`

	// The name of the Resource or array member. [RO]
	Name string `json:"Name"`

	// The number of items in a collection. [RO]
	MemberOdataCount int `json:"Member@odata.count"`

	// The members of this collection. [RO]
	Members []CommonOid

	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`
}
