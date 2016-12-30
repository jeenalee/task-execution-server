package tesTaskEngineWorker

// Worker configuration.
type Config struct {
	MasterAddr  string
	WorkDir     string
	Timeout     int
	NumWorkers  int
	AllowedDirs []string
	LogPath     string
}