package main

import (
	"flag"
	"fmt"
)

func main() {
	var keyPath, ip string
	flag.StringVar(&keyPath, "P", "empty", "input path of public key file")
	flag.StringVar(&keyPath, "path", "empty", "input path of public key file")

	flag.StringVar(&ip, "ip", "empty", "input ip address of the target host")

	flag.Parse()

	fmt.Printf("path: %v\n", keyPath)
	fmt.Printf("ip: %v\n", ip)
	fmt.Printf("args: %v\n", flag.Args())
}
