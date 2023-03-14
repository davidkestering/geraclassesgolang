package CategoriaTipoTransacao

import "time"

type CategoriaTipoTransacaoParent struct {
    Id int64
	Descricao string
	DtCadastro time.Time
	Publicado int
	Ativo int
	
}

func (c *CategoriaTipoTransacaoParent) GetId() int64 {
	return c.Id
}

func (c *CategoriaTipoTransacaoParent) SetId(nId int64) {
	c.Id = nId
}

func (c *CategoriaTipoTransacaoParent) GetDescricao() string {
	return c.Descricao
}

func (c *CategoriaTipoTransacaoParent) SetDescricao(sDescricao string) {
	c.Descricao = sDescricao
}

func (c *CategoriaTipoTransacaoParent) GetDtCadastro() time.Time {
	return c.DtCadastro
}

func (c *CategoriaTipoTransacaoParent) SetDtCadastro(dDtCadastro time.Time) {
	c.DtCadastro = dDtCadastro
}

func (c *CategoriaTipoTransacaoParent) GetPublicado() int {
	return c.Publicado
}

func (c *CategoriaTipoTransacaoParent) SetPublicado(bPublicado int) {
	c.Publicado = bPublicado
}

func (c *CategoriaTipoTransacaoParent) GetAtivo() int {
	return c.Ativo
}

func (c *CategoriaTipoTransacaoParent) SetAtivo(bAtivo int) {
	c.Ativo = bAtivo
}

