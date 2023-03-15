package Usuario

import (
	"fmt"
	"time"
)
import "strconv"
import "geraclasses/classes/conexao"

type UsuarioBDParent struct {
	Conexao *conexao.Conexao
}

func (u *UsuarioBDParent) GetConexao() *conexao.Conexao {
	return u.Conexao
}

func (u *UsuarioBDParent) Recupera(nId int64) (*UsuarioParent, error) {
	oConexao := u.GetConexao()
	sSql := fmt.Sprintf("select * from seg_usuario where id = %s", strconv.FormatInt(nId, 10))
	oConexao.Execute(sSql)
	oReg := oConexao.FetchObject()
	if oReg.Next() {
		var oUsuario UsuarioParent
		err := oReg.Scan(&oUsuario.Id, &oUsuario.IdGrupoUsuario, &oUsuario.NmUsuario, &oUsuario.Login, &oUsuario.Senha, &oUsuario.Email, &oUsuario.Logado, &oUsuario.DtCadastro, &oUsuario.Publicado, &oUsuario.Ativo)
		vData := oUsuario.GetDtCadastro()
		data, _ := time.Parse("2006-01-02T15:04:05.999Z", vData)
		dDataCadastro := data.Format("2006-01-02 15:04:05")
		oUsuario.SetDtCadastro(dDataCadastro)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		return &oUsuario, nil
	}
	return nil, nil
}

func (u *UsuarioBDParent) Presente(nId int64) int64 {
	oConexao := u.GetConexao()
	sSql := fmt.Sprintf("select * from seg_usuario where id = %s", strconv.FormatInt(nId, 10))
	oConexao.Execute(sSql)
	nQtd, _ := oConexao.RecordCount()
	return nQtd
}

func (u *UsuarioBDParent) Insere(oUsuario *UsuarioParent) (int64, error) {
	oConexao := u.GetConexao()
	//vData := oUsuario.GetDtCadastro().Format("2006-01-02 15:04:05.000000")
	//data, err := time.Parse("2006-01-02 15:04:05.000000", vData)
	//if err != nil {
	//	return 0, err
	//}
	//dDataCadastro := data.Format("2006-01-02 15:04:05")
	//oUsuario.SetDtCadastro(dDataCadastro)
	sSql := fmt.Sprintf(`insert into seg_usuario (id_grupo_usuario, nm_usuario, login, email, logado, dt_cadastro, publicado, ativo) values ('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v') SELECT CONVERT(bigint,SCOPE_IDENTITY());`,
		oUsuario.GetIdGrupoUsuario(),
		oConexao.EscapeString(oUsuario.GetNmUsuario()),
		oConexao.EscapeString(oUsuario.GetLogin()),
		oConexao.EscapeString(oUsuario.GetEmail()),
		oUsuario.GetLogado(),
		oUsuario.GetDtCadastro(),
		oUsuario.GetPublicado(),
		oUsuario.GetAtivo(),
	)
	//fmt.Println(sSql)
	err := oConexao.ExecuteInsert(sSql)
	if err != nil {
		fmt.Println("ERRO: ", err)
		return 0, err
	}
	nId := oConexao.GetLastId()
	if nId == 0 {
		fmt.Println("SEM ID")
		return 0, nil
	}
	return nId, nil
}
