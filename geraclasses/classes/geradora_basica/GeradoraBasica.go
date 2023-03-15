package geradora_basica

import (
	"fmt"
	constructor "geraclasses/classes/constructor"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type GeradoraBasica struct {
	Construtor              *constructor.Construtor
	NomeArquivo             string
	NomeArquivoParent       string
	GetSet                  string
	AtributosConstrutor     string
	ListaAtributos          string
	InicializacaoConstrutor string
	SetsConstrutor          string
	FinalizacaoConstrutor   string
}

func NewGeradoraBasica(Construtor *constructor.Construtor) *GeradoraBasica {
	return &GeradoraBasica{
		Construtor:        Construtor,
		NomeArquivo:       Construtor.Classe + ".go",
		NomeArquivoParent: Construtor.Classe + "Parent.go",
	}
}

func (gerador *GeradoraBasica) GeraAtributos() {
	vAtributos := gerador.Construtor.Atributos
	if len(vAtributos) > 0 {
		for _, sAtributo := range vAtributos {
			sNomeAtributo := strings.Title(sAtributo.Atributo[1:])

			gerador.AtributosConstrutor += fmt.Sprintf("%s %s,", sAtributo.Atributo, sAtributo.Tipo)
			gerador.ListaAtributos += fmt.Sprintf("%s %s\n\t", sNomeAtributo, sAtributo.Tipo)
			gerador.GetSet += fmt.Sprintf("func (%s *%sParent) Get%s() %s {\n\treturn %s.%s\n}\n\n", strings.ToLower(gerador.Construtor.Classe[0:1]), gerador.Construtor.Classe, sNomeAtributo, sAtributo.Tipo, strings.ToLower(gerador.Construtor.Classe[0:1]), sNomeAtributo)
			gerador.GetSet += fmt.Sprintf("func (%s *%sParent) Set%s(%s %s) {\n\t%s.%s = %s\n}\n\n", strings.ToLower(gerador.Construtor.Classe[0:1]), gerador.Construtor.Classe, sNomeAtributo, sAtributo.Atributo, sAtributo.Tipo, strings.ToLower(gerador.Construtor.Classe[0:1]), sNomeAtributo, sAtributo.Atributo)
			gerador.SetsConstrutor += fmt.Sprintf("o%s.Set%s(%s)\n\t", gerador.Construtor.Classe, sNomeAtributo, sAtributo.Atributo)
		}
		gerador.InicializacaoConstrutor = fmt.Sprintf("o%s := &%s{}\n\t", gerador.Construtor.Classe, gerador.Construtor.Classe)
		gerador.FinalizacaoConstrutor = fmt.Sprintf("return o%s\n", gerador.Construtor.Classe)
		gerador.AtributosConstrutor = gerador.AtributosConstrutor[:len(gerador.AtributosConstrutor)-1]
	}
}

func (gerador *GeradoraBasica) Gera() {
	gerador.GeraAtributos()
	var sArquivo string
	var sArquivoParent string
	var sPastaClasse string
	sPastaClasse = gerador.Construtor.Classe
	dir := filepath.Join("classes_geradas", sPastaClasse)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	var sCaminhoArquivo = filepath.Join("classes_geradas", sPastaClasse, gerador.Construtor.Classe+".go")
	var sCaminhoArquivoParent = filepath.Join("classes_geradas", sPastaClasse, gerador.Construtor.Classe+"Parent.go")

	f, _ := ioutil.ReadFile(filepath.Join("modelos", "modelo_classe_basica.txt"))
	vModelo := strings.Split(string(f), "\n")
	sArquivo = strings.Join(vModelo, "\n")

	sArquivo = strings.ReplaceAll(sArquivo, "#NOME_CLASSE#", gerador.Construtor.Classe)
	sArquivo = strings.ReplaceAll(sArquivo, "#LISTA_ATRIBUTOS_STRUCT#", gerador.ListaAtributos)
	sArquivo = strings.ReplaceAll(sArquivo, "#LISTA_ATRIBUTOS_CONSTRUTOR#", gerador.AtributosConstrutor)
	sArquivo = strings.ReplaceAll(sArquivo, "#INICIALIZACAO_CONSTRUTOR#", gerador.InicializacaoConstrutor)
	sArquivo = strings.ReplaceAll(sArquivo, "#SETS_CONSTRUTOR#", gerador.SetsConstrutor)
	sArquivo = strings.ReplaceAll(sArquivo, "#FINALIZACAO_CONSTRUTOR#", gerador.FinalizacaoConstrutor)

	//fmt.Println(sArquivo)

	_, err := os.Stat(sCaminhoArquivo)
	if err != nil {
		os.Remove(sCaminhoArquivo)
	}

	err = ioutil.WriteFile(sCaminhoArquivo, []byte(sArquivo), 0644)
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}

	//f, err = os.Create(sCaminhoArquivo)
	//if err != nil {
	//	// handle error
	//}
	//defer f.Close()
	//f.WriteString(sArquivo)

	f, _ = ioutil.ReadFile(filepath.Join("modelos", "modelo_classe_basica_parent.txt"))
	vModeloParent := strings.Split(string(f), "\n")
	sArquivoParent = strings.Join(vModeloParent, "\n")

	sArquivoParent = strings.ReplaceAll(sArquivoParent, "#NOME_CLASSE#", gerador.Construtor.Classe)
	sArquivoParent = strings.ReplaceAll(sArquivoParent, "#LISTA_ATRIBUTOS_STRUCT#", gerador.ListaAtributos)
	sArquivoParent = strings.ReplaceAll(sArquivoParent, "#LISTA_FUNC_GET_SET#", gerador.GetSet)

	_, err = os.Stat(sCaminhoArquivoParent)
	if err != nil {
		os.Remove(sCaminhoArquivoParent)
	}

	err = ioutil.WriteFile(sCaminhoArquivoParent, []byte(sArquivoParent), 0644)
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}

	//f, _ = os.Create(sCaminhoArquivoParent)
	//if err != nil {
	//	// handle error
	//}
	//defer f.Close()
	//f.WriteString(sArquivoParent)
}
