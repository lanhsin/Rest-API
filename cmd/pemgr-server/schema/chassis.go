package schema

type Chassis struct {
	// The type of a resource. [RO]
	OdataType string `json:"@odata.type"`

	// The identifier that uniquely identifies the Resource within
	// the collection of similar Resources. [RO]
	Id string `json:"Id"`

	// The name of the Resource or array member. [RO]
	Name string `json:"Name"`

	// The type of physical form factor of the chassis. [RO]
	// Valid values:
	// Blade:	An enclosed or semi-enclosed, typically vertically-oriented, system chassis that must be plugged into a multi-system chassis to function normally.
	// Card:	        A loose device or circuit board intended to be installed in a system or other enclosure.
	// Cartridge:	    A small self-contained system intended to be plugged into a multi-system chassis.
	// Component:	    A small chassis, card, or device that contains devices for a particular subsystem or function.
	// Drawer:	        An enclosed or semi-enclosed, typically horizontally-oriented, system chassis that may be slid into a multi-system chassis.
	// Enclosure:	    A generic term for a chassis that does not fit any other description.
	// Expansion:	    A chassis that expands the capabilities or capacity of another chassis.
	// IPBasedDrive:	A chassis in a drive form factor with IP-based network connections.
	// Module:	        A small, typically removable, chassis or card that contains devices for a particular subsystem or function.
	// Other:	        A chassis that does not fit any of these definitions.
	// Pod:	            A collection of equipment racks in a large, likely transportable, container.
	// Rack:	        An equipment rack, typically a 19-inch wide freestanding unit.
	// RackGroup:	    A group of racks that form a single entity or share infrastructure.
	// RackMount:	    A single-system chassis designed specifically for mounting in an equipment rack.
	// Row:	            A collection of equipment racks.
	// Shelf:	        An enclosed or semi-enclosed, typically horizontally-oriented, system chassis that must be plugged into a multi-system chassis to function normally.
	// Sidecar:	        A chassis that mates mechanically with another chassis to expand its capabilities or capacity.
	// Sled:	        An enclosed or semi-enclosed, system chassis that must be plugged into a multi-system chassis to function normally similar to a blade type chassis.
	// StandAlone:	    A single, free-standing system, commonly called a tower or desktop chassis.
	// StorageEnclosure:A chassis that encloses storage.
	// Zone:	        A logical division or portion of a physical chassis that contains multiple devices or systems that cannot be physically separated.
	ChassisType string `json:"ChassisType"`

	// The user-assigned asset tag of this chassis. [RO]
	AssetTag string `json:"AssetTag,omitempty"`

	// The manufacturer of this chassis. [RO]
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The model number of the chassis. [RO]
	Model string `json:"Model,omitempty"`

	// The SKU of the chassis. [RO]
	SKU string `json:"SKU,omitempty"`

	// The serial number of the chassis. [RO]
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The part number of the chassis. [RO]
	PartNumber string `json:"PartNumber,omitempty"`

	// The current power state of the chassis. [RO]
	// Valid values:
	// Off:	        The components within the chassis have no power, except some components may continue to have AUX power, such as the management controller.
	// On:	        The components within the chassis have power.
	// PoweringOff:	A temporary state between on and off. The components within the chassis can take time to process the power off action.
	// PoweringOn:	A temporary state between off and on. The components within the chassis can take time to process the power on action.
	PowerState string `json:"PowerState"`

	// The state of the indicator LED, which identifies the chassis. [RW]
	// Valid values:
	// Blinking:	The indicator LED is blinking.
	// Lit:	        The indicator LED is lit.
	// Off:	        The indicator LED is off.
	// Unknown:	    The state of the indicator LED cannot be determined.
	IndicatorLED string `json:"IndicatorLED"`

	// The height of the chassis. [RO]
	HeightMm float64 `json:"HeightMm,omitempty"`

	// The width of the chassis. [RO]
	WidthMm float64 `json:"WidthMm,omitempty"`

	// The depth of the chassis. [RO]
	DepthMm float64 `json:"DepthMm,omitempty"`

	// The weight of the chassis. [RO]
	WeightKg float64 `json:"WeightKg,omitempty"`

	// The location of a Resource. [RW]
	Location ChassisLocation `json:"Location,omitempty"`

	// The status and health of a Resource and its children. [RW]
	Status CommonStatus `json:"Status"`

	// The Thermal schema describes temperature monitoring and thermal management subsystems,
	// such as cooling fans, for a computer system or similiar devices contained within a chassis. [RO]
	Thermal CommonOid `json:"Thermal"`

	// The Power schema describes power metrics and represents the properties for power consumption and power limiting. [RO]
	Power CommonOid `json:"Power"`

	// The OEM extension. [RW]
	Oem ChassisOem `json:"Oem"`

	// The links to other Resources that are related to this Resource. [RW]
	Links ChassisLinks `json:"Links"`

	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`
}

type ChassisLocation struct {
	// The postal address for a Resource. [RW]
	PostalAddress ChassisLocationPostalAddress `json:"PostalAddress"`

	// The placement within the addressed location. [RW]
	Placement ChassisLocationPlacement `json:"Placement"`
}

type ChassisLocationPostalAddress struct {
	// The country. [RW]
	Country string `json:"Country,omitempty"`

	// A top-level subdivision within a country. [RW]
	Territory string `json:"Territory,omitempty"`

	// City, township, or shi (JP). [RW]
	City string `json:"City,omitempty"`

	// Street name. [RW]
	Street string `json:"Street,omitempty"`

	// The numeric portion of house number. [RW]
	HouseNumber int `json:"HouseNumber,omitempty"`

	// The name. [RW]
	Name string `json:"Name,omitempty"`

	// PostalCode [RW]
	PostalCode string `json:"PostalCode"`
}

type ChassisLocationPlacement struct {
	// The name of the row. [RW]
	Row string `json:"Row,omitempty"`

	// The name of a rack location within a row. [RW]
	Rack string `json:"Rack,omitempty"`

	// The type of rack unit in use. [RW]
	// Valid values:
	// EIA_310:	A rack unit that is equal to 1.75 in (44.45 mm).
	// OpenU:	A rack unit that is equal to 48 mm (1.89 in).
	RackOffsetUnits string `json:"RackOffsetUnits,omitempty"`

	// The vertical location of the item, in terms of RackOffsetUnits. [RW]
	RackOffset int `json:"RackOffset,omitempty"`
}

type ChassisLinks struct {
	// An array of links to the computer systems that this chassis directly and wholly contains. [RO]
	ComputerSystems []CommonOid

	// An array of links to the Managers responsible for managing this chassis. [RO]
	ManagedBy []CommonOid

	// An array of links to the managers located in this chassis. [RO]
	ManagersInChassis []CommonOid
}

type ChassisOem struct {
	Custom ChassisOemCustom `json:"Custom"`
}

type ChassisOemCustom struct {
	Misc CommonOid `json:"Misc"`
	// Led CommonOid `json:"Led"`
	// Button CommonOid `json:"Button"`
	// The set of LEDs for this chassis. [RW]
	// Leds []ThermalLeds

	// The set of Buttons for this chassis. [RW]
	// Buttons []ThermalButtons
}
