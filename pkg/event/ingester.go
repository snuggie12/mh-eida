package event

import(
	context "context"
)

type Server struct {
	UnimplementedEmitterServer
}

func (s *Server) Emit(ctx context.Context, e *Event) (*EmitResponse, error) {
	return &EmitResponse{TestReply: "yo dawg"}, nil
}
