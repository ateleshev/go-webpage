package webpage

import (
	"fmt"
)

const (
	ERROR_WARNING = "warning"
	ERROR_DANGER  = "danger"
)

// [PageError]
type PageError struct {
	Type    string
	Name    string
	Message string
}

func (this *PageError) String() string { // {{{
	return fmt.Sprintf("%s: %s", this.Name, this.Message)
} // }}}
