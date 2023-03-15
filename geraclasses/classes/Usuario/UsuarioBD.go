package Usuario

import "geraclasses/classes/conexao"


type UsuarioBD struct {
    UsuarioBDParent
}

func _ (sBanco string) *UsuarioBD {
    return &UsuarioBD{UsuarioBDParent: UsuarioBDParent{Conexao: conexao.NewConexao(sBanco)}}
}