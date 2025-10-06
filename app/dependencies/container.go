package dependencies

import (
	"database/sql"
	"sync"
)

type Container struct {
	mu sync.RWMutex
	db *sql.DB

	// Domain-specific containers
	BookContainer *BookContainer
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) SetDB(db *sql.DB) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.db = db
}

func (c *Container) GetDB() *sql.DB {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.db
}

// Global container instance
var globalContainer *Container
var once sync.Once

// GetContainer returns the global dependency container
func GetContainer() *Container {
	once.Do(func() {
		globalContainer = NewContainer()
	})
	return globalContainer
}

// InitializeContainer initializes the global container with database
func InitializeContainer(db *sql.DB) {
	container := GetContainer()
	container.SetDB(db)
}
