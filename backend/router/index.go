package router

import (
	"net/http"
)

type Router struct {
	*http.ServeMux
	prefix string
}

func New(prefix string) *Router {
	return &Router{
		ServeMux: http.NewServeMux(),
		prefix:   prefix,
	}
}

func (r *Router) Get(path string, handler http.HandlerFunc) {
	r.HandleFunc(r.prefix+path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, req)
	})
}

func (r *Router) Post(path string, handler http.HandlerFunc) {
	r.HandleFunc(r.prefix+path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, req)
	})
}

func (r *Router) Put(path string, handler http.HandlerFunc) {
	r.HandleFunc(r.prefix+path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPut {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, req)
	})
}

func (r *Router) Delete(path string, handler http.HandlerFunc) {
	r.HandleFunc(r.prefix+path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, req)
	})
}

func (r *Router) Patch(path string, handler http.HandlerFunc) {
	r.HandleFunc(r.prefix+path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPatch {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, req)
	})
}
