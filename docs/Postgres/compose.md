```yaml
version: "3.8"
services:
  db:
    image: postgres
    container_name: local_pgdb
    restart: always
    healthcheck:
      test: ['CMD', 'pg_isready', '-q', '-d', 'postgres', '-U', 'root']
      timeout: 45s
      interval: 10s
      retries: 10
    ports:
      - "54320:5432"
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: admin
      POSTGRES_DB: alg_database
    volumes:
      - local_pgdata:/var/lib/postgresql/data
  pgadmin:
    image: dpage/pgadmin4
    container_name: pgadmin4_container
    restart: always
    ports:
      - "5050:80"
    environment:
      PGADMIN_DEFAULT_EMAIL: raj@nola.com
      PGADMIN_DEFAULT_PASSWORD: admin
    volumes:
      - pgadmin-data:/var/lib/pgadmin

volumes:
  local_pgdata:
  pgadmin-data:
```

