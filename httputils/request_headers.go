package httputils

type RequestHeader map[string]string

func (header RequestHeader) Add(k, v string) RequestHeader {
	header[k] = v
	return header
}

func (header RequestHeader) Cookie(cookie string) RequestHeader {
	header["Cookie"] = cookie
	return header
}

func (header RequestHeader) Referer(referer string) RequestHeader {
	header["Referer"] = referer
	return header
}

func (header RequestHeader) UserAgent(userAgent string) RequestHeader {
	header["User-Agent"] = userAgent
	return header
}

func (header RequestHeader) ContentType(contentType string) RequestHeader {
	header["Content-Type"] = contentType
	return header
}
