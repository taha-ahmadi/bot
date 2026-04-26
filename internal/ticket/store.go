package ticket

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Store struct {
	mu      sync.Mutex
	path    string
	tickets []Ticket
	drafts  map[int64]*Draft
}

func NewStore(path string) (*Store, error) {
	s := &Store{path: path, drafts: make(map[int64]*Draft)}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	if err := s.load(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Store) load() error {
	b, err := os.ReadFile(s.path)
	if errors.Is(err, os.ErrNotExist) {
		s.tickets = []Ticket{}
		return nil
	}
	if err != nil {
		return err
	}
	if len(b) == 0 {
		s.tickets = []Ticket{}
		return nil
	}
	return json.Unmarshal(b, &s.tickets)
}

func (s *Store) flush() error {
	b, err := json.MarshalIndent(s.tickets, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, b, 0o644)
}

func newID() string {
	b := make([]byte, 4)
	_, _ = rand.Read(b)
	return "DC-" + hex.EncodeToString(b)
}

func (s *Store) Create(t Ticket) (Ticket, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	t.ID = newID()
	now := time.Now().UTC()
	t.CreatedAt = now
	t.UpdatedAt = now
	t.Status = StatusOpen
	s.tickets = append(s.tickets, t)
	if err := s.flush(); err != nil {
		return Ticket{}, err
	}
	return t, nil
}

func (s *Store) ListByUser(userID int64) []Ticket {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Ticket, 0)
	for _, t := range s.tickets {
		if t.UserID == userID {
			out = append(out, t)
		}
	}
	return out
}

func (s *Store) Get(id string) (Ticket, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, t := range s.tickets {
		if t.ID == id {
			return t, true
		}
	}
	return Ticket{}, false
}

func (s *Store) Draft(userID int64) *Draft {
	s.mu.Lock()
	defer s.mu.Unlock()
	d, ok := s.drafts[userID]
	if !ok {
		d = &Draft{UserID: userID}
		s.drafts[userID] = d
	}
	return d
}

func (s *Store) ResetDraft(userID int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.drafts, userID)
}
