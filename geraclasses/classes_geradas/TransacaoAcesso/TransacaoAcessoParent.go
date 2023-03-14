package TransacaoAcesso

import "time"

type TransacaoAcessoParent struct {
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

func (t *TransacaoAcessoParent) GetId() int64 {
	return t.Id
}

func (t *TransacaoAcessoParent) SetId(nId int64) {
	t.Id = nId
}

func (t *TransacaoAcessoParent) GetIdTipoTransacao() int64 {
	return t.IdTipoTransacao
}

func (t *TransacaoAcessoParent) SetIdTipoTransacao(nIdTipoTransacao int64) {
	t.IdTipoTransacao = nIdTipoTransacao
}

func (t *TransacaoAcessoParent) GetIdUsuario() int64 {
	return t.IdUsuario
}

func (t *TransacaoAcessoParent) SetIdUsuario(nIdUsuario int64) {
	t.IdUsuario = nIdUsuario
}

func (t *TransacaoAcessoParent) GetObjeto() string {
	return t.Objeto
}

func (t *TransacaoAcessoParent) SetObjeto(sObjeto string) {
	t.Objeto = sObjeto
}

func (t *TransacaoAcessoParent) GetCampo() string {
	return t.Campo
}

func (t *TransacaoAcessoParent) SetCampo(sCampo string) {
	t.Campo = sCampo
}

func (t *TransacaoAcessoParent) GetValorAntigo() string {
	return t.ValorAntigo
}

func (t *TransacaoAcessoParent) SetValorAntigo(sValorAntigo string) {
	t.ValorAntigo = sValorAntigo
}

func (t *TransacaoAcessoParent) GetValorNovo() string {
	return t.ValorNovo
}

func (t *TransacaoAcessoParent) SetValorNovo(sValorNovo string) {
	t.ValorNovo = sValorNovo
}

func (t *TransacaoAcessoParent) GetIp() string {
	return t.Ip
}

func (t *TransacaoAcessoParent) SetIp(sIp string) {
	t.Ip = sIp
}

func (t *TransacaoAcessoParent) GetDtCadastro() time.Time {
	return t.DtCadastro
}

func (t *TransacaoAcessoParent) SetDtCadastro(dDtCadastro time.Time) {
	t.DtCadastro = dDtCadastro
}

func (t *TransacaoAcessoParent) GetPublicado() int {
	return t.Publicado
}

func (t *TransacaoAcessoParent) SetPublicado(bPublicado int) {
	t.Publicado = bPublicado
}

func (t *TransacaoAcessoParent) GetAtivo() int {
	return t.Ativo
}

func (t *TransacaoAcessoParent) SetAtivo(bAtivo int) {
	t.Ativo = bAtivo
}

