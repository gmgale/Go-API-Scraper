package main

import (
	"flag"
)

var portPtr = flag.String("port", "10000", "Port for server setup. Default is 8080.")
var port = ":" + *portPtr

