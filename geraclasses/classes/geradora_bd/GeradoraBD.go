package geradora_bd

import (
	"fmt"
	constructor "geraclasses/classes/constructor"
	ct "geraclasses/constantes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type GeradoraBD struct {
	Construtor                      *constructor.Construtor
	NomeArquivo                     string
	NomeArquivoParent               string
	AtributosConstrutor             string
	ListaAtributos                  []Atributos
	ListaAtributosChave             []Atributos
	ListaCamposReg                  []Atributos
	ListaCamposChave                []Atributos
	ListaCamposNaoChave             []Atributos
	ListaValoresNaoChave            []Atributos
	TextoListaAtributos             string
	TextoListaAtributosChave        string
	TextoListaCamposReg             string
	TextoListaCamposChave           string
	TextoListaCamposNaoChave        string
	TextoListaValoresNaoChave       string
	ComparacaoChaveAtributoQuery    string
	ComparacaoChaveAtributoQueryEsp string
	ComparacaoChaveAtributoParam    string
	ComparacaoChaveAtributoParamEsp string
	ConteudoClasse                  string
	AtribuicaoNaoChave              string
	Tabela                          interface{}
	NomeProjeto                     string
}

type Atributos struct {
	Atributo string
	Tipo     string
}

func NewGeradoraBD(Construtor *constructor.Construtor) *GeradoraBD {
	return &GeradoraBD{
		Construtor:                      Construtor,
		NomeArquivo:                     Construtor.Classe + "BD.go",
		NomeArquivoParent:               Construtor.Classe + "BDParent.go",
		AtributosConstrutor:             "",
		ListaAtributos:                  make([]Atributos, 0),
		ListaAtributosChave:             make([]Atributos, 0),
		ListaCamposReg:                  make([]Atributos, 0),
		ListaCamposChave:                make([]Atributos, 0),
		ListaCamposNaoChave:             make([]Atributos, 0),
		ListaValoresNaoChave:            make([]Atributos, 0),
		TextoListaAtributos:             "",
		TextoListaAtributosChave:        "",
		TextoListaCamposReg:             "",
		TextoListaCamposChave:           "",
		TextoListaCamposNaoChave:        "",
		TextoListaValoresNaoChave:       "",
		ComparacaoChaveAtributoQuery:    "",
		ComparacaoChaveAtributoQueryEsp: "",
		ComparacaoChaveAtributoParam:    "",
		ComparacaoChaveAtributoParamEsp: "",
		ConteudoClasse:                  "",
		AtribuicaoNaoChave:              "",
		Tabela:                          nil,
		NomeProjeto:                     "",
	}
}

func (g *GeradoraBD) GeraAtributos() {
	vAtributos := g.Construtor.Atributos
	if len(vAtributos) > 0 {
		for _, sAtributo := range vAtributos {
			//sNomeAtributo := sAtributo.Atributo[1:]
			g.AtributosConstrutor += sAtributo.Atributo + ","
			atributo := Atributos{
				Atributo: sAtributo.Atributo,
				Tipo:     sAtributo.Tipo,
			}
			g.ListaAtributos = append(g.ListaAtributos, atributo)
			g.TextoListaAtributos += fmt.Sprintf("%s %s,", sAtributo.Atributo, sAtributo.Tipo)
		}
		g.TextoListaAtributos = g.TextoListaAtributos[:len(g.TextoListaAtributos)-1]
		g.AtributosConstrutor = g.AtributosConstrutor[:len(g.AtributosConstrutor)-1]
	}
}

func (g *GeradoraBD) GeraCampos() {
	vCampos := g.Construtor.Campos
	vAtributos := g.Construtor.Atributos
	if len(vCampos) > 0 {
		var vAtribuicaoNaoChave []string
		var vCamposReg []string
		var vValoresNaoChave []string
		var vComparacaoQuery []string
		var vComparacaoQueryEsp []string
		var vComparacaoParam []string
		var vComparacaoParamEsp []string
		var vAtributosChave []string
		var vCamposChave []string
		var vCamposNaoChave []string
		//chaveA := "{"
		//chaveF := "}"
		for nIndice, oCampo := range vCampos {
			if oCampo.Pri == "PK" {
				sTeste := fmt.Sprintf("%v", vAtributos[nIndice].Atributo)
				sTeste2 := vAtributos[nIndice].Atributo
				sTeste3 := fmt.Sprintf("%v", sTeste)
				//vAtributosChave = append(vAtributosChave, sTeste)
				sTextoAtributosChaves := fmt.Sprintf("%s %s,", vAtributos[nIndice].Atributo, vAtributos[nIndice].Tipo)
				vAtributosChave = append(vAtributosChave, sTextoAtributosChaves[:len(sTextoAtributosChaves)-1])
				vCamposChave = append(vCamposChave, oCampo.ColumnName)
				vComparacaoQuery = append(vComparacaoQuery, fmt.Sprintf("%v = %s", oCampo.ColumnName, "%s"))
				vComparacaoQueryEsp = append(vComparacaoQueryEsp, fmt.Sprintf("%v = '%v'", oCampo.ColumnName, "{o"+g.Construtor.Classe+".get"+sTeste2[1:]+"()}"))
				vComparacaoParam = append(vComparacaoParam, fmt.Sprintf("strconv.FormatInt(%v,10)", sTeste3))
				vComparacaoParamEsp = append(vComparacaoParamEsp, fmt.Sprintf("%v = '%v'", oCampo.ColumnName, "{o"+g.Construtor.Classe+".get"+sTeste2[1:]+"()}"))
			} else {
				vCamposNaoChave = append(vCamposNaoChave, oCampo.ColumnName)
				if string(vAtributos[nIndice].Atributo[0:1]) == "s" {
					vAtribuicaoNaoChave = append(vAtribuicaoNaoChave, fmt.Sprintf("%v = '%v'", oCampo.ColumnName, "{oConexao.escapeString(o"+g.Construtor.Classe+".get"+vAtributos[nIndice].Atributo[1:]+"())}"))
					vValoresNaoChave = append(vValoresNaoChave, fmt.Sprintf("'%v'", "{Conexao.escapeString(o"+g.Construtor.Classe+".get"+vAtributos[nIndice].Atributo[1:]+"())}"))
				} else {
					vAtribuicaoNaoChave = append(vAtribuicaoNaoChave, fmt.Sprintf("%v = '%v'", oCampo.ColumnName, "{o"+g.Construtor.Classe+".get"+vAtributos[nIndice].Atributo[1:]+"()}"))
					vValoresNaoChave = append(vValoresNaoChave, fmt.Sprintf("'%v'", "{o"+g.Construtor.Classe+".get"+vAtributos[nIndice].Atributo[1:]+"()}"))
				}
			}
			//if string(vAtributos[nIndice].Atributo[0:1]) == "s" {
			//	vCamposReg = append(vCamposReg, fmt.Sprintf("oConexao.unescapeString(oReg.%v)", oCampo.ColumnName))
			//} else {
			//	vCamposReg = append(vCamposReg, fmt.Sprintf("oReg.%v"), oCampo.ColumnName)
			//}
			vCamposReg = append(vCamposReg, fmt.Sprintf("&o%s.%s", g.Construtor.Classe, vAtributos[nIndice].Atributo[1:]))

			g.TextoListaAtributosChave = strings.Join(vAtributosChave, ",")
			//g.ListaCamposChave = strings.Join(vCamposChave, ",")
			//g.ListaCamposNaoChave = strings.Join(vCamposNaoChave, ",")
			//g.ListaValoresNaoChave = strings.Join(vValoresNaoChave, ",")
			g.ComparacaoChaveAtributoQuery = strings.Join(vComparacaoQuery, " and ")
			//g.ComparacaoChaveAtributoEsp = strings.Join(vComparacaoEsp, " and ")
			g.ComparacaoChaveAtributoParam = strings.Join(vComparacaoParam, "")
			//g.AtribuicaoNaoChave = strings.Join(vAtribuicaoNaoChave, ", ")
			g.TextoListaCamposReg = strings.Join(vCamposReg, ",")
		}
	}
}

func (g *GeradoraBD) Gera() {
	g.GeraAtributos()
	g.GeraCampos()
	g.NomeProjeto = ct.CAMINHO

	var sArquivo string
	//var sArquivoParent string
	var sPastaClasse string
	sPastaClasse = g.Construtor.Classe
	dir := filepath.Join("classes_geradas", sPastaClasse)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	var sCaminhoArquivo = filepath.Join("classes_geradas", sPastaClasse, g.Construtor.Classe+"BD.go")
	var sCaminhoArquivoParent = filepath.Join("classes_geradas", sPastaClasse, g.Construtor.Classe+"BDParent.go")

	f, _ := ioutil.ReadFile(filepath.Join("modelos", "modelo_classe_bd.txt"))
	vModelo := strings.Split(string(f), "\n")
	sArquivo = strings.Join(vModelo, "\n")

	sArquivo = strings.ReplaceAll(sArquivo, "#NOME_PROJETO#", g.NomeProjeto)
	sArquivo = strings.ReplaceAll(sArquivo, "#NOME_CLASSE#", g.Construtor.Classe)

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

	f, _ = ioutil.ReadFile(filepath.Join("modelos", "modelo_classe_bd_parent.txt"))
	vModeloParent := strings.Split(string(f), "\n")
	sArquivoParent := strings.Join(vModeloParent, "\n")

	sArquivoParent = strings.ReplaceAll(sArquivoParent, "#NOME_PROJETO#", g.NomeProjeto)
	sArquivoParent = strings.ReplaceAll(sArquivoParent, "#NOME_CLASSE#", g.Construtor.Classe)
	sArquivoParent = strings.ReplaceAll(sArquivoParent, "#NOME_CLASSE_LOWER#", strings.ToLower(g.Construtor.Classe))
	sArquivoParent = strings.ReplaceAll(sArquivoParent, "#PRIMEIRA_LETRA_NOME_CLASSE#", strings.ToLower(g.Construtor.Classe[0:1]))
	sArquivoParent = strings.ReplaceAll(sArquivoParent, "#LISTA_ATRIBUTOS_CHAVE#", g.TextoListaAtributosChave)
	sArquivoParent = strings.ReplaceAll(sArquivoParent, "#NOME_TABELA#", strings.ToLower(g.Construtor.Tabela))
	sArquivoParent = strings.ReplaceAll(sArquivoParent, "#COMPARACAO_CHAVE_ATRIBUTO_QUERY#", g.ComparacaoChaveAtributoQuery)
	sArquivoParent = strings.ReplaceAll(sArquivoParent, "#COMPARACAO_CHAVE_ATRIBUTO_PARAM#", g.ComparacaoChaveAtributoParam)
	sArquivoParent = strings.ReplaceAll(sArquivoParent, "#LISTA_CAMPOS_REG#", g.TextoListaCamposReg)
	//sArquivoParent = strings.ReplaceAll(sArquivoParent, "#LISTA_ATRIBUTOS_STRUCT#", g.ListaAtributos)
	//sArquivoParent = strings.ReplaceAll(sArquivoParent, "#LISTA_FUNC_GET_SET#", g.GetSet)

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

}
