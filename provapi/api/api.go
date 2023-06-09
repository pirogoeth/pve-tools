package api

import (
	"context"

	"github.com/gin-gonic/gin"

	apiCfg "github.com/pirogoeth/pve-tools/provapi/config"
)

func Init(ctx context.Context, cfg *apiCfg.Config, engine *gin.Engine) error {
	initHealthRoutes(engine.Group("/v1/health"))
	initImageFamiliesRoutes(engine.Group("/v1/imagefamilies"))
	initInstancesRoutes(engine.Group("/v1/instances"))

	return nil
}
