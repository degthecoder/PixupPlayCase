package player_handler

import (
	"PixupPlayCaseTrial/src/app"
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
)

type Player struct {
	PlayerID    string  `json:"player_id"`
	WalletID    string  `json:"wallet_id"`
	Balance     float64 `json:"balance"`
	DateCreated string  `json:"date_created"`
}

// GET /players
func GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	query := `SELECT player_id, wallet_id, balance FROM players;`
	rows, err := app.Db.Query(query)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var players []Player
	for rows.Next() {
		var p Player
		if err := rows.Scan(&p.PlayerID, &p.WalletID, &p.Balance); err != nil {
			http.Error(w, "DB error", http.StatusInternalServerError)
			return
		}
		players = append(players, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}

// GET /wallet/{player_id}
func GetPlayerWallet(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/wallet/")
	playerID := strings.Trim(path, "/")

	if playerID == "" {
		http.Error(w, "Missing player_id", http.StatusBadRequest)
		return
	}

	var p Player
	query := `SELECT player_id, wallet_id, balance FROM players WHERE player_id = ?`

	err := app.Db.QueryRow(query, playerID).
		Scan(&p.PlayerID, &p.WalletID, &p.Balance)

	if err == sql.ErrNoRows {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
