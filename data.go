package main

import (
	"log"
	"net"
	"os/exec"
	"strings"
)

// Get preferred outbound ip of this machine
func getOutboundIP() string {
	conn, err := net.Dial("udp", "1.1.1.1:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

func getSSID() string {
	cmd := exec.Command("iwgetid", "-r")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}

func getAllOpenPorts() string {
	// Run the netstat command
	cmd := exec.Command("netstat", "-tulpn")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	// Parse the out using awk command to extract the send-q column
	cmd = exec.Command("awk", "/tcp /{print $4}")
	cmd.Stdin = strings.NewReader(string(out))
	out, err = cmd.Output()
	if err != nil {
		panic(err)
	}

	return string(out)
}

func getSSHUsers() string {
	cmd := exec.Command("who")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(out), "\n")
	var users []string
	for _, line := range lines {
		if strings.Contains(line, "pts") {
			users = append(users, line)
		}
	}

	return string(strings.Join(users, "\n"))
}
