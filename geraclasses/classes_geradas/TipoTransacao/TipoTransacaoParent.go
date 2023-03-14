package TipoTransacao

import "time"

type TipoTransacaoParent struct {
    Id int64
	IdCategoriaTipoTransacao int64
	Transacao string
	DtCadastro time.Time
	Publicado int
	Ativo int
	
}

func (t *TipoTransacaoParent) GetId() int64 {
	return t.Id
}

func (t *TipoTransacaoParent) SetId(nId int64) {
	t.Id = nId
}

func (t *TipoTransacaoParent) GetIdCategoriaTipoTransacao() int64 {
	return t.IdCategoriaTipoTransacao
}

func (t *TipoTransacaoParent) SetIdCategoriaTipoTransacao(nIdCategoriaTipoTransacao int64) {
	t.IdCategoriaTipoTransacao = nIdCategoriaTipoTransacao
}

func (t *TipoTransacaoParent) GetTransacao() string {
	return t.Transacao
}

func (t *TipoTransacaoParent) SetTransacao(sTransacao string) {
	t.Transacao = sTransacao
}

func (t *TipoTransacaoParent) GetDtCadastro() time.Time {
	return t.DtCadastro
}

func (t *TipoTransacaoParent) SetDtCadastro(dDtCadastro time.Time) {
	t.DtCadastro = dDtCadastro
}

func (t *TipoTransacaoParent) GetPublicado() int {
	return t.Publicado
}

func (t *TipoTransacaoParent) SetPublicado(bPublicado int) {
	t.Publicado = bPublicado
}

func (t *TipoTransacaoParent) GetAtivo() int {
	return t.Ativo
}

func (t *TipoTransacaoParent) SetAtivo(bAtivo int) {
	t.Ativo = bAtivo
}

