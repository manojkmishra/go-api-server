package handlers
import "github.com/go-chi/chi"
type Server struct{
}
func NewServer() *Server{
	return &Server{}
}
func SetupRouter() *chi.Mux{
   server := NewServer()
   r :=chi.NewRouter()
   server.setupEndpoints(r)
   return r
}