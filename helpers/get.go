package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/pooled"
	"github.com/byuoitav/common/status"
)

// GetPower .
func GetPower(address string) (status.Power, error) {
	var power status.Power

	work := func(conn pooled.Conn) error {
		conn.Log().Infof("Getting power state")

		cmd := []byte("PWR?")
		cmd = append(cmd, 0x0d)
		checker, err := writeAndRead(conn, cmd, 5*time.Second, ':')
		if err != nil {
			return fmt.Errorf("There was an error getting power status: %v", err)
		}

		switch checker {
		case "PWR=00":
			// Standby
			power.Power = "standby"
		case "PWR=01":
			// On
			power.Power = "on"
		case "PWR=02":
			// Warming up
			power.Power = "standby"
		case "PWR=03":
			// Cooling down
			power.Power = "standby"
		case "PWR=04":
			// Standby (network offline)
			power.Power = "standby"
		case "PWR=05":
			// Standby (abnormal)
			power.Power = "standby"
		case "PWR=09":
			// Standby (A/V standby)
			power.Power = "standby"
		default:
			return fmt.Errorf("unknown power state '%s'", checker)
		}
		conn.Log().Debugf("received response: %s\n", checker)
		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return power, err
	}

	return power, nil
}

// GetInput .
func GetInput(address string) (status.Input, error) {
	var input status.Input

	work := func(conn pooled.Conn) error {

		cmd := []byte("SOURCE?")
		cmd = append(cmd, 0x0d)
		checker, err := writeAndRead(conn, cmd, 5*time.Second, ':')
		if err != nil {
			return fmt.Errorf("There was an error getting the input: %v", err)
		}

		switch checker {
		case "SOURCE=30":
			input.Input = "HDMI1"
		case "SOURCE=A0":
			input.Input = "HDMI2"
		case "SOURCE=C0":
			input.Input = "HDMI3"
		case "SOURCE=11":
			input.Input = "computer"
		case "SOURCE=41":
			input.Input = "VGA"
		case "SOURCE=51":
			input.Input = "USB Display"
		case "SOURCE=52":
			input.Input = "USB1"
		case "SOURCE=54":
			input.Input = "USB2"
		case "SOURCE=53":
			input.Input = "LAN"
		default:
			return fmt.Errorf("unknown source response '%s'", checker)
		}
		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return input, err
	}

	return input, nil
}

// GetVolume .
func GetVolume(address string) (status.Volume, error) {
	var volume status.Volume

	work := func(conn pooled.Conn) error {

		cmd := []byte("VOL?")
		cmd = append(cmd, 0x0d)
		checker, err := writeAndRead(conn, cmd, 5*time.Second, ':')
		if err != nil {
			return fmt.Errorf("There was an error getting the volume: %v", err)
		}
		checker = strings.Split(checker, "=")[1]
		checker = strings.Split(checker, "\r")[0]

		//convert and divide by 12 because they have it on a scale of 0-255
		num, err := strconv.Atoi(checker)
		if err != nil {
			log.L.Warnf("Error converting to int %v\n", err)
		}

		num = num / 12

		log.L.Infof("The volume level is %v", num)

		if num > 20 || num < 0 {
			log.L.Warnf("Volume out of range: %d", num)
			return fmt.Errorf("volume out of range: %d", num)
		}

		volume.Volume = num
		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return volume, err
	}

	return volume, nil
}

// GetBlanked .
func GetBlanked(address string) (status.Blanked, error) {
	var blanked status.Blanked

	work := func(conn pooled.Conn) error {

		cmd := []byte("VOL?")
		cmd = append(cmd, 0x0d)
		checker, err := writeAndRead(conn, cmd, 5*time.Second, ':')
		if err != nil {
			return fmt.Errorf("There was an error getting blanked: %v", err)
		}
		switch checker {
		case "MUTE=ON":
			blanked.Blanked = true
		case "MUTE=OFF":
			blanked.Blanked = false
		default:
			return fmt.Errorf("unknown blanked state '%s'", checker)
		}
		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return blanked, err
	}

	return blanked, nil
}
