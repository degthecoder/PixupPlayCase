# PixupPlay Casino Test Case

This project is a test case for PixupPLay, that handles transactions on a casino wallet.

---

## Tech Stack

- **Language:** Go (Golang)
- **Database:** MySQL
- **Driver:** `github.com/go-sql-driver/mysql`
- **HTTP Server:** `net/http` (no external router)

---

## Project Structure

/handlers → HTTP handler functions
/app → DB & settings config
/make_handle → Panic-safe middleware wrapper
/main.go → App entry point
/.env → Environment variables


## How to Run Locally

### 1. Create a MySQL database

Create a database named `casino_wallet` and run (dumped SQL file can be found on the github repository) :

```sql
CREATE TABLE players (
    player_id VARCHAR(64) PRIMARY KEY,
    wallet_id VARCHAR(64) UNIQUE NOT NULL,
    balance DECIMAL(18,2) NOT NULL DEFAULT 0.00,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    req_id VARCHAR(64) NOT NULL UNIQUE,
    type ENUM('bet', 'result') NOT NULL,
    player_id VARCHAR(64) NOT NULL,
    wallet_id VARCHAR(64) NOT NULL,
    round_id VARCHAR(64) NOT NULL,
    session_id VARCHAR(64) NOT NULL,
    game_code VARCHAR(64),
    amount DECIMAL(18, 2) NOT NULL,
    currency VARCHAR(8) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```
### 2. Add the environment valuables

HOST=localhost
PORT=8080

DB_HOST=localhost
DB_PORT=3306
DB_NAME=casino_wallet
DB_USER=root
DB_PASSWORD=your_password


### 3. Build and Run the code
```
go mod tidy
go run main.go
```


## Testing

The API end points can be tested with the given examples from the pdf. I used Postman for the testing.
