package page

type Page struct {
	Name        string
	Description string

	parent  *Page
	errors  *PageErrors
	notices *PageNotices
}

// [Parent]

func (this *Page) Parent() *Page { // {{{
	return this.Parent
} // }}}

func (this *Page) HasParent() bool { // {{{
	return this.Parent != nil
} // }}}

// [Errors]

func (this *Page) Errors() *PageErrors { // {{{
	if this.errors == nil {
		this.errors = NewPageErrors()
	}

	return this.errors
} // }}}

func (this *Page) HasErrors() bool { // {{{
	return this.Errors().Count() > 0
} // }}}

func (this *Page) addError(errorType string, name string, message string) { // {{{
	this.Errors().Append(&PageError{errorType, name, message})
} // }}}

func (this *Page) AddErrorWarning(name string, message string) { // {{{
	this.addError(ERROR_WARNING, name, message)
} // }}}

func (this *Page) AddErrorDanger(name string, message string) { // {{{
	this.addError(ERROR_DANGER, name, message)
} // }}}

// [Notices]

func (this *Page) Notices() *PageNotices { // {{{
	if this.notices == nil {
		this.notices = NewPageNotices()
	}

	return this.notices
} // }}}

func (this *Page) HasNotices() bool { // {{{
	return this.Notices().Count() > 0
} // }}}

func (this *Page) addNotice(noticeType string, name string, message string) { // {{{
	this.Notices().Append(&PageNotice{noticeType, name, message})
} // }}}

func (this *Page) AddNoticeSuccess(name string, message string) { // {{{
	this.addNotice(NOTICE_SUCCESS, name, message)
} // }}}

func (this *Page) AddNoticeInfo(name string, message string) { // {{{
	this.addNotice(NOTICE_INFO, name, message)
} // }}}
