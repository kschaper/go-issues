package whatever

// User represents a user.
type User struct {
	ID    int
	Email string
	// ... more fields like Username, CreatedAt etc.
}

// UserService manages users.
type UserService interface {
	GetUser(id int) (*User, error)
	// ... many more methods like:
	// CreateUser(email string) (*User, error)
	// UpdateUser(email string) (*User, error)
	// DestroyUser(id int) error
	// GetUserByEmail(email string) (*User, error)
	// GetUsers(ids []int) ([]*User, error)
}

// DatabaseClient represents a database client.
type DatabaseClient interface {
	// Open() error
	// Close() error
	UserService() UserService
	// ... more methods to access other services
}

// Env provides access to common stuff.
type Env struct {
	DB DatabaseClient
}
