package models

import (
	"time"
)

type Email struct {
	MessageID    string    `json:"message_id"`
	Date         time.Time `json:"date"`
	DateSubEmail string    `json:"date_subemail"`
	From         string    `json:"from"`
	To           string    `json:"to"`
	Sent         string    `json:"sent"`
	Cc           string    `json:"cc"`
	// XFrom                   string    `json:"x_from"`
	// XTo                     string    `json:"x_to"`
	// XCc                     string    `json:"x_cc"`
	// XBcc                    string    `json:"x_bcc"`
	// XFolder                 string    `json:"c_folder"`
	// XOrigin                 string    `json:"x_origin"`
	// XFileName               string    `json:"x_file_name"`
	Subject                 string `json:"subject"`
	MimeVersion             string `json:"mime_version"`
	ContentType             string `json:"content_type"`
	ContentTransferEncoding string `json:"content_transfer_encoding"`
	Body                    string `json:"body"`
}

type ZincRequest struct {
	Index   string  `json:"index"`
	Records []Email `json:"records"`
}

type EmailIndexer struct {
	Name        string        `json:"name"`
	StorageType string        `json:"storage_type"`
	Mappings    EmailMappings `json:"mappings"`
	ShardNum    int           `json:"shard_num"`
}

type EmailMappings struct {
	Properties EmailProperties `json:"properties"`
}

type EmailProperties struct {
	Body                    TextProperties `json:"body"`
	CFolder                 TextProperties `json:"c_folder"`
	CC                      TextProperties `json:"cc"`
	ContentTransferEncoding TextProperties `json:"content_transfer_encoding"`
	ContentType             TextProperties `json:"content_type"`
	Date                    TextProperties `json:"date"`
	DateSubEmail            TextProperties `json:"date_subemail"`
	From                    TextProperties `json:"from"`
	MessageID               TextProperties `json:"message_id"`
	MimeVersion             TextProperties `json:"mime_version"`
	Sent                    TextProperties `json:"sent"`
	Subject                 TextProperties `json:"subject"`
	To                      TextProperties `json:"to"`
	// XBcc                    TextProperties `json:"x_bcc"`
	// XCc                     TextProperties `json:"x_cc"`
	// XFileName               TextProperties `json:"x_file_name"`
	// XFrom                   TextProperties `json:"x_from"`
	// XOrigin                 TextProperties `json:"x_origin"`
	// XTo                     TextProperties `json:"x_to"`
}

type TextProperties struct {
	Type          string `json:"type"`
	Index         bool   `json:"index"`
	Store         bool   `json:"store"`
	Sortable      bool   `json:"sortable"`
	Aggregatable  bool   `json:"aggregatable"`
	Highlightable bool   `json:"highlightable"`
}
