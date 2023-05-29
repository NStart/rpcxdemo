package example

import (
	"context"
	"fmt"

	"github.com/smallnest/rpcx/share"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith int

func (a *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reqMeta := ctx.Value(share.ReqMetaDataKey).(map[string]string)
	resMeta := ctx.Value(share.ResMetaDataKey).(map[string]string)

	fmt.Println(reqMeta, resMeta)
	reply.C = args.A + args.B

	return nil

}
