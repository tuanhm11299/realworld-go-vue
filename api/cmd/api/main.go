// Command api is the Conduit backend.
//
// Right now it is a *walking skeleton*: it boots, connects to Postgres, serves
// health probes, and shuts down cleanly — and nothing else. Every /api/** route
// is unimplemented on purpose, which is why the whole Hurl conformance suite is
// red. Turning it green, milestone by milestone, is the exercise. See LEARNING.md.
//
// Note what this file deliberately does NOT do: it reads env vars inline rather
// than using internal/config, and wires routes inline rather than using
// internal/httpx. Building those packages properly is milestone M0/M2 work —
// this is just enough to prove the harness runs.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const shutdownGrace = 10 * time.Second

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	if err := run(context.Background(), logger); err != nil {
		logger.Error("fatal", "err", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, logger *slog.Logger) error {
	// Stop accepting work as soon as the operator (or the container runtime)
	// asks us to. ctx is cancelled on SIGINT/SIGTERM.
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	databaseURL := env("DATABASE_URL", "postgres://conduit:conduit@localhost:5432/conduit?sslmode=disable")
	port := env("PORT", "8080")

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return fmt.Errorf("create connection pool: %w", err)
	}
	defer pool.Close()

	// Fail fast and loudly: a server that starts without its database just
	// moves the error to the first request, where it is harder to read.
	pingCtx, cancelPing := context.WithTimeout(ctx, 5*time.Second)
	defer cancelPing()
	if err := pool.Ping(pingCtx); err != nil {
		return fmt.Errorf("ping database (is `make up` running?): %w", err)
	}
	logger.Info("database connected")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthz)
	mux.HandleFunc("GET /readyz", readyz(pool))

	// TODO(M2+): mount the Conduit API here. Suggested shape once you build
	// internal/httpx and internal/handler:
	//
	//	mux.Handle("/api/", http.StripPrefix("/api", httpx.NewRouter(deps)))
	//
	// Until then every /api/** request 404s, and every Hurl test fails.

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		logger.Info("listening", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		return fmt.Errorf("serve: %w", err)
	case <-ctx.Done():
		logger.Info("shutdown signal received, draining", "grace", shutdownGrace.String())
	}

	// Give in-flight requests a chance to finish before we drop them.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownGrace)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("graceful shutdown: %w", err)
	}
	logger.Info("shutdown complete")
	return nil
}

// healthz reports that the process is alive. It must not touch dependencies:
// a liveness probe that fails when the database blips gets your pod restarted
// for no reason.
func healthz(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// readyz reports that the process can actually serve traffic, dependencies
// included. This is the one a load balancer should poll.
func readyz(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		if err := pool.Ping(ctx); err != nil {
			writeJSON(w, http.StatusServiceUnavailable, map[string]string{
				"status": "unavailable",
				"reason": "database",
			})
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"status": "ready"})
	}
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		slog.Error("write response", "err", err)
	}
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
