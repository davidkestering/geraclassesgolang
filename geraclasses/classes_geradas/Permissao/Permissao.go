package Permissao

import "time"

type Permissao struct {
    PermissaoParent
    IdTipoTransacao int64
	IdGrupoUsuario int64
	DtCadastro time.Time
	Publicado int
	Ativo int
	
}

func _ (nIdTipoTransacao int64,nIdGrupoUsuario int64,dDtCadastro time.Time,bPublicado int,bAtivo int) *Permissao {
    oPermissao := &Permissao{}
	
    oPermissao.SetIdTipoTransacao(nIdTipoTransacao)
	oPermissao.SetIdGrupoUsuario(nIdGrupoUsuario)
	oPermissao.SetDtCadastro(dDtCadastro)
	oPermissao.SetPublicado(bPublicado)
	oPermissao.SetAtivo(bAtivo)
	
    return oPermissao

}