package main

import (
	"fmt"
	"geraclasses/classes/Usuario"
	"geraclasses/classes/conexao"
	ct "geraclasses/constantes"
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

	//RECUPERA OK
	oConexao := conexao.NewConexao(ct.BANCO)
	oFachadaUsuario := Usuario.UsuarioBDParent{Conexao: oConexao}
	oUsuario, _ := oFachadaUsuario.Recupera(1)
	//fmt.Println(oUsuario.GetIdGrupoUsuario())

	//PRESENTE
	//oConexao := conexao.NewConexao(ct.BANCO)
	//oFachadaUsuario := Usuario.UsuarioBDParent{Conexao: oConexao}
	//nQtd := oFachadaUsuario.Presente(10)
	//fmt.Println(nQtd)

	//INSERT
	//oConexao := conexao.NewConexao(ct.BANCO)
	//oFachadaUsuario := Usuario.UsuarioBDParent{Conexao: oConexao}
	oUsuario.SetNmUsuario("D'agua")
	oUsuario.SetLogin("novo")
	nIdNovoUsuario, _ := oFachadaUsuario.Insere(oUsuario)
	fmt.Println(nIdNovoUsuario)

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
