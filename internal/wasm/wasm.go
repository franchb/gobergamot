package wasm

import (
	_ "embed"
)

//go:embed bergamot-translator-worker.wasm
var bergamotTranslatorWorkerWASM []byte
