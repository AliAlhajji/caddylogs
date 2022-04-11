package models

import (
	"net/url"
)

type Request struct {
	RemoteAddress string              `json:"remote_addr"`
	Method        string              `json:"method"`
	Host          string              `json:"host"`
	URI           string              `json:"uri"`
	Headers       map[string][]string `json:"headers"`
}

func (r *Request) URL() string {
	url := url.URL{
		Scheme: "https",
		Host:   r.Host,
		Path:   r.URI,
	}
	return url.String()
}

func (r *Request) Referer() string {
	if referers, ok := r.Headers["Referer"]; ok {
		return referers[0]
	}

	return ""
}

func (r *Request) UserAgent() string {
	if agents, ok := r.Headers["User-Agent"]; ok {
		return agents[0]
	}

	return ""
}
