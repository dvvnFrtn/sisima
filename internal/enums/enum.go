package enums

import (
	"fmt"
	"strings"
)

type DeleteOptions string

const (
	Normal   DeleteOptions = "NORMAL"
	Rollback DeleteOptions = "ROLLBACK"
	Hard     DeleteOptions = "HARD"
)

func ParseDeleteOption(s string) (DeleteOptions, error) {
	switch strings.ToUpper(s) {
	case string(Normal):
		return Normal, nil
	case string(Rollback):
		return Rollback, nil
	case string(Hard):
		return Hard, nil
	default:
		return "", fmt.Errorf("invalid DeleteOptions: %q", s)
	}
}
