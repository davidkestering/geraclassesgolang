package Permissao

import "time"

type PermissaoParent struct {
    IdTipoTransacao int64
	IdGrupoUsuario int64
	DtCadastro time.Time
	Publicado int
	Ativo int
	
}

func (p *PermissaoParent) GetIdTipoTransacao() int64 {
	return p.IdTipoTransacao
}

func (p *PermissaoParent) SetIdTipoTransacao(nIdTipoTransacao int64) {
	p.IdTipoTransacao = nIdTipoTransacao
}

func (p *PermissaoParent) GetIdGrupoUsuario() int64 {
	return p.IdGrupoUsuario
}

func (p *PermissaoParent) SetIdGrupoUsuario(nIdGrupoUsuario int64) {
	p.IdGrupoUsuario = nIdGrupoUsuario
}

func (p *PermissaoParent) GetDtCadastro() time.Time {
	return p.DtCadastro
}

func (p *PermissaoParent) SetDtCadastro(dDtCadastro time.Time) {
	p.DtCadastro = dDtCadastro
}

func (p *PermissaoParent) GetPublicado() int {
	return p.Publicado
}

func (p *PermissaoParent) SetPublicado(bPublicado int) {
	p.Publicado = bPublicado
}

func (p *PermissaoParent) GetAtivo() int {
	return p.Ativo
}

func (p *PermissaoParent) SetAtivo(bAtivo int) {
	p.Ativo = bAtivo
}

