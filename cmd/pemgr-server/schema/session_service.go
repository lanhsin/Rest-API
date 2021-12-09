package schema

type SessionService struct {
	// The type of a resource. [RO]
	OdataType string `json:"@odata.type"`

	// The identifier that uniquely identifies the Resource within
	// the collection of similar Resources. [RO]
	Id string `json:"Id"`

	// The name of the Resource or array member. [RO]
	Name string `json:"Name"`

	// The description of this Resource.
	// Used for commonality in the schema definitions. [RO]
	Description string `json:"Description"`

	// The status and health of a Resource and its children. [RW]
	Status CommonStatus `json:"Status"`

	// An indication of whether this service is enabled.
	// If `true`, this service is enabled.
	// If `false`, it is disabled, and new sessions cannot be created,
	// old sessions cannot be deleted, and established sessions may continue operating. [RW]
	ServiceEnabled bool `json:"ServiceEnabled"`

	// The number of seconds of inactivity that a session may have before
	// the Session Service closes the session due to inactivity. [RW]
	SessionTimeout int `json:"SessionTimeout"`

	// The SessionCollection schema describes a collection of session instances. [RO]
	Sessions CommonOid `json:"Tasks"`

	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`
}
