package remark

import (
	"html/template"
	"net/http"
)

const (
	DefaultListenAddr = "localhost:8080"
	DefaultSrcPath    = "index.md"
	DefaultCssPath    = ""
)

type Server struct {
	ListenAddr    string
	SrcPath       string
	CustomCSSPath string

	tmplIndexHtml *template.Template
	tmplIndexJs   *template.Template
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
		ListenAddr:    DefaultListenAddr,
		SrcPath:       DefaultSrcPath,
		CustomCSSPath: DefaultCssPath,
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
	{
		data, err := Asset("template/index.html")
		if err != nil {
			panic(err)
		}

		s.tmplIndexHtml = template.Must(template.New("index.html").Parse(string(data)))
	}

	{
		data, err := Asset("template/index.js")
		if err != nil {
			panic(err)
		}

		s.tmplIndexJs = template.Must(template.New("index.js").Parse(string(data)))
	}

	return
}

func (s *Server) Serve() error {
	http.HandleFunc("/", s.rootHandler)
	http.HandleFunc("/index.js", s.jsHandler)
	http.HandleFunc("/"+s.SrcPath, s.staticHandler)

	if s.CustomCSSPath != "" {
		http.HandleFunc("/"+s.CustomCSSPath, s.staticHandler)
	}

	return http.ListenAndServe(s.ListenAddr, nil)
}

func (s *Server) rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	w.Header().Set("Cache-Control", "no-cache")
	w.WriteHeader(http.StatusOK)

	err := s.tmplIndexHtml.Execute(w, tmplParamsIndex{
		SrcPath:       s.SrcPath,
		CustomCSSPath: s.CustomCSSPath,
	})
	if err != nil {
		panic(err)
	}
}

func (s *Server) jsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-Type", "text/javascript")
	w.WriteHeader(http.StatusOK)

	err := s.tmplIndexJs.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func (s *Server) staticHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")

	http.ServeFile(w, r, "."+r.URL.Path)
}
