package main

import (
	"fmt"
	"pemgr/database"
	"pemgr/hal"
	. "pemgr/logger"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func updateSsdDb(ssd hal.SsdDriver, format database.SsdFormatter, db database.Connector) error {
	num, err := ssd.GetNumber()
	if err != nil {
		return errors.WithStack(err)
	}

	key := format.GetNumberKey()
	err = db.Set(key, strconv.FormatInt(num, 10))
	if err != nil {
		return errors.WithStack(err)
	}

	if zerolog.GlobalLevel() == zerolog.TraceLevel {
		key = format.GetNumberKey()
		getNumber, err := db.Get(key)
		if err != nil {
			return errors.WithStack(err)
		}

		formatType := fmt.Sprintf("%T", format)
		Log.Trace().
			Str("number", getNumber).
			Msg(formatType)
	}

	for i := int64(0); i < num; i++ {
		sensorNumber = sensorNumber + 1
		key := format.GetMemoryIdKey(i)
		err = db.Set(key, strconv.FormatInt(i, 10))
		if err != nil {
			Lig.Error().Msgf("%v", err)
		}
		key = format.GetSensorNumberKey(i)
		err = db.Set(key, strconv.FormatInt(sensorNumber, 10))
		if err != nil {
			Lig.Error().Msgf("%v", err)
		}

		name, err := ssd.GetName(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			name = ""
		}
		key = format.GetNameKey(i)
		err = db.Set(key, name)
		if err != nil {
			Lig.Error().Msgf("%v", err)
		}

		manufacturer, err := ssd.GetManufacturer(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			manufacturer = ""
		}
		key = format.GetManufacturerKey(i)
		if err := db.Set(key, manufacturer); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		model, err := ssd.GetModel(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			model = ""
		}
		key = format.GetModelKey(i)
		if err := db.Set(key, model); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		sn, err := ssd.GetSerialNumber(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			sn = ""
		}
		key = format.GetSerialNumberKey(i)
		if err := db.Set(key, sn); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		temp, err := ssd.GetTemperature(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			temp = -1
		}
		key = format.GetTemperatureKey(i)
		if err := db.Set(key, strconv.FormatFloat(temp, 'g', 5, 64)); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		if zerolog.GlobalLevel() == zerolog.TraceLevel {
			key = format.GetMemoryIdKey(i)
			memoryId, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				memoryId = ""
			}

			key = format.GetNameKey(i)
			getName, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getName = ""
			}

			key = format.GetSensorNumberKey(i)
			getSensorNumber, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getSensorNumber = ""
			}

			key = format.GetManufacturerKey(i)
			getManu, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getManu = ""
			}

			key = format.GetModelKey(i)
			getModel, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getModel = ""
			}

			key = format.GetSerialNumberKey(i)
			getSn, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getSn = ""
			}

			key = format.GetTemperatureKey(i)
			getTemp, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getTemp = ""
			}

			Log.Trace().
				Str("memoryId", memoryId).
				Str("name", getName).
				Str("sensorNumber", getSensorNumber).
				Str("manufacturer", getManu).
				Str("model", getModel).
				Str("sn", getSn).
				Str("temperature", getTemp).
				Msgf("i=%v", i)
		}
	}
	return nil
}

func applySsdDb(ssd hal.SsdDriver, format database.SsdFormatter, db database.Connector) error {
	num, err := ssd.GetNumber()
	if err != nil {
		return errors.WithStack(err)
	}

	for i := int64(0); i < num; i++ {
		key := format.GetActionKey(i)
		action, err := db.Get(key)
		if err != nil {
			continue
		}

		if strings.ToLower(action) == "reset" {
			Log.Trace().
				Str("action", action).
				Msgf("Reset SSD i=%v:", i)
			// Reset SSD
			if err := ssd.SetReset(i); err != nil {
				return errors.WithStack(err)
			}
			db.Set(key, "normal")
		}
	}
	return nil
}

