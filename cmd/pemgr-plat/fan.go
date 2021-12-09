package main

import (
	"fmt"
	"pemgr/database"
	"pemgr/hal"
	. "pemgr/logger"
	"strconv"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func updateFanDb(fan hal.FanDriver, format database.FanFormatter, db database.Connector) error {
	num, err := fan.GetNumber()
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

	// TODO: We should change the hardcode sensor number
	var fanNumber int64 = 100
	for i := int64(0); i < num; i++ {
		fanNumber = fanNumber + 1
		key := format.GetMemoryIdKey(i)
		err = db.Set(key, strconv.FormatInt(i, 10))
		if err != nil {
			Lig.Error().Msgf("%v", err)
		}
		key = format.GetSensorNumberKey(i)
		err = db.Set(key, strconv.FormatInt(fanNumber, 10))
		if err != nil {
			Lig.Error().Msgf("%v", err)
		}

		name, err := fan.GetName(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			name = ""
		}
		key = format.GetNameKey(i)
		err = db.Set(key, name)
		if err != nil {
			Lig.Error().Msgf("%v", err)
		}

		speed, err := fan.GetSpeed(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			speed = -1
		}
		key = format.GetSpeedKey(i)
		err = db.Set(key, strconv.FormatInt(speed, 10))
		if err != nil {
			Lig.Error().Msgf("%v", err)
		}

		physicalContext, err := fan.GetPhysicalContext(i)
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

			key = format.GetSpeedKey(i)
			getSpeed, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getSpeed = ""
			}

			Log.Trace().
				Str("memoryId", memoryId).
				Str("name", getName).
				Str("sensorNumber", getSensorNumber).
				Str("speed", getSpeed).
				Msgf("i=%v", i)
		}
	}
	return nil
}

func applyFanDb(fan hal.FanDriver, format database.FanFormatter, db database.Connector) error {
	num, err := fan.GetNumber()
	if err != nil {
		return errors.WithStack(err)
	}

	for i := int64(0); i < num; i++ {
		key := format.GetApplySpeedKey(i)
		getApplySpeed, err := db.Get(key)
		if err != nil {
			continue
		}

		key = format.GetSpeedKey(i)
		getSpeed, err := db.Get(key)
		if err != nil {
			continue
		}

		if getApplySpeed == getSpeed {
			continue
		}

		applySpeedInt, err := strconv.ParseInt(getApplySpeed, 10, 64)
		if err != nil {
			continue
		}
		if err := fan.SetSpeed(i, applySpeedInt); err != nil {
			return errors.WithStack(err)
		}
		Log.Trace().
			Str("speed", getSpeed).
			Int64("applySpeed", applySpeedInt).
			Msgf("Set Speed i=%v:", i)
	}
	return nil
}

