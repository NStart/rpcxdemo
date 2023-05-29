package example

import "context"

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith int

func (a *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A*10 + args.B*10
	return nil
}

type Arith2 int

func (a *Arith2) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}
