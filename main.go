package main

import (
    "fmt"
    "log"
    "net"
    "os"
)

func main() {
    port, ok := os.LookupEnv("PORT")
    if !ok {
        port = "8000"
    }

    // PORT is conventionally a number passed as an environment variable
    // However, the golang standard library takes a listener string
    //
    // "host:port" - to listen on a specific host and port tuple.
    // ":port" to listen on all (unspecified) hosts.
    listenOn := fmt.Sprintf(":%s", port)
    server, err := net.Listen("tcp", listenOn)
    if err != nil {
        log.Fatalf("Failed to bind to: %s %v", listenOn, err)
    }

    log.Printf("Waiting for tcp connections on: %s", server.Addr().String())
    for {
        // Block until we establish a connection
        client, err := server.Accept()
        if err != nil {
            // Depending on the application, crashing might not be
            // the most appropriate thing to do.
            log.Fatalf("Failed to accept incomming connections: %v", err)
        }

        // The `Conn` object (in this scope it is the `client` variable)
        // is documented here: https://golang.org/pkg/net/#Conn
        // It has functions like `Read()`, `Write()` and `Close()`.
        // For our purposes, we'll use `RemoteAddr()`.
        addr := client.RemoteAddr()
        log.Printf("Incoming connection from: %s", addr.String())

        // We won't do anything else with the connection.
        // An application can spawn a goroutine or something fancy.
        client.Close()
    }
}
