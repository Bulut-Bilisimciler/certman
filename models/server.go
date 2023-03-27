package models

import "time"

// Server
type Server struct {
	// core opts.
	ServerName   string   `json:"server_name" example:"www.google.com"`
	Port         int      `json:"port" example:"443"`
	ProtocolType string   `json:"protocol_type" example:"tcp, tls, http, https"`
	Debug        bool     `json:"debug" example:"true, false"`
	Rules        []string `json:"rules" example:"[\"http.status_code == 200\"]"`

	// information opts.
	Title         string     `json:"title" example:"Google"`
	Description   string     `json:"description" example:"Google"`
	Keywords      string     `json:"keywords" example:"google, search engine"`
	Status        string     `json:"status" example:"up, down, unknown, timeout"`
	LastCheckedAt *time.Time `json:"last_checked_at" example:"2020-01-01 00:00:00"`

	// trigger functions
	ValidateFn         func() error `json:"validate_fn,omitempty"`
	OnBeforeConnection func() error `json:"on_before_connection,omitempty"`
	OnAfterConnection  func() error `json:"on_after_connection,omitempty"`
	OnCheckSuccess     func() error `json:"on_check_success,omitempty"`
	OnCheckFailure     func() error `json:"on_check_failure,omitempty"`
}
