package schema

type UpdateService struct {
	// The type of a resource. [RO]
	OdataType string `json:"@odata.type"`

	// The identifier that uniquely identifies the Resource within
	// the collection of similar Resources. [RO]
	Id string `json:"Id"`

	// The name of the Resource or array member.
	Name string `json:"Name"`

	// The status and health of a Resource and its children. [RW]
	Status UpdateServiceStatus `json:"Status"`

	// An indication of whether this service is enabled. [RW]
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// The URI used to perform an HTTP or HTTPS push update to the Update Service.
	// The format of the message is vendor-specific. [RO]
	HttpPushUri string `json:"HttpPushUri"`

	// The settings for HttpPushUri-provided software updates. [RW]
	HttpPushUriOptions UpdateServiceHttpPushUriOptions `json:"HttpPushUriOptions"`

	// An indication of whether a client has reserved the HttpPushUriOptions properties for software updates. [RW]
	HttpPushUriOptionsBusy bool `json:"HttpPushUriOptionsBusy,omitempty"`

	// The SoftwareInventoryCollection schema contains a collection of software inventory instances. [RO]
	FirmwareInventory CommonOid `json:"FirmwareInventory"`

	// The available actions for this Resource. [RW]
	Actions UpdateServiceActions `json:"Actions"`

	// The OEM extension. [RW]
	Oem UpdateServiceOem `json:"Oem,omitempty"`

	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`
}

type UpdateServiceStatus struct {
	// The known state of the Resource, such as, enabled. [RO]
	// Valid values:
	// Absent:	            This function or Resource is either not present or detected.
	// Deferring:	        The element does not process any commands but queues new requests.
	// Disabled:	        This function or Resource is disabled.
	// Enabled:	            This function or Resource is enabled.
	// InTest:	            This function or Resource is undergoing testing, or is in the process of capturing information for debugging.
	// Qualified:	        The element quality is within the acceptable range of operation.
	// Quiesced:	        The element is enabled but only processes a restricted set of commands.
	// StandbyOffline:	    This function or Resource is enabled but awaits an external action to activate it.
	// StandbySpare:	    This function or Resource is part of a redundancy set and awaits a failover or other external action to activate it.
	// Starting:	        This function or Resource is starting.
	// UnavailableOffline:	This function or Resource is present but cannot be used.
	// Updating:	        The element is updating and may be unavailable or degraded.
	State string `json:"State"`

	// The health state of this Resource in the absence of its dependent Resources. [RO]
	// Valid values:
	// Critical:A critical condition requires immediate attention.
	// OK:	    Normal.
	// Warning:	A condition requires attention.
	Health string `json:"Health"`

	// The overall health state from the view of this Resource. [RO]
	// Valid values:
	// Critical:	A critical condition requires immediate attention.
	// OK:	        Normal.
	// Warning:	    A condition requires attention.
	HealthRollup string `json:"HealthRollup"`
}

type UpdateServiceHttpPushUriOptions struct {
	// The settings for when to apply HttpPushUri-provided software. [RW]
	HttpPushUriApplyTime UpdateServiceHttpPushUriOptionsHttpPushUriApplyTime `json:"HttpPushUriApplyTime"`
}

type UpdateServiceHttpPushUriOptionsHttpPushUriApplyTime struct {
	// The time when to apply the HttpPushUri-provided software update. [RW]
	// Valid values:
	// AtMaintenanceWindowStart:	Apply during an administrator-specified maintenance window.
	// Immediate:	                Apply immediately.
	// InMaintenanceWindowOnReset:	Apply after a reset but within an administrator-specified maintenance window.
	// OnReset:	                    Apply on a reset.
	ApplyTime                       string   `json:"ApplyTime"`
	ApplyTimeRedfishAllowableValues []string `json:"ApplyTime@Redfish.AllowableValues"`

	// The start time of a maintenance window. [RW]
	MaintenanceWindowStartTime string `json:"MaintenanceWindowStartTime"`

	// The expiry time, in seconds, of the maintenance window. [RW]
	MaintenanceWindowDurationInSeconds int `json:"MaintenanceWindowDurationInSeconds"`
}

type UpdateServiceActions struct {
	UpdateServiceSimpleUpdate UpdateServiceActionsUpdateServiceSimpleUpdate `json:"#UpdateService.SimpleUpdate"`

	// The available OEM-specific actions for this Resource. [RW]
	Oem UpdateServiceActionsOem `json:"Oem,omitempty"`
}

type UpdateServiceActionsUpdateServiceSimpleUpdate struct {
	Target            string `json:"target"`
	RedfishActionInfo string `json:"@Redfish.ActionInfo"`
}

type UpdateServiceActionsOem struct {
}

type UpdateServiceOem struct {
}
