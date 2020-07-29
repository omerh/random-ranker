package helpers

import "log"

// IsRequiredFlagsOK to check root required flags
func IsRequiredFlagsOK(p1 float64) bool {
	if p1 == 0 {
		log.Printf("Missing --p1 flag, example: --p1 10092")
		return false
	}
	return true
}
