package pkg

import "context"

func GetUsernameFromCtx(ctx context.Context) string {
	return ctx.Value("username").(string)
}
