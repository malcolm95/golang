package config

var (
	PORT uint32
)

func Load() {
	PORT = 4001 // initialize PORT
}
