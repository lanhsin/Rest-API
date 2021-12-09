package schema

type CommonOid struct {
	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`
}

type CommonStatus struct {
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
}

type CommonStatus2 struct {
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
	// Critical	A critical condition requires immediate attention.
	// OK	Normal.
	// Warning	A condition requires attention.
	HealthRollup string `json:"HealthRollup"`
}

type CommonError struct {
}

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
