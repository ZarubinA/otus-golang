package hellonow

import (
	"github.com/beevik/ntp"
	"time"
)

const ntpServer = "0.beevik-ntp.pool.ntp.org"

// GetTime ...
func GetTime() (time.Time, error) {
	resp, err := ntp.Query(ntpServer)
	return resp.Time, err
}
