package persistence

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"minhdq/internal/model"
	"minhdq/internal/ulti"
)

const clientTable = "clients"

type clientPSQL struct {
	conn *pgxpool.Pool
}

func (c clientPSQL) FindAll(ctx context.Context) (clients []*model.Client, err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.Select("id", "Secret", "RedirectURIs", "GrantTypes", "ResponseTypes", "Scopes", "Audience", "Public").From(clientTable)

	query, args, err := builder.ToSql()
	if err != nil {
		return
	}

	rows, err := c.conn.Query(ctx, query, args...)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var client model.Client
		var secretString string
		err = rows.Scan(&client.ID, &secretString, &client.RedirectURIs, &client.GrantTypes, &client.ResponseTypes, &client.Scopes, &client.Audience, &client.Public)

		client.Secret = []byte(secretString)

		if err != nil {
			return
		}

		clients = append(clients, &client)
	}

	return
}

func (c clientPSQL) FindOneByID(ctx context.Context, id string) (client *model.Client, err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.Select("id", "Secret", "RedirectURIs", "GrantTypes", "ResponseTypes", "Scopes", "Audience", "Public").From(clientTable).Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return
	}

	row := c.conn.QueryRow(ctx, query, args...)
	var secretString string
	err = row.Scan(&client.ID, &secretString, &client.RedirectURIs, &client.GrantTypes, &client.ResponseTypes, &client.Scopes, &client.Audience, &client.Public)

	client.Secret = []byte(secretString)

	if err != nil {
		err = errors.New("Cant find record with provided client ID")
		return
	}

	return
}

func (c clientPSQL) CreateByID(ctx context.Context, id string, redirectURIs []string, responeType []string, scopes []string, grantTypes []string, public bool, audiens []string) (secret string, err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	secret = ulti.RandStringRunes(30)
	builder := psql.Insert(clientTable).Columns("id", "Secret", "RedirectURIs", "GrantTypes", "ResponseTypes", "Scopes", "Audience", "Public").
		Values(id, secret, redirectURIs, grantTypes, responeType, scopes, audiens, public)
	query, args, err := builder.ToSql()
	if err != nil {
		return
	}

	_, err = c.conn.Query(ctx, query, args...)
	return
}

func newClientGroupRepoPSQL(ctx context.Context, conn *pgxpool.Pool) (repo *clientPSQL, err error) {
	return &clientPSQL{conn: conn}, nil
}
