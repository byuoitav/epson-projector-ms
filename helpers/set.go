package helpers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/pooled"
	"github.com/byuoitav/common/status"
)

// SetPower sets the power for an epson projector
func SetPower(address string, power status.Power) error {
	switch power.Power {
	case "standby":
		power.Power = "off"
	}
	work := func(conn pooled.Conn) error {
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

		checker, err := writeAndRead(conn, cmd, 5*time.Second, ':')
		if err != nil {
			return fmt.Errorf("There was an error setting the power status: %v", err)
		}

		bytes := fmt.Sprintf("%x", checker)

		if bytes != "003a" {
			return fmt.Errorf("There was an error executing the command - %s", bytes)
		}

		conn.Log().Infof("Power state changed: %v", power.Power)
		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return err
	}

	//TODO is the sleep still necessary???
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

	work := func(conn pooled.Conn) error {

		cmd := []byte(fmt.Sprintf("SOURCE %s", str))
		cmd = append(cmd, 0x0d)
		checker, err := writeAndRead(conn, cmd, 5*time.Second, ':')
		if err != nil {
			return fmt.Errorf("There was an error setting the input: %v", err)
		}

		bytes := fmt.Sprintf("%x", checker)

		if bytes != "003a" {
			return fmt.Errorf("There was an error executing the command - %s", bytes)
		}

		conn.Log().Infof("input changed: %v", input.Input)
		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return err
	}

	//TODO remove?
	time.Sleep(25 * time.Second)
	return nil
}

// SetVolume sets the volume on an epson projector
func SetVolume(address string, volume status.Volume) error {
	work := func(conn pooled.Conn) error {

		word := "VOL "
		bigVolume := volume.Volume*12 + 3
		newVolume := strconv.Itoa(bigVolume)
		word += newVolume
		cmd := []byte(word)
		cmd = append(cmd, 0x0d)
		checker, err := writeAndRead(conn, cmd, 5*time.Second, ':')
		if err != nil {
			return fmt.Errorf("There was an error setting the volume: %v", err)
		}

		bytes := fmt.Sprintf("%x", checker)

		if bytes != "003a" {
			return fmt.Errorf("There was an error executing the command - %s", bytes)
		}

		conn.Log().Infof("Volume set to %d", volume.Volume)
		return nil

	}

	err := pool.Do(address, work)
	if err != nil {
		return err
	}

	return nil
}

// SetBlanked sets the blank status on an epson projector
func SetBlanked(address string, blanked status.Blanked) error {
	work := func(conn pooled.Conn) error {
		var str string
		switch blanked.Blanked {
		case true:
			str = "ON"
		case false:
			str = "OFF"
		default:
			return fmt.Errorf("unexpected blank state '%v'", blanked.Blanked)
		}

		cmd := []byte(fmt.Sprintf("MUTE %s", str))
		cmd = append(cmd, 0x0d)
		checker, err := writeAndRead(conn, cmd, 5*time.Second, ':')
		if err != nil {
			return fmt.Errorf("There was an error setting blank status: %v", err)
		}

		bytes := fmt.Sprintf("%x", checker)

		if bytes != "003a" {
			return fmt.Errorf("There was an error executing the command - %s", bytes)
		}

		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return err
	}

	log.L.Infof("blanking screen")
	return nil
}

// SetMuted sets the blank status on an epson projector
func SetMuted(address string, muted status.Mute) error {
	work := func(conn pooled.Conn) error {
		switch muted.Muted {
		case true:
			cmd := []byte("VOL 0")
			cmd = append(cmd, 0x0d)
			checker, err := writeAndRead(conn, cmd, 5*time.Second, ':')
			if err != nil {
				return fmt.Errorf("There was an error muting: %v", err)
			}

			bytes := fmt.Sprintf("%x", checker)

			if bytes != "003a" {
				return fmt.Errorf("There was an error executing the command - %s", bytes)
			}

			return nil
		case false:
			cmd := []byte("VOL 1")
			cmd = append(cmd, 0x0d)
			checker, err := writeAndRead(conn, cmd, 5*time.Second, ':')
			if err != nil {
				return fmt.Errorf("There was an error unmuting: %v", err)
			}

			bytes := fmt.Sprintf("%x", checker)

			if bytes != "003a" {
				return fmt.Errorf("There was an error executing the command - %s", bytes)
			}

			return nil
		default:
			return fmt.Errorf("unexpected mute state '%v'", muted.Muted)
		}
	}

	err := pool.Do(address, work)
	if err != nil {
		return err
	}

	return nil
}
