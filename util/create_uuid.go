package util

import (
	"strings"

	"github.com/google/uuid"
)

func CreateUuid() string {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuid
}
