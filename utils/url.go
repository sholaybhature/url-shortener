package utils

import (
	"net"
	"net/http"
	"net/url"
	"strings"
)

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// X-FORWARDED-FOR sends a chain of IP address, ip1, ip2, ip3...etc
// RemoteAddr will be the last ip address of a physical machine (inaccurate?)
func getIp(r *http.Request) string {
	ip := r.Header.Get("X-FORWARDED-FOR")
	if ip != "" {
		splitIps := strings.Split(ip, ",")
		for _, ip := range splitIps {
			err := net.ParseIP(ip)
			if err != nil {
				return ip
			}
		}

	}
	return r.RemoteAddr
}
