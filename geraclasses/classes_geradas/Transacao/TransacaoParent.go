package Transacao

import "time"

type TransacaoParent struct {
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

func (t *TransacaoParent) GetId() int64 {
	return t.Id
}

func (t *TransacaoParent) SetId(nId int64) {
	t.Id = nId
}

func (t *TransacaoParent) GetIdTipoTransacao() int64 {
	return t.IdTipoTransacao
}

func (t *TransacaoParent) SetIdTipoTransacao(nIdTipoTransacao int64) {
	t.IdTipoTransacao = nIdTipoTransacao
}

func (t *TransacaoParent) GetIdUsuario() int64 {
	return t.IdUsuario
}

func (t *TransacaoParent) SetIdUsuario(nIdUsuario int64) {
	t.IdUsuario = nIdUsuario
}

func (t *TransacaoParent) GetObjeto() string {
	return t.Objeto
}

func (t *TransacaoParent) SetObjeto(sObjeto string) {
	t.Objeto = sObjeto
}

func (t *TransacaoParent) GetCampo() string {
	return t.Campo
}

func (t *TransacaoParent) SetCampo(sCampo string) {
	t.Campo = sCampo
}

func (t *TransacaoParent) GetValorAntigo() string {
	return t.ValorAntigo
}

func (t *TransacaoParent) SetValorAntigo(sValorAntigo string) {
	t.ValorAntigo = sValorAntigo
}

func (t *TransacaoParent) GetValorNovo() string {
	return t.ValorNovo
}

func (t *TransacaoParent) SetValorNovo(sValorNovo string) {
	t.ValorNovo = sValorNovo
}

func (t *TransacaoParent) GetIp() string {
	return t.Ip
}

func (t *TransacaoParent) SetIp(sIp string) {
	t.Ip = sIp
}

func (t *TransacaoParent) GetDtCadastro() time.Time {
	return t.DtCadastro
}

func (t *TransacaoParent) SetDtCadastro(dDtCadastro time.Time) {
	t.DtCadastro = dDtCadastro
}

func (t *TransacaoParent) GetPublicado() int {
	return t.Publicado
}

func (t *TransacaoParent) SetPublicado(bPublicado int) {
	t.Publicado = bPublicado
}

func (t *TransacaoParent) GetAtivo() int {
	return t.Ativo
}

func (t *TransacaoParent) SetAtivo(bAtivo int) {
	t.Ativo = bAtivo
}

