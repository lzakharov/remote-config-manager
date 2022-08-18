package transport

import (
	"io"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/maxence-charriere/go-app/v9/pkg/errors"

	"github.com/lzakharov/remote-config-manager/pkg/provider"
)

// Server is a remote config server.
type Server struct {
	provider provider.Provider
}

// NewServer creates a new remote config server.
func NewServer(provider provider.Provider) *Server {
	return &Server{provider: provider}
}

// ListKeys handles keys listing request.
func (s *Server) ListKeys(w http.ResponseWriter, r *http.Request) {
	keys, err := s.provider.ListKeys(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Respond(w, r, &ListKeysResp{Keys: keys})
}

// Get handles get request.
func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	value, err := s.provider.Get(r.Context(), key)
	if err != nil {
		if errors.Is(err, provider.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	_, err = w.Write([]byte(value))
	if err != nil {
		log.Printf("write get response: %v", err)
	}
}

// Put handles put request.
func (s *Server) Put(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("read put request body: %v", err)
		return
	}
	defer func() {
		err = r.Body.Close()
		if err != nil {
			log.Printf("close body on put request: %v", err)
		}
	}()

	err = s.provider.Put(r.Context(), key, string(data))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
