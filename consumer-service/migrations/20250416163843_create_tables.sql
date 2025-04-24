-- +goose Up
create table if not exists orders (
    id int primary key,
    status varchar not null,
    updated_time timestamp not null
);

-- +goose Down
drop table orders;
