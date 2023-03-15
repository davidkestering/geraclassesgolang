package CategoriaTipoTransacao

import "geraclasses/classes/conexao"


type CategoriaTipoTransacaoBD struct {
    CategoriaTipoTransacaoBDParent
}

func _ (sBanco string) *CategoriaTipoTransacaoBD {
    return &CategoriaTipoTransacaoBD{CategoriaTipoTransacaoBDParent: CategoriaTipoTransacaoBDParent{Conexao: conexao.NewConexao(sBanco)}}
}