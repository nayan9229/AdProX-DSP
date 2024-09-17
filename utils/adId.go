package utils

import (
	"github.com/rs/xid"
)

func AdID() string {
	id := xid.New()
	return id.String()
}
