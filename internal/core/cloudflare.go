// PROJECT NAME: TNH-AI-MASTER-TEMPLATE
// FOLDER PATH: /internal/core/
// FILE NAME: cloudflare.go

package core

import (
    "context"
    "time"
    "github.com/cloudflare/cloudflare-go"
    "github.com/cloudflare/cloudflare-go/option"
)

func NewCloudflareClient(token string) *cloudflare.Client {
    return cloudflare.NewClient(
        option.WithAPIToken(token),
        option.WithMaxRetries(2),
    )
}

func GetZoneContext() (context.Context, context.CancelFunc) {
    // ปรับ Timeout ให้กริบตามกฎ 0.xxms ของเจ้านาย
    return context.WithTimeout(context.Background(), 5*time.Second)
}
