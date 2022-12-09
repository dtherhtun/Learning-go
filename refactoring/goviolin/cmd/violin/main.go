package violin

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ardanlabs/conf/v3"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	log := log.New(os.Stdout, "VIOLIN : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	var cfg struct {
		Web struct {
			APIHOST         string        `conf:"default:0.0.0.0:8080"`
			ReadTimeout     time.Duration `conf:"default:5s"`
			WriteTimeout    time.Duration `conf:"default:5s"`
			ShutdownTimeout time.Duration `conf:"default:5s"`
		}
	}

	if _, err := conf.Parse("VIOLIN", &cfg); err != nil {
		if err == conf.ErrHelpWanted {
			usage, err := conf.UsageInfo("VIOLIN", &cfg)
			if err != nil {
				return fmt.Errorf("generating config usage: %w", err)
			}
			fmt.Println(usage)
			return nil
		}
		return fmt.Errorf("parsing config: %w", err)
	}
	return nil
}
