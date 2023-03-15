package TipoTransacao

import "geraclasses/classes/conexao"


type TipoTransacaoBD struct {
    TipoTransacaoBDParent
}

func _ (sBanco string) *TipoTransacaoBD {
    return &TipoTransacaoBD{TipoTransacaoBDParent: TipoTransacaoBDParent{Conexao: conexao.NewConexao(sBanco)}}
}