create table if not exists sensor_groups
(
    id         serial primary key,
    group_name varchar(255) unique
);
create table if not exists sensors
(
    id              serial primary key,
    sensor_group_id int not null references sensor_groups (id),
    sensor_index    int
);
create table if not exists temperature_readings
(
    id            serial primary key,
    sensor_id     int   not null references sensors (id),
    temperature_C float not null
);

