package main

import (
	"html/template"
	"log"
	"os"

	"github.com/CSpecht/rest2s3/models"
	routes "github.com/CSpecht/rest2s3/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	entrypoint := GetEnv("MINIO_ENDPOINT", "play.min.io")
	port := GetEnv("MINIO_PORT", "443")
	accesskey := GetEnv("MINIO_ACCESSKEY", "minio")
	secretkey := GetEnv("MINIO_SECRETKEY", "minio")
	mode := GetEnv("mode", "cli")
	bucket := GetEnv("MINIO_BUCKET", "upload")
	publicUrl := GetEnv("PUBLIC_URL", "http://localhost:8080")
	apiPort := GetEnv("API_PORT", "8080")
	var tmplFile = "idx.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	models.CurrentConfig = models.Config{Entrypoint: entrypoint + ":" + port, Bucket: bucket, AccessKey: accesskey, SecretKey: secretkey, PublicUrl: publicUrl, Mode: mode, Index: tmpl}
	log.Printf("Using endpoint: %s ", entrypoint)
	gin.SetMode(gin.ReleaseMode)
	r := routes.Routes()
	
	log.Printf("Running on port: %s ", apiPort)
	r.Run(":" + apiPort)

}

func GetEnv(name string, def string) string {
	if len(os.Getenv(name)) > 0 {
		return os.Getenv(name)
	} else {
		return def
	}
}
