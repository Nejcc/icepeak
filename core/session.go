package core

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// SessionManager handles session storage.
type SessionManager struct {
	sessions map[string]Session
	mu       sync.Mutex
}

// Session stores individual session data.
type Session struct {
	ID      string
	Values  map[string]interface{}
	Expires time.Time
}

// NewSessionManager creates a new SessionManager.
func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]Session),
	}
}

// CreateSession creates a new session.
func (sm *SessionManager) CreateSession(w http.ResponseWriter) Session {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	id := generateSessionID()
	session := Session{
		ID:      id,
		Values:  make(map[string]interface{}),
		Expires: time.Now().Add(24 * time.Hour), // 24-hour expiration
	}

	sm.sessions[id] = session
	http.SetCookie(w, &http.Cookie{
		Name:    "session_id",
		Value:   id,
		Expires: session.Expires,
	})
	return session
}

// GetSession retrieves a session from a request.
func (sm *SessionManager) GetSession(r *http.Request) (Session, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	cookie, err := r.Cookie("session_id")
	if err != nil {
		return Session{}, false
	}

	session, exists := sm.sessions[cookie.Value]
	return session, exists
}

// DeleteSession deletes a session.
func (sm *SessionManager) DeleteSession(sessionID string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.sessions, sessionID)
}

// generateSessionID creates a new session ID.
func generateSessionID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
