# Vue app with go rest api

## Setup
1. Database setup
```bash
CREATE DATABASE <databasename>;
GRANT ALL PRIVILEGES ON DATABASE <databasename> TO <postgresuser>;
\c <databasename>
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO <postgresuser>;
```

2. Run schema against database
```bash
psql -U <postgresuser> -d <databasename> -a -f server/database/schema.sql
```

## TODO
