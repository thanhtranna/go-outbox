package outbox

import (
	time2 "time"

	"github.com/thanhtranna/outbox/internal/time"
)

type recordUnlocker struct {
	store                   Store
	time                    time.Provider
	maxLockTimeDurationMins time2.Duration
}

func newRecordUnlocker(store Store, maxLockTimeDurationMins time2.Duration) recordUnlocker {
	return recordUnlocker{
		maxLockTimeDurationMins: maxLockTimeDurationMins,
		store:                   store,
		time:                    time.NewTimeProvider(),
	}
}

func (d recordUnlocker) UnlockExpiredMessages() error {
	expiryTime := d.time.Now().UTC().Add(-d.maxLockTimeDurationMins)
	clearErr := d.store.ClearLocksWithDurationBeforeDate(expiryTime)
	if clearErr != nil {
		return clearErr
	}

	return nil
}
