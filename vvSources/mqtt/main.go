package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const protocol = "tcp"
const broker = ""
const port = 1883
const topic = "devops/1"
const username = ""
const password = ""

func main() {
	client := createMqttClient()
	for {
		subscribe(client)
		time.Sleep(time.Second * 1)
	} // we use goroutine to run the subscription function
	// pause 1s to wait for the subscription function to be ready
	//publish(client)
}

func createMqttClient() mqtt.Client {
	connectAddress := "mqtts://ssq.staging.amity.co:443"
	rand.Seed(time.Now().UnixNano())
	clientID := fmt.Sprintf("go-client-%d", rand.Int())

	fmt.Println("connect address: ", connectAddress)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(connectAddress)
	opts.SetUsername("65f81dc0f521d4ff7db14a21")
	opts.SetPassword("eyJhbGciOiJSUzI1NiIsImtpZCI6IkQ4cVQzRHA5dlNvX0w5d2l4YmF1QlFkNHFLbk5neFhqWHJHakhlTGxVaDAifQ.eyJ1c2VyIjp7InVzZXJJZCI6IjY1ZjgxZGMwZjUyMWQ0ZmY3ZGIxNGEyMSIsInB1YmxpY1VzZXJJZCI6ImR0aGVyIiwiZGV2aWNlSW5mbyI6eyJraW5kIjoid2ViIiwibW9kZWwiOiJnZWNrbyNNb3ppbGxhLzUuMCAoTWFjaW50b3NoOyBJbnRlbCBNYWMgT1MgWCAxMF8xNV83KSBBcHBsZVdlYktpdC81MzcuMzYgKEtIVE1MLCBsaWtlIEdlY2tvKSBDaHJvbWUvMTIyLjAuMC4wIFNhZmFyaS81MzcuMzYiLCJzZGtWZXJzaW9uIjoidjYuMTkuMC1lc20ifSwibmV0d29ya0lkIjoiNjMxZjE2YmUxZTQ0MDQwMGRhNTY2M2IwIiwiZGlzcGxheU5hbWUiOiJkdGhlciIsInJlZnJlc2hUb2tlbiI6Ijc5ZTY1ODk3OTU4YTg4MjZjNTg2ZGIzZWJmNTI2NjdhNjBmNDQ3YTA0ZDZmZDNkOGNjZDVjYWRlNTA3NWIwYmI5MzE4MzE1ZWQxMGVkNGQxIn0sInN1YiI6IjY1ZjgxZGMwZjUyMWQ0ZmY3ZGIxNGEyMSIsImlzcyI6Imh0dHBzOi8vYXBpLnN0YWdpbmcuYW1pdHkuY28iLCJpYXQiOjE3MTA3NTkzNjAsImV4cCI6MTcxMzM1MTM2MH0.U4hGxT15Pxe4z34ogg8BkGK9M6B9lkraUBqc4Hv_7-Io0qI2K-lCg7n1rNAADQbPcuGmqNx64v2dPG2_Nf7NkDJxmq7xdvuKvFHPD8n150aD13iZoY9eVhogbokQoeOKolvcbpfMajWyRyJt_JSlau0z26pDjSmOEHiSYZaZKA-ZjOQBSuxWm610SS5IhxRXsVr86gBGxlvQR_dCuC8xDOGC86qrjYSHkR_t2h7tCtO4jyDN74-1bUhLq6nJ2x1rAG9QcojLZ4ZQZnOEH1KBwyhhZ8oF85PwS3bUjPBIu-XCYrT4O1Zl5yo4iVRX5qwYAsRcyrlyITZUVg-epefs1w")
	opts.SetClientID(clientID)
	opts.SetKeepAlive(time.Second * 60)

	// Optional: set server CA
	// opts.SetTLSConfig(loadTLSConfig("caFilePath"))

	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.WaitTimeout(3*time.Second) && token.Error() != nil {
		log.Fatal(token.Error())
	}
	return client
}

func publish(client mqtt.Client) {
	qos := 0
	msgCount := 0
	for {
		payload := fmt.Sprintf("message: %d!", msgCount)
		if token := client.Publish(topic, byte(qos), false, payload); token.Wait() && token.Error() != nil {
			fmt.Printf("publish failed, topic: %s, payload: %s\n", topic, payload)
		} else {
			fmt.Printf("publish success, topic: %s, payload: %s\n", topic, payload)
		}
		msgCount++
		time.Sleep(time.Second * 1)
	}
}

func subscribe(client mqtt.Client) {
	qos := 0
	client.Subscribe(topic, byte(qos), func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received `%s` from `%s` topic\n", msg.Payload(), msg.Topic())
	})
}

func loadTLSConfig(caFile string) *tls.Config {
	// load tls config
	var tlsConfig tls.Config
	tlsConfig.InsecureSkipVerify = false
	if caFile != "" {
		certpool := x509.NewCertPool()
		ca, err := ioutil.ReadFile(caFile)
		if err != nil {
			log.Fatal(err.Error())
		}
		certpool.AppendCertsFromPEM(ca)
		tlsConfig.RootCAs = certpool
	}
	return &tlsConfig
}
