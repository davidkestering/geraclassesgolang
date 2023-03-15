package Usuario

type Usuario struct {
	UsuarioParent
	Id             int64
	IdGrupoUsuario int64
	NmUsuario      string
	Login          string
	Senha          string
	Email          string
	Logado         int
	DtCadastro     string
	Publicado      int
	Ativo          int
}

func _(nId int64, nIdGrupoUsuario int64, sNmUsuario string, sLogin string, sSenha string, sEmail string, bLogado int, dDtCadastro string, bPublicado int, bAtivo int) *Usuario {
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
