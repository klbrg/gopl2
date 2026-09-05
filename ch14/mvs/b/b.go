package b

import "github.com/klbrg/gopl2/ch14/mvs/c"

// C reports the version of c that b was built against.
func C() string { return c.Version }
