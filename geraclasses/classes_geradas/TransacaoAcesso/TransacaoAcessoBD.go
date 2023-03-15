package TransacaoAcesso

import "geraclasses/classes/conexao"


type TransacaoAcessoBD struct {
    TransacaoAcessoBDParent
}

func _ (sBanco string) *TransacaoAcessoBD {
    return &TransacaoAcessoBD{TransacaoAcessoBDParent: TransacaoAcessoBDParent{Conexao: conexao.NewConexao(sBanco)}}
}