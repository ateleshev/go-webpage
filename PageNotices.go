package webpage

// [PageNotices]
type PageNotices []*PageNotice

func (this *PageNotices) Count() int { // {{{
	return len(*this)
} // }}}

func (this *PageNotices) Append(pageNotice *PageNotice) { // {{{
	*this = append(*this, pageNotice)
} // }}}

func NewPageNotices() *PageNotices { // {{{
	pageNotices := make(PageNotices, 0)

	return &pageNotices
} // }}}
