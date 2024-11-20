package http

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/mohammadne/takhir/internal/http/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Server struct {
	logger   *zap.Logger
	handlers *Handlers

	monitorApp *fiber.App
	clientApp  *fiber.App
}

type Handlers struct {
	Healthz handlers.Healthz
	Items   handlers.Items
}

func New(log *zap.Logger, handlers *Handlers) *Server {
	server := &Server{logger: log, handlers: handlers}
	fiberConfig := fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal}

	// ----------------------------------------- Monitor Endpoints

	server.monitorApp = fiber.New(fiberConfig)

	server.monitorApp.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	healthzGroup := server.monitorApp.Group("healthz")
	healthzGroup.Get("/liveness", server.handlers.Healthz.Liveness)
	healthzGroup.Get("/readiness", server.handlers.Healthz.Readiness)

	// ----------------------------------------- Client Endpoints

	server.clientApp = fiber.New(fiberConfig)

	v1 := server.clientApp.Group("api/v1")

	itemsGroup := v1.Group("items")
	itemsGroup.Get("/", server.handlers.Items.List)

	// auth := v1.Group("auth")
	// auth.Post("/register", server.register)
	// auth.Post("/login", server.login)

	// contacts := v1.Group("contacts", server.fetchUserId)
	// contacts.Get("/", server.getContacts)
	// contacts.Post("/", server.createContact)
	// contacts.Get("/:id", server.getContact)
	// contacts.Put("/:id", server.updateContact)
	// contacts.Delete("/:id", server.deleteContact)

	return server
}

func (server *Server) Serve(monitor, client int) {
	runnables := []struct {
		port        int
		app         *fiber.App
		description string
	}{
		{monitor, server.monitorApp, "monitor server"},
		{client, server.clientApp, "client server"},
	}

	for _, runnable := range runnables {
		go func() {
			address := fmt.Sprintf("0.0.0.0:%d", runnable.port)
			fields := []zapcore.Field{zap.String("address", address),
				zap.String("description", runnable.description)}

			server.logger.Info("starting server", fields...)
			err := runnable.app.Listen(address)
			fields = append(fields, zap.Error(err))
			server.logger.Fatal("error resolving server", fields...)
		}()
	}
}
