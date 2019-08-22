package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/byuoitav/common/log"
)

func setPowerOn(address string) string {
	conn := getConnection(address)

	cmd := []byte("PWR ON")
	cmd = append(cmd, 0x0d)
	n, err := conn.Write(cmd)
	switch {
	case err != nil:
		log.L.Debugf("There was an error setting power status: %v", err)
		return ""
	case n != len(cmd):
		log.L.Debugf("Error setting power status: only sent %v/%v bytes\n", n, len(cmd))
		return ""
	}
	bytes, err := conn.ReadUntil(':', 5*time.Second)
	if err != nil {
		log.L.Debugf("Error reading the response: %v", err)
		return ""
	}

	checker := fmt.Sprintf("%x", bytes)

	if checker != "003a" {
		log.L.Debugf("There was an error executing the command - %s", bytes)
	}

	//wanna return powering on or somehting probably
	log.L.Infof("Powering on")
	time.Sleep(25 * time.Second)
	return ""
}

func setPowerOff(address string) string {
	conn := getConnection(address)

	cmd := []byte("PWR OFF")
	cmd = append(cmd, 0x0d)
	n, err := conn.Write(cmd)
	switch {
	case err != nil:
		log.L.Debugf("There was an error setting power status: %v", err)
		return ""
	case n != len(cmd):
		log.L.Debugf("Error setting power status: only sent %v/%v bytes\n", n, len(cmd))
		return ""
	}
	bytes, err := conn.ReadUntil(':', 5*time.Second)
	if err != nil {
		log.L.Debugf("Error reading the response: %v", err)
		return ""
	}

	checker := fmt.Sprintf("%x", bytes)

	if checker != "003a" {
		log.L.Debugf("There was an error executing the command - %s", bytes)
	}

	//return powering off or something probably
	log.L.Infof("Powering off")
	time.Sleep(10 * time.Second)
	return ""
}

func setVolume(address string, volume int) string {
	conn := getConnection(address)

	word := "VOL "
	bigVolume := volume*12 + 3
	newVolume := strconv.Itoa(bigVolume)
	word += newVolume
	cmd := []byte(word)
	cmd = append(cmd, 0x0d)
	n, err := conn.Write(cmd)
	switch {
	case err != nil:
		log.L.Debugf("There was an error getting power status: %v", err)
		return ""
	case n != len(cmd):
		log.L.Debugf("Error gettng power status: only sent %v/%v bytes\n", n, len(cmd))
		return ""
	}
	bytes, err := conn.ReadUntil(':', 5*time.Second)
	if err != nil {
		log.L.Debugf("Error reading the response: %v", err)
		return ""
	}

	checker := fmt.Sprintf("%x", bytes)

	if checker != "003a" {
		log.L.Debugf("There was an error executing the command - %s", bytes)
	}
	finalVolume := fmt.Sprintf("Volume set to %d", volume)
	return finalVolume
}

func blank(address string) string {
	conn := getConnection(address)

	cmd := []byte("MUTE ON")
	cmd = append(cmd, 0x0d)
	n, err := conn.Write(cmd)
	switch {
	case err != nil:
		log.L.Debugf("There was an error blanking the projector: %v", err)
		return ""
	case n != len(cmd):
		log.L.Debugf("Error blanking projector: only sent %v/%v bytes\n", n, len(cmd))
		return ""
	}
	bytes, err := conn.ReadUntil(':', 5*time.Second)
	if err != nil {
		log.L.Debugf("Error reading the response: %v", err)
		return ""
	}

	checker := fmt.Sprintf("%x", bytes)

	if checker != "003a" {
		log.L.Debugf("There was an error executing the command - ")
	}

	//return powering off or something probably
	return "Blanking screen"
}
