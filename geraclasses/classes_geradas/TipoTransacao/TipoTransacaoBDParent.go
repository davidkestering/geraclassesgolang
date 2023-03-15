package TipoTransacao

import "fmt"
import "time"
import "strconv"
import "geraclasses/classes/conexao"

type TipoTransacaoBDParent struct{
    Conexao *conexao.Conexao
}

func (t *TipoTransacaoBDParent) GetConexao() *conexao.Conexao {
	return t.Conexao
}

func (t *TipoTransacaoBDParent) Recupera(nId int64) (*TipoTransacaoParent, error) {
	oConexao := t.GetConexao()
	sSql := fmt.Sprintf("select * from seg_tipo_transacao where id = %s",strconv.FormatInt(nId,10))
	oConexao.Execute(sSql)
	oReg := oConexao.FetchObject()
	if oReg.Next() {
		var oTipoTransacao TipoTransacaoParent
		err := oReg.Scan(&oTipoTransacao.Id,&oTipoTransacao.IdCategoriaTipoTransacao,&oTipoTransacao.Transacao,&oTipoTransacao.DtCadastro,&oTipoTransacao.Publicado,&oTipoTransacao.Ativo)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		return &oTipoTransacao, nil
	}
	return nil, nil
}

func (t *TipoTransacaoBDParent) Presente(nId int64) int64 {
	oConexao := t.GetConexao()
	sSql := fmt.Sprintf("select * from seg_tipo_transacao where id = %s",strconv.FormatInt(nId,10))
	oConexao.Execute(sSql)
	nQtd, _ := oConexao.RecordCount()
	return nQtd
}