package Permissao

import "geraclasses/classes/conexao"


type PermissaoBD struct {
    PermissaoBDParent
}

func _ (sBanco string) *PermissaoBD {
    return &PermissaoBD{PermissaoBDParent: PermissaoBDParent{Conexao: conexao.NewConexao(sBanco)}}
}