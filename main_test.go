package main

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		require.NotNil(t, err)
		fmt.Println(err)
		return
	}
	defer db.Close()

	clientID := 1

	cl, err := selectClient(db, clientID)
	if err != nil {
		fmt.Println(err)
		return
	}
	require.Equal(t, cl.ID, clientID)
	assert.NotEmpty(t, cl.FIO, cl.Email, cl.Login, cl.Birthday)
	require.NoError(t, err)

}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	clientID := -1
	cl, err := selectClient(db, clientID)
	assert.Empty(t, cl.ID, cl.FIO, cl.Email, cl.Login, cl.Birthday)
	require.Equal(t, sql.ErrNoRows, err)
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	cl.ID, err = insertClient(db, cl)
	require.NoError(t, err)
	require.NotEmpty(t, cl.ID)

	client, err := selectClient(db, cl.ID)
	require.NoError(t, err)
	assert.Equal(t, client, cl)

}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	id, err := insertClient(db, cl)
	require.NotEmpty(t, id)

	client, err := selectClient(db, id)
	require.NoError(t, err)
	deleteClient(db, client.ID)

	_, err = selectClient(db, id)
	require.Equal(t, sql.ErrNoRows, err)

}
