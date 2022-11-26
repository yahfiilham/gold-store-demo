// This file is safe to edit. Once it exists it will not be overwritten

package apis

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/yahfiilham/gold-store-demo/configs"
	"github.com/yahfiilham/gold-store-demo/internal/apis/operations"
	"github.com/yahfiilham/gold-store-demo/internal/apis/operations/balance"
	"github.com/yahfiilham/gold-store-demo/internal/apis/operations/buyback"
	"github.com/yahfiilham/gold-store-demo/internal/apis/operations/health_check"
	"github.com/yahfiilham/gold-store-demo/internal/apis/operations/price"
	"github.com/yahfiilham/gold-store-demo/internal/apis/operations/topup"
	"github.com/yahfiilham/gold-store-demo/internal/apis/operations/transaction"
	"github.com/yahfiilham/gold-store-demo/internal/handlers"
	"github.com/yahfiilham/gold-store-demo/pkg/models"
)

//go:generate swagger generate server --target ../../../gold-store-demo --name GoldStoreDemo --spec ../../api/swagger.yml --model-package pkg/models --server-package internal/apis --principal interface{}

func configureFlags(api *operations.GoldStoreDemoAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.GoldStoreDemoAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.HealthCheckGetHealthCheckHandler == nil {
		api.HealthCheckGetHealthCheckHandler = health_check.GetHealthCheckHandlerFunc(func(params health_check.GetHealthCheckParams) middleware.Responder {
			return middleware.NotImplemented("operation health_check.GetHealthCheck has not yet been implemented")
		})
	}

	cfg := configs.NewConfig()

	// health check
	api.HealthCheckGetHealthCheckHandler = health_check.GetHealthCheckHandlerFunc(func(ghcp health_check.GetHealthCheckParams) middleware.Responder {
		return health_check.NewGetHealthCheckOK().WithPayload(&models.BaseResponse{
			Code:    http.StatusOK,
			Message: "server is running",
		})
	})

	// price
	api.PriceSavePriceHandler = price.SavePriceHandlerFunc(func(spp price.SavePriceParams) middleware.Responder {
		data, err := handlers.SavePrice(cfg, spp.Data)
		if err != nil {
			return price.NewSavePriceDefault(http.StatusInternalServerError).WithPayload(
				&models.BaseResponse{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				})
		}

		return price.NewSavePriceCreated().WithPayload(&price.SavePriceCreatedBody{
			Data: data,
		})
	})

	api.PriceGetPriceHandler = price.GetPriceHandlerFunc(func(spp price.GetPriceParams) middleware.Responder {
		data, err := handlers.GetPrice(cfg)
		if err != nil {
			return price.NewGetPriceDefault(http.StatusInternalServerError).WithPayload(
				&models.BaseResponse{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				})
		}

		return price.NewGetPriceOK().WithPayload(&price.GetPriceOKBody{
			Data: data,
		})
	})

	// topup
	api.TopupSaveTopupGoldHandler = topup.SaveTopupGoldHandlerFunc(func(stgp topup.SaveTopupGoldParams) middleware.Responder {
		if err := handlers.SaveTopupGold(cfg, stgp.Data); err != nil {
			return topup.NewSaveTopupGoldDefault(http.StatusInternalServerError).WithPayload(
				&models.BaseResponse{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				})
		}

		return topup.NewSaveTopupGoldCreated().WithPayload(&models.BaseResponse{
			Code:    http.StatusCreated,
			Message: fmt.Sprintf("success top up gold with account no %s", stgp.Data.AccountNumber),
		})
	})

	// buyback
	api.BuybackSaveBuybackHandler = buyback.SaveBuybackHandlerFunc(func(sbp buyback.SaveBuybackParams) middleware.Responder {
		if err := handlers.SaveBuybackGold(cfg, sbp.Data); err != nil {
			return buyback.NewSaveBuybackDefault(http.StatusInternalServerError).WithPayload(
				&models.BaseResponse{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				})
		}

		return buyback.NewSaveBuybackCreated().WithPayload(&models.BaseResponse{
			Code:    http.StatusCreated,
			Message: fmt.Sprintf("success buyback gold with account no %s", sbp.Data.AccountNumber),
		})
	})

	// balance
	api.BalanceGetBalanceHandler = balance.GetBalanceHandlerFunc(func(gbp balance.GetBalanceParams) middleware.Responder {
		data, err := handlers.GetAccount(cfg, gbp.AccountNo)
		if err != nil {
			return balance.NewGetBalanceDefault(http.StatusInternalServerError).WithPayload(
				&models.BaseResponse{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				})
		}

		return balance.NewGetBalanceOK().WithPayload(&balance.GetBalanceOKBody{
			Data: data,
		})
	})

	// balance
	api.TransactionGetMutationHandler = transaction.GetMutationHandlerFunc(func(gmp transaction.GetMutationParams) middleware.Responder {
		data, err := handlers.GetMutation(cfg, gmp.AccountNo, gmp.StartDate, gmp.EndDate)
		if err != nil {
			return transaction.NewGetMutationDefault(http.StatusInternalServerError).WithPayload(
				&models.BaseResponse{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				})
		}

		return transaction.NewGetMutationOK().WithPayload(&transaction.GetMutationOKBody{
			Data: data,
		})
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
