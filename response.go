package trimet

import (
	"time"
)

type Response struct {
	QueryTime *time.Time `json:"queryTime"`
	Error     string     `json:"errorMessage"`
}
