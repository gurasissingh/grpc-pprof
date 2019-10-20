package sleep

import (
	"time"
)

func SleepFunc() {
	for i := 0; i < 500; i++ {
		time.Sleep(1 * time.Millisecond)
	}
}
