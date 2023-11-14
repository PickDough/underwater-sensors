create table if not exists sensor_groups
(
    id         serial primary key,
    group_name varchar(255) unique
);
create table if not exists sensors
(
    id              serial primary key,
    sensor_group_id int not null references sensor_groups (id),
    sensor_index    int not null,
    unique (sensor_group_id, sensor_index)
);
create table if not exists sensor_readings
(
    id            serial primary key,
    sensor_id     int   not null references sensors (id),
    temperature_C float not null,
    created_at    timestamp default current_timestamp
);
create table if not exists fish_readings
(
    id                 serial primary key,
    sensor_readings_id int          not null references sensor_readings (id),
    fish               varchar(255) not null,
    count              int          not null
)

