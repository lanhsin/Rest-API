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

func updatePsuDb(psu hal.PsuDriver, format database.PsuFormatter, db database.Connector) error {
	num, err := psu.GetNumber()
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

		name, err := psu.GetName(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			name = ""
		}
		key = format.GetNameKey(i)
		err = db.Set(key, name)
		if err != nil {
			Lig.Error().Msgf("%v", err)
		}

		powerSupplyType, err := psu.GetPowerSupplyType(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			powerSupplyType = ""
		}
		key = format.GetPowerSupplyTypeKey(i)
		if err := db.Set(key, powerSupplyType); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		lineInputVoltageType, err := psu.GetLineInputVoltageType(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			lineInputVoltageType = ""
		}
		key = format.GetLineInputVoltageTypeKey(i)
		if err := db.Set(key, lineInputVoltageType); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		lineInputVoltage, err := psu.GetLineInputVoltage(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			lineInputVoltage = -1
		}
		key = format.GetLineInputVoltageKey(i)
		if err := db.Set(key, strconv.FormatFloat(lineInputVoltage, 'g', 5, 64)); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		powerCapacityWatts, err := psu.GetPowerCapacityWatts(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			powerCapacityWatts = -1
		}
		key = format.GetPowerCapacityWattsKey(i)
		if err := db.Set(key, strconv.FormatFloat(powerCapacityWatts, 'g', 5, 64)); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		lastPowerOutputWatts, err := psu.GetLastPowerOutputWatts(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			lastPowerOutputWatts = -1
		}
		key = format.GetLastPowerOutputWattsKey(i)
		if err := db.Set(key, strconv.FormatFloat(lastPowerOutputWatts, 'g', 5, 64)); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		powerInputWatts, err := psu.GetPowerInputWatts(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			powerInputWatts = -1
		}
		key = format.GetPowerInputWattsKey(i)
		if err := db.Set(key, strconv.FormatFloat(powerInputWatts, 'g', 5, 64)); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		powerOutputWatts, err := psu.GetPowerOutputWatts(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			powerOutputWatts = -1
		}
		key = format.GetPowerOutputWattsKey(i)
		if err := db.Set(key, strconv.FormatFloat(powerOutputWatts, 'g', 5, 64)); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		model, err := psu.GetModel(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			model = ""
		}
		key = format.GetModelKey(i)
		if err := db.Set(key, model); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		manufacturer, err := psu.GetManufacturer(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			manufacturer = ""
		}
		key = format.GetManufacturerKey(i)
		if err := db.Set(key, manufacturer); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		version, err := psu.GetFirmwareVersion(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			version = ""
		}
		key = format.GetFirmwareVersionKey(i)
		if err := db.Set(key, version); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		sn, err := psu.GetSerialNumber(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			sn = ""
		}
		key = format.GetSerialNumberKey(i)
		if err := db.Set(key, sn); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		pn, err := psu.GetPartNumber(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			pn = ""
		}
		key = format.GetPartNumberKey(i)
		if err := db.Set(key, pn); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		spn, err := psu.GetSparePartNumber(i)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			spn = ""
		}
		key = format.GetSparePartNumberKey(i)
		if err := db.Set(key, spn); err != nil {
			Lig.Error().Msgf("%v", err)
		}

		// Update InputRange array
		inputRanges, err := psu.GetInputRanges(i)
		inputRangesNumber := len(inputRanges)
		if err != nil {
			Lig.Error().Msgf("%v", err)
			inputRangesNumber = -1
		}
		key = format.GetInputRangesNumberKey(i)
		if err := db.Set(key, strconv.Itoa(inputRangesNumber)); err != nil {
			Lig.Error().Msgf("%v", err)
		}
		if inputRangesNumber >= 0 {
			for r, data := range inputRanges {
				key = format.GetInputRangesInputTypeKey(i, int64(r))
				if err := db.Set(key, data.InputType); err != nil {
					Lig.Error().Msgf("%v", err)
				}
				key = format.GetInputRangesMinimumVoltageKey(i, int64(r))
				if err := db.Set(key, strconv.FormatFloat(data.MinVoltage, 'g', 5, 64)); err != nil {
					Lig.Error().Msgf("%v", err)
				}
				key = format.GetInputRangesMaximumVoltageKey(i, int64(r))
				if err := db.Set(key, strconv.FormatFloat(data.MaxVoltage, 'g', 5, 64)); err != nil {
					Lig.Error().Msgf("%v", err)
				}
				key = format.GetInputRangesMinimumFrequencyHzKey(i, int64(r))
				if err := db.Set(key, strconv.FormatFloat(data.MinHz, 'g', 5, 64)); err != nil {
					Lig.Error().Msgf("%v", err)
				}
				key = format.GetInputRangesMaximumFrequencyHzKey(i, int64(r))
				if err := db.Set(key, strconv.FormatFloat(data.MaxHz, 'g', 5, 64)); err != nil {
					Lig.Error().Msgf("%v", err)
				}
				key = format.GetInputRangesOutputWattageKey(i, int64(r))
				if err := db.Set(key, strconv.FormatFloat(data.OutputWatt, 'g', 5, 64)); err != nil {
					Lig.Error().Msgf("%v", err)
				}
			}
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

			key = format.GetPowerSupplyTypeKey(i)
			getPowerSupplyType, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getPowerSupplyType = ""
			}

			key = format.GetLineInputVoltageTypeKey(i)
			getLineInputVoltageType, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getLineInputVoltageType = ""
			}

			key = format.GetLineInputVoltageKey(i)
			getLineInputVoltage, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getLineInputVoltage = ""
			}

			key = format.GetPowerCapacityWattsKey(i)
			getPowerCapacityWatts, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getPowerCapacityWatts = ""
			}

			key = format.GetLastPowerOutputWattsKey(i)
			getLastPowerOutputWatts, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getLastPowerOutputWatts = ""
			}

			key = format.GetPowerInputWattsKey(i)
			getPowerInputWatts, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getPowerInputWatts = ""
			}

			key = format.GetPowerOutputWattsKey(i)
			getPowerOutputWatts, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getPowerOutputWatts = ""
			}

			key = format.GetModelKey(i)
			getModel, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getModel = ""
			}

			key = format.GetManufacturerKey(i)
			getManu, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getManu = ""
			}

			key = format.GetFirmwareVersionKey(i)
			getVersion, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getVersion = ""
			}

			key = format.GetSerialNumberKey(i)
			getSn, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getSn = ""
			}

			key = format.GetPartNumberKey(i)
			getPn, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getPn = ""
			}

			key = format.GetSparePartNumberKey(i)
			getSpn, err := db.Get(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				getSpn = ""
			}

			Log.Trace().
				Str("memoryId", memoryId).
				Str("name", getName).
				Str("powerSupplyType", getPowerSupplyType).
				Str("lineInputVoltageType", getLineInputVoltageType).
				Str("lineInputVoltage", getLineInputVoltage).
				Str("powerCapacityWatts", getPowerCapacityWatts).
				Str("lastPowerOutputWatts", getLastPowerOutputWatts).
				Str("powerInputWatts", getPowerInputWatts).
				Str("powerOutputWatts", getPowerOutputWatts).
				Str("model", getModel).
				Str("manufacturer", getManu).
				Str("version", getVersion).
				Str("sn", getSn).
				Str("pn", getPn).
				Str("spn", getSpn).
				Msgf("i=%v", i)

			// Get InputRanges array
			key = format.GetInputRangesNumberKey(i)
			number, err := db.GetInt64(key)
			if err != nil {
				Lig.Error().Msgf("%v", err)
				number = 0
			}

			Log.Trace().
				Int64("InputRangesNumber", number).
				Msgf("")

			for r := int64(0); r < number; r++ {
				key = format.GetInputRangesInputTypeKey(i, r)
				typeKey, err := db.Get(key)
				if err != nil {
					Lig.Error().Msgf("%v", err)
					typeKey = ""
				}

				key = format.GetInputRangesMinimumVoltageKey(i, r)
				minVolt, err := db.Get(key)
				if err != nil {
					Lig.Error().Msgf("%v", err)
					minVolt = ""
				}

				key = format.GetInputRangesMaximumVoltageKey(i, r)
				maxVolt, err := db.Get(key)
				if err != nil {
					Lig.Error().Msgf("%v", err)
					maxVolt = ""
				}

				key = format.GetInputRangesMinimumFrequencyHzKey(i, r)
				minHz, err := db.Get(key)
				if err != nil {
					Lig.Error().Msgf("%v", err)
					minVolt = ""
				}

				key = format.GetInputRangesMaximumFrequencyHzKey(i, r)
				maxHz, err := db.Get(key)
				if err != nil {
					Lig.Error().Msgf("%v", err)
					maxVolt = ""
				}

				key = format.GetInputRangesOutputWattageKey(i, r)
				outputWatt, err := db.Get(key)
				if err != nil {
					Lig.Error().Msgf("%v", err)
					outputWatt = ""
				}

				Log.Trace().
					Str("typeKey", typeKey).
					Str("minVolt", minVolt).
					Str("maxVolt", maxVolt).
					Str("minHz", minHz).
					Str("maxHz", maxHz).
					Str("outputWatt", outputWatt).
					Msgf("r=%v", r)
			}
		}
	}
	return nil
}
