package modelos

// id e token do usuario logado
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
