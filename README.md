# Init project commands

```bash
go mod init github.com/LaJase/go-gin-ntfy
go get -u github.com/gin-gonic/gin
```

# Live module reload install

```bash
go install github.com/codegangsta/gin@latest
```

# Launch application

## Create env-file

```
NOTION_SECRET_TOKEN=<your-notion-token>
NOTION_DB_ID=<your-db-id>
```

## Use docker compose

```bash
docker compose up --build
```
