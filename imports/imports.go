package imports

import (
	"context"
	"log"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/emscripten"
)

// Instantiate instantiates the "env" module used by Emscripten into the
// runtime default namespace.
//
// # Notes
//
//   - Closing the wazero.Runtime has the same effect as closing the result.
//   - To add more functions to the "env" module, use FunctionExporter.
//   - To instantiate into another wazero.Namespace, use FunctionExporter.
func Instantiate(ctx context.Context, r wazero.Runtime) (api.Closer, error) {
	builder := r.NewHostModuleBuilder("env")
	emscripten.NewFunctionExporter().ExportFunctions(builder)
	NewFunctionExporter().ExportFunctions(builder)
	return builder.Instantiate(ctx)
}

// FunctionExporter configures the functions in the "env" module used by
// Emscripten.
type FunctionExporter interface {
	// ExportFunctions builds functions to export with a wazero.HostModuleBuilder
	// named "env".
	ExportFunctions(builder wazero.HostModuleBuilder)
}

// NewFunctionExporter returns a FunctionExporter object with trace disabled.
func NewFunctionExporter() FunctionExporter {
	return &functionExporter{}
}

type functionExporter struct{}

// ExportFunctions implements FunctionExporter.ExportFunctions
func (e *functionExporter) ExportFunctions(b wazero.HostModuleBuilder) {
	b.NewFunctionBuilder().WithGoModuleFunction(invoke_dummy{}, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeF64, api.ValueTypeF64, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{}).Export("invoke_viiiddiiiiii")
	b.NewFunctionBuilder().WithGoModuleFunction(invoke_dummy{}, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{}).Export("invoke_viiiiiiiii")
	b.NewFunctionBuilder().WithGoModuleFunction(invoke_dummy{}, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{}).Export("invoke_viiiii")
	b.NewFunctionBuilder().WithGoModuleFunction(invoke_dummy{}, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeF64, api.ValueTypeF64, api.ValueTypeF64, api.ValueTypeF64, api.ValueTypeF64, api.ValueTypeF64}, []api.ValueType{}).Export("invoke_vidddddd")
	b.NewFunctionBuilder().WithGoModuleFunction(invoke_dummy{}, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeF64, api.ValueTypeF64, api.ValueTypeF64, api.ValueTypeF64, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeF64, api.ValueTypeI32}, []api.ValueType{}).Export("invoke_viddddiidi")
	b.NewFunctionBuilder().WithGoModuleFunction(invoke_dummy{}, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeF64, api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{}).Export("invoke_viiiiiidii")
	b.NewFunctionBuilder().WithGoModuleFunction(invoke_dummy{}, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{api.ValueTypeF64}).Export("invoke_dii")
	b.NewFunctionBuilder().WithGoModuleFunction(invoke_dummy{}, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{api.ValueTypeF64}).Export("invoke_di")
	b.NewFunctionBuilder().WithGoModuleFunction(invoke_dummy{}, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{api.ValueTypeI32}).Export("invoke_iiiiii")
	b.NewFunctionBuilder().WithGoModuleFunction(invoke_dummy{}, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{api.ValueTypeI32}).Export("invoke_iiiiiiiiii")
	b.NewFunctionBuilder().WithGoModuleFunction(invoke_dummy{}, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{api.ValueTypeI32}).Export("invoke_iiji")
}

type invoke_dummy struct {
}

func (i invoke_dummy) Call(ctx context.Context, mod api.Module, stack []uint64) {
	log.Println("invoke_dummy was called")
}
