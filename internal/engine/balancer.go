package engine

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

// Represents a single instance of backend
type Backend struct {
	URL          *url.URL
	ReverseProxy *httputil.ReverseProxy
}

// Pool of backends to load balance
type ServerPool struct {
	backends []*Backend
	current  uint64
	mu       sync.Mutex
}

// Registers a new backend to the Server Pool
func (s *ServerPool) AddBackend(backendURL string) error {
	parsedURL, err := url.Parse(backendURL)
	if err != nil {
		return err
	}

	proxy := httputil.NewSingleHostReverseProxy(parsedURL)

	s.backends = append(s.backends, &Backend{
		URL:          parsedURL,
		ReverseProxy: proxy,
	})
	return nil
}

// Using mutex for now, will switch to atomic operations after learning more about them
func (s *ServerPool) NextPeer() *Backend {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.current++

	idx := s.current % uint64(len(s.backends))
	return s.backends[idx]
}

// Implements the http.Handler interface to serve incoming requests
// Without health checks and retrying to other possible peers for now, will add those features later
func (s *ServerPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	peer := s.NextPeer()
	if peer != nil {
		peer.ReverseProxy.ServeHTTP(w, r)
		return
	}
	http.Error(w, "Service not available", http.StatusServiceUnavailable)
}
