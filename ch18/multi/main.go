// Multi sends every record to two handlers at once.
package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/klbrg/gopl2/ch18/maphandler"
	"github.com/klbrg/gopl2/ch18/notime"
)

func main() {
	text := slog.NewTextHandler(os.Stdout, notime.Options(nil))
	mem := maphandler.New(slog.LevelDebug)
	logger := slog.New(slog.NewMultiHandler(text, mem))

	logger.Info("started", "pid", 4242)
	logger.Debug("cache warm", "entries", 3)

	fmt.Printf("%d records retained in memory\n", len(mem.Records()))
}
