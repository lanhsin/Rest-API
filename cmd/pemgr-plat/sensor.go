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

func updateSensorDb(sensor hal.SensorDriver, format database.SensorFormatter, db database.Connector) error {
	num, err := sensor.GetNumber()
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

		name, err := sensor.GetName(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			name = ""
		}
		key = format.GetNameKey(i)
		err = db.Set(key, name)
		if err != nil {
			Lig.Error().Msgf("%v", err)
		}

		raw, unit, err := sensor.GetRaw(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			raw = -1
			unit = ""
		}
		key = format.GetReadingKey(i)
		err = db.Set(key, strconv.FormatFloat(raw, 'g', 5, 64))
		if err != nil {
			Lig.Error().Msgf("%v", err)
		}
		key = format.GetUnitKey(i)
		err = db.Set(key, unit)
		if err != nil {
			Lig.Error().Msgf("%v", err)
		}

		physicalContext, err := sensor.GetPhysicalContext(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			name = ""
		}
		key = format.GetPhysicalContextKey(i)
		err = db.Set(key, physicalContext)
		if err != nil {
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

			key = format.GetReadingKey(i)
			getReading, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getReading = ""
			}

			key = format.GetUnitKey(i)
			getUnit, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getUnit = ""
			}

			Log.Trace().
				Str("memoryId", memoryId).
				Str("name", getName).
				Str("sensorNumber", getSensorNumber).
				Str("reading", getReading).
				Str("unit", getUnit).
				Msgf("i=%v", i)
		}
	}
	return nil
}

func applyLedDb(led hal.SensorDriver, format database.LedFormatter, db database.Connector) error {
	num, err := led.GetNumber()
	if err != nil {
		return errors.WithStack(err)
	}

	for i := int64(0); i < num; i++ {
		key := format.GetActionKey(i)
		action, err := db.Get(key)
		if err != nil {
			continue
		}

		key = format.GetReadingKey(i)
		reading, err := db.Get(key)
		if err != nil {
			continue
		}

		if strings.ToLower(action) == "on" && reading == "0" {
			Log.Trace().
				Str("action", action).
				Str("reading", reading).
				Msgf("Set LED i=%v:", i)
			// Turn on LED
			if err := led.SetRaw(i, 1); err != nil {
				return errors.WithStack(err)
			}
		} else if strings.ToLower(action) == "off" && reading == "1" {
			Log.Trace().
				Str("action", action).
				Str("reading", reading).
				Msgf("Set LED i=%v:", i)
			// Turn off LED
			if err := led.SetRaw(i, 0); err != nil {
				return errors.WithStack(err)
			}
		}
	}
	return nil
}

