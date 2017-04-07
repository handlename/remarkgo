package remark

import (
	"html/template"
	"net/http"
)

type Server struct {
	ListenAddr string
	SrcPath    string

	tmplIndex *template.Template
}

type ServerOption func(*Server) error

func ServerOptionSrcPath(path string) ServerOption {
	return func(s *Server) error {
		s.SrcPath = path
		return nil
	}
}

func NewServer(addr string, options ...ServerOption) (*Server, error) {
	s := Server{
		ListenAddr: addr,
		SrcPath:    "index.md",
	}

	for _, o := range options {
		if err := o(&s); err != nil {
			return nil, err
		}
	}

	s.initTemplates()

	return &s, nil
}

func (s *Server) initTemplates() {
	s.tmplIndex = template.Must(template.New("index").Parse(tmplIndex))

	return
}

func (s *Server) Serve() error {
	http.HandleFunc("/", s.rootHandler)
	http.Handle("/"+s.SrcPath, http.FileServer(http.Dir(".")))

	return http.ListenAndServe(s.ListenAddr, nil)
}

func (s *Server) rootHandler(w http.ResponseWriter, r *http.Request) {
	s.tmplIndex.Execute(w, map[string]string{
		"SrcPath": s.SrcPath,
	})
}
