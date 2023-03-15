package TransacaoAcesso

import "fmt"
import "time"
import "strconv"
import "geraclasses/classes/conexao"

type TransacaoAcessoBDParent struct{
    Conexao *conexao.Conexao
}

func (t *TransacaoAcessoBDParent) GetConexao() *conexao.Conexao {
	return t.Conexao
}

func (t *TransacaoAcessoBDParent) Recupera(nId int64) (*TransacaoAcessoParent, error) {
	oConexao := t.GetConexao()
	sSql := fmt.Sprintf("select * from seg_transacao_acesso where id = %s",strconv.FormatInt(nId,10))
	oConexao.Execute(sSql)
	oReg := oConexao.FetchObject()
	if oReg.Next() {
		var oTransacaoAcesso TransacaoAcessoParent
		err := oReg.Scan(&oTransacaoAcesso.Id,&oTransacaoAcesso.IdTipoTransacao,&oTransacaoAcesso.IdUsuario,&oTransacaoAcesso.Objeto,&oTransacaoAcesso.Campo,&oTransacaoAcesso.ValorAntigo,&oTransacaoAcesso.ValorNovo,&oTransacaoAcesso.Ip,&oTransacaoAcesso.DtCadastro,&oTransacaoAcesso.Publicado,&oTransacaoAcesso.Ativo)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		return &oTransacaoAcesso, nil
	}
	return nil, nil
}

func (t *TransacaoAcessoBDParent) Presente(nId int64) int64 {
	oConexao := t.GetConexao()
	sSql := fmt.Sprintf("select * from seg_transacao_acesso where id = %s",strconv.FormatInt(nId,10))
	oConexao.Execute(sSql)
	nQtd, _ := oConexao.RecordCount()
	return nQtd
}