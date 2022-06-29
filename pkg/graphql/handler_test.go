package graphql

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gavv/httpexpect"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestHandler_create(t *testing.T) {
	//ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	handler := NewHandler(sqlxDB)

	expect := httpexpect.New(t, httptest.NewServer(handler).URL)

	t.Run("create category", func(t *testing.T) {

		queryDb := regexp.QuoteMeta("INSERT INTO categorias (id, nome) VALUES (?, ?)")

		mock.ExpectExec(queryDb).
			WithArgs(0, "intimas").
			WillReturnResult(sqlmock.NewResult(1, 1))

		queryApi := `
			mutation Mutation{
				createCategory (input:{
					name: "intimas"
				} )
			}
		`

		responseApi := `
			{
			  "data": {
				"createCategory": true
			  }
			}
		`

		response := expect.POST("/").WithJSON(map[string]interface{}{
			"query":         queryApi,
			"operationName": "Mutation",
		}).Expect()
		require.JSONEq(t, responseApi, response.Body().Raw())

	})

	t.Run("create store", func(t *testing.T) {

		queryDb := regexp.QuoteMeta(
			"INSERT INTO lojas (nome, cnpj, telefone, endereco, id_categoria) VALUES (?, ?, ?, ?, ?)")

		mock.ExpectExec(queryDb).
			WithArgs("loja_teste", "0001", "31313131", "rua torta", 1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		queryApi := `
			mutation createStore{
			createStore (input:{
				name:"loja_teste"
				cnpj:"0001",
				phoneNumber:"31313131",
				address:"rua torta",
				category: 1
			} ){
				id
			}
		}
		`

		responseApi := `{
		  "data": {
			"createStore": {
			  "id": 1
			}
		  }
		}`

		response := expect.POST("/").WithJSON(map[string]interface{}{
			"query":         queryApi,
			"operationName": "createStore",
		}).Expect()
		require.JSONEq(t, responseApi, response.Body().Raw())
	})

	t.Run("create cloth", func(t *testing.T) {

		queryDb := regexp.QuoteMeta(
			"INSERT INTO produtos (id_loja, id_categoria, nome, preco, quantidade, tamanho) VALUES (?, ?, ?, ?, ?, ?)")

		mock.ExpectExec(queryDb).
			WithArgs(1, 1, "loja_teste", float64(100), 2, "M").
			WillReturnResult(sqlmock.NewResult(1, 1))

		queryApi := `
			mutation createCloth {
				createClothes (input :{
					id_store:1,
					name: "loja_teste",
					category:1,
					size:"M",
					price:100,
					quantity:2
				}){
					id_store
					name
				}
			}
		`

		responseApi := `
		{
		  "data": {
			"createClothes": {
			  "id_store": 1,
			  "name": "loja_teste"
			}
		  }
		}
		`

		response := expect.POST("/").WithJSON(map[string]interface{}{
			"query":         queryApi,
			"operationName": "createCloth",
		}).Expect()
		require.JSONEq(t, responseApi, response.Body().Raw())
	})

}
