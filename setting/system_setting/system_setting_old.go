package system_setting

var ServerAddress = "https://api.flare-sh.tech"
var WorkerUrl = ""
var WorkerValidKey = ""
var WorkerAllowHttpImageRequestEnabled = false

func EnableWorker() bool {
	return WorkerUrl != ""
}
