package hellonow

import (
	"github.com/beevik/ntp"
	"time"
)

const ntpServer = "0.beevik-ntp.pool.ntp.org"

// GetTime returns a response from remote NTP server.
// It contains current time and some error if occurred
func GetTime() (time.Time, error) {
	resp, err := ntp.Query(ntpServer)
	return resp.Time, err
}
