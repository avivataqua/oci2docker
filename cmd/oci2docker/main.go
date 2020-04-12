package main

import (
	"context"
	"fmt"
	"github.com/avivataqua/oci2docker/pkg/image"
	"os"
)

func main() {
	ctx := context.Background()
	if len(os.Args) < 3 {
		fmt.Println(fmt.Errorf("exactly src and dst arguments needed"))
		os.Exit(-1)
	}
	err := image.Copy(ctx, os.Args[1], os.Args[2] )
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}
}
