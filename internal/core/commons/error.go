package commons

import (
	"github.com/joomcode/errorx"
	"strings"
)

const duplicatedEntryMsgSubstring = "duplicate entry"

func IsDuplicatedEntryError(err error) bool {
	return strings.Contains(strings.ToLower(err.Error()), duplicatedEntryMsgSubstring)
}

var (
	NotFound     = errorx.CommonErrors.NewType("not_found", errorx.NotFound())
	BadRequest   = errorx.CommonErrors.NewType("bad_request")
	Unauthorized = errorx.CommonErrors.NewType("unauthorized")
)
