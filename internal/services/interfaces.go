package services

import (
	"gocoop/internal/protocols"
	"gocoop/pkg/coop"
)

//------------------------------------------------------------------------------
// Interfaces
//------------------------------------------------------------------------------

// CoopService is the interface
type CoopService interface {
	Get() *coop.Coop
	Update(protocols.CoopUpdateRequestController) error
	Open() error
	Close() error
}

// JwtService is the service for JWT.
type JwtService interface {
	Create(string) (map[string]string, error)
	Get(string) (string, error)
}
