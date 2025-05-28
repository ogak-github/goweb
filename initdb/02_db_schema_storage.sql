--
-- PostgreSQL database setup for PowerSync storage
-- This is the target database that will receive data from the source
--

-- Create the storage user with password
CREATE USER powersync_storage_user WITH PASSWORD 'secure_password';

-- Grant necessary privileges
-- The user should only have access to create and manage its own schema
GRANT CONNECT, CREATE ON DATABASE postgres TO powersync_storage_user;

-- Enable the uuid-ossp extension for UUID generation
-- This is required by PowerSync
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

-- Note: We're not creating any tables here because PowerSync will create its own schema
-- and tables based on the sync rules defined in sync_rules.yml
-- The user has been granted CREATE privilege, so PowerSync can create what it needs

