package order

import (
	"consumer-service/iternal/model"
	"consumer-service/iternal/repository"
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	tableName = "orders"

	idColumn     = "id"
	statusColumn = "status"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.OrderRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, order model.OrderData) error {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(
			idColumn,
			statusColumn,
		).
		Values(
			order.Id,
			order.Status,
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	r.db.QueryRow(ctx, query, args...)

	return nil
}

func (r *repo) Get(ctx context.Context, id int64) (model.OrderData, error) {
	builder := sq.Select(idColumn, statusColumn).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return model.OrderData{}, err
	}

	var orderInfo model.OrderData
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&orderInfo.Id,
		&orderInfo.Status,
	)
	if err != nil {
		return model.OrderData{}, err
	}

	return orderInfo, nil
}

func (r *repo) Update(ctx context.Context, orderData model.OrderData) error {
	builder := sq.Update(tableName).
		SetMap(map[string]interface{}{
			statusColumn: orderData.Status,
		}).Where(sq.Eq{idColumn: orderData.Id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	r.db.QueryRow(ctx, query, args...)

	return nil
}
