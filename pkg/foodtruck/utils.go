package foodtruck

import (
	"fmt"
	"time"
)

// helper function to return current HHMM in any format
// default format used by SODAS API is `HH:MM`
func getCurrentHHMM(format string) string {
	return fmt.Sprintf(format, time.Now().Hour(), time.Now().Minute())
}
