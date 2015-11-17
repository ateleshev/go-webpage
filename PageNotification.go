package webpage

import (
	"fmt"
)

// [PageNotification]
type PageNotification struct {
	Type    string
	Name    string
	Message string
}

func (this *PageNotification) String() string { // {{{
	return fmt.Sprintf("%s: %s", this.Name, this.Message)
} // }}}
