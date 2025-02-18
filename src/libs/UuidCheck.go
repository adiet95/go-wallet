package libs

import "github.com/gofrs/uuid/v5"

func IsZero(id uuid.UUID) bool {
	for x := 0; x < 16; x++ {
		if id[x] != 0 {
			return false
		}
	}
	return true
}
