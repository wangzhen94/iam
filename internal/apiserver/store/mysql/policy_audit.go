package mysql

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type policyAudit struct {
	db *gorm.DB
}

func newPolicyAudits(ds *datastore) *policyAudit {
	return &policyAudit{db: ds.db}
}

func (p *policyAudit) ClearOutdated(ctx context.Context, maxReserveDays int) (int64, error) {
	date := time.Now().AddDate(0, 0, -maxReserveDays).Format("2006-01-02 15:04:05")

	d := p.db.Exec("delete from policy_audit where deleteAt < ?", date)

	return d.RowsAffected, d.Error
}
