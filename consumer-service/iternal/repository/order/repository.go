package order

import (
	"consumer-service/iternal/model"
	"consumer-service/iternal/repository"
	"context"
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	tableName = "orders"

	idColumn          = "id"
	statusColumn      = "status"
	updatedTimeColumn = "updated_time"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.OrderRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, order model.Order) error {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(
			idColumn,
			statusColumn,
			updatedTimeColumn,
		).
		Values(
			order.OrderInfo.Id,
			order.OrderInfo.Status,
			order.UpdateTime,
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	r.db.QueryRow(ctx, query, args...)

	return nil
}

func (r *repo) Get(ctx context.Context, id int64) (model.Order, error) {
	builder := sq.Select(idColumn, statusColumn, updatedTimeColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return model.Order{}, err
	}

	var order model.Order
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&order.OrderInfo.Id,
		&order.OrderInfo.Status,
		&order.UpdateTime,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Order{}, model.ErrObjectNotExists
	}
	if err != nil {
		return model.Order{}, model.ErrDb
	}

	return order, nil
}

func (r *repo) Update(ctx context.Context, order model.Order) error {
	builder := sq.Update(tableName).
		SetMap(map[string]interface{}{
			statusColumn:      order.OrderInfo.Status,
			updatedTimeColumn: order.UpdateTime,
		}).Where(sq.Eq{idColumn: order.OrderInfo.Id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	r.db.QueryRow(ctx, query, args...)

	return nil
}
