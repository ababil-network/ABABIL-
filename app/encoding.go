package app

// EncodingConfig stores encoding configuration.
//
// This is a temporary structure.
// It will be replaced with the Cosmos SDK encoding
// configuration in the production implementation.
type EncodingConfig struct{}

// MakeEncodingConfig creates the default encoding configuration.
func MakeEncodingConfig() EncodingConfig {
	return EncodingConfig{}
}
