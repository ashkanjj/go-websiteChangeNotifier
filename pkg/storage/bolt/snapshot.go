package bolt

import (
	"time"
)

type Snapshot struct {
	WebsiteID string
	Date      time.Time
	Value     string
}
