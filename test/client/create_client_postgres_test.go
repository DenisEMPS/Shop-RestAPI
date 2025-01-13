package client

import (
	"school21_project1/pkg/repository"
	"school21_project1/types"
	"testing"

	_ "github.com/go-playground/assert/v2"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestClientPostgres_Create(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		logrus.Fatalf("error to open connect to stub database: %s", err)
	}
	defer db.Close()

	r := repository.NewClientPostgres(db)

	type mockBehavior func(client types.CreateClient, id int)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		client       types.CreateClient
		id           int
		wantErr      bool
	}{
		{
			name: "OK",
			client: types.CreateClient{
				Name:     "Evgeniy",
				Surname:  "Isaev",
				Birthday: "1985-11-11",
				Gender:   true,
				Country:  "USA",
				City:     "Florida",
			},
			mockBehavior: func(client types.CreateClient, id int) {
				mock.ExpectBegin()

				rows := sqlxmock.NewRows([]string{"adress_id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO adress").WithArgs(client.Country, client.City, client.Street).WillReturnRows(rows)

				rows = sqlxmock.NewRows([]string{"client_id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO client").WithArgs(client.Name, client.Surname, client.Birthday, client.Gender, id).WillReturnRows(rows)

				mock.ExpectCommit()
			},
			id: 1,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.client, testCase.id)

			got, err := r.Create(testCase.client)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.id, got)
			}
		})
	}
}
