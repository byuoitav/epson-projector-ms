package helpers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/status"
)

// SetPower sets the power for an epson projector
func SetPower(address string, power status.Power) error {
	switch power.Power {
	case "standby":
		power.Power = "off"
	}

	conn := getConnection(address)

	var cmd []byte

	switch power.Power {
	case "on":
		cmd = []byte("PWR ON")
	case "off":
		cmd = []byte("PWR OFF")
	default:
		return fmt.Errorf("unexpected power state: %v", power.Power)
	}
	cmd = append(cmd, 0x0d)
	n, err := conn.Write(cmd)
	switch {
	case err != nil:
		log.L.Warnf("There was an error setting power status: %v", err)
		return err
	case n != len(cmd):
		log.L.Warnf("Error setting power status: only sent %v/%v bytes\n", n, len(cmd))
		return err
	}
	bytes, err := conn.ReadUntil(':', 5*time.Second)
	if err != nil {
		log.L.Warnf("Error reading the response: %v", err)
		return err
	}

	checker := fmt.Sprintf("%x", bytes)

	if checker != "003a" {
		return fmt.Errorf("There was an error executing the command - %s", bytes)
	}

	//wanna return powering on or something probably
	log.L.Infof("Power state changed: %v", power.Power)
	time.Sleep(25 * time.Second)
	return nil
}

// SetInput sets the power for an epson projector
func SetInput(address string, input status.Input) error {
	var str string
	switch input.Input {
	case "HDMI1":
		str = "30"
	case "HDMI2":
		str = "A0"
	case "HDMI3":
		str = "C0"
	case "computer":
		str = "11"
	case "VGA":
		str = "41"
	case "USB Display":
		str = "51"
	case "USB1":
		str = "52"
	case "USB2":
		str = "54"
	case "LAN":
		str = "53"
	default:
		return fmt.Errorf("unknown source input '%s'", input.Input)
	}

	conn := getConnection(address)

	cmd := []byte(fmt.Sprintf("SOURCE %s", str))
	cmd = append(cmd, 0x0d)
	n, err := conn.Write(cmd)
	switch {
	case err != nil:
		log.L.Warnf("there was an error setting input: %v", err)
		return err
	case n != len(cmd):
		log.L.Warnf("error setting input: only sent %v/%v bytes\n", n, len(cmd))
		return err
	}
	bytes, err := conn.ReadUntil(':', 5*time.Second)
	if err != nil {
		log.L.Warnf("error reading the response: %v", err)
		return err
	}

	checker := fmt.Sprintf("%x", bytes)

	if checker != "003a" {
		return fmt.Errorf("there was an error executing the command - %s", bytes)
	}

	log.L.Infof("input changed: %v", input.Input)
	time.Sleep(25 * time.Second)
	return nil
}

// SetVolume sets the volume on an epson projector
func SetVolume(address string, volume status.Volume) error {
	conn := getConnection(address)

	word := "VOL "
	bigVolume := volume.Volume*12 + 3
	newVolume := strconv.Itoa(bigVolume)
	word += newVolume
	cmd := []byte(word)
	cmd = append(cmd, 0x0d)
	n, err := conn.Write(cmd)
	switch {
	case err != nil:
		log.L.Warnf("There was an error getting power status: %v", err)
		return err
	case n != len(cmd):
		log.L.Warnf("Error gettng power status: only sent %v/%v bytes\n", n, len(cmd))
		return err
	}
	bytes, err := conn.ReadUntil(':', 5*time.Second)
	if err != nil {
		log.L.Warnf("Error reading the response: %v", err)
		return err
	}

	checker := fmt.Sprintf("%x", bytes)

	if checker != "003a" {
		return fmt.Errorf("There was an error executing the command - %s", bytes)
	}
	log.L.Infof("Volume set to %d", volume.Volume)

	return nil
}

// SetBlanked sets the blank status on an epson projector
func SetBlanked(address string, blanked status.Blanked) error {
	var str string
	switch blanked.Blanked {
	case true:
		str = "ON"
	case false:
		str = "OFF"
	default:
		return fmt.Errorf("unexpected blank state '%v'", blanked.Blanked)
	}

	conn := getConnection(address)

	cmd := []byte(fmt.Sprintf("MUTE %s", str))
	cmd = append(cmd, 0x0d)
	n, err := conn.Write(cmd)
	switch {
	case err != nil:
		log.L.Warnf("There was an error blanking the projector: %v", err)
		return err
	case n != len(cmd):
		log.L.Warnf("Error blanking projector: only sent %v/%v bytes\n", n, len(cmd))
		return err
	}
	bytes, err := conn.ReadUntil(':', 5*time.Second)
	if err != nil {
		log.L.Warnf("Error reading the response: %v", err)
		return err
	}

	checker := fmt.Sprintf("%x", bytes)

	if checker != "003a" {
		return fmt.Errorf("There was an error executing the command - %s", bytes)
	}

	//return powering off or something probably
	log.L.Infof("blanking screen")
	return nil
}
