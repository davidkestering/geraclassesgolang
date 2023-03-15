package GrupoUsuario

import "geraclasses/classes/conexao"


type GrupoUsuarioBD struct {
    GrupoUsuarioBDParent
}

func _ (sBanco string) *GrupoUsuarioBD {
    return &GrupoUsuarioBD{GrupoUsuarioBDParent: GrupoUsuarioBDParent{Conexao: conexao.NewConexao(sBanco)}}
}