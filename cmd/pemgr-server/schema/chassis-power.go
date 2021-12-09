package schema

type Power struct {
	// The type of a resource. [RO]
	OdataType string `json:"@odata.type"`

	// The identifier that uniquely identifies the Resource within
	// the collection of similar Resources. [RO]
	Id string `json:"Id"`

	// The name of the Resource or array member. [RO]
	Name string `json:"Name"`

	// The set of voltage sensors for this chassis. [RW]
	Voltages []PowerVoltage `json:"Voltages"`

	// The set of fans for this chassis. [RW]
	PowerSupplies []PowerPowerSupply `json:"PowerSupplies"`

	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`
}

type PowerVoltage struct {
	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`

	// The identifier for the member within the collection. [RO]
	MemberId string `json:"MemberId"`

	// The temperature sensor name. [RO]
	Name string `json:"Name,omitempty"`

	// The numerical identifier of the temperature sensor. [RO]
	SensorNumber int64 `json:"SensorNumber,omitempty"`

	// The status and health of a Resource and its children. [RW]
	Status CommonStatus `json:"Status"`

	// The reading of the voltage sensor. [RO]
	ReadingVolts float64 `json:"ReadingVolts"`

	// The value at which the reading is above normal range. [RO]
	UpperThresholdNonCritical float64 `json:"UpperThresholdNonCritical,omitempty"`

	// The value at which the reading is above normal range but not yet fatal. [RO]
	UpperThresholdCritical float64 `json:"UpperThresholdCritical,omitempty"`

	// The value at which the reading is above normal range and fatal. [RO]
	UpperThresholdFatal float64 `json:"UpperThresholdFatal,omitempty"`

	// The value at which the reading is below normal range. [RO]
	LowerThresholdNonCritical float64 `json:"LowerThresholdNonCritical,omitempty"`

	// The value at which the reading is below normal range but not yet fatal. [RO]
	LowerThresholdCritical float64 `json:"LowerThresholdCritical,omitempty"`

	// The value at which the reading is below normal range and fatal. [RO]
	LowerThresholdFatal float64 `json:"LowerThresholdFatal,omitempty"`

	// Minimum value for this sensor. [RO]
	MinReadingRangeTemp float64 `json:"MinReadingRangeTemp,omitempty"`

	// Maximum value for this sensor. [RO]
	MaxReadingRangeTemp float64 `json:"MaxReadingRangeTemp,omitempty"`

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

	// The areas or devices to which this temperature applies. [RO]
	RelatedItem []CommonOid
}

type PowerPowerSupply struct {
	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`

	// The identifier for the member within the collection. [RO]
	MemberId string `json:"MemberId"`

	// The temperature sensor name. [RO]
	Name string `json:"Name,omitempty"`

	// The numerical identifier of the temperature sensor. [RO]
	SensorNumber int64 `json:"SensorNumber,omitempty"`

	// The status and health of a Resource and its children. [RW]
	Status CommonStatus `json:"Status"`

	// The power supply type (AC or DC). [RO]
	PowerSupplyType string `json:"PowerSupplyType"`

	// The line voltage type supported as an input to this power supply. [RO]
	//Valid values:
	//AC120V	        AC 120V nominal input.
	//AC240V	        AC 240V nominal input.
	//AC277V	        AC 277V nominal input.
	//ACHighLine	    277V AC input.
	//ACLowLine	        100-127V AC input.
	//ACMidLine	        200-240V AC input.
	//ACWideRange	    Wide range AC input.
	//ACandDCWideRange	Wide range AC or DC input.
	//DC240V	        DC 240V nominal input.
	//DC380V	        High Voltage DC input (380V).
	//DCNeg48V	        -48V DC input.
	//Unknown	        The power supply line input voltage type cannot be determined.
	LineInputVoltageType string `json:"LineInputVoltageType"`

	// The line input voltage at which the power supply is operating. [RO]
	LineInputVoltage float64 `json:"LineInputVoltage"`

	// The maximum capacity of this power supply. [RO]
	PowerCapacityWatts float64 `json:"PowerCapacityWatts"`

	// The average power output of this power supply. [RO]
	LastPowerOutputWatts float64 `json:"LastPowerOutputWatts"`

	// The power input of this power supply. [RO]
	PowerInputWatts float64 `json:"PowerInputWatts"`

	// The power output of this power supply. [RO]
	PowerOutputWatts float64 `json:"PowerOutputWatts"`

	// The model number for this power supply. [RO]
	Model string `json:"Model"`

	// The manufacturer of this power supply. [RO]
	Manufacturer string `json:"Manufacturer"`

	// The firmware version for this power supply. [RO]
	FirmwareVersion string `json:"FirmwareVersion"`

	// The serial number for this power supply. [RO]
	SerialNumber string `json:"SerialNumber"`

	// The part number for this power supply. [RO]
	PartNumber string `json:"PartNumber"`

	// The spare part number for this power supply. [RO]
	SparePartNumber string `json:"SparePartNumber"`

	// The input ranges that the power supply can use. [RW]
	InputRanges []PowerSupplyInputRange `json:"InputRanges"`

	// The areas or devices to which this temperature applies. [RO]
	RelatedItem []CommonOid
}

type PowerSupplyInputRange struct {
	// The Input type (AC or DC). [RO]
	InputType string `json:"InputType"`

	// The minimum line input voltage at which this power supply input range is effective. [RO]
	MinimumVoltage float64 `json:"MinimumVoltage"`

	// The maximum line input voltage at which this power supply input range is effective. [RO]
	MaximumVoltage float64 `json:"MaximumVoltage"`

	// The minimum input frequency. [RO]
	MinimumFrequencyHz float64 `json:"MinimumFrequencyHz"`

	// The maximum input frequency. [RO]
	MaximumFrequencyHz float64 `json:"MaximumFrequencyHz"`

	// The maximum capacity of this power supply when operating in this input range. [RO]
	OutputWattage float64 `json:"OutputWattage"`
}
