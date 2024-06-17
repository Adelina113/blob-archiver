package flags

import (
	opservice "github.com/ethereum-optimism/optimism/op-service"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/urfave/cli/v2"
)

const EnvVarPrefix = "BLOB_VALIDATOR"

var (
	BeaconClientTimeoutFlag = &cli.StringFlag{
		Name:    "beacon-client-timeout",
		Usage:   "The timeout duration for the beacon client",
		Value:   "10s",
		EnvVars: opservice.PrefixEnvVar(EnvVarPrefix, "CLIENT_TIMEOUT"),
	}
	L1BeaconClientUrlFlag = &cli.StringFlag{
		Name:     "l1-beacon-http",
		Usage:    "URL for a L1 Beacon-node API",
		Required: true,
		EnvVars:  opservice.PrefixEnvVar(EnvVarPrefix, "L1_BEACON_HTTP"),
	}
	BlobApiClientUrlFlag = &cli.StringFlag{
		Name:     "blob-api-http",
		Usage:    "URL for a Blob API",
		Required: true,
		EnvVars:  opservice.PrefixEnvVar(EnvVarPrefix, "BLOB_API_HTTP"),
	}
	BlocksPerMinuteClientFlag = &cli.IntFlag{
		Name:     "blocks-per-minute",
		Usage:    "The number of blocks per minute",
		Value:    5,
		Required: true,
		EnvVars:  opservice.PrefixEnvVar(EnvVarPrefix, "BLOCKS_PER_MINUTE"),
	}
	HoursOfBlobDataClientFlag = &cli.IntFlag{
		Name:     "hours-of-blob-data",
		Usage:    "The number of hours of blob data to fetch",
		Value:    2,
		Required: true,
		EnvVars:  opservice.PrefixEnvVar(EnvVarPrefix, "HOURS_OF_BLOB_DATA"),
	}
)

func init() {
	Flags = append(Flags, oplog.CLIFlags(EnvVarPrefix)...)
	Flags = append(Flags, BeaconClientTimeoutFlag, L1BeaconClientUrlFlag, BlobApiClientUrlFlag)
}

// Flags contains the list of configuration options available to the binary.
var Flags []cli.Flag
