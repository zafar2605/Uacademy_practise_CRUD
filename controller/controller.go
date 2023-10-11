package controller

import (
	"practise_CRUD/config"
	"practise_CRUD/storage"
)

type Con struct {
	cfg     *config.Config
	storage storage.StorageI
}

func NewController(cfg *config.Config, strg storage.StorageI) *Con {
	return &Con{
		cfg:     cfg,
		storage: strg,
	}
}
