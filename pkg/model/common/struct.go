package commonmodel

import (
	"fmt"
	"math/rand"
	"strings"
)

// EdgeDocument is a minimal document for use in edge collection.
// You can use this in your own edge document structures completely use your own.
// If you use your own, make sure to include a `_from` and `_to` field.
//this was separated out as it was becoming harder to keep the business logic and data layer separated
type Edge struct {
	From string `json:"_from,omitempty"`
	To   string `json:"_to,omitempty"`
}

// Collection returns the collection part of the ID.
func (e Edge) ToCollection() string {
	parts := strings.Split(e.To, "/")
	return parts[0]
}

// Collection returns the collection part of the ID.
func (e Edge) FromCollection() string {
	parts := strings.Split(e.From, "/")
	return parts[0]
}

type Meta struct {
	Key string `json:"_key,omitempty"`
	ID  string `json:"_id,omitempty"`
	Rev string `json:"_rev,omitempty"`
}

//Data - the interface structure which holds platform specific information.
type Data struct {
	Data interface{} `json:"data,omitempty"`
}

//TODO:: reimplement similar arangodb logic

// Validate validates the given id.
func (m Meta) Validate() error {
	if m.ID == "" {
		return fmt.Errorf("documentID is empty")
	}
	parts := strings.Split(m.ID, "/")
	if len(parts) != 2 {
		return fmt.Errorf("Expected 'collection/key', got '%s'", m.ID)
	}
	if parts[0] == "" {
		return fmt.Errorf("Collection part of '%s' is empty", m.ID)
	}
	if parts[1] == "" {
		return fmt.Errorf("Key part of '%s' is empty", m.ID)
	}
	return nil
}

// ValidateOrEmpty validates the given id unless it is empty.
// In case of empty, nil is returned.
func (m Meta) ValidateOrEmpty() error {
	if m.ID == "" {
		return nil
	}
	if err := m.Validate(); err != nil {
		return err
	}
	return nil
}

// IsEmpty returns true if the given ID is empty, false otherwise.
func (m Meta) IsEmpty() bool {
	return m.ID == ""
}

// Collection returns the collection part of the ID.
func (m Meta) Collection() string {
	parts := strings.Split(m.ID, "/")
	return parts[0]
}

type QueryEntryPositions []int

func (positions QueryEntryPositions) Shuffle(seed int64) QueryEntryPositions {
	rand.Seed(seed)
	rand.Shuffle(len(positions), func(i, j int) { positions[i], positions[j] = positions[j], positions[i] })
	return positions
}

func (positions QueryEntryPositions) RemoveExcess(totalEntries int) (trimmed QueryEntryPositions) {
	for _, position := range positions {
		if position < totalEntries {
			trimmed = append(trimmed, position)
		}
	}
	return
}
