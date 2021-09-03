CREATE KEYSPACE IF NOT EXISTS catalog WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' };

CREATE TABLE IF NOT EXISTS catalog.brands (
userid text PRIMARY KEY,
item_count int,
last_update_timestamp timestamp
);