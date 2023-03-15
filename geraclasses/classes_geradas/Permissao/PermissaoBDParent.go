package Permissao

import "fmt"
import "time"
import "strconv"
import "geraclasses/classes/conexao"

type PermissaoBDParent struct{
    Conexao *conexao.Conexao
}

func (p *PermissaoBDParent) GetConexao() *conexao.Conexao {
	return p.Conexao
}

func (p *PermissaoBDParent) Recupera(nIdTipoTransacao int64,nIdGrupoUsuario int64) (*PermissaoParent, error) {
	oConexao := p.GetConexao()
	sSql := fmt.Sprintf("select * from seg_permissao where id_tipo_transacao = %s and id_grupo_usuario = %s",strconv.FormatInt(nIdTipoTransacao,10)strconv.FormatInt(nIdGrupoUsuario,10))
	oConexao.Execute(sSql)
	oReg := oConexao.FetchObject()
	if oReg.Next() {
		var oPermissao PermissaoParent
		err := oReg.Scan(&oPermissao.IdTipoTransacao,&oPermissao.IdGrupoUsuario,&oPermissao.DtCadastro,&oPermissao.Publicado,&oPermissao.Ativo)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		return &oPermissao, nil
	}
	return nil, nil
}

func (p *PermissaoBDParent) Presente(nIdTipoTransacao int64,nIdGrupoUsuario int64) int64 {
	oConexao := p.GetConexao()
	sSql := fmt.Sprintf("select * from seg_permissao where id_tipo_transacao = %s and id_grupo_usuario = %s",strconv.FormatInt(nIdTipoTransacao,10)strconv.FormatInt(nIdGrupoUsuario,10))
	oConexao.Execute(sSql)
	nQtd, _ := oConexao.RecordCount()
	return nQtd
}