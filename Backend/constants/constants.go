package constants

import (
	models "github.com/alde326/EmailSearch/Backend/models"
)

const FILE_NAME string = "./enron_mail_20110402/maildir"

//const FILE_NAME string = "./allen-p"

const REFERENCE_FIRST_DATE string = "Mon, 01 Jan 2006 15:04:05 -0700 (MST)"
const REFERENCE_SECOND_DATE string = "Mon, 2 Jan 2006 15:04:05 -0700 (MST)"

//const EMAILS_AMOUNT int = 658252

//const EMAILS_AMOUNT int = 3402

const EMAILS_AMOUNT int = 1000

const MAX_WORKERS int = 5

const USER_NAME string = "admin"
const PASSWORD string = "Complexpass#123"
const INDEXER_NAME string = "email-indexer"

const SERVER string = "http://localhost:4080"
const ENDPOINT string = "/api/_bulkv2"
const ENDPOINT_INDEX string = "/api/index"
const METHOD_POST string = "POST"

var TEXT_VALUE = models.TextProperties{
	Type:          "text",
	Index:         true,
	Store:         false,
	Sortable:      false,
	Aggregatable:  false,
	Highlightable: false,
}

var DATE_VALUE = models.TextProperties{
	Type:          "date",
	Index:         true,
	Store:         false,
	Sortable:      true,
	Aggregatable:  false,
	Highlightable: false,
}

var PROPERTIES_VALUES = models.EmailProperties{
	Body:                    TEXT_VALUE,
	CFolder:                 TEXT_VALUE,
	CC:                      TEXT_VALUE,
	ContentTransferEncoding: TEXT_VALUE,
	ContentType:             TEXT_VALUE,
	Date:                    DATE_VALUE,
	DateSubEmail:            TEXT_VALUE,
	From:                    TEXT_VALUE,
	MessageID:               TEXT_VALUE,
	MimeVersion:             TEXT_VALUE,
	Sent:                    TEXT_VALUE,
	Subject:                 TEXT_VALUE,
	To:                      TEXT_VALUE,
	// XBcc:                    TEXT_VALUE,
	// XCc:                     TEXT_VALUE,
	// XFileName:               TEXT_VALUE,
	// XFrom:                   TEXT_VALUE,
	// XOrigin:                 TEXT_VALUE,
	// XTo:                     TEXT_VALUE,
}

var MAPPINGS_VALUES = models.EmailMappings{
	Properties: PROPERTIES_VALUES,
}

var EMAIL_INDEXER_INDEX = models.EmailIndexer{
	Name:        INDEXER_NAME,
	ShardNum:    5,
	StorageType: "disk",
	Mappings:    MAPPINGS_VALUES,
}
