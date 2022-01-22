package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	stripeClient "github.com/stripe/stripe-go/client"
)

var (
	client   *stripeClient.API
	fsClient *firestore.Client
)

func main() {
	mainMux := http.NewServeMux()

	mainMux.HandleFunc("/create-subscription", createUserSubscriptionHandler)
	mainSrv := &http.Server{
		Addr:    "4321",
		Handler: mainMux,
	}

	client = stripeClient.New(os.Getenv("STRIPE_API_KEY"), nil)
	cli, err := firestore.NewClient(context.Background(), os.Getenv("GCP_PROJECT"))
	if err != nil {
		log.Fatalf("Failed to create firestore client: %v", err)
	}
	fsClient = cli

	if err := mainSrv.ListenAndServe(); err != nil {
		return
	}
}
