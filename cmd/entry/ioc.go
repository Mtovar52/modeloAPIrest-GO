package entry

import (
	"api-gateway/internal/domain/usecase"
	"api-gateway/internal/infra/cache"
	"api-gateway/internal/infra/jwt"
)

func genIoc() usecase.UserUseCase {
	clientJwt := jwt.NewJwtClient()
	cacheProvider := cache.GetCacheProvider()
	cacheProvider.Set("chupi", "plum")
	return usecase.NewUserUserCase(clientJwt, cacheProvider)
}
