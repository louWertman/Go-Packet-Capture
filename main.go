package main

import (
	"flag"
	"os"
)

func main(){
	device := flags.String("interface", "eth0", "interface to capture from")
	seconds := flags.Float64("seconds", 3, "How many seconds to capture for, float");
	filePath := flags.String("filePath", "capture.txt", "Where you would like to output your capture");

	flags.Parse();
	packetCaptureStart(*device, *seconds, *file);
}
