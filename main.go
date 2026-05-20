package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const page = `<!doctype html><html><head><meta charset="utf-8"><title>FlazHost Gin App</title><style>*{margin:0;padding:0;box-sizing:border-box}body{font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,sans-serif;background:linear-gradient(135deg,#1e293b,#0f172a);min-height:100vh;display:flex;align-items:center;justify-content:center;color:#e2e8f0}.card{background:rgba(255,255,255,.05);border:1px solid rgba(255,255,255,.1);border-radius:16px;padding:3rem 2.5rem;text-align:center;max-width:480px}.card h1{font-size:1.6rem;margin-bottom:1rem}.card p{margin-bottom:.6rem;color:#cbd5e1;line-height:1.6}.meta{margin-top:1.5rem;padding-top:1.2rem;border-top:1px solid rgba(255,255,255,.08);font-size:.8rem;color:#64748b}</style></head><body><div class="card"><h1>Gin berjalan di FlazHost</h1><p>Aplikasi Gin Anda berhasil di-deploy.</p><p>Edit kode lalu push ke Git untuk redeploy otomatis.</p><div class="meta">Powered by FlazHost Runner</div></div></body></html>`

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(page))
	})
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	r.Run(":" + port)
}
