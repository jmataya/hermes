package models

import "github.com/jmataya/hermes/utils"

// Model is an interface that provides the basic structure for interacting
// with database tables. One can safely assume that a struct conforming to
// this interface has a corresponding database table.
type Model interface {
	ID() int
	Create(utils.DB) error
	Update(utils.DB) error
	Delete(utils.DB) error
}
