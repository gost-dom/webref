package scripting

import (
	"errors"

	. "github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

// ESWrapper serves as a helper for building v8 wrapping code around go objects.
// Generated code assumes that a wrapper type is used with specific helper
// methods implemented.
type ESWrapper[T Entity] struct {
	host *ScriptHost
}

func (w ESWrapper[T]) GetInstance(info *v8.FunctionCallbackInfo) (result T, err error) {
	if ctx, ok := w.host.GetContext(info.Context()); ok {
		if instance, ok := ctx.GetCachedNode(info.This()); ok {
			if typedInstance, ok := instance.(T); ok {
				return typedInstance, nil
			}
		}
		err = v8.NewTypeError(ctx.host.iso, "Not an instance of NamedNodeMap")
		return
	}
	err = errors.New("Could not get context")
	return
}

func (w ESWrapper[T]) GetArgDOMString(args []*v8.Value, idx int) (result string, err error) {
	if idx >= len(args) {
		err = errors.New("Index out of range")
		return
	}
	result = args[idx].String()
	return
}

func (w ESWrapper[T]) GetArgByteString(args []*v8.Value, idx int) (result string, err error) {
	return w.GetArgDOMString(args, idx)
}
func (w ESWrapper[T]) GetArgUSVString(args []*v8.Value, idx int) (result string, err error) {
	return w.GetArgDOMString(args, idx)
}
