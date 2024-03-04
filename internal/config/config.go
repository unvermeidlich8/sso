package config

type config struct {
	Env         string     `yaml:"env" env-default:"local"`
	StoragePath string     `yaml:"storage_path" env-require:"true"`
	GRPC        GRPCConfig `yaml`
}
