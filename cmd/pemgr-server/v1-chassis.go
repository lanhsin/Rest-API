package main

import (
	"net/http"
	"pemgr/cmd/pemgr-server/schema"
	db "pemgr/database"
	. "pemgr/logger"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func getChassis(chassisId string) schema.Chassis {
	var res schema.Chassis
	res.OdataType = "#Chassis.v1_11_0.Chassis"
	res.Id = chassisId
	res.Name = "EBOF System Chassis"
	res.Thermal = schema.CommonOid{OdataId: "redfish/v1/Chassis/" + chassisId + "/Thermal"}
	res.Power = schema.CommonOid{OdataId: "redfish/v1/Chassis/" + chassisId + "/Power"}
	res.Oem.Custom.Misc = schema.CommonOid{OdataId: "redfish/v1/Chassis/" + chassisId + "/Misc"}
	res.OdataId = "redfish/v1/Chassis/" + chassisId
	return res
}

func getTemperature(chassisId string, tempId string) (schema.ThermalTemperatures, error) {
	i, err := strconv.ParseInt(tempId, 10, 64)
	if err != nil {
		return schema.ThermalTemperatures{}, errors.WithStack(err)
	}

	format := &db.TemperatureFormat{}
	state := "Enabled"
	health := "OK"

	key := format.GetSensorNumberKey(i)
	sensorNumber, err := db.PemgrDB.GetInt64(key)
	if err != nil {
		return schema.ThermalTemperatures{}, errors.WithStack(err)
	}

	key = format.GetNameKey(i)
	name, err := db.PemgrDB.Get(key)
	if err != nil {
		return schema.ThermalTemperatures{}, errors.WithStack(err)
	}

	key = format.GetReadingKey(i)
	reading, err := db.PemgrDB.GetFloat64(key)
	if err != nil {
		return schema.ThermalTemperatures{}, errors.WithStack(err)
	}

	key = format.GetPhysicalContextKey(i)
	physicalContext, err := db.PemgrDB.Get(key)
	if err != nil {
		return schema.ThermalTemperatures{}, errors.WithStack(err)
	}

	if sensorNumber < 0 || name == "" || reading < 0 {
		state = "Absent"
		health = "Critical"
	}

	temp := schema.ThermalTemperatures{
		OdataId:         "/redfish/v1/Chassis/" + chassisId + "/Thermal#/Tempereatures/" + strconv.FormatInt(i, 10),
		MemberId:        strconv.FormatInt(i, 10),
		Name:            name,
		SensorNumber:    sensorNumber,
		Status:          schema.CommonStatus{State: state, Health: health},
		ReadingCelsius:  reading,
		PhysicalContext: physicalContext,
	}
	return temp, nil
}

func getTemperatures(chassisId string) ([]schema.ThermalTemperatures, error) {
	temperatures := make([]schema.ThermalTemperatures, 0)

	format := &db.TemperatureFormat{}
	key := format.GetNumberKey()
	num, err := db.PemgrDB.GetInt64(key)
	if err != nil {
		return temperatures, errors.WithStack(err)
	}

	for i := int64(0); i < num; i++ {
		temp, err := getTemperature(chassisId, strconv.FormatInt(i, 10))
		if err != nil {
			return temperatures, errors.WithStack(err)
		}
		temperatures = append(temperatures, temp)
	}
	return temperatures, nil
}

func getFan(chassisId string, fanId string) (schema.ThermalFans, error) {
	i, err := strconv.ParseInt(fanId, 10, 64)
	if err != nil {
		return schema.ThermalFans{}, errors.WithStack(err)
	}

	format := &db.FanFormat{}
	state := "Enabled"
	health := "OK"

	key := format.GetSensorNumberKey(i)
	sensorNumber, err := db.PemgrDB.GetInt64(key)
	if err != nil {
		if err.Error() != "redis: nil" {
			return schema.ThermalFans{}, errors.WithStack(err)
		}
	}

	key = format.GetNameKey(i)
	name, err := db.PemgrDB.Get(key)
	if err != nil {
		if err.Error() != "redis: nil" {
			return schema.ThermalFans{}, errors.WithStack(err)
		}
	}

	key = format.GetSpeedKey(i)
	speed, err := db.PemgrDB.GetInt64(key)
	if err != nil {
		if err.Error() != "redis: nil" {
			return schema.ThermalFans{}, errors.WithStack(err)
		}
	}

	key = format.GetPhysicalContextKey(i)
	physicalContext, err := db.PemgrDB.Get(key)
	if err != nil {
		return schema.ThermalFans{}, errors.WithStack(err)
	}

	if sensorNumber < 0 || name == "" || speed < 0  {
		state = "Absent"
		health = "Critical"
	}

	temp := schema.ThermalFans{
		OdataId:         "/redfish/v1/Chassis/" + chassisId + "/Thermal#/Fans/" + strconv.FormatInt(i, 10),
		MemberId:        strconv.FormatInt(i, 10),
		Name:            name,
		SensorNumber:    sensorNumber,
		Status:          schema.CommonStatus{State: state, Health: health},
		Reading:         speed,
		ReadingUnits:    "Percentage",
		PhysicalContext: physicalContext,
		Oem:             schema.FanOem{Custom: schema.FanOemCustom{Duty: strconv.FormatInt(speed, 10)}},
	}
	return temp, err
}

func getFans(chassisId string) ([]schema.ThermalFans, error) {
	fans := make([]schema.ThermalFans, 0)

	format := &db.FanFormat{}
	key := format.GetNumberKey()
	num, err := db.PemgrDB.GetInt64(key)
	if err != nil {
		return fans, errors.WithStack(err)
	}

	for i := int64(0); i < num; i++ {
		temp, err := getFan(chassisId, strconv.FormatInt(i, 10))
		if err != nil {
			return fans, errors.WithStack(err)
		}
		fans = append(fans, temp)
	}
	return fans, nil
}

func setFan(chassisId string, fanId string, speed string) error {
	var (
		i, num int64
		err error
	)

	i, err = strconv.ParseInt(fanId, 10, 64)
	if err != nil {
		return errors.WithStack(err)
	}

	format := &db.FanFormat{}
	num, err = db.PemgrDB.GetInt64(format.GetNumberKey())
	if err != nil {
		return errors.WithStack(err)
	}
	if i >= num || i < 0 {
		return  errors.New("Unknown id")
	}
	key := format.GetApplySpeedKey(i)
	if err := db.PemgrDB.Set(key, speed); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func getLed(chassisId string, ledId string) (schema.MiscLeds, error) {
	i, err := strconv.ParseInt(ledId, 10, 64)
	if err != nil {
		if err.Error() != "redis: nil" {
			return schema.MiscLeds{}, errors.WithStack(err)
		}
	}

	format := &db.LedFormat{}
	state := "Enabled"
	health := "OK"

	key := format.GetSensorNumberKey(i)
	sensorNumber, err := db.PemgrDB.GetInt64(key)
	if err != nil {
		if err.Error() != "redis: nil" {
			return schema.MiscLeds{}, errors.WithStack(err)
		}
	}

	key = format.GetNameKey(i)
	name, err := db.PemgrDB.Get(key)
	if err != nil {
		if err.Error() != "redis: nil" {
			return schema.MiscLeds{}, errors.WithStack(err)
		}
	}

	key = format.GetReadingKey(i)
	reading, err := db.PemgrDB.GetFloat64(key)
	if err != nil {
		if err.Error() != "redis: nil" {
			return schema.MiscLeds{}, errors.WithStack(err)
		}
	}

	key = format.GetUnitKey(i)
	unit, err := db.PemgrDB.Get(key)
	if err != nil {
		if err.Error() != "redis: nil" {
			return schema.MiscLeds{}, errors.WithStack(err)
		}
	}

	key = format.GetActionKey(i)
	action, err := db.PemgrDB.Get(key)
	if err != nil {
		if err.Error() != "redis: nil" {
			return schema.MiscLeds{}, errors.WithStack(err)
		}
	}

	key = format.GetPhysicalContextKey(i)
	physicalContext, err := db.PemgrDB.Get(key)
	if err != nil {
		return schema.MiscLeds{}, errors.WithStack(err)
	}

	if sensorNumber < 0 || name == "" || reading < 0 {
		state = "Absent"
		health = "Critical"
	}

	temp := schema.MiscLeds{
		OdataId:         "/redfish/v1/Chassis/" + chassisId + "/Misc#/Leds/" + strconv.FormatInt(i, 10),
		MemberId:        strconv.FormatInt(i, 10),
		Name:            name,
		SensorNumber:    sensorNumber,
		Status:          schema.CommonStatus{State: state, Health: health},
		Reading:         int64(reading),
		ReadingUnits:    unit,
		PhysicalContext: physicalContext,
		Oem:             schema.LedOem{Custom: schema.LedOemCustom{Action: action}},
	}
	return temp, nil
}

func getLeds(chassisId string) ([]schema.MiscLeds, error) {
	leds := make([]schema.MiscLeds, 0)

	format := &db.LedFormat{}
	key := format.GetNumberKey()
	num, err := db.PemgrDB.GetInt64(key)
	if err != nil {
		return leds, errors.WithStack(err)
	}

	for i := int64(0); i < num; i++ {
		temp, err := getLed(chassisId, strconv.FormatInt(i, 10))
		if err != nil {
			return leds, errors.WithStack(err)
		}
		leds = append(leds, temp)
	}
	return leds, nil
}

func setLed(chassisId string, ledId string, action string) error {
	var (
		i, num int64
		err error
	)

	i, err = strconv.ParseInt(ledId, 10, 64)
	if err != nil {
		return errors.WithStack(err)
	}

	format := &db.LedFormat{}
	num, err = db.PemgrDB.GetInt64(format.GetNumberKey())
	if err != nil {
		return errors.WithStack(err)
	}
	if i >= num || i < 0 {
		return  errors.New("Unknown id")
	}
	key := format.GetActionKey(i)
	err = db.PemgrDB.Set(key, action)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func getButton(chassisId string, ledId string) (schema.MiscButtons, error) {
	i, err := strconv.ParseInt(ledId, 10, 64)
	if err != nil {
		return schema.MiscButtons{}, errors.WithStack(err)
	}

	format := &db.ButtonFormat{}
	state := "Enabled"
	health := "OK"

	key := format.GetSensorNumberKey(i)
	sensorNumber, err := db.PemgrDB.GetInt64(key)
	if err != nil {
		return schema.MiscButtons{}, errors.WithStack(err)
	}

	key = format.GetNameKey(i)
	name, err := db.PemgrDB.Get(key)
	if err != nil {
		return schema.MiscButtons{}, errors.WithStack(err)
	}

	key = format.GetReadingKey(i)
	reading, err := db.PemgrDB.GetFloat64(key)
	if err != nil {
		return schema.MiscButtons{}, errors.WithStack(err)
	}

	key = format.GetUnitKey(i)
	unit, err := db.PemgrDB.Get(key)
	if err != nil {
		return schema.MiscButtons{}, errors.WithStack(err)
	}

	key = format.GetPhysicalContextKey(i)
	physicalContext, err := db.PemgrDB.Get(key)
	if err != nil {
		return schema.MiscButtons{}, errors.WithStack(err)
	}

	if sensorNumber < 0 || name == "" || reading < 0 {
		state = "Absent"
		health = "Critical"
	}

	temp := schema.MiscButtons{
		OdataId:         "/redfish/v1/Chassis/" + chassisId + "/Misc#/Buttons/" + strconv.FormatInt(i, 10),
		MemberId:        strconv.FormatInt(i, 10),
		Name:            name,
		SensorNumber:    sensorNumber,
		Status:          schema.CommonStatus{State: state, Health: health},
		Reading:         int64(reading),
		ReadingUnits:    unit,
		PhysicalContext: physicalContext,
	}
	return temp, err
}

func getButtons(chassisId string) ([]schema.MiscButtons, error) {
	btns := make([]schema.MiscButtons, 0)

	format := &db.ButtonFormat{}
	key := format.GetNumberKey()
	num, err := db.PemgrDB.GetInt64(key)
	if err != nil {
		return btns, errors.WithStack(err)
	}

	for i := int64(0); i < num; i++ {
		temp, err := getButton(chassisId, strconv.FormatInt(i, 10))
		if err != nil {
			return btns, errors.WithStack(err)
		}
		btns = append(btns, temp)
	}
	return btns, nil
}

func getVoltage(chassisId string, voltId string) (schema.PowerVoltage, error) {
	i, err := strconv.ParseInt(voltId, 10, 64)
	if err != nil {
		return schema.PowerVoltage{}, errors.WithStack(err)
	}

	format := &db.VoltageFormat{}
	state := "Enabled"
	health := "OK"

	key := format.GetSensorNumberKey(i)
	sensorNumber, err := db.PemgrDB.GetInt64(key)
	if err != nil {
		return schema.PowerVoltage{}, errors.WithStack(err)
	}

	key = format.GetNameKey(i)
	name, err := db.PemgrDB.Get(key)
	if err != nil {
		return schema.PowerVoltage{}, errors.WithStack(err)
	}

	key = format.GetReadingKey(i)
	reading, err := db.PemgrDB.GetFloat64(key)
	if err != nil {
		return schema.PowerVoltage{}, errors.WithStack(err)
	}

	key = format.GetPhysicalContextKey(i)
	physicalContext, err := db.PemgrDB.Get(key)
	if err != nil {
		return schema.PowerVoltage{}, errors.WithStack(err)
	}

	if sensorNumber < 0 || name == "" || reading < 0 {
		state = "Absent"
		health = "Critical"
	}

	volt := schema.PowerVoltage{
		OdataId:         "/redfish/v1/Chassis/" + chassisId + "/Power#/Voltages/" + strconv.FormatInt(i, 10),
		MemberId:        strconv.FormatInt(i, 10),
		Name:            name,
		SensorNumber:    sensorNumber,
		Status:          schema.CommonStatus{State: state, Health: health},
		ReadingVolts:    reading,
		PhysicalContext: physicalContext,
	}
	return volt, nil
}

func getVoltages(chassisId string) ([]schema.PowerVoltage, error) {
	voltages := make([]schema.PowerVoltage, 0)

	format := &db.VoltageFormat{}
	key := format.GetNumberKey()
	num, err := db.PemgrDB.GetInt64(key)
	if err != nil {
		return voltages, errors.WithStack(err)
	}

	for i := int64(0); i < num; i++ {
		volt, err := getVoltage(chassisId, strconv.FormatInt(i, 10))
		if err != nil {
			return voltages, errors.WithStack(err)
		}
		voltages = append(voltages, volt)
	}
	return voltages, nil
}

func getPsu(chassisId string, psuId string) (schema.PowerPowerSupply, error) {
	i, err := strconv.ParseInt(psuId, 10, 64)
	if err != nil {
		return schema.PowerPowerSupply{}, errors.WithStack(err)
	}

	format := &db.PsuFormat{}
	state := "Enabled"
	health := "OK"

	key := format.GetNameKey(i)
	name, err := db.PemgrDB.Get(key)
	if err != nil {
		return schema.PowerPowerSupply{}, errors.WithStack(err)
	}

	key = format.GetPowerSupplyTypeKey(i)
	powerSupplyType, err := db.PemgrDB.Get(key)
	if err != nil {
		return schema.PowerPowerSupply{}, errors.WithStack(err)
	}

	key = format.GetLineInputVoltageTypeKey(i)
	lineInputVoltageType, err := db.PemgrDB.Get(key)
	if err != nil {
		return schema.PowerPowerSupply{}, errors.WithStack(err)
	}

	key = format.GetLineInputVoltageKey(i)
	lineInputVoltage, err := db.PemgrDB.GetFloat64(key)
	if err != nil {
		return schema.PowerPowerSupply{}, errors.WithStack(err)
	}

	key = format.GetPowerCapacityWattsKey(i)
	powerCapacityWatts, err := db.PemgrDB.GetFloat64(key)
	if err != nil {
		return schema.PowerPowerSupply{}, errors.WithStack(err)
	}

	key = format.GetLastPowerOutputWattsKey(i)
	lastPowerOutputWatts, err := db.PemgrDB.GetFloat64(key)
	if err != nil {
		return schema.PowerPowerSupply{}, errors.WithStack(err)
	}

	key = format.GetPowerInputWattsKey(i)
	powerInputWatts, err := db.PemgrDB.GetFloat64(key)
	if err != nil {
		return schema.PowerPowerSupply{}, errors.WithStack(err)
	}

	key = format.GetPowerOutputWattsKey(i)
	powerOutputWatts, err := db.PemgrDB.GetFloat64(key)
	if err != nil {
		return schema.PowerPowerSupply{}, errors.WithStack(err)
	}

	key = format.GetModelKey(i)
	model, err := db.PemgrDB.Get(key)
	if err != nil {
		return schema.PowerPowerSupply{}, errors.WithStack(err)
	}

	key = format.GetManufacturerKey(i)
	manu, err := db.PemgrDB.Get(key)
	if err != nil {
		return schema.PowerPowerSupply{}, errors.WithStack(err)
	}

	key = format.GetFirmwareVersionKey(i)
	version, err := db.PemgrDB.Get(key)
	if err != nil {
		return schema.PowerPowerSupply{}, errors.WithStack(err)
	}

	key = format.GetSerialNumberKey(i)
	sn, err := db.PemgrDB.Get(key)
	if err != nil {
		return schema.PowerPowerSupply{}, errors.WithStack(err)
	}

	key = format.GetPartNumberKey(i)
	pn, err := db.PemgrDB.Get(key)
	if err != nil {
		return schema.PowerPowerSupply{}, errors.WithStack(err)
	}

	key = format.GetSparePartNumberKey(i)
	spn, err := db.PemgrDB.Get(key)
	if err != nil {
		return schema.PowerPowerSupply{}, errors.WithStack(err)
	}

	// Get InputRanges array
	var inputRange schema.PowerSupplyInputRange
	inputRanges := make([]schema.PowerSupplyInputRange, 0)

	key = format.GetInputRangesNumberKey(i)
	number, err := db.PemgrDB.GetInt64(key)
	if err != nil {
		Lig.Error().Msgf("%v", err)
		number = 0
	}

	for r := int64(0); r < number; r++ {
		key = format.GetInputRangesInputTypeKey(i, r)
		inputRange.InputType, err = db.PemgrDB.Get(key)
		if err != nil {
			return schema.PowerPowerSupply{}, errors.WithStack(err)
		}

		key = format.GetInputRangesMinimumVoltageKey(i, r)
		inputRange.MinimumVoltage, err = db.PemgrDB.GetFloat64(key)
		if err != nil {
			return schema.PowerPowerSupply{}, errors.WithStack(err)
		}

		key = format.GetInputRangesMaximumVoltageKey(i, r)
		inputRange.MaximumVoltage, err = db.PemgrDB.GetFloat64(key)
		if err != nil {
			return schema.PowerPowerSupply{}, errors.WithStack(err)
		}

		key = format.GetInputRangesMinimumFrequencyHzKey(i, r)
		inputRange.MinimumFrequencyHz, err = db.PemgrDB.GetFloat64(key)
		if err != nil {
			return schema.PowerPowerSupply{}, errors.WithStack(err)
		}

		key = format.GetInputRangesMaximumFrequencyHzKey(i, r)
		inputRange.MaximumFrequencyHz, err = db.PemgrDB.GetFloat64(key)
		if err != nil {
			return schema.PowerPowerSupply{}, errors.WithStack(err)
		}

		key = format.GetInputRangesOutputWattageKey(i, r)
		inputRange.OutputWattage, err = db.PemgrDB.GetFloat64(key)
		if err != nil {
			return schema.PowerPowerSupply{}, errors.WithStack(err)
		}

		inputRanges = append(inputRanges, inputRange)
	}

	if name == "" {
		state = "Absent"
		health = "Critical"
	}

	psu := schema.PowerPowerSupply{
		OdataId:              "/redfish/v1/Chassis/" + chassisId + "/Power#/PowerSupplies/" + strconv.FormatInt(i, 10),
		MemberId:             strconv.FormatInt(i, 10),
		Name:                 name,
		Status:               schema.CommonStatus{State: state, Health: health},
		PowerSupplyType:      powerSupplyType,
		LineInputVoltageType: lineInputVoltageType,
		LineInputVoltage:     lineInputVoltage,
		PowerCapacityWatts:   powerCapacityWatts,
		LastPowerOutputWatts: lastPowerOutputWatts,
		PowerInputWatts:      powerInputWatts,
		PowerOutputWatts:     powerOutputWatts,
		Model:                model,
		Manufacturer:         manu,
		FirmwareVersion:      version,
		SerialNumber:         sn,
		PartNumber:           pn,
		SparePartNumber:      spn,
		InputRanges:          inputRanges,
	}
	return psu, nil
}

func getPsus(chassisId string) ([]schema.PowerPowerSupply, error) {
	psus := make([]schema.PowerPowerSupply, 0)

	format := &db.PsuFormat{}
	key := format.GetNumberKey()
	num, err := db.PemgrDB.GetInt64(key)
	if err != nil {
		return psus, errors.WithStack(err)
	}

	for i := int64(0); i < num; i++ {
		psu, err := getPsu(chassisId, strconv.FormatInt(i, 10))
		if err != nil {
			return psus, errors.WithStack(err)
		}
		psus = append(psus, psu)
	}
	return psus, nil
}

func getThermal(chassisId string) (schema.Thermal, error) {
	var res schema.Thermal
	res.OdataType = "#Thermal.v1_6_0.Thermal"
	res.Id = "Thermal"
	res.Name = "Thermal"

	temperatures, err := getTemperatures(chassisId)
	if err != nil {
		return res, errors.WithStack(err)
	}

	fans, err := getFans(chassisId)
	if err != nil {
		return res, err
	}

	res.Temperatures = temperatures
	res.Fans = fans
	res.OdataId = "redfish/v1/Chassis/" + chassisId + "/Thermal"
	return res, nil
}

func getMisc(chassisId string) (schema.Misc, error) {
	var res schema.Misc
	res.OdataType = "#Misc.v0_0_0.Misc"
	res.Id = "Misc"
	res.Name = "Misc"

	leds, err := getLeds(chassisId)
	if err != nil {
		return res, errors.WithStack(err)
	}

	buttons, err := getButtons(chassisId)
	if err != nil {
		return res, errors.WithStack(err)
	}

	res.Leds = leds
	res.Buttons = buttons
	res.OdataId = "redfish/v1/Chassis/" + chassisId + "/Thermal"
	return res, nil
}

func getPower(chassisId string) (schema.Power, error) {
	var res schema.Power
	res.OdataType = "#Power.v1_6_0.Power"
	res.Id = "Power"
	res.Name = "Power"

	voltages, err := getVoltages(chassisId)
	if err != nil {
		return res, errors.WithStack(err)
	}

	psus, err := getPsus(chassisId)
	if err != nil {
		return res, err
	}

	res.Voltages = voltages
	res.PowerSupplies = psus
	res.OdataId = "redfish/v1/Chassis/" + chassisId + "/Power"
	return res, nil
}

// handleChassis godoc
// @Summary Retrieve list of physical components
// @Tags Chassis
// @Accept json
// @Produce json
// @Success 200 {object} schema.ChassisCollection
// @Failure 400 {object} schema.HttpError
// @Failure 404 {object} schema.HttpError
// @Failure 500 {object} schema.HttpError
// @Router /redfish/v1/Chassis [get]
func handleCasssis(c *gin.Context) {
	var res schema.ChassisCollection

	members := make([]schema.CommonOid, 0)
	for i := 0; i < len(template.ChassisIds); i++ {
		members = append(members, schema.CommonOid{OdataId: "/redfish/v1/Chassis/" + template.ChassisIds[i]})
	}

	res.OdataType = "#ChassisCollection.ChassisCollection"
	res.Name = "Chassis Collection"
	res.MemberOdataCount = 1
	res.Members = members
	res.OdataId = "redfish/v1/Chassis"

	c.JSON(http.StatusOK, res)
}

// handleChassisId godoc
// @Summary Retrieve list of physical components
// @Tags Chassis
// @Accept json
// @Produce json
// @Param  chassisId path string true "Chassis ID (Default = 1)" default(1)
// @Success 200 {object} schema.Chassis
// @Failure 400 {object} schema.HttpError
// @Failure 404 {object} schema.HttpError
// @Failure 500 {object} schema.HttpError
// @Router /redfish/v1/Chassis/{chassisId} [get]
func handleChassisId(c *gin.Context) {
	found := false
	chassisId := c.Param("chassisId")

	for i := 0; i < len(template.ChassisIds); i++ {
		if chassisId == template.ChassisIds[i] {
			res := getChassis(chassisId)
			c.JSON(http.StatusOK, res)
			found = true
			break
		}
	}

	if found == false {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    -1,
			"message": "Not found chassis id:" + chassisId,
		})
	}
}

// handleThermal godoc
// @Summary Retrieve thermal characteristics of a chassis
// @Tags Chassis
// @Accept json
// @Produce json
// @Param  chassisId path string true "Chassis ID (Default = 1)" default(1)
// @Success 200 {object} schema.Thermal
// @Failure 400 {object} schema.HttpError
// @Failure 404 {object} schema.HttpError
// @Failure 500 {object} schema.HttpError
// @Router /redfish/v1/Chassis/{chassisId}/Thermal [get]
func handleThermal(c *gin.Context) {
	found := false
	chassisId := c.Param("chassisId")

	for i := 0; i < len(template.ChassisIds); i++ {
		if chassisId == template.ChassisIds[i] {
			res, err := getThermal(chassisId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    -1,
					"message": err.Error(),
				})
				Lig.Error().Msgf("%+v", err)
				return
			}
			c.JSON(http.StatusOK, res)
			found = true
			break
		}
	}

	if found == false {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    -1,
			"message": "Not found chassis id:" + chassisId,
		})
	}
}

// handleTemperatures godoc
// @Summary Retrieve temperature devices of a chassis
// @Tags Chassis
// @Accept json
// @Produce json
// @Param  chassisId path string true "Chassis ID (Default = 1)" default(1)
// @Success 200 {object} schema.ThermalTemperatures
// @Failure 400 {object} schema.HttpError
// @Failure 404 {object} schema.HttpError
// @Failure 500 {object} schema.HttpError
// @Router /redfish/v1/Chassis/{chassisId}/Thermal/Temperatures [get]
func handleTemperatures(c *gin.Context) {
	found := false
	chassisId := c.Param("chassisId")

	for i := 0; i < len(template.ChassisIds); i++ {
		if chassisId == template.ChassisIds[i] {
			temperatures, err := getTemperatures(chassisId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    -1,
					"message": err.Error(),
				})
				Lig.Error().Msgf("%+v", err)
				return
			}
			c.JSON(http.StatusOK, temperatures)
			found = true
			break
		}
	}

	if found == false {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    -1,
			"message": "Not found chassis id:" + chassisId,
		})
	}
}

// handleFans godoc
// @Summary Retrieve fan devices of a chassis
// @Tags Chassis
// @Accept json
// @Produce json
// @Param  chassisId path string true "Chassis ID (Default = 1)" default(1)
// @Success 200 {array} schema.ThermalFans
// @Failure 400 {object} schema.HttpError
// @Failure 404 {object} schema.HttpError
// @Failure 500 {object} schema.HttpError
// @Router /redfish/v1/Chassis/{chassisId}/Thermal/Fans [get]
func handleFans(c *gin.Context) {
	found := false
	chassisId := c.Param("chassisId")

	for i := 0; i < len(template.ChassisIds); i++ {
		if chassisId == template.ChassisIds[i] {
			fans, err := getFans(chassisId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    -1,
					"message": err.Error(),
				})
				Lig.Error().Msgf("%+v", err)
				return
			}
			c.JSON(http.StatusOK, fans)
			found = true
			break
		}
	}

	if found == false {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    -1,
			"message": "Not found chassis id:" + chassisId,
		})
	}
}

// handleFanId godoc
// @Summary Retrieve a specific fan device
// @Tags Chassis
// @Accept json
// @Produce json
// @Param  chassisId path string true "Chassis ID (Default = 1)" default(1)
// @Param  fanId path string true "Fan ID" default(0)
// @Success 200 {object} schema.ThermalFans
// @Failure 400 {object} schema.HttpError
// @Failure 404 {object} schema.HttpError
// @Failure 500 {object} schema.HttpError
// @Router /redfish/v1/Chassis/{chassisId}/Thermal/Fans/{fanId} [get]
func handleFanId(c *gin.Context) {
	found := false
	chassisId := c.Param("chassisId")
	fanId := c.Param("fanId")

	for i := 0; i < len(template.ChassisIds); i++ {
		if chassisId == template.ChassisIds[i] {
			res, err := getFan(chassisId, fanId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    -1,
					"message": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, res)
			found = true
			break
		}
	}

	if found == false {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    -1,
			"message": "Not found chassis id:" + chassisId,
		})
	}
}

// handleFanIdPatch godoc
// @Summary Update duty cycle of a fan device
// @Tags Chassis
// @Accept json
// @Produce json
// @Param  chassisId path string true "Chassis ID (Default = 1)" default(1)
// @Param  fanId path string true "Fan ID" default(0)
// @Param  Oem body schema.FanOemCustom true "Update duty cycles"
// @Success 200 {object} schema.FanOemCustom
// @Failure 400 {object} schema.HttpError
// @Failure 404 {object} schema.HttpError
// @Failure 500 {object} schema.HttpError
// @Router /redfish/v1/Chassis/{chassisId}/Thermal/Fans/{fanId} [patch]
func handleFanIdPatch(c *gin.Context) {
	chassisId := c.Param("chassisId")
	fanId := c.Param("fanId")
	fanOem := schema.FanOemCustom{}
	if err := c.ShouldBindJSON(&fanOem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	if err := fanOem.Validation(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	if err := setFan(chassisId, fanId, fanOem.Duty); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, fanOem)
}

// handleMisc godoc
// @Summary Retrieve miscellaneous OEM devices of a chassis
// @Tags Chassis - Custom OEM
// @Accept json
// @Produce json
// @Param  chassisId path string true "Chassis ID (Default = 1)" default(1)
// @Success 200 {object} schema.Misc
// @Failure 400 {object} schema.HttpError
// @Failure 404 {object} schema.HttpError
// @Failure 500 {object} schema.HttpError
// @Router /redfish/v1/Chassis/{chassisId}/Misc [get]
func handleMisc(c *gin.Context) {
	found := false
	chassisId := c.Param("chassisId")

	for i := 0; i < len(template.ChassisIds); i++ {
		if chassisId == template.ChassisIds[i] {
			res, err := getMisc(chassisId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    -1,
					"message": err.Error(),
				})
				Lig.Error().Msgf("%+v", err)
				return
			}
			c.JSON(http.StatusOK, res)
			found = true
			break
		}
	}

	if found == false {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    -1,
			"message": "Not found chassis id:" + chassisId,
		})
	}
}

// handleLeds godoc
// @Summary Retrieve LED devices of a chassis
// @Tags Chassis - Custom OEM
// @Accept json
// @Produce json
// @Param  chassisId path string true "Chassis ID (Default = 1)" default(1)
// @Success 200 {array} schema.MiscLeds
// @Failure 400 {object} schema.HttpError
// @Failure 404 {object} schema.HttpError
// @Failure 500 {object} schema.HttpError
// @Router /redfish/v1/Chassis/{chassisId}/Misc/Leds [get]
func handleLeds(c *gin.Context) {
	found := false
	chassisId := c.Param("chassisId")

	for i := 0; i < len(template.ChassisIds); i++ {
		if chassisId == template.ChassisIds[i] {
			res, err := getLeds(chassisId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    -1,
					"message": err.Error(),
				})
				Lig.Error().Msgf("%+v", err)
				return
			}
			c.JSON(http.StatusOK, res)
			found = true
			break
		}
	}

	if found == false {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    -1,
			"message": "Not found chassis id:" + chassisId,
		})
	}
}

// handleLedId godoc
// @Summary Retrieve a specific LED device
// @Tags Chassis - Custom OEM
// @Accept json
// @Produce json
// @Param  chassisId path string true "Chassis ID (Default = 1)" default(1)
// @Param  ledId path string true "Led ID" default(0)
// @Success 200 {object} schema.MiscLeds
// @Failure 400 {object} schema.HttpError
// @Failure 404 {object} schema.HttpError
// @Failure 500 {object} schema.HttpError
// @Router /redfish/v1/Chassis/{chassisId}/Misc/Leds/{ledId} [get]
func handleLedId(c *gin.Context) {
	found := false
	chassisId := c.Param("chassisId")
	ledId := c.Param("ledId")

	for i := 0; i < len(template.ChassisIds); i++ {
		if chassisId == template.ChassisIds[i] {
			res, err := getLed(chassisId, ledId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    -1,
					"message": err.Error(),
				})
				c.String(http.StatusInternalServerError, "Fail to get fans, chassisId= %s", chassisId)
				return
			}
			c.JSON(http.StatusOK, res)
			found = true
			break
		}
	}

	if found == false {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    -1,
			"message": "Not found chassis id:" + chassisId,
		})
	}
}

// handleLedIdPatch godoc
// @Summary Update status of a LED device
// @Tags Chassis - Custom OEM
// @Accept json
// @Produce json
// @Param  chassisId path string true "Chassis ID (Default = 1)" default(1)
// @Param  ledId path string true "LED ID" default(0)
// @Param  Oem body schema.LedOemCustom true "Update LED status (on/off)"
// @Success 200 {object} schema.LedOemCustom
// @Failure 400 {object} schema.HttpError
// @Failure 404 {object} schema.HttpError
// @Failure 500 {object} schema.HttpError
// @Router /redfish/v1/Chassis/{chassisId}/Misc/Leds/{ledId} [patch]
func handleLedIdPatch(c *gin.Context) {
	chassisId := c.Param("chassisId")
	ledId := c.Param("ledId")
	ledOem := schema.LedOemCustom{}
	if err := c.ShouldBindJSON(&ledOem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	if err := ledOem.Validation(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	if err := setLed(chassisId, ledId, ledOem.Action); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ledOem)
}

// handleButtons godoc
// @Summary Retrieve buttons of a chassis
// @Tags Chassis - Custom OEM
// @Accept json
// @Produce json
// @Param  chassisId path string true "Chassis ID (Default = 1)" default(1)
// @Success 200 {object} schema.MiscButtons
// @Failure 400 {object} schema.HttpError
// @Failure 404 {object} schema.HttpError
// @Failure 500 {object} schema.HttpError
// @Router /redfish/v1/Chassis/{chassisId}/Misc/Buttons [get]
func handleButtons(c *gin.Context) {
	found := false
	chassisId := c.Param("chassisId")

	for i := 0; i < len(template.ChassisIds); i++ {
		if chassisId == template.ChassisIds[i] {
			res, err := getButtons(chassisId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    -1,
					"message": err.Error(),
				})
				Lig.Error().Msgf("%+v", err)
				return
			}
			c.JSON(http.StatusOK, res)
			found = true
			break
		}
	}

	if found == false {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    -1,
			"message": "Not found chassis id:" + chassisId,
		})
	}
}

// handlePower godoc
// @Summary Retrieve power characteristics of a chassis
// @Tags Chassis
// @Accept json
// @Produce json
// @Param  chassisId path string true "Chassis ID (Default = 1)" default(1)
// @Success 200 {object} schema.Power
// @Failure 400 {object} schema.HttpError
// @Failure 404 {object} schema.HttpError
// @Failure 500 {object} schema.HttpError
// @Router /redfish/v1/Chassis/{chassisId}/Power [get]
func handlePower(c *gin.Context) {
	found := false
	chassisId := c.Param("chassisId")

	for i := 0; i < len(template.ChassisIds); i++ {
		if chassisId == template.ChassisIds[i] {
			res, err := getPower(chassisId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    -1,
					"message": err.Error(),
				})
				Lig.Error().Msgf("%+v", err)
				return
			}
			c.JSON(http.StatusOK, res)
			found = true
			break
		}
	}

	if found == false {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    -1,
			"message": "Not found chassis id:" + chassisId,
		})
	}
}
