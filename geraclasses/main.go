package main

import (
	"geraclasses/views"
	"log"
	"net/http"
)

//type Usuario struct {
//	ID int64
//}

func main() {
	http.HandleFunc("/", views.IndexPage)
	http.HandleFunc("/listar_tabelas", views.ListarTabelasPage)
	http.HandleFunc("/processa", views.ProcessaPage)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}

	//env := os.Environ()
	//for _, v := range env {
	//	fmt.Println(v)
	//}

	//fmt.Println(ct.RetornaIpUsuario())
	//fmt.Println(ct.RetornaIpServidor())
	//
	//oConexao := conexao.NewConexao(ct.BANCO)
	//sSql := "SELECT id FROM seg_usuario"
	//oConexao.Execute(sSql)
	//oConexao.ExecuteInsert(sSql)
	//var voUsuario []Usuario
	//vAResult := oConexao.GetConsulta()
	//for vAResult.Next() {
	//	var Usuario Usuario
	//	err := vAResult.Scan(&Usuario.ID)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	voUsuario = append(voUsuario, Usuario)
	//}
	//
	//for i, Usuario := range voUsuario {
	//	fmt.Println(i)
	//	fmt.Println(Usuario.ID)
	//}
	//
	//defer vAResult.Close()
	//
	//nQtdLinhas, err := oConexao.RecordCountInsert()
	//fmt.Println("AQUI")
	//fmt.Println(nQtdLinhas, err)
	//
	//nQtdLinhas2, err := oConexao.RecordCountInsert()
	//fmt.Println("AQUI")
	//fmt.Println(nQtdLinhas2, err)
	//
	//oConexao.InsereErroSql(sSql)

}
