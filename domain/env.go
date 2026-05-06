package domain

type SlskdEnv struct {
	HOST            string
	PORT            string
	USER            string
	PASSWORD        string
	SCRAPE_INTERVAL string
}

type ExposerEnv struct {
	PORT string
}

type Env struct {
	Slskd   SlskdEnv
	Exposer ExposerEnv
}
