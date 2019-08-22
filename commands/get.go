package commands

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/byuoitav/common/log"
)

func getPowerStatus(address string) string {
	conn := getConnection(address)

	cmd := []byte("PWR?")
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

	checker := fmt.Sprintf("%s", bytes)

	switch {
	case checker == "PWR=00":
		return "The projector is in standby mode (network offline)"

	case checker == "PWR=01":
		return "The projector is on"
	case checker == "PWR=02":
		return "The projector is warming up"
	case checker == "PWR=03":
		return "The projector is cooling down"
	case checker == "PWR=04":
		return "The projector is in standby mode (network offline)"
	case checker == "PWR=05":
		return "The projector is in standby mode (abnormal (I don't know what that means))"
	case checker == "PWR=09":
		return "The projector is in standby mode (A/V standby)"
	}

	//if it gets here then something is wrong

	fmt.Printf("received response: %s\n", bytes)
	return ""
}

func getSource(address string) string {
	conn := getConnection(address)

	cmd := []byte("SOURCE?")
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

	checker := fmt.Sprintf("%s", bytes)

	switch {
	case checker == "SOURCE=30":
		return "The projector source is HDMI1"
	case checker == "SOURCE=A0":
		return "The projector source is HDMI2"
	case checker == "SOURCE=C0":
		return "The projector source is HDMI3"
	case checker == "SOURCE=11":
		return "The projector source is computer"
	case checker == "SOURCE=41":
		return "The projector source is VGA"
	case checker == "SOURCE=51":
		return "The projector source is USB Display"
	case checker == "SOURCE=52":
		return "The projector source is USB1"
	case checker == "SOURCE=54":
		return "The projector source is USB2"
	case checker == "SOURCE=53":
		return "The projector source is LAN"
	}

	fmt.Printf("received response: %s\n", bytes)
}

func getVolume(address string) {
	conn := getConnection(address)

	cmd := []byte("VOL?")
	cmd = append(cmd, 0x0d)

	n, err := conn.Write(cmd)

	switch {
	case err != nil:
		log.L.Debugf("There was an error sending the command: %v", err)
		return
	case n != len(cmd):
		log.L.Debugf("Only sent %v/%v bytes\n", n, len(cmd))
		return
	}

	bytes, err := conn.ReadUntil(':', 5*time.Second)
	if err != nil {
		log.L.Debugf("Error reading the response: %v", err)
		return
	}

	//checker should be 0x00"VOL=xx"\r so we need to get just the number
	checker := fmt.Sprintf("%s", bytes)
	checker = strings.Split(checker, "=")[1]
	checker = strings.Split(checker, "\r")[0]

	//convert and divide by 12 because they have it on a scale of 0-255
	num, err := strconv.Atoi(checker)
	if err != nil {
		log.L.Debugf("Error converting to int %v\n", err)
	}

	num = num / 12

	log.L.Infof("The volume level is %v", num)
}

func getBlank(address string) string {
	conn := getConnection(address)

	cmd := []byte("MUTE?")
	cmd = append(cmd, 0x0d)

	n, err := conn.Write(cmd)

	switch {
	case err != nil:
		log.L.Debugf("There was an error sending the command: %v", err)
		return ""
	case n != len(cmd):
		log.L.Debugf("Only sent %v/%v bytes\n", n, len(cmd))
		return ""
	}

	bytes, err := conn.ReadUntil(':', 5*time.Second)
	if err != nil {
		log.L.Debugf("Error reading the response: %v", err)
		return ""
	}

	checker := fmt.Sprintf("%s", bytes)

	switch {
	case checker == "MUTE=ON":
		return "The projector is blanked"
	case checker == "MUTE=OFF":
		return "The projector is not blanked"
	}

	return ""
}
