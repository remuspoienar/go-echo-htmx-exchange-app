package model

import (
	"gorm.io/gorm"
)

type Asset struct {
	gorm.Model
	ExternalId       string
	Name             string
	Status           string
	Type             string
	MinSize          string
	MaxPrecision     string
	DefaultNetworkId *int

	DefaultNetwork    *Network
	SupportedNetworks []Network `gorm:"many2many:assets_networks;" json:"-"`
	Users             []User    `gorm:"many2many:users_pins;" json:"-"`
}

type AssetIn struct {
	ExternalId        string      `json:"id"`
	Name              string      `json:"name"`
	Status            string      `json:"status"`
	MinSize           string      `json:"min_size"`
	MaxPrecision      string      `json:"max_precision"`
	DefaultNetwork    string      `json:"default_network"`
	SupportedNetworks []NetworkIn `json:"supported_networks"`
	Details           struct {
		Type string `json:"type"`
	} `json:"details"`
}
