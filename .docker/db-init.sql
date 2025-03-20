CREATE DATABASE mediaservice;
CREATE USER mediaservice WITH PASSWORD 'mediaservice';
GRANT ALL PRIVILEGES ON DATABASE mediaservice TO mediaservice;
\connect mediaservice
ALTER DATABASE mediaservice OWNER TO mediaservice;
ALTER SCHEMA public OWNER TO mediaservice;
