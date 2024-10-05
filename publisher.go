package outbox

import (
	"database/sql"

	"github.com/thanhtranna/outbox/internal/time"
	"github.com/thanhtranna/outbox/internal/uuid"
)

// Publisher encapsulates the save functionality of the outbox pattern
type Publisher struct {
	store Store
	time  time.Provider
	uuid  uuid.Provider
}

// NewPublisher is the Publisher constructor
func NewPublisher(store Store) Publisher {
	return Publisher{store: store, time: time.NewTimeProvider(), uuid: uuid.NewUUIDProvider()}
}

// Send stores the provided Message within the provided sql.Tx
func (o Publisher) Send(msg Message, tx *sql.Tx) error {
	newID := o.uuid.NewUUID()
	record := Record{
		ID:          newID,
		Message:     msg,
		State:       PendingDelivery,
		CreatedOn:   o.time.Now().UTC(),
		LockID:      nil,
		LockedOn:    nil,
		ProcessedOn: nil,
	}

	return o.store.AddRecordTx(record, tx)
}
