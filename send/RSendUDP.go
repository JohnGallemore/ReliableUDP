package send

import "net"

var fileName string
var port int
var mode int = 0          //The mode in which we are running. 0 is stop and wait, 1 is sliding window.
var modeParam int64 = 256 //The size of the window in bytes. 256 is default.
var receiver net.UDPAddr
var timeout int64 = 1000 //The timeout in milliseconds. 1 second is default

func getFilename() string {
	return fileName
}

func getLocalPort() int {
	return port
}

func getMode() int {
	return mode
}

func getModeParameter() int64 {
	return modeParam
}

func getReceiver() net.UDPAddr {
	return receiver
}

func getTimeout() int64 {
	return timeout
}

func sendFile() bool {
	return true
}

func setFilename(fName string) {
	fileName = fName
}

func setLocalPort(newPort int) bool {
	port = newPort
	return true
}

func setMode(newMode int) bool {
	mode = newMode
	return true
}

func setModeParameter(n int64) bool {
	if mode == 0 {
		return false
	} else {
		modeParam = n
		return true
	}
}

func setReceiver(newReceiver *net.UDPAddr) bool {
	receiver = *newReceiver
	return true
}

func setTimeout(newTimeout int64) bool {
	timeout = newTimeout
	return true
}
