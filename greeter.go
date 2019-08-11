package greeter

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/twitchtv/twirp"

	pb "github.com/firstthumb/twirp-gcloud-greeter/proto"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	if req.Name == "" {
		return nil, twirp.RequiredArgumentError("name")
	}

	return &pb.GreetResponse{Greeting: "Hello, " + req.Name + "!"}, nil
}

var mux = NewMux()

func Greeter(w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w, r)
}

func NewMux() *http.ServeMux {
	s := &server{}
	twirpHandler := pb.NewGreeterServer(s, nil)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", HealthCheck)
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)

	return mux
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"message": "OK",
	})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	json.HTMLEscape(&buf, body)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf.Bytes())
}
