package main

import (
	"flag"
)

func main(){
	protocol := flags.String("protocol", "tcp","Protocol from which you would like to capture")
	seconds = flags.Float64("seconds", 3, "How many seconds to capture for")
	filePath = flags.String("filePath", "", "Where you would like to output your capture")

	validateFlags(*protocol, *seconds, *filePath);
	flags.Parse();

	packetCaptureStart(*protocol, *seconds, *filePath);
}

func validateFlags(){

}