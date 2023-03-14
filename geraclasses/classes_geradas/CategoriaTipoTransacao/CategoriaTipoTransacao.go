package CategoriaTipoTransacao

import "time"

type CategoriaTipoTransacao struct {
    CategoriaTipoTransacaoParent
    Id int64
	Descricao string
	DtCadastro time.Time
	Publicado int
	Ativo int
	
}

func NewCategoriaTipoTransacao(nId int64,sDescricao string,dDtCadastro time.Time,bPublicado int,bAtivo int) *CategoriaTipoTransacao {
    oCategoriaTipoTransacao := &CategoriaTipoTransacao{}
	
    oCategoriaTipoTransacao.SetId(nId)
	oCategoriaTipoTransacao.SetDescricao(sDescricao)
	oCategoriaTipoTransacao.SetDtCadastro(dDtCadastro)
	oCategoriaTipoTransacao.SetPublicado(bPublicado)
	oCategoriaTipoTransacao.SetAtivo(bAtivo)
	
    return oCategoriaTipoTransacao

}