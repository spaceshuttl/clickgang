package main

import (
	"log"

    "clickgang/server"
)

func main() {
    log.SetFlags(log.LstdFlags|log.Lshortfile)

    go func (){
        log.Printf("starting game server at %s")
        srv, err := server.New("localhost:3117")
        if err != nil {
            log.Fatal(err)
        }
        log.Fatal(srv.Start())
    }()

    go func (){
        srv2, err := server.NewWeb("localhost:4040")
        if err != nil {
            log.Fatal(err)
        }
        log.Fatal(srv2.Start())
    }()

    select{}
}