package server

import (
	"context"
	"fmt"

	"github.com/lutfiharidha/google-trace/internal/config"
	userRepo "github.com/lutfiharidha/google-trace/pkg/adapter/repository/db"
	controller "github.com/lutfiharidha/google-trace/pkg/infrasturcture/server/controller/user"
	"github.com/lutfiharidha/google-trace/pkg/shared/tracing"
	userUc "github.com/lutfiharidha/google-trace/pkg/usecase/user"
	"go.opentelemetry.io/otel"

	"log"
	"net/http"
)

func RunServer() {
	cfg := config.GetConfig()
	ctx := context.Background()
	projectID := cfg.Client.Google.ProjectID

	tracer, errReport := tracing.Init(ctx, projectID, "Gtrace")
	defer tracer.Shutdown(ctx) // flushes any pending spans, and closes connections.
	defer errReport.Close()
	otel.SetTracerProvider(tracer)

	// Create a trace span around your HTTP handler.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Your handler logic goes here.
		w.Write([]byte("Hello, World!"))
	})

	//user
	userRepo := userRepo.NewUserRepository()
	userUc := userUc.NewUserUsecase(userRepo)
	userController := controller.NewUserController(userUc)
	http.HandleFunc("/user", userController.GetUser)

	// Start the HTTP server.
	port := fmt.Sprintf(":%d", cfg.Server.Gtrace.Port)
	log.Fatal(http.ListenAndServe(port, nil))
}
