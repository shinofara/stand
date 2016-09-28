package config

const (
	DEFAULT_COMPRESSION_FORMAT = "zip"
	DEFAULT_COMPRESSION_PREFIX = ""
	DEFAULT_ZIP_PASSWORD = ""	
)

type CompressionConfig struct {
	Prefix string `yaml:"prefix"` // prefix of the compression file name.
	Format string `yaml:"format"` // format of the compression file.
	Password string `yaml:"password"`
}

func mergeDefaultCompressionConfig(cfg *CompressionConfig) *CompressionConfig {

	defaultCompressionConfig := &CompressionConfig{
		Prefix: DEFAULT_COMPRESSION_PREFIX,
		Format: DEFAULT_COMPRESSION_FORMAT,
		Password: DEFAULT_ZIP_PASSWORD,
	}

	if cfg == nil {
		return defaultCompressionConfig
	}

	if cfg.Prefix != "" {
		defaultCompressionConfig.Prefix = cfg.Prefix
	}
	if cfg.Format != "" {
		defaultCompressionConfig.Format = cfg.Format
	}
	if cfg.Password != "" {
		defaultCompressionConfig.Password = cfg.Password
	}

	return defaultCompressionConfig
}
