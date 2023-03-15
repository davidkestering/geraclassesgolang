package listar_tabelas

import (
	"fmt"
	"geraclasses/classes/conexao"
	construtor "geraclasses/classes/constructor"
	"geraclasses/classes/geradora_basica"
	"geraclasses/classes/geradora_bd"
	ct "geraclasses/constantes"
	"html/template"
	"net/http"
	"path/filepath"
)

func ListarTabelasHandler(w http.ResponseWriter, r *http.Request) {
	// Construa o caminho absoluto para o arquivo index.html
	pathIndex := filepath.Join("templates", "listar_tabelas", "lista_tabelas.html")

	tmpl, err := template.ParseFiles(pathIndex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Conexao := conexao.NewConexao(ct.BANCO)
	sNomeBanco := ""
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sNomeBanco = r.FormValue("fNome")
	}

	//fmt.Println(r.Method)

	//MAPEANDO POST
	//if r.Method == http.MethodPost {
	//	// Parse o formulário enviado
	//	err := r.ParseForm()
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//
	//	// Itere sobre todos os parâmetros
	//	for nome, valores := range r.Form {
	//		// Imprima o nome do parâmetro
	//		fmt.Fprintf(w, "%s: ", nome)
	//
	//		// Imprima todos os valores do parâmetro
	//		for _, valor := range valores {
	//			fmt.Fprintf(w, "%s\n", valor)
	//		}
	//	}
	//}

	//MAPEANDO GET
	//if r.Method == http.MethodGet {
	//	// Parse o formulário enviado
	//	err := r.ParseForm()
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//
	//	// Itere sobre todos os parâmetros
	//	for nome, valores := range r.Form {
	//		// Imprima o nome do parâmetro
	//		fmt.Fprintf(w, "%s: ", nome)
	//
	//		// Imprima todos os valores do parâmetro
	//		for _, valor := range valores {
	//			fmt.Fprintf(w, "%s\n", valor)
	//		}
	//	}
	//}

	//type Colunas struct {
	//	Coluna string
	//}

	var AsColunas []string
	Tabelas := Conexao.PegaTabelas()
	for _, linha := range Tabelas {
		//var novaColuna Colunas
		for _, coluna := range linha {
			//novaColuna.Coluna = append(novaColuna.Coluna, fmt.Sprintf("%v", coluna))
			AsColunas = append(AsColunas, fmt.Sprintf("%v", coluna))
		}
		//AsColunas = append(AsColunas, novaColuna)
	}

	data := struct {
		QtdTabelas int64
		NomeBanco  string
		Tabelas    []string
	}{
		QtdTabelas: Conexao.GetQtdTabelas(),
		NomeBanco:  sNomeBanco,
		Tabelas:    AsColunas,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Processa(w http.ResponseWriter, r *http.Request) {
	// Construa o caminho absoluto para o arquivo index.html
	pathIndex := filepath.Join("templates", "listar_tabelas", "lista_tabelas.html")

	tmpl, err := template.ParseFiles(pathIndex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Conexao := conexao.NewConexao(ct.BANCO)
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sNomeBanco := r.FormValue("fNomeBanco")
		vTabelas := r.PostForm["fTabelas[]"]

		//oConexao := conexao.NewConexao(ct.BANCO)
		for _, sTabela := range vTabelas {
			oConstrutor := construtor.NewConstrutor(sNomeBanco, sTabela)
			oGeradoraBasica := geradora_basica.NewGeradoraBasica(oConstrutor)
			oGeradoraBasica.Gera()
			oGeradoraBD := geradora_bd.NewGeradoraBD(oConstrutor)
			oGeradoraBD.Gera()
		}

		//fmt.Println(sNomeBanco)
		//fmt.Println(vTabelas)

		//// Itere sobre todos os parâmetros
		//for nome, valores := range r.Form {
		//	// Imprima o nome do parâmetro
		//	fmt.Fprintf(w, "%s: ", nome)
		//
		//	// Imprima todos os valores do parâmetro
		//	for _, valor := range valores {
		//		fmt.Fprintf(w, "%s\n", valor)
		//	}
		//}
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
