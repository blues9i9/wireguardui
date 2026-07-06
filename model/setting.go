package model

import (
	"time"
)

// GlobalSetting model
type GlobalSetting struct {
	EndpointAddress     string    `json:"endpoint_address"`
	DNSServers          []string  `json:"dns_servers"`
	MTU                 int       `json:"mtu,string"`
	PersistentKeepalive int       `json:"persistent_keepalive,string"`
	FirewallMark        string    `json:"firewall_mark"`
	Table               string    `json:"table"`
	ConfigFilePath      string    `json:"config_file_path"`
	SmtpHostname        string    `json:"smtp_hostname"`
	SmtpPort            int       `json:"smtp_port,string"`
	SmtpUsername        string    `json:"smtp_username"`
	SmtpPassword        string    `json:"smtp_password"`
	SmtpAuthType        string    `json:"smtp_auth_type"`
	SmtpEncryption      string    `json:"smtp_encryption"`
	SmtpNoTLSCheck      bool      `json:"smtp_no_tls_check"`
	EmailFrom           string    `json:"email_from"`
	EmailFromName       string    `json:"email_from_name"`
	EnableAutoEmail     bool      `json:"enable_auto_email"`
	UpdatedAt           time.Time `json:"updated_at"`
}
