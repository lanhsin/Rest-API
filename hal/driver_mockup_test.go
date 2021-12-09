package hal

import (
	"os"
	"pemgr/cmd/pemgr-server/schema"
	"testing"
)

func testSensor(t *testing.T, sensor SensorDriver) {
	num, err := sensor.GetNumber()
	if err != nil {
		t.Errorf("GetNumber() error! err:%v", err)
	}

	for i := int64(0); i < num; i++ {
		raw, unit, err := sensor.GetRaw(i)
		if err != nil {
			t.Errorf("GetRaw() error! err:%v id:%v", err, i)
			continue
		}
		t.Logf("GetRaw() id:%v raw:%v unit:%v", i, raw, unit)
	}
}

func testFan(t *testing.T, fan FanDriver) {
	num, err := fan.GetNumber()
	if err != nil {
		t.Errorf("GetNumber() error! err:%v", err)
	}

	for i := int64(0); i < num; i++ {
		testSpeed := int64(10)
		err = fan.SetSpeed(i, testSpeed)
		if err != nil {
			t.Errorf("SetSpeed() error! err:%v id:%v", err, i)
			continue
		}
		speed, err := fan.GetSpeed(i)
		if err != nil {
			t.Errorf("GetSpeed() error! err:%v id:%v", err, i)
			continue
		}

		if speed != testSpeed {
			t.Errorf("Speed error! err:%v id:%v speed:%v testSpeed:%v", err, i, speed, testSpeed)
			continue
		}

		t.Logf("Fan id:%v speed:%v", i, speed)
	}
}

func testSsd(t *testing.T, ssd SsdDriver) {
	num, err := ssd.GetNumber()
	if err != nil {
		t.Errorf("GetNumber() error! err:%v", err)
	}

	for i := int64(0); i < num; i++ {
		manufacturer, err := ssd.GetManufacturer(i)
		if err != nil {
			t.Errorf("GetManufacturer() error! err:%v id:%v", err, i)
			continue
		}
		model, err := ssd.GetModel(i)
		if err != nil {
			t.Errorf("GetModel() error! err:%v id:%v", err, i)
			continue
		}
		sn, err := ssd.GetSerialNumber(i)
		if err != nil {
			t.Errorf("GetSerialNumber() error! err:%v id:%v", err, i)
			continue
		}
		temp, err := ssd.GetTemperature(i)
		if err != nil {
			t.Errorf("GetTemperature() error! err:%v id:%v", err, i)
			continue
		}
		t.Logf("SSD id:%v manufacturer:%v model:%v sn:%v temp:%v", i, manufacturer, model, sn, temp)
	}
}

func testPsu(t *testing.T, psu PsuDriver) {
	num, err := psu.GetNumber()
	if err != nil {
		t.Errorf("GetNumber() error! err:%v", err)
	}

	for i := int64(0); i < num; i++ {
		name, err := psu.GetName(i)
		if err != nil {
			t.Errorf("GetName() error! err:%v id:%v", err, i)
			continue
		}

		supplyType, err := psu.GetPowerSupplyType(i)
		if err != nil {
			t.Errorf("GetPowerSuppolyType() error! err:%v id:%v", err, i)
			continue
		}

		lineInputVoltageType, err := psu.GetLineInputVoltageType(i)
		if err != nil {
			t.Errorf("GetLineInputVoltageType() error! err:%v id:%v", err, i)
			continue
		}

		lineInputVoltage, err := psu.GetLineInputVoltage(i)
		if err != nil {
			t.Errorf("GetLineInputVoltage() error! err:%v id:%v", err, i)
			continue
		}

		powerCapacityWatts, err := psu.GetPowerCapacityWatts(i)
		if err != nil {
			t.Errorf("GetPowerCapacityWatts() error! err:%v id:%v", err, i)
			continue
		}

		lastPowerOutputWatts, err := psu.GetLastPowerOutputWatts(i)
		if err != nil {
			t.Errorf("GetLastPowerOutputWatts() error! err:%v id:%v", err, i)
			continue
		}

		powerInputWatts, err := psu.GetPowerInputWatts(i)
		if err != nil {
			t.Errorf("GetPowerInputWatts() error! err:%v id:%v", err, i)
			continue
		}

		powerOutputWatts, err := psu.GetPowerOutputWatts(i)
		if err != nil {
			t.Errorf("GetPowerOutputWatts() error! err:%v id:%v", err, i)
			continue
		}

		model, err := psu.GetModel(i)
		if err != nil {
			t.Errorf("GetModel() error! err:%v id:%v", err, i)
			continue
		}

		manufacturer, err := psu.GetManufacturer(i)
		if err != nil {
			t.Errorf("GetManufacturer() error! err:%v id:%v", err, i)
			continue
		}

		version, err := psu.GetFirmwareVersion(i)
		if err != nil {
			t.Errorf("GetFirmwareVersion() error! err:%v id:%v", err, i)
			continue
		}

		sn, err := psu.GetSerialNumber(i)
		if err != nil {
			t.Errorf("GetSerialNumber() error! err:%v id:%v", err, i)
			continue
		}

		pn, err := psu.GetPartNumber(i)
		if err != nil {
			t.Errorf("GetPartNumber() error! err:%v id:%v", err, i)
			continue
		}

		spn, err := psu.GetSparePartNumber(i)
		if err != nil {
			t.Errorf("GetSparePartNumber() error! err:%v id:%v", err, i)
			continue
		}

		ranges, err := psu.GetInputRanges(i)
		if err != nil {
			t.Errorf("GetInputRanges() error! err:%v id:%v", err, i)
			continue
		}

		t.Logf("=====Psu id:%v name:%v=====", i, name)
		t.Logf("=supplyType:%v", supplyType)
		t.Logf("=lineInputVoltageType:%v", lineInputVoltageType)
		t.Logf("=lineInputVoltage:%v", lineInputVoltage)
		t.Logf("=powerCapacityWatts:%v", powerCapacityWatts)
		t.Logf("=lastPowerOutputWatts:%v", lastPowerOutputWatts)
		t.Logf("=powerInputWatts:%v", powerInputWatts)
		t.Logf("=powerOutputWatts:%v", powerOutputWatts)
		t.Logf("=model:%v", model)
		t.Logf("=manufacturer:%v", manufacturer)
		t.Logf("=version:%v", version)
		t.Logf("=sn:%v", sn)
		t.Logf("=pn:%v", pn)
		t.Logf("=spn:%v", spn)
		t.Logf("=ranges:%v", ranges)
	}
}

func TestHalMockup(t *testing.T) {
	if os.Getenv("TEST_PLATFORM") != "mockup" {
		t.Logf("Skip the test (TEST_PLATFORM=%v)", os.Getenv("TEST_PLATFORM"))
		return
	}

	schema.InitProperty()

	var sensor SensorDriver
	t.Logf("Test Temperature")
	sensor = &MockupTemperature{}
	testSensor(t, sensor)

	t.Logf("Test LED")
	sensor = &MockupLed{}
	testSensor(t, sensor)

	t.Logf("Test Button")
	sensor = &MockupButton{}
	testSensor(t, sensor)

	var fan FanDriver
	t.Logf("Test Fan")
	fan = &MockupFan{}
	testFan(t, fan)

	var ssd SsdDriver
	t.Logf("Test SSD")
	ssd = &MockupSsd{}
	testSsd(t, ssd)

	t.Logf("Test Voltage")
	sensor = &MockupVoltage{}
	testSensor(t, sensor)

	var psu PsuDriver
	t.Logf("Test PSU")
	psu = &MockupPsu{}
	testPsu(t, psu)
}
