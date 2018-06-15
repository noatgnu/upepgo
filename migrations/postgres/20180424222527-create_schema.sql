
-- +migrate Up
CREATE SCHEMA if NOT EXISTS upep;

-- +migrate Down
DROP SCHEMA upep CASCADE ;
