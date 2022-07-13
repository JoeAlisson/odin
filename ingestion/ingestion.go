package ingestion

import (
	"context"
	"time"
)

type Kind int32

const (
	CREATE Kind = 0
	UPDATE Kind = 1
	DELETE Kind = 2
)

var (
	kindMap = map[Kind]string{
		CREATE: "CREATE",
		UPDATE: "UPDATE",
		DELETE: "DELETE",
	}
)

func (k *Kind) String() string {
	v, ok := kindMap[*k]
	if !ok {
		return "UNKNOWN"
	}
	return v
}

type User struct {
	UID      string `json:"uid"`
	Username string `json:"username"`
	IP       string `json:"ip"`
}

type Changes struct {
	Entity      string            `json:"entity"`
	Description string            `json:"description"`
	Kind        Kind              `json:"Kind"`
	Values      map[string]string `json:"values"`
}

type AuditLog struct {
	System    string    `json:"system"`
	Timestamp time.Time `json:"timestamp"`
	User      User      `json:"user"`
	Changes   Changes   `json:"changes"`
}

type Ingestor interface {
	Ingest(context.Context, *AuditLog) error
}
