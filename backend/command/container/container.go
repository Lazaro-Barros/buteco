package container

import (
	"log"
	"sync"
	"context"
	"time"
	"github.com/Lazaro-Barros/buteco/command/application/product"
	entities "github.com/Lazaro-Barros/buteco/command/domain/entities/product"
	"github.com/Lazaro-Barros/buteco/command/infra/drivens/repository"
	"github.com/Lazaro-Barros/buteco/command/infra/drivers/db"
	handler "github.com/Lazaro-Barros/buteco/command/infra/drivers/http/handler/product"
	"github.com/Lazaro-Barros/buteco/command/infra/drivers/http/router"
)

var container sync.Map
var timeoutToPingDB = 5

func Init() {
	container = sync.Map{}
}

func getService(alias string, f func() interface{}) interface{} {
	/* Reuse instance */
	if val, ok := container.Load(alias); ok {
		return val
	}

	service := f()
	/* Save instance in container */
	container.Store(alias, service)
	return service
}

func GetProductMemoryRepository() entities.ProductRepository {
	return getService("product.repository.memory", func() interface{} {
		return repository.NewProductMemoryRepository()
	},
	).(entities.ProductRepository)
}

func GetProductService() *product.ProductService {
	return getService("product.service", func() interface{} {
		return product.NewProductService(GetProductMemoryRepository())
	}).(*product.ProductService)
}

func GetProductHandler() *handler.ProductHandler {
	return getService("product.handler", func() interface{} {
		return handler.NewProductHandler(GetProductService())
	}).(*handler.ProductHandler)
}

func GetRouter() router.Router {
	return getService("router", func() interface{} {
		r := router.NewGinRouter()

		//product
		productHandler := GetProductHandler()
		r.PUT("/product", productHandler.UpdateProductHandler)
		r.POST("/product", productHandler.UpdateProductHandler)
		r.DELETE("/product", productHandler.DeleteProductHandler)

		return r
	}).(router.Router)
}

func GetDatabase() db.DB {
	return getService("database", func() interface{} {
		db, err := db.NewPgxAdapter("postgres://root:root@localhost:5432/root")
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutToPingDB)*time.Second)
		defer cancel()

		err = db.Ping(ctx)
		if err != nil {
			log.Fatalf("Database ping failed: %v", err)
		}
		return db
	}).(db.DB)
}
