package main

type room struct {
	// is a channle that holds incoming messages
	// that should be forwarded to the other clients
	forward chan []byte
}
