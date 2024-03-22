package server

func httpLink(addr string, secure bool) string {
	if addr[0] == ':' {
		addr = "localhost" + addr
	}
	if secure {
		return "https://" + addr
	}
	return "http://" + addr
}
