package webpage

// [PageErrors]
type PageErrors []*PageError

func (this *PageErrors) Count() int { // {{{
	return len(*this)
} // }}}

func (this *PageErrors) Append(pageError *PageError) { // {{{
	*this = append(*this, pageError)
} // }}}

func NewPageErrors() *PageErrors { // {{{
	pageErrors = make(PageErrors, 0)

	return &pageErrors
} // }}}
