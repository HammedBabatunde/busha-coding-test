package movie

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/emekarr/coding-test-busha/logger"
	"github.com/emekarr/coding-test-busha/network"
	redisRepo "github.com/emekarr/coding-test-busha/repository/redis"
	"go.uber.org/zap"
)

var MovieService = movieService{network: network.NetworkController{BaseUrl: "https://swapi.dev/api/films"}}

type movieService struct {
	network network.NetworkController
}

func (ms *movieService) FetchMovies() (*[]Movie, error) {
	var movies *[]Movie
	movies, err := FectchAllCachedMovies()
	if err == nil && len(*movies) > 0 {
		return movies, nil
	}
	response := ms.network.Get("/", nil, nil)
	if response.Error != nil {
		logger.Error(errors.New("error fetching movies from movie service"), zap.Error(response.Error))
		return nil, response.Error
	}
	var parsedResponse MovieResponse
	json.Unmarshal(*response.Body, &parsedResponse)
	CacheMovies(&parsedResponse.Results)
	return &parsedResponse.Results, nil
}

func (ms *movieService) SearchMovies(term string) (*Movie, error) {
	response := ms.network.Get(fmt.Sprintf("/%s", term), nil, nil)
	if response.Error != nil {
		logger.Error(errors.New("error searching movies from movie service"), zap.Error(response.Error))
		return nil, response.Error
	}
	var parsedResponse Movie
	json.Unmarshal(*response.Body, &parsedResponse)
	return &parsedResponse, nil
}

var cacheKey = "movies-cache"

func FectchAllCachedMovies() (*[]Movie, error) {
	result, err := redisRepo.RedisRepo.FindSet(cacheKey)
	if err != nil {
		logger.Error(errors.New("could not fetch movie cache"), zap.Error(err))
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
	unmarshaledResult := []Movie{}
	movieChan := make(chan Movie, len(*result))

	// unmarshal results using a loop because the json package cannot unmarshal a string array
	for _, m := range *result {
		go func(mCached string, mChan chan Movie) {
			var unmarshaledMovie Movie
			err = json.Unmarshal([]byte(mCached), &unmarshaledMovie)
			if err != nil {
				mChan <- Movie{}
				logger.Error(errors.New("could not unmarshal cached movie"), zap.Error(err), zap.Any("cached_movie", result))
				return
			}
			mChan <- unmarshaledMovie
		}(m, movieChan)
		unmarshaledResult = append(unmarshaledResult, <-movieChan)
	}
	return &unmarshaledResult, nil
}

// this cache last for 12 hours before being expired by redis
func CacheMovies(payload *[]Movie) (bool, error) {
	var wg sync.WaitGroup

	for i, m := range *payload {
		wg.Add(1)
		go func(m Movie, i int) {
			defer func() {
				wg.Done()
			}()
			marshaledPayload, err := json.Marshal(m)
			if err != nil {
				logger.Error(errors.New("could not marshal movie payload"), zap.Error(err), zap.Any("payload", *payload))
				return
			}
			if err != nil {
				logger.Error(errors.New("could not process movie score for redis"), zap.Error(err))
				return
			}
			// cache for 12 hours
			success, err := redisRepo.RedisRepo.CreateInSet(cacheKey, float64(i), marshaledPayload)
			if err != nil || !success {
				logger.Error(errors.New("could not cache movies"), zap.Error(err))
				return
			}
		}(m, i)
	}
	wg.Wait()
	return true, nil
}
