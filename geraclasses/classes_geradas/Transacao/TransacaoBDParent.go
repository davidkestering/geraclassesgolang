package Transacao

import "fmt"
import "time"
import "strconv"
import "geraclasses/classes/conexao"

type TransacaoBDParent struct{
    Conexao *conexao.Conexao
}

func (t *TransacaoBDParent) GetConexao() *conexao.Conexao {
	return t.Conexao
}

func (t *TransacaoBDParent) Recupera(nId int64) (*TransacaoParent, error) {
	oConexao := t.GetConexao()
	sSql := fmt.Sprintf("select * from seg_transacao where id = %s",strconv.FormatInt(nId,10))
	oConexao.Execute(sSql)
	oReg := oConexao.FetchObject()
	if oReg.Next() {
		var oTransacao TransacaoParent
		err := oReg.Scan(&oTransacao.Id,&oTransacao.IdTipoTransacao,&oTransacao.IdUsuario,&oTransacao.Objeto,&oTransacao.Campo,&oTransacao.ValorAntigo,&oTransacao.ValorNovo,&oTransacao.Ip,&oTransacao.DtCadastro,&oTransacao.Publicado,&oTransacao.Ativo)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		return &oTransacao, nil
	}
	return nil, nil
}

func (t *TransacaoBDParent) Presente(nId int64) int64 {
	oConexao := t.GetConexao()
	sSql := fmt.Sprintf("select * from seg_transacao where id = %s",strconv.FormatInt(nId,10))
	oConexao.Execute(sSql)
	nQtd, _ := oConexao.RecordCount()
	return nQtd
}