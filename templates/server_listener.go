package main

import (
	"log"
	"crypto/tls"
	"bufio"
	"net"
	"fmt"
)

const (
	server_crt = "./database/.certificate/server.crt"
	server_key = "./database/.certificate/server.key"
	LN_PORT = ":<port number>"
	network = "tcp"
)

func Initialise_Listener() {
	log.SetFlags(log.Lshortfile)

	certs, err := tls.LoadX509KeyPair(server_crt, server_key) // Load keypair
	if err != nil {
		log.Println(err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{certs}} // Makes tls configuration with server.crt, server.key.
	ln, err := tls.Listen(network, LN_PORT, config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept() // accepting incoming connection with the right keys and certs
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn) // calls the handler with a distinct connection to a client as a gourotine 
	}
}

// Function to handle the traffic to a given port and check if the keys and certs are valid
func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		from_client, err := r.ReadString('\n') // Reads from client
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(from_client)
		to_client := "Hello to client"

		n, err := conn.Write([]byte(to_client + "\n")) // Writes to connection (to client)
		if err != nil {
			log.Println(n, err)
			return
		}

	}
}