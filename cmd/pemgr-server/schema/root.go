package schema

type Root struct {
	// The type of a resource. [RO]
	OdataType string `json:"@odata.type"`

	// The identifier that uniquely identifies the Resource within
	// the collection of similar Resources. [RO]
	Id string `json:"Id"`

	// The name of the Resource or array member. [RO]
	Name string `json:"Name"`

	// The version of the Redfish Service. [RO]
	RedfishVersion string `json:"RedfishVersion"`

	// Unique identifier for a service instance. When SSDP is used,
	// this value should be an exact match of the UUID value returned in a 200 OK
	// from an SSDP M-SEARCH request during discovery. [RO]
	UUID string `json:"UUID"`

	// The collection of ComputerSystem Resource instances. [RO]
	Systems CommonOid `json:"Systems"`

	// The ChassisCollection schema describes a collection of Chassis Resource instances. [RO]
	Chassis CommonOid `json:"Chasis"`

	// A Collection of manager Resource instances. [RO]
	Managers CommonOid `json:"Managers"`

	// The TaskService schema describes a Task Service that
	// enables management of long-duration operations,
	// includes the properties for the Task Service itself,
	// and has links to the actual Resource Collection of Tasks. [RO]
	Tasks CommonOid `json:"Tasks"`

	// The SessionService schema describes the Session Service and its properties,
	// with links to the actual list of sessions. [RO]
	SessionService CommonOid `json:"SessionService"`

	// The AccountService schema defines an Account Service.
	// The properties are common to, and enable management of, all user accounts.
	// The properties include the password requirements and control features, such as account lockout.
	// The schema also contains links to the manager accounts and roles. [RO]
	AccountService CommonOid `json:"AccountService"`

	// The EventService schema contains properties for managing event subcriptions and
	// generates the events sent to subscribers.
	// The Resource has links to the actual collection of subscriptions,
	// which are called event destinations. [RO]
	EventService CommonOid `json:"EventService"`

	// The collection of MessageRegistryFile Resource instances. [RO]
	Registeries CommonOid `json:"Registeries"`

	// The UpdateService schema describes the Update Service and the properties for the Service itself with
	// links to collections of firmware and software inventory.
	// The Update Service also provides methods for
	// updating software and firmware of the Resources in a Redfish Service. [RO]
	UpdateService CommonOid `json:"UpdateService"`

	// The CertificateService schema describes a Certificate Service that
	// represents the actions available to manage certificates and
	// links to the certificates. [RO]
	CertificateService CommonOid `json:"CertificateService"`

	// The links to other Resources that are related to this Resource. [RW]
	Link RootLink `json:"Link"`

	// The OEM extension. [RW]
	Oem RootOem `json:"Oem,omitempty"`

	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`
}

type RootLink struct {
	// The SessionCollection schema describes a collection of session instances. [RO]
	Sessions CommonOid `json:"Sessions"`
}

type RootOem struct {
}
