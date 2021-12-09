package main

import (
	"fmt"
	"net/http"
	"pemgr/cmd/pemgr-server/schema"
	db "pemgr/database"
	. "pemgr/logger"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func getSystems(systemId string) schema.System {
	var res schema.System
	res.OdataType = "#ComputerSystemCollection.ComputerSystemCollection"
	res.Id = systemId
	res.Name = "EBOF"
	res.Status.State = "Enabled"
	res.Status.Health = "OK"
	res.Status.HealthRollup = "OK"
	res.Oem.Custom.Ssd = schema.CommonOid{OdataId: "redfish/v1/Systems/" + systemId + "/Ssd"}
	res.OdataId = "/redfish/v1/Systems/" + systemId
	return res
}

func getSsdCollection(systemId string) (schema.SsdCollection, error) {
	format := db.SsdFormat{}

	key := format.GetNumberKey()
	num, err := db.PemgrDB.GetInt64(key)
	if err != nil {
		fmt.Printf("Fail to get DB\n")
		return schema.SsdCollection{}, errors.WithStack(err)
	}

	var res schema.SsdCollection
	res.OdataType = "#SsdCollection.SsdCollection"
	res.Name = "SSDs Collection"
	res.MemberOdataCount = num
	members := make([]schema.CommonOid, 0)
	for i := int64(0); i < num; i++ {
		members = append(members, schema.CommonOid{OdataId: "/redfish/v1/Systems/" + systemId + "/Ssd/" + strconv.FormatInt(i, 10)})
	}
	res.Members = members
	res.OdataId = "/redfish/v1/Systems/" + systemId + "/Ssd"
	return res, nil
}

func getSsdId(systemId string, id int64) schema.Ssd {
	format := db.SsdFormat{}
	state := "Enabled"
	health := "OK"

	key := format.GetNameKey(id)
	name, err := db.PemgrDB.Get(key)
	if err != nil {
		Lig.Error().Str("key", key).Msgf("%v", err)
		name = ""
	}

	key = format.GetManufacturerKey(id)
	manufacturer, err := db.PemgrDB.Get(key)
	if err != nil {
		Lig.Error().Str("key", key).Msgf("%v", err)
		manufacturer = ""
	}

	key = format.GetModelKey(id)
	model, err := db.PemgrDB.Get(key)
	if err != nil {
		Lig.Error().Str("key", key).Msgf("%v", err)
		model = ""
	}

	key = format.GetSerialNumberKey(id)
	sn, err := db.PemgrDB.Get(key)
	if err != nil {
		Lig.Error().Str("key", key).Msgf("%v", err)
		sn = ""
	}

	key = format.GetTemperatureKey(id)
	temperature, err := db.PemgrDB.GetFloat64(key)
	if err != nil {
		Lig.Error().Str("key", key).Msgf("%v", err)
		temperature = -1
	}

	if manufacturer == "" || model == "" || sn == "" || temperature < 0 {
		state = "Absent"
		health = "Critical"
	}

	var res schema.Ssd
	res.OdataType = "#SsdCollection.SsdCollection"
	res.Id = strconv.FormatInt(id, 10)
	res.Name = name
	res.Manufacturer = manufacturer
	res.Model = model
	res.SerialNumber = sn
	res.ReadingCelsius = temperature
	//res.Status.State = state
	//res.Status.Health = health
	res.Status = schema.CommonStatus{State: state, Health: health}
	res.OdataId = "/redfish/v1/Systems/" + systemId + "/Ssd/" + strconv.FormatInt(id, 10)
	Lig.Info().Msgf("res=%v", res)
	return res
}

func setSsd(systemId string, ssdId string, action string) error {
	var (
		i, num int64
		err error
	)

	i, err = strconv.ParseInt(ssdId, 10, 64)
	if err != nil {
		return errors.WithStack(err)
	}

	format := &db.SsdFormat{}
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

// handleSystems godoc
// @Summary Retrieve list of computer systems
// @Tags Systems
// @Accept json
// @Produce json
// @Success 200 {object} schema.SystemCollection
// @Failure 400 {object} schema.HttpError
// @Failure 404 {object} schema.HttpError
// @Failure 500 {object} schema.HttpError
// @Router /redfish/v1/Systems [get]
func handleSystems(c *gin.Context) {
	var res schema.SystemCollection

	members := make([]schema.CommonOid, 0)
	for i := 0; i < len(template.SystemsIds); i++ {
		members = append(members, schema.CommonOid{OdataId: "/redfish/v1/Systems/" + template.SystemsIds[i]})
	}

	res.OdataType = "#ComputerSystemCollection.ComputerSystemCollection"
	res.Name = "Computer System Collection"
	res.MemberOdataCount = 1
	res.Members = members
	res.OdataId = "/redfish/v1/Systems"

	c.JSON(http.StatusOK, res)
}

// handleSystemId godoc
// @Summary Retrieve information for a specific computer system
// @Tags Systems
// @Accept json
// @Produce json
// @Param  systemsId path string true "System ID (Default = Irene)" default(Irene)
// @Success 200 {object} schema.System
// @Failure 400 {object} schema.HttpError
// @Failure 404 {object} schema.HttpError
// @Failure 500 {object} schema.HttpError
// @Router /redfish/v1/Systems/{systemsId} [get]
func handleSystemId(c *gin.Context) {
	found := false
	systemId := c.Param("systemId")

	for i := 0; i < len(template.SystemsIds); i++ {
		if systemId == template.SystemsIds[i] {
			res := getSystems(systemId)
			c.JSON(http.StatusOK, res)
			found = true
			break
		}
	}

	if found == false {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    -1,
			"message": "Not found system id:" + systemId,
		})
	}
}

// handleSsd godoc
// @Summary Retrieve SSD devices of a chassis
// @Tags Systems - Custom OEM
// @Accept json
// @Produce json
// @Param  systemsId path string true "System ID (Default = Irene)" default(Irene)
// @Success 200 {object} schema.SsdCollection
// @Failure 400 {object} schema.HttpError
// @Failure 404 {object} schema.HttpError
// @Failure 500 {object} schema.HttpError
// @Router /redfish/v1/Systems/{systemsId}/Ssd [get]
func handleSsd(c *gin.Context) {
	found := false
	systemId := c.Param("systemId")

	for i := 0; i < len(template.SystemsIds); i++ {
		if systemId == template.SystemsIds[i] {
			res, err := getSsdCollection(systemId)
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
			"message": "Not found system id:" + systemId,
		})
	}
}

// handleSsdId godoc
// @Summary Retrieve a specific SSD device
// @Tags Systems - Custom OEM
// @Accept json
// @Produce json
// @Param  systemsId path string true "System ID (Default = Irene)" default(Irene)
// @Param  ssdId path int true "Ssd ID" default(0)
// @Success 200 {object} schema.Ssd
// @Failure 400 {object} schema.HttpError
// @Failure 404 {object} schema.HttpError
// @Failure 500 {object} schema.HttpError
// @Router /redfish/v1/Systems/{systemsId}/Ssd/{ssdId} [get]
func handleSsdId(c *gin.Context) {
	systemId := c.Param("systemId")
	ssdId := c.Param("ssdId")
	found := false
	format := db.SsdFormat{}
	key := format.GetNumberKey()
	num, err := db.PemgrDB.GetInt64(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		Lig.Error().Msgf("%+v", err)
		return
	}
	ssdIdInt, err := strconv.ParseInt(ssdId, 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		Lig.Error().Msgf("%+v", err)
		return
	}

	for i := 0; i < len(template.SystemsIds); i++ {
		if systemId == template.SystemsIds[i] {
			if ssdIdInt < 0 || ssdIdInt > num {
				continue
			}
			res := getSsdId(systemId, ssdIdInt)
			c.JSON(http.StatusOK, res)
			found = true
			break
		}
	}

	if found == false {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    -1,
			"message": "Not found system id:" + systemId,
		})
	}
}

// handleSsdIdPatch godoc
// @Summary Update status of a SSD device
// @Tags Chassis - Custom OEM
// @Accept json
// @Produce json
// @Param  systemsId path string true "System ID (Default = Irene)" default(Irene)
// @Param  ssdId path int true "Ssd ID" default(0)
// @Param  Oem body schema.SsdOemCustom true "Reset SSD (reset)"
// @Success 200 {object} schema.SsdOemCustom
// @Failure 400 {object} schema.HttpError
// @Failure 404 {object} schema.HttpError
// @Failure 500 {object} schema.HttpError
// @Router /redfish/v1/Systems/{systemsId}/Ssd/{ssdId} [patch]
func handleSsdIdPatch(c *gin.Context) {
	systemId := c.Param("systemId")
	ssdId := c.Param("ssdId")
	ssdOem := schema.SsdOemCustom{}
	if err := c.ShouldBindJSON(&ssdOem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	if err := ssdOem.Validation(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	if err := setSsd(systemId, ssdId, ssdOem.Action); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ssdOem)
}