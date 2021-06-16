package server

import (
	"context"
	"fmt"
	"github.com/creatly/leads-api/internal/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	allowedOrigins = []string{"https://creatly.me", "http://localhost", "http://localhost:3000"}
)

type CRM interface {
	SaveLead(lead models.Lead) error
}

type Server struct {
	server *http.Server
	crm    CRM
}

func New(port int, crm CRM) *Server {
	return &Server{
		server: &http.Server{
			Addr: fmt.Sprintf(":%d", port),
		},
		crm: crm,
	}
}

func (s *Server) Init() error {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowMethods: []string{"POST"},
	}))
	r.POST("/leads", s.saveLead)
	s.server.Handler = r

	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) saveLead(c *gin.Context) {
	var lead models.Lead
	if err := c.BindJSON(&lead); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := s.crm.SaveLead(lead); err != nil {
		log.Printf("SaveLead() error: %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
