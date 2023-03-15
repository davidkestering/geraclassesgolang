package Erros

import "time"

type Erros struct {
    ErrosParent
    Id int64
	Erro string
	IdUsuario int64
	Ip string
	DtCadastro time.Time
	Publicado int
	Ativo int
	
}

func _ (nId int64,sErro string,nIdUsuario int64,sIp string,dDtCadastro time.Time,bPublicado int,bAtivo int) *Erros {
    oErros := &Erros{}
	
    oErros.SetId(nId)
	oErros.SetErro(sErro)
	oErros.SetIdUsuario(nIdUsuario)
	oErros.SetIp(sIp)
	oErros.SetDtCadastro(dDtCadastro)
	oErros.SetPublicado(bPublicado)
	oErros.SetAtivo(bAtivo)
	
    return oErros

}