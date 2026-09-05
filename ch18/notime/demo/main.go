// Demo shows the effect of the notime handler options.
package main

import (
	"log/slog"
	"os"

	"github.com/klbrg/gopl2/ch18/notime"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, notime.Options(nil)))
	logger.Info("open", "path", "/etc/passwd", "readonly", true)
	logger.Warn("slow read", "elapsed", "1.5s")
}
