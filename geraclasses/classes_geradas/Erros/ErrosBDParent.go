package Erros

import "fmt"
import "time"
import "strconv"
import "geraclasses/classes/conexao"

type ErrosBDParent struct{
    Conexao *conexao.Conexao
}

func (e *ErrosBDParent) GetConexao() *conexao.Conexao {
	return e.Conexao
}

func (e *ErrosBDParent) Recupera(nId int64) (*ErrosParent, error) {
	oConexao := e.GetConexao()
	sSql := fmt.Sprintf("select * from seg_erros_sql where id = %s",strconv.FormatInt(nId,10))
	oConexao.Execute(sSql)
	oReg := oConexao.FetchObject()
	if oReg.Next() {
		var oErros ErrosParent
		err := oReg.Scan(&oErros.Id,&oErros.Erro,&oErros.IdUsuario,&oErros.Ip,&oErros.DtCadastro,&oErros.Publicado,&oErros.Ativo)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		return &oErros, nil
	}
	return nil, nil
}

func (e *ErrosBDParent) Presente(nId int64) int64 {
	oConexao := e.GetConexao()
	sSql := fmt.Sprintf("select * from seg_erros_sql where id = %s",strconv.FormatInt(nId,10))
	oConexao.Execute(sSql)
	nQtd, _ := oConexao.RecordCount()
	return nQtd
}