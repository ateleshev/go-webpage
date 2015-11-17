package webpage

import (
	"fmt"
)

const (
	NOTICE_SUCCESS = "success"
	NOTICE_INFO    = "info"
)

// [PageNotice]
type PageNotice struct {
	Type    string
	Name    string
	Message string
}

func (this *PageNotice) String() string { // {{{
	return fmt.Sprintf("%s: %s", this.Name, this.Message)
} // }}}
