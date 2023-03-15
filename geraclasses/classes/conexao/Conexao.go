package conexao

import "C"
import (
	"database/sql"
	"errors"
	"fmt"
	ct "geraclasses/constantes"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"strings"
)

type Conexao struct {
	Conexao        *sql.DB
	ConsultaInsert sql.Result
	Consulta       *sql.Rows
	Banco          *sql.DB
	Erro           error
	SqlError       error
	QtdTabelas     int64
	QtdCampos      int64
	Servidor       string
	LastId         int64
}

type Campos struct {
	TableCatalog, TableSchema, TableName, ColumnName, DataType, Pri string
}

func NewConexao(Servidor string) *Conexao {
	conexao := &Conexao{}
	conexao.Servidor = Servidor
	conexao.SetServidor(Servidor)

	if Servidor == "LOCAL" {
		err := conexao.ConectaBD("WINDOWS_SERVER\\MSSQLSERVER01", "sa", "fb123", "novodbteste")
		if err != nil {
			return nil
		}
	} else if Servidor == "BANCO" {
		err := conexao.ConectaBD("localhost", "", "", "")
		if err != nil {
			return nil
		}
	} else {
		return nil
	}

	return conexao
}

func (c *Conexao) SetServidor(Servidor string) {
	c.Servidor = Servidor
}

func (c *Conexao) GetServidor() string {
	return c.Servidor
}

func (c *Conexao) ConectaBD(sHost, sUser, sSenha, sBanco string) error {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", sHost, sUser, sSenha, sBanco)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return err
	}

	c.Conexao = db
	return nil
}

func (c *Conexao) ExecuteInsert(sSql string) error {
	tx, err := c.Conexao.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	//result, err := tx.Exec(sSql)
	err = tx.QueryRow(sSql).Scan(&c.LastId)
	//rows, err := tx.Query(sSql)
	if err != nil {
		c.SqlError = err
		c.Erro = fmt.Errorf("Ocorreu o seguinte erro na consulta: %s Query: %s", c.SqlError, sSql)
		c.InsereErroSql(sSql)
		return c.GetErro()
	}
	//for rows.Next() {
	//	rows.Scan(c.LastId)
	//}

	//row, err := c.Conexao.Query(sSql)
	//if err != nil {
	//	c.SqlError = err
	//	c.Erro = fmt.Errorf("Ocorreu o seguinte erro na consulta: %s Query: %s", c.SqlError, sSql)
	//	c.InsereErroSql(sSql)
	//	return c.GetErro()
	//} else {
	//	row.Next()
	//	row.Scan(&c.LastId)
	//}
	//c.LastId, _ = result.LastInsertId()

	err = tx.Commit()
	if err != nil {
		return err
	}

	//c.ConsultaInsert = result

	return nil
}

func (c *Conexao) Execute(sSql string) error {
	var err error
	c.Consulta, err = c.Conexao.Query(sSql)
	if err != nil {
		c.SqlError = err
		c.Erro = fmt.Errorf("Ocorreu o seguinte erro na consulta: %s Query: %s", c.SqlError, sSql)
		c.InsereErroSql(sSql)
		return c.GetErro()
	}
	return nil
}

func (c *Conexao) InsereErroSql(sSql string) error {
	tx, err := c.Conexao.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	sMsgErro := ""
	if c.GetErro() != nil {
		sMsgErro = c.EscapeString(c.GetErro().Error())
		//fmt.Println(sMsgErro)
	}

	sSqlErroExecucao := fmt.Sprintf(`insert into seg_erros_sql (erro,ip,publicado) values ('%s','%s',1)`, sMsgErro, ct.IP_SERVIDOR)
	//fmt.Println(sSqlErroExecucao)
	_, err = tx.Exec(sSqlErroExecucao)
	if err != nil {
		//fmt.Println("NAO SALVOU O ERRO")
		//fmt.Println(err)
		c.SqlError = err
		c.Erro = errors.New(fmt.Sprintf("Ocorreu o seguinte erro na inserção do erro na tabela seg_erros_sql: %s <br> Query: %s", c.SqlError.Error(), sSql))
		return fmt.Errorf("Ocorreu o seguinte erro na inserção do erro na tabela seg_erros_sql: %s <br> Query: %s", err.Error(), sSqlErroExecucao)
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c *Conexao) GetErro() error {
	return c.Erro
}

func (c *Conexao) RecordCountInsert() (int64, error) {
	numRows, err := c.ConsultaInsert.RowsAffected()
	if err != nil {
		return 0, err
	}

	return numRows, nil
}

func (c *Conexao) RecordCount() (int64, error) {
	var count int64
	for c.Consulta.Next() {
		count++
	}
	err := c.Consulta.Err()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (c *Conexao) FetchObject2() sql.Result {
	return c.ConsultaInsert
}

func (c *Conexao) FetchObject() *sql.Rows {
	return c.Consulta
}

func (c *Conexao) Close() {
	c.Conexao.Close()
}

func (c *Conexao) GetErroSql() error {
	return c.SqlError
}

func (c *Conexao) GetConexao() *sql.DB {
	return c.Conexao
}

func (c *Conexao) GetConsulta2() sql.Result {
	return c.ConsultaInsert
}

func (c *Conexao) GetConsulta() *sql.Rows {
	return c.Consulta
}

func (c *Conexao) EscapeString(sAtributo string) string {
	return strings.ReplaceAll(sAtributo, "'", "''")
}

func (c *Conexao) UnescapeString(sEscapedString string) string {
	return strings.ReplaceAll(sEscapedString, "''", "'")
}

func (c *Conexao) GetLastId() int64 {
	//var lastID int64
	//rows, _ := c.Conexao.Query("SELECT SCOPE_IDENTITY()")
	//for rows.Next() {
	//	fmt.Println(rows.Err())
	//	rows.Scan(&lastID)
	//	fmt.Println(lastID)
	//	fmt.Println(rows.Err())
	//}
	//return lastID
	return c.LastId
}

func (c *Conexao) SetQtdTabelas(nQtdTabelas int64) {
	c.QtdTabelas = nQtdTabelas
}

func (c *Conexao) GetQtdTabelas() int64 {
	nQtdTabelas := c.CarregaQtdTabelas()
	return nQtdTabelas
}

func (c *Conexao) SetQtdCampos(nQtdCampos int64) {
	c.QtdCampos = nQtdCampos
}

func (c *Conexao) GetQtdCampos() int64 {
	return c.QtdCampos
}

func (c *Conexao) CarregaQtdTabelas() int64 {
	var err error
	c.Consulta, err = c.Conexao.Query("SELECT count(0) as qtd FROM INFORMATION_SCHEMA.TABLES")
	if err != nil {
		// Trata o erro
		log.Fatal(err)
	}
	defer c.Consulta.Close()

	if c.Consulta.Next() {
		//var nQtdTabelas int64
		err := c.Consulta.Scan(&c.QtdTabelas)
		if err != nil {
			// Trata o erro
			log.Fatal(err)
		}
		//c.SetQtdTabelas(nQtdTabelas)
	}

	return c.QtdTabelas
}

func (c *Conexao) PegaTabelas() [][]interface{} {
	var err error
	c.Consulta, err = c.Conexao.Query("SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES")
	if err != nil {
		// Trata o erro
		log.Fatal(err)
	}
	defer c.Consulta.Close()

	// Obter informações sobre as colunas
	colunas, err := c.Consulta.Columns()
	if err != nil {
		// Trata o erro
		log.Fatal(err)
	}

	//fmt.Println(colunas)

	// Criar slice de interface{} para armazenar os valores das colunas
	valores := make([]interface{}, len(colunas))
	for i := range valores {
		valores[i] = new(interface{})
	}

	// Criar slice de slices de interface{} para armazenar as linhas
	linhas := make([][]interface{}, 0)

	// Iterar sobre os resultados
	for c.Consulta.Next() {
		// Escanear os valores das colunas em um slice de interface{}
		err := c.Consulta.Scan(valores...)
		if err != nil {
			// Trata o erro
			log.Fatal(err)
		}

		// Criar um slice de interface{} para armazenar os valores desta linha
		linha := make([]interface{}, len(colunas))
		for i, v := range valores {
			linha[i] = *v.(*interface{})
		}

		// Adicionar esta linha à matriz de linhas
		linhas = append(linhas, linha)
	}

	// Retornar a matriz de linhas
	return linhas
}

func (c *Conexao) CarregaQtdCampos(sNomeBanco string, sNomeTabela string) int64 {
	var err error
	c.Consulta, err = c.Conexao.Query(fmt.Sprintf("SELECT COUNT(COLUMN_NAME) AS qtd FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_CATALOG = '%s' AND TABLE_NAME = '%s'", sNomeBanco, sNomeTabela))
	if err != nil {
		// Trata o erro
		log.Fatal(err)
	}
	//var qtd int64
	if c.Consulta.Next() {
		err := c.Consulta.Scan(&c.QtdCampos)
		if err != nil {
			// Trata o erro
			log.Fatal(err)
		}
	}

	//c.SetQtdCampos(qtd)
	return c.QtdCampos
}

func (c *Conexao) PegaCampos(sNomeBanco, sNomeTabela string) []Campos {
	sSql := fmt.Sprintf(`select T.TABLE_CATALOG, T.TABLE_SCHEMA, T.TABLE_NAME, C.COLUMN_NAME, C.DATA_TYPE, SUBSTRING(CCU.CONSTRAINT_NAME,1,2) as PRI from INFORMATION_SCHEMA.TABLES T
            join INFORMATION_SCHEMA.COLUMNS C
                on C.TABLE_CATALOG = T.TABLE_CATALOG
                and C.TABLE_NAME = T.TABLE_NAME
                and C.TABLE_SCHEMA = T.TABLE_SCHEMA
            left join INFORMATION_SCHEMA.CONSTRAINT_COLUMN_USAGE CCU
                on CCU.TABLE_CATALOG = C.TABLE_CATALOG
                and CCU.TABLE_NAME = C.TABLE_NAME
                and CCU.TABLE_SCHEMA = C.TABLE_SCHEMA
                and CCU.COLUMN_NAME = C.COLUMN_NAME
                and SUBSTRING(CCU.CONSTRAINT_NAME,1,2) = 'PK'
            where T.TABLE_CATALOG = '%s' and T.TABLE_NAME = '%s' order by C.ORDINAL_POSITION`, sNomeBanco, sNomeTabela)
	//fmt.Println(sSql)
	rows, err := c.Conexao.Query(sSql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var voObjeto []Campos
	for rows.Next() {
		var Objeto Campos
		var TableCatalog, TableSchema, TableName, ColumnName, DataType, Pri sql.NullString
		if err := rows.Scan(&TableCatalog, &TableSchema, &TableName, &ColumnName, &DataType, &Pri); err != nil {
			log.Fatal(err)
		}
		Objeto.TableCatalog = TableCatalog.String
		Objeto.TableSchema = TableSchema.String
		Objeto.TableName = TableName.String
		Objeto.ColumnName = ColumnName.String
		Objeto.DataType = DataType.String
		if Pri.Valid {
			Objeto.Pri = Pri.String
		}
		voObjeto = append(voObjeto, Objeto)
	}
	return voObjeto
}
