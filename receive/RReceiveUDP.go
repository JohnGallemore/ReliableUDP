package receive

var fileName string
var port int
var mode int = 0          //The mode in which we are running. 0 is stop and wait, 1 is sliding window.
var modeParam int64 = 256 //The size of the window in bytes. 256 is default.

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

func receiveFile() bool {
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
