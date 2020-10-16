package rawg

// Config represents a configuration of RAWG client
type Config struct {
	AppName  string // Name of your application. It will be set as User-Agent header
	Language string // Language name (ISO 639-1)
	Rps      int    // Max allowed count of requests per second (default: 5)
}
