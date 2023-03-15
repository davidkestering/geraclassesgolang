package Transacao

import "geraclasses/classes/conexao"


type TransacaoBD struct {
    TransacaoBDParent
}

func _ (sBanco string) *TransacaoBD {
    return &TransacaoBD{TransacaoBDParent: TransacaoBDParent{Conexao: conexao.NewConexao(sBanco)}}
}