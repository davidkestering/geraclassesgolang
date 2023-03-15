package GrupoUsuario

import "fmt"
import "time"
import "strconv"
import "geraclasses/classes/conexao"

type GrupoUsuarioBDParent struct{
    Conexao *conexao.Conexao
}

func (g *GrupoUsuarioBDParent) GetConexao() *conexao.Conexao {
	return g.Conexao
}

func (g *GrupoUsuarioBDParent) Recupera(nId int64) (*GrupoUsuarioParent, error) {
	oConexao := g.GetConexao()
	sSql := fmt.Sprintf("select * from seg_grupo_usuario where id = %s",strconv.FormatInt(nId,10))
	oConexao.Execute(sSql)
	oReg := oConexao.FetchObject()
	if oReg.Next() {
		var oGrupoUsuario GrupoUsuarioParent
		err := oReg.Scan(&oGrupoUsuario.Id,&oGrupoUsuario.NmGrupoUsuario,&oGrupoUsuario.DtCadastro,&oGrupoUsuario.Publicado,&oGrupoUsuario.Ativo)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		return &oGrupoUsuario, nil
	}
	return nil, nil
}

func (g *GrupoUsuarioBDParent) Presente(nId int64) int64 {
	oConexao := g.GetConexao()
	sSql := fmt.Sprintf("select * from seg_grupo_usuario where id = %s",strconv.FormatInt(nId,10))
	oConexao.Execute(sSql)
	nQtd, _ := oConexao.RecordCount()
	return nQtd
}