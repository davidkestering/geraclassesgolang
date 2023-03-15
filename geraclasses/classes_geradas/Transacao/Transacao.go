package Transacao

import "time"

type Transacao struct {
    TransacaoParent
    Id int64
	IdTipoTransacao int64
	IdUsuario int64
	Objeto string
	Campo string
	ValorAntigo string
	ValorNovo string
	Ip string
	DtCadastro time.Time
	Publicado int
	Ativo int
	
}

func _ (nId int64,nIdTipoTransacao int64,nIdUsuario int64,sObjeto string,sCampo string,sValorAntigo string,sValorNovo string,sIp string,dDtCadastro time.Time,bPublicado int,bAtivo int) *Transacao {
    oTransacao := &Transacao{}
	
    oTransacao.SetId(nId)
	oTransacao.SetIdTipoTransacao(nIdTipoTransacao)
	oTransacao.SetIdUsuario(nIdUsuario)
	oTransacao.SetObjeto(sObjeto)
	oTransacao.SetCampo(sCampo)
	oTransacao.SetValorAntigo(sValorAntigo)
	oTransacao.SetValorNovo(sValorNovo)
	oTransacao.SetIp(sIp)
	oTransacao.SetDtCadastro(dDtCadastro)
	oTransacao.SetPublicado(bPublicado)
	oTransacao.SetAtivo(bAtivo)
	
    return oTransacao

}