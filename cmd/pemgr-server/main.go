package main

import (
	"net/http"
	"pemgr/cmd/pemgr-server/schema"
	"pemgr/config"
	db "pemgr/database"
	"pemgr/logger"
	. "pemgr/logger"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

var (
	template Template
	pConfig  config.Config
	pInfo    config.Info
)

// Solve CORS issue from: https://github.com/gin-contrib/cors/issues/29
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// @title Enclosure Management API
// @version 0.2 (Date: 2020/06/09)
// @description Enclosure Management API based on RedFish

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.basic BasicAuth
func main() {
	logger.InitZeroLog(zerolog.TraceLevel)

	pConfig.Load()

	pInfo.Load()

	Lig.Info().Msgf("Version=%v", pInfo.Server.Version)
	Lig.Info().Msgf("Commit=%v", pInfo.Server.Commit)
	Lig.Info().Msgf("Date=%v", pInfo.Server.Date)

	schema.InitProperty()

	db.InitializePemgr(pConfig.Redis.Addr)

	template.setDefaultTemplate()

	r := gin.Default()
	r.Use(CORS())

	authV1 := r.Group("/redfish/v1", gin.BasicAuth(gin.Accounts{
		pConfig.Server.BasicAuthUser: pConfig.Server.BasicAuthPassword,
	}))
	{
		authV1.GET("", handleRoot)
		authV1.GET("/", handleRoot)

		authV1.GET("/Systems", handleSystems)
		authV1.GET("/Systems/:systemId", handleSystemId)
		authV1.GET("/Systems/:systemId/Ssd", handleSsd)
		authV1.GET("/Systems/:systemId/Ssd/:ssdId", handleSsdId)
		authV1.PATCH("/Systems/:systemId/Ssd/:ssdId", handleSsdIdPatch)

		authV1.GET("/Chassis", handleCasssis)
		authV1.GET("/Chassis/:chassisId", handleChassisId)

		authV1.GET("/Chassis/:chassisId/Thermal", handleThermal)
		authV1.GET("/Chassis/:chassisId/Thermal/Temperatures", handleTemperatures)
		authV1.GET("/Chassis/:chassisId/Thermal/Temperatures/:tempId", handleTemperatures)

		authV1.GET("/Chassis/:chassisId/Thermal/Fans", handleFans)
		authV1.GET("/Chassis/:chassisId/Thermal/Fans/:fanId", handleFanId)
		authV1.PATCH("/Chassis/:chassisId/Thermal/Fans/:fanId", handleFanIdPatch)

		authV1.GET("/Chassis/:chassisId/Misc", handleMisc)
		authV1.GET("/Chassis/:chassisId/Misc/Leds", handleLeds)
		authV1.GET("/Chassis/:chassisId/Misc/Leds/:ledId", handleLedId)
		authV1.PATCH("/Chassis/:chassisId/Misc/Leds/:ledId", handleLedIdPatch)

		authV1.GET("/Chassis/:chassisId/Misc/Buttons", handleButtons)
		authV1.GET("/Chassis/:chassisId/Misc/Buttons/:btnId", handleButtons)

		authV1.GET("/Chassis/:chassisId/Power", handlePower)
	}

	authPrivate := r.Group("/private", gin.BasicAuth(gin.Accounts{
		pConfig.Server.BasicAuthUser: pConfig.Server.BasicAuthPassword,
	}))
	{
		authPrivate.GET("/version", handleVersion)
		//authPrivate.GET("/db", handleDb)
		authPrivate.GET("/logsocket", logsocket)
	}

	if pConfig.Server.HttpsEnable {
		crt := pConfig.Server.Certs + "/cert.pem"
		key := pConfig.Server.Certs + "/key.pem"
		go r.RunTLS(pConfig.Server.HttpsAddr, crt, key)
	}

	if pConfig.Server.HttpEnable {
		r.Run(pConfig.Server.HttpAddr)
	}
}

// rootHandler godoc
// @Summary Retrieve list of root-level resources
// @Tags /Root
// @Accept json
// @Produce json
// @Success 200 {object} schema.Root
// @Failure 400 {object} schema.HttpError
// @Failure 404 {object} schema.HttpError
// @Failure 500 {object} schema.HttpError
// @Router /redfish/v1 [get]
func handleRoot(c *gin.Context) {
	var res schema.Root

	res.OdataType = "#ServiceRoot.v1_0_2.ServiceRoot"
	res.Id = "RootService"
	res.Name = "Root Service"
	res.RedfishVersion = "1.6.0"
	res.UUID = "2e390294-285e-4f4d-8473-149599bf699b"
	res.Systems.OdataId = "/redfish/v1/Systems"
	res.Chassis.OdataId = "/redfish/v1/Chassis"
	res.Oem = schema.RootOem{}
	res.OdataId = "/redfish/v1"

	c.JSON(http.StatusOK, res)
}
