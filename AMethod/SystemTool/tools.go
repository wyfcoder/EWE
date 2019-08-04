package SystemTool

import "runtime"

func  OsSystem() int{
	sysType := runtime.GOOS

	if sysType == "linux" {
		return 1
	}

	if sysType == "windows" {

		return 2

	}

	return -1
}
