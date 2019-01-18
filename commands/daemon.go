package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Worker interface {
	Start() error
	API() error
	Stop() error
}
type WorkerFactory func(version, db string) (Worker, error)

func daemonize(db string, wf WorkerFactory) {
	w, err := wf(Version, db)
	if err != nil {
		log.Fatal("ERROR: Failed to initialize controller. Error:", err)
	}
	if err := w.Start(); err != nil {
		log.Println("ERROR: Failed to start controller. Error:", err)
	}
	if err := w.API(); err != nil {
		log.Println("ERROR: Failed to start API server. Error:", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGUSR2, syscall.SIGTERM)
	for {
		select {
		case s := <-ch:
			switch s {
			case syscall.SIGTERM:
				w.Stop()
				return
			case os.Interrupt:
				w.Stop()
				return
			case syscall.SIGUSR2:
			}
		}
	}
}
