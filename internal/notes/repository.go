package notes

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository { return &Repository{db: db} }

func (r *Repository) Create(note *Note) error {
	return r.db.QueryRow(
		"INSERT INTO notes (title, content) VALUES ($1, $2) RETURNING id, created_at",
		note.Title, note.Content,
	).Scan(&note.ID, &note.CreatedAt)
}

func (r *Repository) GetAll() ([]Note, error) {
	rows, err := r.db.Query("SELECT id, title, content, created_at FROM notes ORDER BY created_at DESC")
	if err != nil { return nil, err }
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var n Note
		if err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt); err != nil { return nil, err }
		notes = append(notes, n)
	}
	return notes, nil
}

func (r *Repository) GetByID(id int) (*Note, error) {
	var n Note
	err := r.db.QueryRow("SELECT id, title, content, created_at FROM notes WHERE id=$1", id).
		Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt)
	if err != nil { return nil, err }
	return &n, nil
}

func (r *Repository) Update(note *Note) error {
	_, err := r.db.Exec("UPDATE notes SET title=$1, content=$2 WHERE id=$3", note.Title, note.Content, note.ID)
	return err
}

func (r *Repository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM notes WHERE id=$1", id)
	return err
}