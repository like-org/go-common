package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

type MyConfig struct {
	Version int
	Name    string
	Tags    []string
}

func main() {
	log.Print("app interface starting....")

	/*

		0.0.0.0:80/acc/*  ===> 0.0.0.0:5051/* (account http server)

	*/
	l, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatal(err)
	}

	for {
		client, err := l.Accept()
		if err != nil {
			log.Print(err)
			break
		}
		go handleClientRequest(client)
	}
}

func handleClientRequest(client net.Conn) {
	var data [1024]byte
	n, err := client.Read(data[:])
	if err != nil {
		log.Print(err)
		return
	}
	info := strings.Split(string(data[:]), "\n")
	fmt.Println(info[0])
	infos := strings.Split(info[0], " ")
	path := infos[1]

	paths := strings.Split(path, "/")

	fmt.Println(paths[1])
	if paths[1] == "acc" {
		server, err := net.Dial("tcp", ":5051")
		if err != nil {
			log.Print(err)
			return
		}
		server.Write(data[:n])
		go io.Copy(server, client)
		io.Copy(client, server)
	}
}
