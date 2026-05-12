package slskd

import (
	"net"
	"net/url"
	"strings"
)

type Routes struct {
	Session   string
	Uploads   string
	Downloads string
}

const ( // Routes
	api       = "/api/v0"
	session   = api + "/session"
	transfers = api + "/transfers"
	uploads   = transfers + "/uploads"
	downloads = transfers + "/downloads"
)

func joinRoutes(base url.URL, path string) string {
	u := base
	u.Path = strings.TrimRight(u.Path, "/") + "/" + strings.TrimLeft(path, "/")
	return u.String()
}

func NewRoutes(host string, port string) Routes {
	authority := net.JoinHostPort(host, port) // e.g.: 192.168.1.2:8080

	u := url.URL{
		Scheme: "https",
		Host:   authority,
	}

	routes := Routes{
		Session:   joinRoutes(u, session),
		Uploads:   joinRoutes(u, uploads),
		Downloads: joinRoutes(u, downloads),
	}

	return routes
}
