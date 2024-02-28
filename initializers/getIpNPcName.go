package initializers

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/user"
)

func GetIpNPcName() {
	log.Println("My UserName is >>>>>>>>>>>>>")
	u, _ := user.Current()
	hostname, _ := os.Hostname()

	log.Println(u)
	log.Println(hostname)

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			fmt.Println(ipNet.IP)
		}
	}

}
