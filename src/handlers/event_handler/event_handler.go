package event_handler

import (
	"PixupPlayCaseTrial/src/app"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Event struct {
	Amount    string `json:"amount"`
	Currency  string `json:"currency"`
	GameCode  string `json:"game_code"`
	PlayerID  string `json:"player_id"`
	WalletID  string `json:"wallet_id"`
	ReqID     string `json:"req_id"`
	RoundID   string `json:"round_id"`
	SessionID string `json:"session_id"`
	Type      string `json:"type"` // bet-result
}

// POST /event -> {event-type}
func EventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var e Event
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if e.Type != "bet" && e.Type != "result" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "rejected",
			"reason": "Invalid type, must be 'bet' or 'result'",
		})
		return
	}

	switch e.Type {
	case "bet":
		fmt.Printf("Processing BET: %+v\n", e)
		if err := ProcessBet(e); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"status": "rejected",
				"reason": err.Error(),
			})
			return
		}

	case "result":
		fmt.Printf("Processing RESULT: %+v\n", e)
		if err := ProcessResult(e); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"status": "rejected",
				"reason": err.Error(),
			})
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func ProcessBet(e Event) error {
	amount, err := strconv.ParseFloat(e.Amount, 64)
	if err != nil || amount <= 0 {
		return fmt.Errorf("invalid bet amount")
	}

	tx, err := app.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var currentBalance float64
	query := `SELECT balance FROM players WHERE player_id = ? AND wallet_id = ? FOR UPDATE`
	err = tx.QueryRow(query, e.PlayerID, e.WalletID).Scan(&currentBalance)
	if err == sql.ErrNoRows {
		return fmt.Errorf("Player or Wallet not found.")
	} else if err != nil {
		return err
	}

	if currentBalance < amount {
		return fmt.Errorf("insufficient balance")
	}

	newBalance := currentBalance - amount
	_, err = tx.Exec("UPDATE players SET balance = ? WHERE player_id = ?", newBalance, e.PlayerID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO transactions
		(req_id, type, player_id, wallet_id, round_id, session_id, game_code, amount, currency)
		VALUES (?, 'bet', ?, ?, ?, ?, ?, ?, ?)`,
		e.ReqID, e.PlayerID, e.WalletID, e.RoundID, e.SessionID, e.GameCode, amount, e.Currency)
	if err != nil {
		return fmt.Errorf("transaction insert failed: %v", err)
	}

	return tx.Commit()
}

func ProcessResult(e Event) error {
	amount, err := strconv.ParseFloat(e.Amount, 64)

	if err != nil || amount <= 0 {
		return fmt.Errorf("Invalid result amount")
	}

	tx, err := app.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var exists int
	err = tx.QueryRow(`
		SELECT COUNT(*) FROM transactions 
		WHERE round_id = ? AND player_id = ? AND type = 'bet'
	`, e.RoundID, e.PlayerID).Scan(&exists)
	if err != nil {
		return err
	}
	if exists == 0 {
		return fmt.Errorf("No matching bet found for result")
	}

	err = tx.QueryRow(`
		SELECT COUNT(*) FROM transactions 
		WHERE round_id = ? AND player_id = ? AND type = 'result'
	`, e.RoundID, e.PlayerID).Scan(&exists)
	if err != nil {
		return err
	}
	if exists > 0 {
		return fmt.Errorf("Result already exists for this round")
	}

	_, err = tx.Exec(`
		UPDATE players SET balance = balance + ? WHERE player_id = ? AND wallet_id = ?
	`, amount, e.PlayerID, e.WalletID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO transactions
		(req_id, type, player_id, wallet_id, round_id, session_id, game_code, amount, currency)
		VALUES (?, 'result', ?, ?, ?, ?, ?, ?, ?)`,
		e.ReqID, e.PlayerID, e.WalletID, e.RoundID, e.SessionID, e.GameCode, amount, e.Currency)
	if err != nil {
		return fmt.Errorf("transaction insert failed: %v", err)
	}

	return tx.Commit()
}
