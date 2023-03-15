package Usuario

type UsuarioParent struct {
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

func (u *UsuarioParent) GetId() int64 {
	return u.Id
}

func (u *UsuarioParent) SetId(nId int64) {
	u.Id = nId
}

func (u *UsuarioParent) GetIdGrupoUsuario() int64 {
	return u.IdGrupoUsuario
}

func (u *UsuarioParent) SetIdGrupoUsuario(nIdGrupoUsuario int64) {
	u.IdGrupoUsuario = nIdGrupoUsuario
}

func (u *UsuarioParent) GetNmUsuario() string {
	return u.NmUsuario
}

func (u *UsuarioParent) SetNmUsuario(sNmUsuario string) {
	u.NmUsuario = sNmUsuario
}

func (u *UsuarioParent) GetLogin() string {
	return u.Login
}

func (u *UsuarioParent) SetLogin(sLogin string) {
	u.Login = sLogin
}

func (u *UsuarioParent) GetSenha() string {
	return u.Senha
}

func (u *UsuarioParent) SetSenha(sSenha string) {
	u.Senha = sSenha
}

func (u *UsuarioParent) GetEmail() string {
	return u.Email
}

func (u *UsuarioParent) SetEmail(sEmail string) {
	u.Email = sEmail
}

func (u *UsuarioParent) GetLogado() int {
	return u.Logado
}

func (u *UsuarioParent) SetLogado(bLogado int) {
	u.Logado = bLogado
}

func (u *UsuarioParent) GetDtCadastro() string {
	return u.DtCadastro
}

func (u *UsuarioParent) SetDtCadastro(dDtCadastro string) {
	u.DtCadastro = dDtCadastro
}

func (u *UsuarioParent) GetPublicado() int {
	return u.Publicado
}

func (u *UsuarioParent) SetPublicado(bPublicado int) {
	u.Publicado = bPublicado
}

func (u *UsuarioParent) GetAtivo() int {
	return u.Ativo
}

func (u *UsuarioParent) SetAtivo(bAtivo int) {
	u.Ativo = bAtivo
}
