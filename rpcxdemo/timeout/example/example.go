package example

import (
	"context"
	"log"
	"time"
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
	time.Sleep(10 * time.Second)

	select {
	case <-ctx.Done():
		log.Println("Request canceled or time out")
		return ctx.Err()
	default:
		reply.C = args.A + args.B
		return nil
	}

}
