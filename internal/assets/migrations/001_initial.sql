-- +migrate Up
create table if not exists transfers
(
    id bigserial primary key not null,
    from_address char(42) not null,
    to_address char(42) not null,
    value numeric(78, 0) not null
    );
create index if not exists transfers_index on transfers (from_address, to_address);

-- +migrate Down
drop index if exists transfers_index;
drop table if exists transfers;