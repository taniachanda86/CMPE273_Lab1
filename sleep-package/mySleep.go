package sleep

import (
    "time"
)

func Sleep(secs int) {
    select {
          case <-time.After(time.Duration(secs) * time.Second):
            return
    }
}