<div align="center">
  <h1>GO-Visit Counter</h1>
  <img src="https://vcounter.archfan.com/counter.svg?url=https://github.com/arch-fan/visit-counter">
  <p>A super-fast visit counter implemented in Go, powered with GORM and SQLite</p>
</div>

# Usage

Deploy with a single command:

```bash
docker run -p 8080:8080 -v ./db.sqlite3:/app/db.sqlite3 -e TOKEN=foo ghcr.io/arch-fan/visit-counter:latest
```

Or with docker compose:

```yml
services:
  counter:
    image: ghcr.io/arch-fan/visit-counter:latest
    ports:
      - 8080:8080
    volumes:
      - ./db.sqlite3:/app/db.sqlite3
    environment:
      TOKEN: foo
```

# Environment Variables

| Name               | Obligatory | Description                                             | DEFAULT      |
| ------------------ | ---------- | ------------------------------------------------------- | ------------ |
| `TOKEN`            | Yes        | A secret token for registering new URLs on the counter. | NO DEFAULT   |
| `ADDRESS`          | No         | The listening address.                                  | `0.0.0.0`    |
| `SQLITE_FILE_PATH` | No         | The path of the SQLite database inside the container    | `db.sqlite3` |
| `PORT`             | No         | The port where the app is going to listen               | `8080`       |

# App endpoints

- `GET /create`: Create a new URL to count visits. Query params:

  - `url`: The URL to count visits from.
  - `token`: The token to authorize the creation.

example: `GET /create?url=https://example.com&token=foo`

- `GET /counter.svg`: Get the SVG of the counter. Query params:

  - `url`: The URL to get the counter from.

example: `GET /counter.svg?url=https://example.com`

# Deploy from source

First, clone the repo

```bash
git clone --depth 1 https://github.com/arch-fan/nvim.git
```

Then, fill up the `.env` file with you data (example at `.env.example`).
Once you have the env file filled, you can run it with `docker compose up -d`
