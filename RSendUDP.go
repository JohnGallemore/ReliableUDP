package main

import "net"

func getFilenameS() string {
	return ""
}

func getLocalPortS() int {
	return 0
}

func getModeS() int {
	return 0
}

func getModeParameterS() int64 {
	return 0
}

func getReceiver() net.UDPAddr {
	return net.UDPAddr{}
}

func getTimeout() int64 {
	return 0
}

func sendFile() bool {
	return true
}

func setFilenameS(fName string) {

}

func setLocalPortS(port int) bool {
	return true
}

func setModeS(mode int) bool {
	return true
}

func setModeParameterS(n int64) bool {
	return true
}

func setReceiver(receiver *net.UDPAddr) bool {
	return true
}

func setTimeout(timeout int64) bool {
	return true
}
