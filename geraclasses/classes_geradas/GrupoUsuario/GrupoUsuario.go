package GrupoUsuario

import "time"

type GrupoUsuario struct {
    GrupoUsuarioParent
    Id int64
	NmGrupoUsuario string
	DtCadastro time.Time
	Publicado int
	Ativo int
	
}

func _ (nId int64,sNmGrupoUsuario string,dDtCadastro time.Time,bPublicado int,bAtivo int) *GrupoUsuario {
    oGrupoUsuario := &GrupoUsuario{}
	
    oGrupoUsuario.SetId(nId)
	oGrupoUsuario.SetNmGrupoUsuario(sNmGrupoUsuario)
	oGrupoUsuario.SetDtCadastro(dDtCadastro)
	oGrupoUsuario.SetPublicado(bPublicado)
	oGrupoUsuario.SetAtivo(bAtivo)
	
    return oGrupoUsuario

}