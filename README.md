<div align="center">
  <h1>GO-Visit Counter</h1>
  <img src="https://github.com/user-attachments/assets/3d4a0648-23d9-4f2e-b610-58b2c64524d0">
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

# Deploy from source

First, clone the repo

```bash
git clone --depth 1 https://github.com/arch-fan/nvim.git
```

Then, fill up the `.env` file with you data (example at `.env.example`).
Once you have the env file filled, you can run it with `docker compose up -d`
