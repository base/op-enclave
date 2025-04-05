package main

import (
	"net/http"

	enclave2 "github.com/base/op-enclave/op-enclave/enclave"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/mdlayher/vsock"
)

func main() {
	oplog.SetupDefaults()

	s := rpc.NewServer()
	serv, err := enclave2.NewServer()
	if err != nil {
		log.Crit("Error creating API server", "error", err)
		return
	}

	err = s.RegisterName(enclave2.Namespace, serv)
	if err != nil {
		log.Crit("Error registering API", "error", err)
		return
	}

	listener, err := vsock.Listen(1234, &vsock.Config{})
	if err != nil {
		log.Warn("Error opening vsock listener, running in HTTP mode", "error", err)
		err = http.ListenAndServe(":1234", s)
		if err != nil {
			log.Crit("Error starting HTTP server", "error", err)
		}
	} else {
		err = s.ServeListener(listener)
		if err != nil {
			log.Crit("Error starting server", "error", err)
		}
	}
}
