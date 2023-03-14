package TipoTransacao

import "time"

type TipoTransacao struct {
    TipoTransacaoParent
    Id int64
	IdCategoriaTipoTransacao int64
	Transacao string
	DtCadastro time.Time
	Publicado int
	Ativo int
	
}

func NewTipoTransacao(nId int64,nIdCategoriaTipoTransacao int64,sTransacao string,dDtCadastro time.Time,bPublicado int,bAtivo int) *TipoTransacao {
    oTipoTransacao := &TipoTransacao{}
	
    oTipoTransacao.SetId(nId)
	oTipoTransacao.SetIdCategoriaTipoTransacao(nIdCategoriaTipoTransacao)
	oTipoTransacao.SetTransacao(sTransacao)
	oTipoTransacao.SetDtCadastro(dDtCadastro)
	oTipoTransacao.SetPublicado(bPublicado)
	oTipoTransacao.SetAtivo(bAtivo)
	
    return oTipoTransacao

}