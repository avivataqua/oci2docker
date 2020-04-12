package image

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/image/transports/alltransports"
	"github.com/containers/image/v5/copy"
	"github.com/containers/image/v5/signature"
	"github.com/containers/image/v5/types"
)

func Copy(ctx context.Context, src, dest string) error {

	policy, err := signature.DefaultPolicy(nil)
	if err != nil {
		return fmt.Errorf("failed to get default policy: %w", err)
	}
	policyCtx, err := signature.NewPolicyContext(policy)
	if err != nil {
		return fmt.Errorf("failed to generate default policy context: %w", err)
	}
	defer policyCtx.Destroy()

	srcRef, err := alltransports.ParseImageName(src)
	if err != nil {
		return fmt.Errorf("Invalid source name %s: %v", src, err)
	}
	destRef, err := alltransports.ParseImageName(dest)
	if err != nil {
		return fmt.Errorf("Invalid destination name %s: %v", dest, err)
	}

	sourceCtx := &types.SystemContext{

	}

	destCtx :=  &types.SystemContext{

	}

	opts := copy.Options{
		RemoveSignatures:false,
		ReportWriter:os.Stdout,
		SourceCtx: sourceCtx,
		DestinationCtx: destCtx,
	}

	_, err = copy.Image(ctx, policyCtx, destRef, srcRef, &opts)
	if err != nil {
		return fmt.Errorf("failed to copy image: %w", err)
	}
	return nil
}
