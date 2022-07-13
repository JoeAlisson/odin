package ingestion

import "context"

type InMemoryIngestor struct {
	logs map[string][]*AuditLog
}

func (i *InMemoryIngestor) Ingest(_ context.Context, log *AuditLog) error {
	i.logs[log.System] = append(i.logs[log.System], log)
	return nil
}
