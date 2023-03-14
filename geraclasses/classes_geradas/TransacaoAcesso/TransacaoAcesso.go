package TransacaoAcesso

import "time"

type TransacaoAcesso struct {
    TransacaoAcessoParent
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

func NewTransacaoAcesso(nId int64,nIdTipoTransacao int64,nIdUsuario int64,sObjeto string,sCampo string,sValorAntigo string,sValorNovo string,sIp string,dDtCadastro time.Time,bPublicado int,bAtivo int) *TransacaoAcesso {
    oTransacaoAcesso := &TransacaoAcesso{}
	
    oTransacaoAcesso.SetId(nId)
	oTransacaoAcesso.SetIdTipoTransacao(nIdTipoTransacao)
	oTransacaoAcesso.SetIdUsuario(nIdUsuario)
	oTransacaoAcesso.SetObjeto(sObjeto)
	oTransacaoAcesso.SetCampo(sCampo)
	oTransacaoAcesso.SetValorAntigo(sValorAntigo)
	oTransacaoAcesso.SetValorNovo(sValorNovo)
	oTransacaoAcesso.SetIp(sIp)
	oTransacaoAcesso.SetDtCadastro(dDtCadastro)
	oTransacaoAcesso.SetPublicado(bPublicado)
	oTransacaoAcesso.SetAtivo(bAtivo)
	
    return oTransacaoAcesso

}