package httputils

type requestHeaders struct {
	Header map[string]string
}

func NewRequestHeaders() requestHeaders {
	return requestHeaders{Header: make(map[string]string)}
}

func (header requestHeaders) Get() map[string]string {
	return header.Header
}

func (header requestHeaders) Add(k, v string) requestHeaders {
	header.Header[k] = v
	return header
}

func (header requestHeaders) Cookie(cookie string) requestHeaders {
	header.Header["Cookie"] = cookie
	return header
}

func (header requestHeaders) Referer(referer string) requestHeaders {
	header.Header["Referer"] = referer
	return header
}

func (header requestHeaders) UserAgent(userAgent string) requestHeaders {
	header.Header["User-Agent"] = userAgent
	return header
}

func (header requestHeaders) ContentType(contentType string) requestHeaders {
	header.Header["Content-Type"] = contentType
	return header
}
