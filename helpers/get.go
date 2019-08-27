package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/status"
)

// GetPower .
func GetPower(address string) (status.Power, error) {
	var power status.Power
	conn := getConnection(address)

	cmd := []byte("PWR?")
	cmd = append(cmd, 0x0d)
	n, err := conn.Write(cmd)
	switch {
	case err != nil:
		log.L.Warnf("There was an error getting power status: %v", err)
		return power, err
	case n != len(cmd):
		log.L.Warnf("Error gettng power status: only sent %v/%v bytes\n", n, len(cmd))
		return power, fmt.Errorf("error gettng power status: only sent %v/%v bytes", n, len(cmd))
	}
	bytes, err := conn.ReadUntil(':', 5*time.Second)
	if err != nil {
		log.L.Warnf("Error reading the response: %v", err)
		return power, err
	}

	checker := fmt.Sprintf("%s", bytes)

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
		return power, fmt.Errorf("unknown power state '%s'", bytes)
	}
	log.L.Debugf("received response: %s\n", bytes)
	return power, nil
}

// GetInput .
func GetInput(address string) (status.Input, error) {
	var input status.Input
	conn := getConnection(address)

	cmd := []byte("SOURCE?")
	cmd = append(cmd, 0x0d)
	n, err := conn.Write(cmd)
	switch {
	case err != nil:
		log.L.Warnf("There was an error getting power status: %v", err)
		return input, err
	case n != len(cmd):
		log.L.Warnf("Error gettng power status: only sent %v/%v bytes\n", n, len(cmd))
		return input, fmt.Errorf("error gettng power status: only sent %v/%v bytes", n, len(cmd))
	}
	bytes, err := conn.ReadUntil(':', 5*time.Second)
	if err != nil {
		log.L.Warnf("Error reading the response: %v", err)
		return input, err
	}

	checker := fmt.Sprintf("%s", bytes)

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
		return input, fmt.Errorf("unknown source response '%s'", bytes)
	}

	log.L.Debugf("received response: %s\n", bytes)
	return input, nil
}

// GetVolume .
func GetVolume(address string) (status.Volume, error) {
	var volume status.Volume

	conn := getConnection(address)

	cmd := []byte("VOL?")
	cmd = append(cmd, 0x0d)

	n, err := conn.Write(cmd)

	switch {
	case err != nil:
		log.L.Warnf("There was an error sending the command: %v", err)
		return volume, err
	case n != len(cmd):
		log.L.Warnf("Only sent %v/%v bytes\n", n, len(cmd))
		return volume, fmt.Errorf("only sent %v/%v bytes", n, len(cmd))
	}

	bytes, err := conn.ReadUntil(':', 5*time.Second)
	if err != nil {
		log.L.Warnf("Error reading the response: %v", err)
		return volume, err
	}

	//checker should be 0x00"VOL=xx"\r so we need to get just the number
	checker := fmt.Sprintf("%s", bytes)
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
		return volume, fmt.Errorf("volume out of range: %d", num)
	}

	volume.Volume = num

	return volume, nil
}

// GetBlanked .
func GetBlanked(address string) (status.Blanked, error) {
	var blanked status.Blanked

	conn := getConnection(address)

	cmd := []byte("MUTE?")
	cmd = append(cmd, 0x0d)

	n, err := conn.Write(cmd)

	switch {
	case err != nil:
		log.L.Warnf("There was an error sending the command: %v", err)
		return blanked, err
	case n != len(cmd):
		log.L.Warnf("Only sent %v/%v bytes\n", n, len(cmd))
		return blanked, fmt.Errorf("only sent %v/%v bytes", n, len(cmd))
	}

	bytes, err := conn.ReadUntil(':', 5*time.Second)
	if err != nil {
		log.L.Warnf("Error reading the response: %v", err)
		return blanked, err
	}

	checker := fmt.Sprintf("%s", bytes)

	switch checker {
	case "MUTE=ON":
		blanked.Blanked = true
	case "MUTE=OFF":
		blanked.Blanked = false
	default:
		return blanked, fmt.Errorf("unknown blanked state '%s'", bytes)
	}

	log.L.Infof("Bytes: %s", bytes)
	return blanked, nil
}
