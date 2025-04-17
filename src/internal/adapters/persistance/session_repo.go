package persistance

import models "movie-crud-application/src/internal/core"

type SessionRepo struct {
	db *Database
}

func NewSessionRepo(d *Database) models.SessionRepoImpl {
	return SessionRepo{db: d}
}

func (u SessionRepo) CreateSession(session models.Session) error {
	_, err := u.db.db.Exec("INSERT INTO sessions (id, user_id, token_hash, expires_at, issued_at) VALUES ($1, $2, $3, $4, $5) ON CONFLICT (user_id) DO UPDATE SET id = EXCLUDED.id, token_hash = EXCLUDED.token_hash, expires_at = EXCLUDED.expires_at, issued_at = EXCLUDED.issued_at", session.Id, session.UserId, session.TokenHash, session.ExpiresAt, session.IssuedAt)
	if err != nil {
		return err
	}
	return nil
}

func (u SessionRepo) GetSessionById(sessionId string) (models.Session, error) {
	var session models.Session

	query := "SELECT id, user_id, token_hash, expires_at, issued_at FROM sessions WHERE id=$1"

	err := u.db.db.QueryRow(query, sessionId).Scan(&session.Id, &session.UserId, &session.TokenHash, &session.ExpiresAt, &session.IssuedAt)
	if err != nil {
		return session, err
	}

	return session, nil
}

func (u SessionRepo) DeleteSession(id int) error {
	query := "DELETE FROM SESSIONS WHERE id=$1"

	_, err := u.db.db.Query(query, id)
	if err != nil {
		return err
	}

	return nil
}
