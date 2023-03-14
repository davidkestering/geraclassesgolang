package Erros

import "time"

type ErrosParent struct {
    Id int64
	Erro string
	IdUsuario int64
	Ip string
	DtCadastro time.Time
	Publicado int
	Ativo int
	
}

func (e *ErrosParent) GetId() int64 {
	return e.Id
}

func (e *ErrosParent) SetId(nId int64) {
	e.Id = nId
}

func (e *ErrosParent) GetErro() string {
	return e.Erro
}

func (e *ErrosParent) SetErro(sErro string) {
	e.Erro = sErro
}

func (e *ErrosParent) GetIdUsuario() int64 {
	return e.IdUsuario
}

func (e *ErrosParent) SetIdUsuario(nIdUsuario int64) {
	e.IdUsuario = nIdUsuario
}

func (e *ErrosParent) GetIp() string {
	return e.Ip
}

func (e *ErrosParent) SetIp(sIp string) {
	e.Ip = sIp
}

func (e *ErrosParent) GetDtCadastro() time.Time {
	return e.DtCadastro
}

func (e *ErrosParent) SetDtCadastro(dDtCadastro time.Time) {
	e.DtCadastro = dDtCadastro
}

func (e *ErrosParent) GetPublicado() int {
	return e.Publicado
}

func (e *ErrosParent) SetPublicado(bPublicado int) {
	e.Publicado = bPublicado
}

func (e *ErrosParent) GetAtivo() int {
	return e.Ativo
}

func (e *ErrosParent) SetAtivo(bAtivo int) {
	e.Ativo = bAtivo
}

