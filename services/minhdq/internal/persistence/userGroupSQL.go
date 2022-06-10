package persistence

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"minhdq/internal/model"
)

const userGroupTable = "user_groups"

type userGroupPSQL struct {
	conn *pgxpool.Pool
}

func (u userGroupPSQL) FindAll(ctx context.Context) (usergroup []*model.UserGroup, err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.Select("group_id", "user_id", "created_at").From(userGroupTable)

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
		var userGroup model.UserGroup

		err = rows.Scan(&userGroup.GroupID, &userGroup.UserID, &userGroup.CreatedAt)

		if err != nil {
			return
		}

		usergroup = append(usergroup, &userGroup)
	}

	return
}

func (u userGroupPSQL) Save(ctx context.Context, userGroup model.UserGroup) (err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.Insert(userGroupTable).Columns("group_id", "user_id", "created_at").Values(userGroup.GroupID, userGroup.UserID, userGroup.CreatedAt)
	query, args, err := builder.ToSql()

	if err != nil {
		return
	}

	_, err = u.conn.Query(ctx, query, args...)
	return
}

func (u userGroupPSQL) Update(ctx context.Context, userGroup model.UserGroup) (record model.UserGroup, err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builer := psql.Update(userGroupTable).Set("created_at", userGroup.CreatedAt).Where(sq.And{sq.Eq{"group_id": userGroup.GroupID}, sq.Eq{"user_id": userGroup.UserID}})

	query, args, err := builer.ToSql()

	if err != nil {
		return
	}

	_, err = u.conn.Query(ctx, query, args...)

	return userGroup, err
}

func (u userGroupPSQL) Delete(ctx context.Context, userID int, groupID int) (err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.Delete(userGroupTable).Where(sq.And{sq.Eq{"user_id": userID}, sq.Eq{"group_id": groupID}})

	query, args, err := builder.ToSql()

	if err != nil {
		return
	}

	_, err = u.conn.Query(ctx, query, args...)

	return
}

func newUserGroupRepoPSQL(ctx context.Context, conn *pgxpool.Pool) (repo *userGroupPSQL, err error) {
	return &userGroupPSQL{conn: conn}, nil
}
