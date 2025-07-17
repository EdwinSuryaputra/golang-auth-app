package config

type Environment string

var (
	Local Environment = "LOCAL"
	Dev   Environment = "DEV"
	Stg   Environment = "STG"
	Prod  Environment = "PROD"
)
