package programador

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Handler struct {
	repo *Repository
	stmt *sql.Stmt
}

func NewHandler(db *sql.DB) *Handler {
	repo := NewRepository(db)
	stmt, err := repo.Insert()
	if err != nil {
		log.Fatalf("Erro ao preparar statement: %v", err)
	}

	return &Handler{
		repo: repo,
		stmt: stmt,
	}
}

func (h *Handler) CadastrarProgramador(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var p Programador
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if err := ValidarProgramador(p); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	id := uuid.New().String()
	_, err := h.stmt.Exec(id, p.Apelido, p.Nome, p.Nascimento, strings.Join(p.Stack, ","))
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			http.Error(w, "Apelido já existe", http.StatusUnprocessableEntity)
		} else {
			log.Println("Erro no insert:", err)
			http.Error(w, "Erro ao salvar programador", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Location", "/programadores/"+id)
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) ContarProgramadores(w http.ResponseWriter, r *http.Request) {
	count, err := h.repo.Count()
	if err != nil {
		http.Error(w, "Erro ao contar programadores", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strconv.Itoa(count)))
}
