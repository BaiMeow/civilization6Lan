package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
)

func main() {
	log.Println("start")
	var listener *net.UDPConn
	for {
		var err error
		port := 62900 + rand.Intn(100)
		listener, err = net.ListenUDP("udp4", &net.UDPAddr{IP: nil, Port: port})
		if err != nil {
			log.Println(err)
		}
		break
	}
	var lip, rip string
	fmt.Println("input local(used to send packet) ip(like 172.16.4.1):")
	fmt.Scanln(&lip)
	if lip == "" {
		addrs, _ := net.InterfaceAddrs()
		lip = addrs[0].String()
	}
	log.Println("local ip:", lip)
	fmt.Println("input server ip(like 172.16.4.1):")
	fmt.Scanln(&rip)
	log.Println("remote ip:", rip)

	for {
		var data [4]byte
		_, addr, err := listener.ReadFromUDPAddrPort(data[:])
		if data != [4]byte{141, 232, 252, 112} {
			continue
		}
		if err != nil {
			log.Println(err)
		}
		log.Println("receive civilization broadcast from:", addr)
		log.Println("sending to remote")
		sender, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.ParseIP(lip)})
		if err != nil {
			log.Println(err)
			continue
		}
		for port := 62900; port < 63000; port++ {
			_, err = sender.WriteTo(data[:], &net.UDPAddr{IP: net.ParseIP(rip), Port: port})
			if err != nil {
				log.Println(err)
				continue
			}
		}
		log.Println("finished")

	}

}
