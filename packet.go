package main

import(
	"time"
	"github.com/google/gopacket"
)

func packetCaptureStart(seconds float64, filePath string, device string) {
	if((os.Stat(filePath).Mode().Perm()&0200 == 0)){
		print("File: " + filePath + " does give you write permissions");
		return;
	}
	capture := capture(device, seconds);
	writeOut(filePath, capture);
	return;
}

func writeOut(filePath string, data string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644);
	_, err = file.WriteString(data);
	if(err!=nil){panic(err)}
	file.Close();
}

func capture(device string, seconds float64) (capture string) {
	handle, err := OpenLive(device string, 1024, false, time.Duration = seconds*time.Second);

}
