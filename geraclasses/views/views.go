package views

import (
	"geraclasses/templates/index"
	"geraclasses/templates/listar_tabelas"
	"net/http"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	index.IndexHandler(w, r)
}

func ListarTabelasPage(w http.ResponseWriter, r *http.Request) {
	listar_tabelas.ListarTabelasHandler(w, r)
}

func ProcessaPage(w http.ResponseWriter, r *http.Request) {
	listar_tabelas.Processa(w, r)
}
