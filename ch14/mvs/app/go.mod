module github.com/klbrg/gopl2/ch14/mvs/app

go 1.26

require (
	github.com/klbrg/gopl2/ch14/mvs/a v1.0.0
	github.com/klbrg/gopl2/ch14/mvs/b v1.0.0
)

require github.com/klbrg/gopl2/ch14/mvs/c v1.2.0 // indirect

replace github.com/klbrg/gopl2/ch14/mvs/a v1.0.0 => ../a

replace github.com/klbrg/gopl2/ch14/mvs/b v1.0.0 => ../b

replace github.com/klbrg/gopl2/ch14/mvs/c v1.1.0 => ../c1.1

replace github.com/klbrg/gopl2/ch14/mvs/c v1.2.0 => ../c1.2
