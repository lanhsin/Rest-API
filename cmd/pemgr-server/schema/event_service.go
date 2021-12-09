package schema

type EventService struct {
	// The type of a resource. [RO]
	OdataType string `json:"@odata.type"`

	// The identifier that uniquely identifies the Resource within
	// the collection of similar Resources. [RO]
	Id string `json:"Id"`

	// The name of the Resource or array member. [RO]
	Name string `json:"Name"`

	// The status and health of a Resource and its children. [RW]
	Status CommonStatus `json:"Status"`

	// An indication of whether this service is enabled. [RW]
	ServiceEnabled bool `json:"ServiceEnabled"`

	// The number of times that the POST of an event is retried before the subscription terminates.
	// This retry occurs at the service level, which means that the HTTP POST to the event destination fails with
	// an HTTP `4XX` or `5XX` status code or an HTTP timeout occurs this many times before
	// the event destination subscription terminates. [RW]
	DeliveryRetryAttempts int `json:"DeliveryRetryAttempts"`

	// The interval, in seconds, between retry attempts for sending any event. [RW]
	DeliveryRetryIntervalSeconds int `json:"DeliveryRetryIntervalSeconds"`

	// The types of events to which a client can subscribe. [RO]
	EventTypesForSubscription []string `json:"EventTypesForSubscription"`

	// A Collection of EventDestination Resource instances. [RO]
	Subscriptions CommonOid `json:"Subscriptions"`

	// The available actions for this Resource. [RW]
	Actions EventServiceActions `json:"Actions"`

	// Settings for SMTP event delivery. [RW]
	SMTP EventServiceSMTP `json:"SMTP"`

	// The OEM extension. [RW]
	Oem EventServiceOem `json:"Oem,omitempty"`

	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`
}

type EventServiceActions struct {
	EventServiceSubmitTestEvent EventServiceActionsEventServiceSubmitTestEvent `json:"#EventService.SubmitTestEvent"`

	// The available OEM-specific actions for this Resource.
	Oem EventServiceActionsOem `json:"Oem,omitempty"`
}

type EventServiceActionsEventServiceSubmitTestEvent struct {
	Target            string `json:"target"`
	RedfishActionInfo string `json:"@Redfish.ActionInfo"`
}

type EventServiceActionsOem struct {
}

type EventServiceSMTP struct {
	// An indication if SMTP for event delivery is enabled. [RW]
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// The destination SMTP port. [RW]
	Port int `json:"Port,omitempty"`

	// The address of the SMTP server. [RW]
	ServerAddress string `json:"ServerAddress,omitempty"`

	// The 'from' email address of the outgoing email. [RW]
	FromAddress string `json:"FromAddress,omitempty"`

	// The connection type to the outgoing SMTP server. [RW]
	// Valid values:
	// AutoDetect:	Auto-detect.
	// None:	    Clear text.
	// StartTLS:	StartTLS.
	// TLS_SSL:	    TLS/SSL.
	ConnectionProtocol string `json:"ConnectionProtocol"`

	// The authentication method for the SMTP server. [RW]
	// Valid values:
	// AutoDetect:	Auto-detect.
	// CRAM_MD5:	CRAM-MD5 authentication.
	// Login:	    LOGIN authentication.
	// None:	    No authentication.
	// Plain:	    PLAIN authentication.
	Authentication string `json:"Authentication"`

	// The username for authentication with the SMTP server. [RW]
	Username string `json:"Username,omitempty"`

	// The password for authentication with the SMTP server.
	// The value is `null` in responses. [RW]
	Password string `json:"Password,omitempty"`
}

type EventServiceOem struct {
}
