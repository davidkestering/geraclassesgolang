package GrupoUsuario

import "time"

type GrupoUsuarioParent struct {
    Id int64
	NmGrupoUsuario string
	DtCadastro time.Time
	Publicado int
	Ativo int
	
}

func (g *GrupoUsuarioParent) GetId() int64 {
	return g.Id
}

func (g *GrupoUsuarioParent) SetId(nId int64) {
	g.Id = nId
}

func (g *GrupoUsuarioParent) GetNmGrupoUsuario() string {
	return g.NmGrupoUsuario
}

func (g *GrupoUsuarioParent) SetNmGrupoUsuario(sNmGrupoUsuario string) {
	g.NmGrupoUsuario = sNmGrupoUsuario
}

func (g *GrupoUsuarioParent) GetDtCadastro() time.Time {
	return g.DtCadastro
}

func (g *GrupoUsuarioParent) SetDtCadastro(dDtCadastro time.Time) {
	g.DtCadastro = dDtCadastro
}

func (g *GrupoUsuarioParent) GetPublicado() int {
	return g.Publicado
}

func (g *GrupoUsuarioParent) SetPublicado(bPublicado int) {
	g.Publicado = bPublicado
}

func (g *GrupoUsuarioParent) GetAtivo() int {
	return g.Ativo
}

func (g *GrupoUsuarioParent) SetAtivo(bAtivo int) {
	g.Ativo = bAtivo
}

