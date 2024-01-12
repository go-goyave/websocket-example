package static

import (
	"github.com/go-goyave/websocket-example/service"
	"goyave.dev/goyave/v5/util/fsutil"
)

// Service for the static resources.
type Service struct {
	fs fsutil.Embed
}

// NewService create a new user Service.
func NewService(fs fsutil.Embed) *Service {
	return &Service{
		fs: fs,
	}
}

// FS returns the static resources filesystem
func (s Service) FS() fsutil.Embed {
	return s.fs
}

// Name returns the service name.
func (s *Service) Name() string {
	return service.Static
}
