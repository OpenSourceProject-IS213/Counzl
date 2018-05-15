package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const ( // Location of the keys
	CA          = "./database/.certificate/CA/ca.crt"
	client_crt  = "./database/.certificate/client.crt"
	client_key  = "./database/.certificate/client.key" 
	server_ip   = "<(IP) address/localhost>"
	server_port = ":<port-number>"
)
// Validating keys and certificates
func CheckCerts_client() {
	log.SetFlags(log.Lshortfile)

	cert, err := tls.LoadX509KeyPair(client_crt, client_key)
	if err != nil {
		log.Fatal(err)
	}
	caCert, err := ioutil.ReadFile(CA) // Reads the CA certificate that validates the server's IP
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool() // Creates a new certification pool
	caCertPool.AppendCertsFromPEM(caCert)

	conf := &tls.Config{ // This configuration are a composition of the CA-cert, client.key and client.crt
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{cert},
	}

	err := dialMUX(conf) // Calls the dial-function, if the programs return error, it will exit with exit code: 1
	if err != nil {
		fmt.Println(cmd.ChangeColor("404", "red") + ": Får ikke kontakt med server. \nDette betyr at du dessverre ikke får opprettet bruker :'(")
		os.Exit(1)
	}
	return id

}

func dialMUX(conf *tls.Config) error {
	// Dials the server, you can see that the third parameter is the configuration from line 36 to validate the connection
	conn, err := tls.Dial("tcp", server_ip+server_port, conf)
	if err != nil {

		return err // return error   
	}
	defer conn.Close()
	var message = "message test"
	// Writes to the established connection (writes to server)
	n, err := conn.Write([]byte(message + "\n"))
	if err != nil {
		log.Println(n, err)
		return err
	}

	// Makes a buffer for incoming messages from server
	buf := make([]byte, 100)
	msg_from_server, err = conn.Read(buf)
	if err != nil {
		log.Println(msg_from_server, err)
		return err
	}

	println((string(buf[:msg_from_server])))
	return nil // return nil if the connection was a success.
}
