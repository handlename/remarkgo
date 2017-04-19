package remark

import (
	"html/template"
	"net/http"
)

type Server struct {
	ListenAddr    string
	SrcPath       string
	CustomCSSPath string

	tmplIndex *template.Template
}

type ServerOption func(*Server) error

func ServerOptionSrcPath(path string) ServerOption {
	return func(s *Server) error {
		s.SrcPath = path
		return nil
	}
}

func ServerOptionCustomCSSPath(path string) ServerOption {
	return func(s *Server) error {
		s.CustomCSSPath = path
		return nil
	}
}

func NewServer(addr string, options ...ServerOption) (*Server, error) {
	s := Server{
		ListenAddr:    addr,
		SrcPath:       "index.md",
		CustomCSSPath: "",
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
	http.HandleFunc("/"+s.SrcPath, s.staticHandler)
	http.HandleFunc("/"+s.CustomCSSPath, s.staticHandler)

	return http.ListenAndServe(s.ListenAddr, nil)
}

func (s *Server) rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	s.tmplIndex.Execute(w, tmplParamsIndex{
		SrcPath:       s.SrcPath,
		CustomCSSPath: s.CustomCSSPath,
	})
}

func (s *Server) staticHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")
	http.ServeFile(w, r, "."+r.URL.Path)
}
