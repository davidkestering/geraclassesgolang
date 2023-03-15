package CategoriaTipoTransacao

import "fmt"
import "time"
import "strconv"
import "geraclasses/classes/conexao"

type CategoriaTipoTransacaoBDParent struct{
    Conexao *conexao.Conexao
}

func (c *CategoriaTipoTransacaoBDParent) GetConexao() *conexao.Conexao {
	return c.Conexao
}

func (c *CategoriaTipoTransacaoBDParent) Recupera(nId int64) (*CategoriaTipoTransacaoParent, error) {
	oConexao := c.GetConexao()
	sSql := fmt.Sprintf("select * from seg_categoria_tipo_transacao where id = %s",strconv.FormatInt(nId,10))
	oConexao.Execute(sSql)
	oReg := oConexao.FetchObject()
	if oReg.Next() {
		var oCategoriaTipoTransacao CategoriaTipoTransacaoParent
		err := oReg.Scan(&oCategoriaTipoTransacao.Id,&oCategoriaTipoTransacao.Descricao,&oCategoriaTipoTransacao.DtCadastro,&oCategoriaTipoTransacao.Publicado,&oCategoriaTipoTransacao.Ativo)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		return &oCategoriaTipoTransacao, nil
	}
	return nil, nil
}

func (c *CategoriaTipoTransacaoBDParent) Presente(nId int64) int64 {
	oConexao := c.GetConexao()
	sSql := fmt.Sprintf("select * from seg_categoria_tipo_transacao where id = %s",strconv.FormatInt(nId,10))
	oConexao.Execute(sSql)
	nQtd, _ := oConexao.RecordCount()
	return nQtd
}