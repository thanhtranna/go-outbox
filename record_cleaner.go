package outbox

import (
	time2 "time"

	"github.com/thanhtranna/outbox/internal/time"
)

type recordCleaner struct {
	store             Store
	time              time.Provider
	maxRecordLifetime time2.Duration
}

func newRecordCleaner(store Store, maxRecordLifetime time2.Duration) recordCleaner {
	return recordCleaner{
		maxRecordLifetime: maxRecordLifetime,
		store:             store,
		time:              time.NewTimeProvider(),
	}
}

func (d recordCleaner) RemoveExpiredMessages() error {
	expiryTime := d.time.Now().UTC().Add(-d.maxRecordLifetime)
	err := d.store.RemoveRecordsBeforeDatetime(expiryTime)
	if err != nil {
		return err
	}

	return nil
}
