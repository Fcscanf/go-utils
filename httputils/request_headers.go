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

func (header RequestHeader) UserAgentWindows() RequestHeader {
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36"
	return header
}

func (header RequestHeader) UserAgentAndroid() RequestHeader {
	header["User-Agent"] = "Mozilla/5.0 (Linux; Android 12; ABR-AL80 Build/HUAWEIABR-AL80; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 "
	return header
}

func (header RequestHeader) ContentType(contentType string) RequestHeader {
	header["Content-Type"] = contentType
	return header
}

func (header RequestHeader) ContentTypeApplicationJSON() RequestHeader {
	header["Content-Type"] = "application/json;charset=UTF-8"
	return header
}

func (header RequestHeader) ContentTypeApplicationXML() RequestHeader {
	header["Content-Type"] = "application/xml;charset=UTF-8"
	return header
}

func (header RequestHeader) ContentTypeTextXML() RequestHeader {
	header["Content-Type"] = "text/xml;charset=UTF-8"
	return header
}
