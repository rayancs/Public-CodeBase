package internals

import "time"

func DB_MAX_OPEN_CONNS() int        { return 500 }
func DB_MAX_IDLE_CONNS() int        { return 50 }
func DB_SET_MAX_CONN_LIFETIME() time.Duration { return 5 * time.Minute }
