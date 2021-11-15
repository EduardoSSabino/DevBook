package router

// Gerar vai retornar um router com as rotas configuradas
import "github.com/gorilla/mux"

func Gerar() *mux.Router { // mux.Router é o tipo que a gente usa pra fazer a criação do router e passar as rotas pra ele depois
	return mux.NewRouter()
}
