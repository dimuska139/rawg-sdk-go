package rawg

// Config represents a configuration of RAWG client
type Config struct {
	ApiKey   string // Your personal API key
	Language string // Language name (ISO 639-1)
	Rps      int    // Max allowed count of requests per second (default: 5)
}
