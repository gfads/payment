package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/microservices-demo/payment"
	"golang.org/x/net/context"
)

const (
	ServiceName = "payment"
)

func main() {
	var (
		port          = flag.String("port", "8080", "Port to bind HTTP listener")
		declineAmount = flag.Float64("decline", 105, "Decline payments over certain amount")
	)
	flag.Parse()
	// Mechanical stuff.
	errc := make(chan error)
	ctx := context.Background()

	handler, logger := payment.WireUp(ctx, float32(*declineAmount), ServiceName)

	// Create and launch the HTTP server.
	go func() {
		logger.Log("transport", "HTTP", "port", *port)
		errc <- http.ListenAndServe(":"+*port, handler)
	}()

	// Capture interrupts.
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("exit", <-errc)
}
