package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalln("expected 1 argument (root directory), but didn't receive it")
	}
	rootDir := args[1]

	handler := http.FileServer(http.Dir(rootDir))

	port := 9000

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if nil != err {
		log.Fatalf("couldn't listen to TCP connections on a port. Error: %s\n", err)
	}

	fmt.Printf("serving from '%s' on port %s\nDon't forget to open your firewall port!\n", rootDir, listener.Addr().String())

	ipAddresses, err := getIPAddresses()
	if err != nil {
		log.Printf("couldn't get your IP addresses. Error: %q\n", err)
	} else {
		fmt.Printf("serving on:\n")
		for _, ipAddress := range ipAddresses {
			fmt.Printf("  http://%s:%d\n", ipAddress, port)
		}
	}

	err = http.Serve(listener, handler)
	if nil != err {
		log.Fatalf("couldn't start web server. Error: %s\n", err)
	}

}

func getIPAddresses() ([]net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var ips []net.IP
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ips = append(ips, v.IP)
			case *net.IPAddr:
				ips = append(ips, v.IP)
			}
		}
	}

	return ips, nil
}
