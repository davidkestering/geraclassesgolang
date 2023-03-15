package constantes

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"
)

// CONSTANTES BANCO
// DADOS BASICOS DO SISTEMA
const (
	PATH                       = "./"
	BANCO                      = "LOCAL"
	ACESSO_NEGADO              = "Você não tem permissão para acessar esta área."
	ACESSO_NEGADO_TRANSACAO    = "ACESSO NEGADO"
	ACESSO_PERMITIDO_TRANSACAO = "ACESSO PERMITIDO"
	ACESSO_TENTATIVA           = "TENTATIVA NAO REALIZADA"
	ID_USUARIO_SISTEMA         = 0
	GRUPO_ADMINISTRADOR        = 1
	ALTERAR_SENHA              = 11
	ACESSO_LOGIN               = 2
	ACESSO_LOGOUT              = 3
)

var (
	//IP_USUARIO  = os.Getenv("REMOTE_ADDR")
	//IP_SERVIDOR = os.Getenv("SERVER_ADDR") + os.Getenv("LOCAL_ADDR") + "SEM IP"
	IP_USUARIO  = RetornaIpUsuario()
	IP_SERVIDOR = RetornaIpServidor()
	CAMINHO     = RetornaNomedoProjeto()
	SITE        = "../../../../admin/"
	DATAHORA    = time.Now().Format("2006-01-02 15:04:05")
)

func RetornaIpUsuario() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return "SEM IP"
	}

	for _, address := range addrs {
		// Verifique se o endereço é um endereço IP e não IPv6
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			//fmt.Println(ipnet.IP.String())
			return ipnet.IP.String()
		}
	}

	return "SEM IP"
}

func RetornaIpServidor() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return RetornaIpUsuario()
	}

	for _, address := range addrs {
		// Verifique se o endereço é um endereço IP e não IPv6
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil && !ipnet.IP.IsPrivate() {
			//fmt.Println(ipnet.IP.String())
			return ipnet.IP.String()
		}
	}

	return RetornaIpUsuario()
}

func RetornaNomedoProjeto() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Obtém o nome do pacote (e do projeto) a partir do nome do diretório
	pkgName := filepath.Base(wd)

	// Cria a importação com o nome do pacote e imprime na tela
	importPath := fmt.Sprintf("%s", pkgName)
	return importPath
}
