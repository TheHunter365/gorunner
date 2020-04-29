package utils

import (
	"log"
	"time"
)

//TimeTrack function
func TimeTrack(start time.Time, name string) {
	since := time.Since(start)
	log.Printf("%s Took %s", name, since)
}


