package model

const (
	ERROR_NO_INPUT_FILE  = "'filepath' path not provided"
	ERROR_CANT_READ_FILE = "can not read file input string, exact error %s"
	ERROR_CANT_READ_FILEPATH = "can't load vocabulary, exact error: %s"

	FLAG_FILENAME    = "filepath"
	FLAG_CONCURRENCY = "concurrency value"

	MESSAGE_START = "\nstarting on concurrency %d with NumCPU %d.."
	MESSAGE_WORK = "working time: %+v, result: %d\n"
)
