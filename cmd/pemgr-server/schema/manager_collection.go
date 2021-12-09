package schema

type ManagerCollection struct {
	// The type of a resource. [RO]
	OdataType string `json:"@odata.type"`

	// The name of the Resource or array member. [RO]
	Name string `json:"Name"`

	// The number of items in a collection. [RO]
	MemberOdataCount int `json:"Member@odata.count"`

	// The members of this collection. [RO]
	Members []CommonOid `json:"Members,omitempty"`

	// The OEM extension. [RW]
	Oem ManagerCollectionOem `json:"Oem,omitempty"`

	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`
}

type ManagerCollectionOem struct {
}
