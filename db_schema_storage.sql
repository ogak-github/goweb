--
-- PostgreSQL database setup for storage user with logical replication
--

-- Set WAL level to logical for replication
ALTER SYSTEM SET wal_level = 'logical';

-- Create the storage user with password
CREATE USER powersync_storage_user WITH PASSWORD 'secure_password';

-- Grant necessary privileges
GRANT CONNECT, CREATE ON DATABASE postgres TO powersync_storage_user;

-- Grant replication role to the user
ALTER USER powersync_storage_user WITH REPLICATION;
