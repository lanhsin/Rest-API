package schema

type SystemCollection struct {
	// The type of a resource.
	OdataType string `json:"@odata.type"`

	// The name of the Resource or array member.
	Name string `json:"Name"`

	// The number of items in a collection.
	MemberOdataCount int `json:"Member@odata.count"`

	// The members of this collection.
	Members []CommonOid

	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`
}
