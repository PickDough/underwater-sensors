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
);

insert into sensor_groups(id, group_name) values (1, 'gamma');

insert into sensors(id, sensor_group_id, sensor_index) values (1, 1, 1);
insert into sensors(id, sensor_group_id, sensor_index) values (2, 1, 2);
insert into sensors(id, sensor_group_id, sensor_index) values (3, 1, 3);

insert into sensor_readings(id, sensor_id, temperature_c) values (1, 1, 11.0);
insert into sensor_readings(id, sensor_id, temperature_c) values (2, 2, 12.0);
insert into sensor_readings(id, sensor_id, temperature_c) values (3, 3, 13.0);

insert into sensor_readings(id, sensor_id, temperature_c) values (4, 1, 14.0);
insert into sensor_readings(id, sensor_id, temperature_c) values (5, 2, 15.0);
insert into sensor_readings(id, sensor_id, temperature_c) values (6, 3, 16.0);