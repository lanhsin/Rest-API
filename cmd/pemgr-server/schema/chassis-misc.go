package schema

import (
	"errors"
	"strings"
)

type Misc struct {
	// The type of a resource. [RO]
	OdataType string `json:"@odata.type"`

	// The identifier that uniquely identifies the Resource within
	// the collection of similar Resources. [RO]
	Id string `json:"Id"`

	// The name of the Resource or array member. [RO]
	Name string `json:"Name"`

	// The set of LEDs for this chassis. [RW]
	Leds []MiscLeds

	// The set of Buttons for this chassis. [RW]
	Buttons []MiscButtons

	// The OEM extension. [RW]
	// Oem ThermalOem `json:"Oem"`

	// The redundancy information for the set of fans in this chassis. [RW]
	Redundancy []ThermalRedundancy

	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`
}

type MiscLeds struct {
	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`

	// The identifier for the member within the collection. [RO]
	MemberId string `json:"MemberId"`

	// The temperature sensor name. [RO]
	Name string `json:"Name,omitempty"`

	// The numerical identifier of the temperature sensor. [RO]
	SensorNumber int64 `json:"SensorNumber,omitempty"`

	// The area or device to which this temperature measurement applies. [RO]
	// Valid values:
	// ACInput:	                    An AC input.
	// ACMaintenanceBypassInput:	An AC maintenance bypass input.
	// ACOutput:	                An AC output.
	// ACStaticBypassInput:	        An AC static bypass input.
	// ACUtilityInput:	            An AC utility input.
	// ASIC:	                    An ASIC device, such as a networking chip or chipset component.
	// Accelerator:	                An accelerator.
	// Back:	                    The back of the chassis.
	// Backplane:	                A backplane within the chassis.
	// CPU:	                        A processor (CPU).
	// CPUSubsystem:	            The entire processor (CPU) subsystem.
	// Chassis:	                    The entire chassis.
	// ComputeBay:	                Within a compute bay.
	// CoolingSubsystem:	        The entire cooling, or air and liquid, subsystem.
	// DCBus:	                    A DC bus.
	// Exhaust:	                    The air exhaust point or points or region of the chassis.
	// ExpansionBay:	            Within an expansion bay.
	// FPGA:	                    An FPGA.
	// Fan:	                        A fan.
	// Front:	                    The front of the chassis.
	// GPU:	                        A graphics processor (GPU).
	// GPUSubsystem:	            The entire graphics processor (GPU) subsystem.
	// Intake:	                    The air intake point or points or region of the chassis.
	// LiquidInlet:	                The liquid inlet point of the chassis.
	// LiquidOutlet:	            The liquid outlet point of the chassis.
	// Lower:	                    The lower portion of the chassis.
	// Memory:	                    A memory device.
	// MemorySubsystem:	            The entire memory subsystem.
	// Motor:	                    A motor.
	// NetworkBay:	                Within a networking bay.
	// NetworkingDevice:	        A networking device.
	// PowerSubsystem:	            The entire power subsystem.
	// PowerSupply:	                A power supply.
	// PowerSupplyBay:	            Within a power supply bay.
	// Rectifier:	                A rectifier device.
	// Room:	                    The room.
	// StorageBay:	                Within a storage bay.
	// StorageDevice:	            A storage device.
	// SystemBoard:	                The system board (PCB).
	// Transformer:	                A transformer.
	// Upper:	                    The upper portion of the chassis.
	// VoltageRegulator:	        A voltage regulator device.
	PhysicalContext string `json:"PhysicalContext"`

	// The status and health of a Resource and its children. [RW]
	Status CommonStatus `json:"Status"`

	// The fan speed. [RO]
	Reading int64 `json:"Reading"`

	// The units in which the fan reading and thresholds are measured. [RO]
	// Valid values:
	// Percent:	The fan reading and thresholds are measured as a percentage.
	// RPM:	    The fan reading and thresholds are measured in rotations per minute.
	ReadingUnits string `json:"ReadingUnits"`

	// The value at which the reading is below normal range and fatal. [RO]
	LowerThresholdFatal int64 `json:"LowerThresholdFatal"`

	// Minimum value for this sensor. [RO]
	MinReadingRange int64 `json:"MinReadingRange"`

	// Maximum value for this sensor. [RO]
	MaxReadingRange int64 `json:"MaxReadingRange"`

	// The set of redundancy groups for this fan. [RW]
	Redundancy []CommonOid

	// The areas or devices to which this temperature applies. [RO]
	RelatedItem []CommonOid

	Oem LedOem `json:"Oem"`
}

type MiscButtons struct {
	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`

	// The identifier for the member within the collection. [RO]
	MemberId string `json:"MemberId"`

	// The temperature sensor name. [RO]
	Name string `json:"Name,omitempty"`

	// The numerical identifier of the temperature sensor. [RO]
	SensorNumber int64 `json:"SensorNumber,omitempty"`

	// The area or device to which this temperature measurement applies. [RO]
	// Valid values:
	// ACInput:	                    An AC input.
	// ACMaintenanceBypassInput:	An AC maintenance bypass input.
	// ACOutput:	                An AC output.
	// ACStaticBypassInput:	        An AC static bypass input.
	// ACUtilityInput:	            An AC utility input.
	// ASIC:	                    An ASIC device, such as a networking chip or chipset component.
	// Accelerator:	                An accelerator.
	// Back:	                    The back of the chassis.
	// Backplane:	                A backplane within the chassis.
	// CPU:	                        A processor (CPU).
	// CPUSubsystem:	            The entire processor (CPU) subsystem.
	// Chassis:	                    The entire chassis.
	// ComputeBay:	                Within a compute bay.
	// CoolingSubsystem:	        The entire cooling, or air and liquid, subsystem.
	// DCBus:	                    A DC bus.
	// Exhaust:	                    The air exhaust point or points or region of the chassis.
	// ExpansionBay:	            Within an expansion bay.
	// FPGA:	                    An FPGA.
	// Fan:	                        A fan.
	// Front:	                    The front of the chassis.
	// GPU:	                        A graphics processor (GPU).
	// GPUSubsystem:	            The entire graphics processor (GPU) subsystem.
	// Intake:	                    The air intake point or points or region of the chassis.
	// LiquidInlet:	                The liquid inlet point of the chassis.
	// LiquidOutlet:	            The liquid outlet point of the chassis.
	// Lower:	                    The lower portion of the chassis.
	// Memory:	                    A memory device.
	// MemorySubsystem:	            The entire memory subsystem.
	// Motor:	                    A motor.
	// NetworkBay:	                Within a networking bay.
	// NetworkingDevice:	        A networking device.
	// PowerSubsystem:	            The entire power subsystem.
	// PowerSupply:	                A power supply.
	// PowerSupplyBay:	            Within a power supply bay.
	// Rectifier:	                A rectifier device.
	// Room:	                    The room.
	// StorageBay:	                Within a storage bay.
	// StorageDevice:	            A storage device.
	// SystemBoard:	                The system board (PCB).
	// Transformer:	                A transformer.
	// Upper:	                    The upper portion of the chassis.
	// VoltageRegulator:	        A voltage regulator device.
	PhysicalContext string `json:"PhysicalContext"`

	// The status and health of a Resource and its children. [RW]
	Status CommonStatus `json:"Status"`

	// The fan speed. [RO]
	Reading int64 `json:"Reading"`

	// The units in which the fan reading and thresholds are measured. [RO]
	// Valid values:
	// Percent:	The fan reading and thresholds are measured as a percentage.
	// RPM:	    The fan reading and thresholds are measured in rotations per minute.
	ReadingUnits string `json:"ReadingUnits"`

	// The value at which the reading is below normal range and fatal. [RO]
	LowerThresholdFatal int64 `json:"LowerThresholdFatal"`

	// Minimum value for this sensor. [RO]
	MinReadingRange int64 `json:"MinReadingRange"`

	// Maximum value for this sensor. [RO]
	MaxReadingRange int64 `json:"MaxReadingRange"`

	// The set of redundancy groups for this fan. [RW]
	Redundancy []CommonOid

	// The areas or devices to which this temperature applies. [RO]
	RelatedItem []CommonOid
}

type LedOem struct {
	Custom LedOemCustom `json:"Custom"`
}

type LedOemCustom struct {
	// The action to control LED. [RW]
	// Valid values:
	// On:	Turn on LED.
	// Off: Turn off LED.
	Action string `json:"Action"`
}

func (led LedOemCustom) Validation() error {
	switch {
	case len(led.Action) == 0:
		return errors.New("The action is empty")
	case strings.ToLower(led.Action) == "on":
	case strings.ToLower(led.Action) == "off":
		break
	default:
		return errors.New("The action is invalied")
	}
	return nil
}
