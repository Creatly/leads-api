package server

import (
	"encoding/json"
	"fmt"
	"github.com/creatly/leads-api/internal/models"
	"net/http"
)

type CRM interface {
	SaveLead(lead models.Lead) error
}

type Server struct {
	server *http.Server
	crm    CRM
}

func New(port string, crm CRM) *Server {
	return &Server{
		server: &http.Server{
			Addr: fmt.Sprintf(":%s", port),
		},
		crm: crm,
	}
}

func (s *Server) Init() error {
	mux := new(http.ServeMux)
	mux.HandleFunc("/leads", s.leadsHandler)
	s.server.Handler = mux

	return s.server.ListenAndServe()
}

func (s *Server) leadsHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	var lead models.Lead
	if err := json.NewDecoder(req.Body).Decode(&lead); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := s.crm.SaveLead(lead); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
