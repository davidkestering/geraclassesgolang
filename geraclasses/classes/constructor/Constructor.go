package construtor

import (
	conexao "geraclasses/classes/conexao"
	ct "geraclasses/constantes"
	"regexp"
	"strings"
)

type Construtor struct {
	NomeBanco       string
	Tabela          string
	Campos          []conexao.Campos
	Atributos       []Atributos
	QtdCamposTabela int64
	Classe          string
	ClasseSimples   string
	Arquivo         string
}

type Atributos struct {
	Atributo string
	Tipo     string
}

func NewConstrutor(NomeBanco, NomeTabela string) *Construtor {
	Conexao := conexao.NewConexao(ct.BANCO)
	c := &Construtor{
		NomeBanco:       NomeBanco,
		Tabela:          NomeTabela,
		Campos:          make([]conexao.Campos, 0),
		Atributos:       make([]Atributos, 0),
		QtdCamposTabela: Conexao.CarregaQtdCampos(NomeBanco, NomeTabela),
	}
	c.GeraNomeClasse()
	if c.QtdCamposTabela > 0 {
		c.Campos = Conexao.PegaCampos(NomeBanco, NomeTabela)
		c.GeraAtributos()
	}
	return c
}

func (c *Construtor) GeraNomeClasse() {
	var Nome []string
	var Separador string
	match1, _ := regexp.MatchString("-", c.Tabela)
	match2, _ := regexp.MatchString("_", c.Tabela)
	if match1 {
		Separador = "-"
	} else if match2 {
		Separador = "_"
	}
	if Separador != "" {
		Nome = strings.Split(c.Tabela, Separador)
		var Auxiliar []string
		for _, nome := range Nome {
			if len(nome) > 3 {
				Auxiliar = append(Auxiliar, strings.Title(strings.ToLower(nome)))
			}
		}
		c.Classe = strings.Join(Auxiliar, "")
		c.Arquivo = strings.ToLower(strings.Join(Auxiliar, "_"))
	} else {
		c.Classe = strings.Title(strings.ToLower(c.Tabela))
		c.ClasseSimples = c.Classe
		c.Arquivo = strings.ToLower(c.Tabela)
	}
}

func (c *Construtor) GeraAtributos() {
	PadraoInteiro := "(int|integer|bigint|smallint|mediumint)"
	PadraoDecimal := "(real|double|float|numeric|decimal)"
	PadraoData := "(time|timestamp|date|datetime)"
	Campos := c.Campos
	//fmt.Println(c.Campos)
	for _, Campo := range Campos {
		NomeAtributo := ""
		TipoAtributo := ""
		PassouEmPadrao := false
		if matched, _ := regexp.MatchString(PadraoInteiro, Campo.DataType); matched {
			NomeAtributo = "n" + strings.Title(strings.ToLower(Campo.ColumnName))
			TipoAtributo = "int64"
			PassouEmPadrao = true
		}
		if matched, _ := regexp.MatchString(PadraoDecimal, Campo.DataType); matched {
			NomeAtributo = "n" + strings.Title(strings.ToLower(Campo.ColumnName))
			TipoAtributo = "float64"
			PassouEmPadrao = true
		}
		if matched, _ := regexp.MatchString(PadraoData, Campo.DataType); matched {
			NomeAtributo = "d" + strings.Title(strings.ToLower(Campo.ColumnName))
			TipoAtributo = "string"
			PassouEmPadrao = true
		}
		if Campo.DataType == "tinyint" {
			NomeAtributo = "b" + strings.Title(strings.ToLower(Campo.ColumnName))
			TipoAtributo = "int"
			PassouEmPadrao = true
		}
		if !PassouEmPadrao {
			NomeAtributo = "s" + strings.Title(strings.ToLower(Campo.ColumnName))
			TipoAtributo = "string"
		}
		VNomeAtributo := strings.Split(strings.ReplaceAll(strings.ReplaceAll(NomeAtributo, "-", "_"), " ", ""), "_")
		if len(VNomeAtributo) > 0 {
			for i := 1; i < len(VNomeAtributo); i++ {
				VNomeAtributo[i] = strings.Title(VNomeAtributo[i])
			}
			NomeAtributo = strings.Join(VNomeAtributo, "")
		}
		atributo := Atributos{
			Atributo: NomeAtributo,
			Tipo:     TipoAtributo,
		}
		c.Atributos = append(c.Atributos, atributo)
	}

	//fmt.Println(c.Atributos)
}
