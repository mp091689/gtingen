package gtingen

import (
	"os"

	"github.com/MykytaPopov/gtingen/internal/router"
)

func Run() {
	r := router.NewRouter(os.Args)
	r.Process()
}
