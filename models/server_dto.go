package models

import "time"

// RegisterServerDto
type RegisterServerDto struct {
	// core opts.
	ServerName    string   `json:"server_name" example:"www.google.com"`
	Port          int      `json:"port" example:"443"`
	ProtocolType  string   `json:"protocol_type" example:"tcp, tls, http, https"`
	Debug         bool     `json:"debug" example:"true, false"`
	Rules         []string `json:"rules" example:"[\"http.status_code == 200\"]"`
	CheckDuration int      `json:"check_duration" example:"60"`

	// information opts.
	Title         string     `json:"title" example:"Google"`
	Description   string     `json:"description" example:"Google"`
	Keywords      string     `json:"keywords" example:"google, search engine"`
	Status        string     `json:"status" example:"up, down, unknown, timeout"`
	LastCheckedAt *time.Time `json:"last_checked_at" example:"2020-01-01 00:00:00"`
}
