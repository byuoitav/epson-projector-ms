package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/byuoitav/common/pooled"
)

func main() {
	netConn, err := net.Dial("tcp", "10.5.34.44:3629")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn := pooled.Wrap(netConn)

	cmd := []byte{0x45, 0x53, 0x43, 0x2F, 0x56, 0x50, 0x2E, 0x6E, 0x65, 0x74, 0x10, 0x03, 0x00, 0x00, 0x00, 0x00}
	fmt.Printf("command: %s\n", cmd)
	fmt.Printf("command: %x\n", cmd)
	n, err := conn.Write(cmd)
	switch {
	case err != nil:
		fmt.Println(err)
		return
	case n != len(cmd):
		fmt.Printf("only sent %v/%v bytes\n", n, len(cmd))
		return
	}
	bytes, err := conn.ReadUntil(' ', 5*time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("receivced response: 0x%x\n", bytes)
	time.Sleep(1 * time.Second)

	// volume := 15
	word := "MUTE?"
	// if word != "VOL?" && word != "VOL INC" {
	// 	// oldword := strings.Split(word, " ")[1]
	// 	// derek, err := strconv.Atoi(oldword)
	// 	// if err != nil {
	// 	// 	fmt.Printf("boo you suck: %v", err)
	// 	// }
	// 	volume = volume*12 + 3
	// 	newword := strconv.Itoa(volume)
	// 	word += newword
	// }

	cmd = []byte(word)

	// byteVolume := []byte(volume)
	// for _, i := range byteVolume {
	// 	cmd = append(cmd, i)
	// }
	cmd = append(cmd, 0x0d)
	fmt.Printf("command: %s\n", cmd)
	fmt.Printf("command: %x\n", cmd)
	n, err = conn.Write(cmd)
	switch {
	case err != nil:
		fmt.Println(err)
		return
	case n != len(cmd):
		fmt.Printf("only sent %v/%v bytes\n", n, len(cmd))
		return
	}
	bytes, err = conn.ReadUntil(':', 5*time.Second)
	if err != nil {
		fmt.Println("there was an errror %v", err)
		return
	}
	// if word == "VOL?" {
	checker := fmt.Sprintf("%s", bytes)
	fmt.Printf("this is the boy: %s\n", checker)
	// checker = strings.Split(checker, "=")[1]
	// checker = strings.Split(checker, "\r")[0]
	// num, err := strconv.Atoi(checker)
	// if err != nil {
	// fmt.Printf("BAD BOY: %v\n", err)
	// }
	// newnum := (float64(num) / 255) * 20
	fmt.Printf("his is checker %v\n", checker)
	// }

	// if checker != "003a" {
	// 	fmt.Printf("BAD\n")
	// }

	fmt.Printf("receivced response: 0x%x\n", bytes)

	word = "VOL?"
	if word != "VOL?" && word != "VOL INC" {
		oldword := strings.Split(word, " ")[1]
		derek, err := strconv.Atoi(oldword)
		if err != nil {
			fmt.Printf("boo you suck: %v", err)
		}
		derek = derek*12 + 3
		newword := strconv.Itoa(derek)
		word = strings.ReplaceAll(word, oldword, newword)
	}

	cmd = []byte(word)
	cmd = append(cmd, 0x0d)
	fmt.Printf("command: %s\n", cmd)
	fmt.Printf("command: %x\n", cmd)
	n, err = conn.Write(cmd)
	switch {
	case err != nil:
		fmt.Println(err)
		return
	case n != len(cmd):
		fmt.Printf("only sent %v/%v bytes\n", n, len(cmd))
		return
	}
	bytes, err = conn.ReadUntil(':', 5*time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	if word == "VOL?" {
		checker = fmt.Sprintf("%s", bytes)
		fmt.Printf("this is the boy: %s\n", checker)
		checker = strings.Split(checker, "=")[1]
		checker = strings.Split(checker, "\r")[0]
		num, err := strconv.Atoi(checker)
		if err != nil {
			fmt.Printf("BAD BOY: %v\n", err)
		}
		num = num / 12
		// newnum := (float64(num) / 255) * 20
		fmt.Printf("his is checker %v\n", num)
	}

	// if checker != "003a" {
	// 	fmt.Printf("BAD\n")
	// }

	fmt.Printf("receivced response: 0x%x\n", bytes)
}
