package schema

type System struct {
	// The type of a resource. [RO]
	OdataType string `json:"@odata.type"`

	// The identifier that uniquely identifies the Resource within
	// the collection of similar Resources. [RO]
	Id string `json:"Id"`

	// The name of the Resource or array member. [RO]
	Name string `json:"Name"`

	// The type of computer system that this Resource represents. [RO]
	// Valid values:
	// Composed:	            A computer system constructed by binding Resource Blocks together.
	// OS:	                    An operating system instance.
	// Physical:	            A computer system.
	// PhysicallyPartitioned:	A hardware-based partition of a computer system.
	// Virtual:	                A virtual machine instance running on this system.
	// VirtuallyPartitioned:	A virtual or software-based partition of a computer system.
	SystemType string `json:"SystemType"`

	// The user-definable tag that can track this computer system for inventory or other client purposes. [RW]
	AssetTag string `json:"AssetTag,omitempty"`

	// The manufacturer or OEM of this system. [RO]
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The product name for this system, without the manufacturer name. [RO]
	Model string `json:"Model,omitempty"`

	// The sub-model for this system. [RO]
	SubModel string `json:"SubModel,omitempty"`

	// The manufacturer SKU for this system. [RO]
	SKU string `json:"SKU,omitempty"`

	// The serial number for this system. [RO]
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The part number for this system. [RO]
	PartNumber string `json:"PartNumber,omitempty"`

	// The description of this Resource. Used for commonality in the schema definitions. [RO]
	Description string `json:"Description"`

	// The UUID for this system. [RO]
	UUID string `json:"UUID"`

	// The DNS host name, without any domain information. [RW]
	HostName string `json:"HostName,omitempty"`

	// The status and health of a Resource and its children. [RW]
	Status SystemStatus `json:"Status"`

	// The hosting roles that this computer system supports. [RO]
	HostingRoles []string `json:"HostingRoles"`

	// The state of the indicator LED, which identifies the system. [RW]
	// Valid values:
	// Blinking:	The indicator LED is blinking.
	// Lit:	        The indicator LED is lit.
	// Off:	        The indicator LED is off.
	// Unknown:	    The state of the indicator LED cannot be determined.
	IndicatorLED string `json:"IndicatorLED"` // Blinking, Lit, Off, Unknown

	// The current power state of the system. [RO]
	// Valid values:
	// Off:	        The system is powered off, although some components may continue to have AUX power such as management controller.
	// On:	        The system is powered on.
	// PoweringOff:	A temporary state between on and off. The power off action can take time while the OS is in the shutdown process.
	// PoweringOn:	A temporary state between off and on. This temporary state can be very short.
	PowerState string `json:"PowerState"` // Off, On, PoweringOff, PoweringOn

	// The boot information for this Resource. [RW]
	Boot SystemBoot `json:"Boot"`

	// An array of trusted modules in the system. [RW]
	TrustedModules []SystemTrustedModules `json:"TrustedModules"`

	// The OEM extension. [RW]
	Oem SystemOem `json:"Oem,omitempty"`

	// The version of the system BIOS or primary system firmware. [RO]
	BiosVersion string `json:"BiosVersion,omitempty"`

	// The central processors of the system in general detail. [RW]
	ProcessorSummary SystemProcessorSummary `json:"ProcessorSummary"`

	// The Bios schema contains properties related to the BIOS Attribute Registry.
	// The Attribute Registry describes the system-specific BIOS attributes and actions for changing to BIOS settings.
	// Changes to the BIOS typically require a system reset before they take effect.
	// It is likely that a client finds the `@Redfish.Settings` term in this Resource, and if it is found,
	// the client makes requests to change BIOS settings by modifying the Resource identified by the `@Redfish.Settings` term. [RO]
	Bios CommonOid `json:"Bios"`

	// The ProcessorCollection contains a collection of processor instances. [RO]
	Processors CommonOid `json:"Processors"`

	// The collection of Memory Resource instances. [RO]
	Memory CommonOid `json:"Memory"`

	// The collection of EthernetInterface Resource instances. [RO]
	EthernetInterfaces CommonOid `json:"EthernetInterfaces"`

	// The SimpleStorageCollection schema contains a collection of simple storage instances. [RO]
	SimpleStorage CommonOid `json:"SimpleStorage"`

	// The LogServiceCollection schema describes a Resource Collection of LogService instances. [RO]
	LogServices CommonOid `json:"LogServices"`

	// The links to other Resources that are related to this Resource.
	Links SystemLinks `json:"Links"`

	// Actions

	// Irene OEM
	// Ssd CommonOid `json:"Ssd"`

	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`
}

type SystemStatus struct {
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
	// Critical:	A critical condition requires immediate attention.
	// OK:	        Normal.
	// Warning:	    A condition requires attention.
	Health string `json:"Health"`

	// The overall health state from the view of this Resource. [RO]
	// Valid values:
	// Critical:	A critical condition requires immediate attention.
	// OK:	        Normal.
	// Warning:	    A condition requires attention.
	HealthRollup string `json:"HealthRollup"`
}

type SystemBoot struct {
	// The state of the boot source override feature. [RW]
	// Valid values:
	// Continuous:	The system boots to the target specified in the BootSourceOverrideTarget property until this property is `Disabled`.
	// Disabled:	The system boots normally.
	// Once:	    On its next boot cycle, the system boots one time to the boot source override target.
	//              Then, the BootSourceOverrideEnabled value is reset to `Disabled`.
	BootSourceOverrideEnabled string `json:"BootSourceOverrideEnabled"`

	// The current boot source to use at the next boot instead of the normal boot device,
	// if BootSourceOverrideEnabled is `true`. [RW]
	// Valid values:
	// BiosSetup:	Boot to the BIOS setup utility.
	// Cd:	        Boot from the CD or DVD.
	// Diags:	    Boot to the manufacturer's diagnostics program.
	// Floppy:	    Boot from the floppy disk drive.
	// Hdd:	        Boot from a hard drive.
	// None:	    Boot from the normal boot device.
	// Pxe:	        Boot from the Pre-Boot EXecution (PXE) environment.
	// RemoteDrive:	Boot from a remote drive, such as an iSCSI target.
	// SDCard:	    Boot from an SD card.
	// UefiBootNext:Boot to the UEFI device that the BootNext property specifies.
	// UefiHttp:	Boot from a UEFI HTTP network location.
	// UefiShell:	Boot to the UEFI Shell.
	// UefiTarget:	Boot to the UEFI device specified in the UefiTargetBootSourceOverride property.
	// Usb:	Boot from a system BIOS-specified USB device.
	// Utilities:	Boot to the manufacturer's utilities program or programs.
	BootSourceOverrideTarget                       string   `json:"BootSourceOverrideTarget"`
	BootSourceOverrideTargetRedfishAllowableValues []string `json:"BootSourceOverrideTarget@Redfish.AllowableValues"`

	// The BIOS boot mode to use when the system boots from the BootSourceOverrideTarget boot source. [RW]
	// Valid values:
	// Legacy:	The system boots in non-UEFI boot mode to the boot source override target.
	// UEFI:	The system boots in UEFI boot mode to the boot source override target.
	BootSourceOverrideMode string `json:"BootSourceOverrideMode"` // Legacy, UEFI

	// The UEFI device path of the device from which to boot when BootSourceOverrideTarget is `UefiTarget`. [RW]
	UefiTargetBootSourceOverride string `json:"UefiTargetBootSourceOverride,omitempty"`
}

type SystemTrustedModules struct {
	// The firmware version of this Trusted Module. [RO]
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The interface type of the Trusted Module. [RO]
	// Valid values:
	// TCM1_0:	Trusted Cryptography Module (TCM) 1.0.
	// TPM1_2:	Trusted Platform Module (TPM) 1.2.
	// TPM2_0:	Trusted Platform Module (TPM) 2.0.
	InterfaceType string `json:"InterfaceType"`

	// The status and health of a Resource and its children. [RW]
	Status SystemTrustedModulesStatus `json:"Status"`
}

type SystemTrustedModulesStatus struct {
	// The known state of the Resource, such as, enabled. [RO]
	// Valid values:
	// Absent:	This function or Resource is either not present or detected.
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
	// Warning: A condition requires attention.
	Health string `json:"Health"`
}

type SystemOem struct {
	Custom SystemOemCustom `json:"Custom"`
}

type SystemOemCustom struct {
	Ssd CommonOid `json:"Ssd"`
}

type SystemProcessorSummary struct {
	// The number of physical processors in the system. [RO]
	Count int64 `json:"Count,omitempty"`

	// The processor model for the primary or majority of processors in this system. [RO]
	Model string `json:"Model,omitempty"`

	// The status and health of a Resource and its children. [RW]
	Status CommonStatus2
}

type SystemMemorySummary struct {
	// The total configured operating system-accessible memory (RAM), measured in GiB. [RO]
	TotalSystemMemoryGiB int64 `json:"TotalSystemMemoryGiB"`

	// The total configured, system-accessible persistent memory, measured in GiB. [RO]
	TotalSystemPersistentMemoryGiB int64 `json:"TotalSystemPersistentMemoryGiB"`

	// The ability and type of memory mirroring that this computer system supports.
	// Valid values:
	// DIMM	The system supports DIMM mirroring at the DIMM level. Individual DIMMs can be mirrored.
	// Hybrid	The system supports a hybrid mirroring at the system and DIMM levels. Individual DIMMs can be mirrored.
	// None	The system does not support DIMM mirroring.
	// System	The system supports DIMM mirroring at the system level. Individual DIMMs are not paired for mirroring in this mode.
	MemoryMirroring string `json:"MemoryMirroring"`

	// The status and health of a Resource and its children. [RW]
	Status CommonStatus2
}

type SystemLinks struct {
	// An array of links to the chassis that contains this system.
	Chassis []CommonOid

	// An array of links to the Managers responsible for managing this chassis. [RO]
	ManagedBy []CommonOid
}
