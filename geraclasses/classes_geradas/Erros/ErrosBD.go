package Erros

import "geraclasses/classes/conexao"


type ErrosBD struct {
    ErrosBDParent
}

func _ (sBanco string) *ErrosBD {
    return &ErrosBD{ErrosBDParent: ErrosBDParent{Conexao: conexao.NewConexao(sBanco)}}
}