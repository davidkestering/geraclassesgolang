package Usuario

import "time"

type Usuario struct {
    UsuarioParent
    Id int64
	IdGrupoUsuario int64
	NmUsuario string
	Login string
	Senha string
	Email string
	Logado int
	DtCadastro time.Time
	Publicado int
	Ativo int
	
}

func NewUsuario(nId int64,nIdGrupoUsuario int64,sNmUsuario string,sLogin string,sSenha string,sEmail string,bLogado int,dDtCadastro time.Time,bPublicado int,bAtivo int) *Usuario {
    oUsuario := &Usuario{}
	
    oUsuario.SetId(nId)
	oUsuario.SetIdGrupoUsuario(nIdGrupoUsuario)
	oUsuario.SetNmUsuario(sNmUsuario)
	oUsuario.SetLogin(sLogin)
	oUsuario.SetSenha(sSenha)
	oUsuario.SetEmail(sEmail)
	oUsuario.SetLogado(bLogado)
	oUsuario.SetDtCadastro(dDtCadastro)
	oUsuario.SetPublicado(bPublicado)
	oUsuario.SetAtivo(bAtivo)
	
    return oUsuario

}