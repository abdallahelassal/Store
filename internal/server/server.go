package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	Port   string
	Router *gin.Engine
	Logger *zap.Logger
}

func NewServer(port string,router *gin.Engine,logger *zap.Logger)*Server{
	return &Server{
		Port: port,
		Router: router,
		Logger: logger,
	}
}

func (s *Server) Start()error{
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s",s.Port),
		Handler: s.Router,
	}

	go func(){
		s.Logger.Info("starting server",zap.String("port", s.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed{
			s.Logger.Fatal("failed start server",zap.Error(err))
		}
	}()

	quit := make(chan os.Signal,1)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<-quit
	s.Logger.Info("Shutting down server.....")

	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second) 
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		s.Logger.Fatal("server forced shutdown",zap.Error(err))
		return err
	}
	s.Logger.Info("server exited successfully")
	return nil 
}