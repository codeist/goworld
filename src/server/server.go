package server

type Server struct {
	Host string
	Port string
}

func NewServer(host, port) *Server {
	s := Server{Host: host, Port: port}
	return &s
}
