package main

import (
	"pemgr/config"
	"pemgr/database"
	"pemgr/hal"
	"pemgr/logger"
	. "pemgr/logger"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

var pConfig config.Config
var sensorNumber int64


func monitorPlatform(wg *sync.WaitGroup, platServer *hal.PlatServer, db database.Connector) {
	for {
		sensorNumber = 0

		if err := updateSensorDb(platServer.Temperature, &database.TemperatureFormat{}, db); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		if err := updateSensorDb(platServer.Led, &database.LedFormat{}, db); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		if err := updateSensorDb(platServer.Button, &database.ButtonFormat{}, db); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		if err := updateSensorDb(platServer.Voltage, &database.VoltageFormat{}, db); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		if err := updateSsdDb(platServer.Ssd, &database.SsdFormat{}, db); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		if err := applyFanDb(platServer.Fan, &database.FanFormat{}, db); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		if err := applyLedDb(platServer.Led, &database.LedFormat{}, db); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		if err := applySsdDb(platServer.Ssd, &database.SsdFormat{}, db); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		time.Sleep(1 * time.Minute)
	}
}

func monitorDelayedDriver(wg *sync.WaitGroup, platServer *hal.PlatServer, db database.Connector) {

	go func() {
		for {
			if err := updateFanDb(platServer.Fan, &database.FanFormat{}, db); err != nil {
				Lig.Error().Msgf("%v", err)
			}
			time.Sleep(1 * time.Minute)
		}
	}()

	go func() {
		for {
			if err := updatePsuDb(platServer.Psu, &database.PsuFormat{}, db); err != nil {
				Lig.Error().Msgf("%v", err)
			}
			time.Sleep(1 * time.Minute)
		}
	}()
}

func main() {
	logger.InitZeroLog(zerolog.InfoLevel)
	//logger.InitZeroLog(zerolog.TraceLevel)

	pConfig.Load()

	database.InitializePemgr(pConfig.Redis.Addr)

	Lig.Info().Msgf("ImpMode: %v", pConfig.Plat.ImpMode)
	var platServer hal.PlatServer
	switch pConfig.Plat.ImpMode {
	default:
		platServer = hal.PlatServer{
			Temperature: &hal.MockupTemperature{},
			Led:         &hal.MockupLed{},
			Button:      &hal.MockupButton{},
			Fan:         &hal.MockupFan{},
			Ssd:         &hal.MockupSsd{},
			Voltage:     &hal.MockupVoltage{},
			Psu:         &hal.MockupPsu{},
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go monitorPlatform(wg, &platServer, &(database.PemgrDB))
	// TODO: This thread is for Irene platform only
	// We should have a common way to do that
	go monitorDelayedDriver(wg, &platServer, &(database.PemgrDB))

	wg.Wait()
}
