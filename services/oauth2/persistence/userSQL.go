package persistence

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"oauth/model"
	"oauth/ulti"
)

const userTbale = "users"

type userPSQL struct {
	conn *pgxpool.Pool
}

func (u *userPSQL) FindAll(ctx context.Context) (users []*model.User, err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.Select("username", "password").From(userTbale)

	query, args, err := builder.ToSql()
	if err != nil {
		return
	}

	rows, err := u.conn.Query(ctx, query, args...)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.Username, &user.Password)

		if err != nil {
			return
		}

		users = append(users, &user)
	}

	return
}

func (u *userPSQL) FindOneByID(ctx context.Context, id string) (user *model.User, err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.Select("username", "password").From(userTbale).Where(sq.Eq{"username": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return
	}

	row := u.conn.QueryRow(ctx, query, args...)
	err = row.Scan(&user.Username, &user.Password)

	if err != nil {
		return
	}

	return
}

func (u *userPSQL) CheckUser(ctx context.Context, username string, password string) (err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	hashed := ulti.MD5(password)
	builder := psql.Select("username", "password").From(userTbale).Where(sq.And{sq.Eq{"username": username}, sq.Eq{"password": hashed}})

	query, args, err := builder.ToSql()
	if err != nil {
		return
	}

	row := u.conn.QueryRow(ctx, query, args...)
	var user model.User
	return row.Scan(&user.Username, &user.Password)

}

func (u *userPSQL) CreateUser(ctx context.Context, username string, password string) (err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	hashed := ulti.MD5(password)
	builder := psql.Insert(userTbale).Columns("username", "password").
		Values(username, hashed)
	query, args, err := builder.ToSql()
	if err != nil {
		return
	}

	_, err = u.conn.Exec(ctx, query, args...)
	return
}

func newUserRepoPSQL(ctx context.Context, conn *pgxpool.Pool) (repo *userPSQL, err error) {
	return &userPSQL{conn: conn}, nil
}
