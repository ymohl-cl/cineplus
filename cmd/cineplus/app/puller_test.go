package app

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/ymohl-cl/cineplus/pkg/ghibli"
	mock_ghibli "github.com/ymohl-cl/cineplus/pkg/ghibli/mock"
)

func TestPuller_Start(t *testing.T) {
	t.Run("Should not update the cache because error occured from the movies getter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cache := &Cache{}
		m := mock_ghibli.NewMockClient(ctrl)
		p := &Puller{
			ghibli:    m,
			tickTimer: 1,
			cache:     cache,
		}
		m.EXPECT().Movies().Return(nil, errors.New("error test")).AnyTimes()
		go p.Start()
		defer p.Close()
		time.Sleep(10 * time.Millisecond)
		cache.Lock()
		defer cache.Unlock()
		assert.Nil(t, cache.movies)
		assert.Nil(t, cache.peoples)
	})
	t.Run("Should not update the cache because error occured from the people getter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cache := &Cache{}
		m := mock_ghibli.NewMockClient(ctrl)
		p := &Puller{
			ghibli:    m,
			tickTimer: 1,
			cache:     cache,
		}
		m.EXPECT().Movies().Return([]ghibli.Movie{ghibli.Movie{ID: "m1"}}, nil).AnyTimes()
		m.EXPECT().Peoples().Return(nil, errors.New("error test")).AnyTimes()
		go p.Start()
		defer p.Close()
		time.Sleep(10 * time.Millisecond)
		cache.Lock()
		defer cache.Unlock()
		assert.Nil(t, cache.movies)
		assert.Nil(t, cache.peoples)
	})
	t.Run("Should be ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cache := &Cache{}
		m := mock_ghibli.NewMockClient(ctrl)
		p := &Puller{
			ghibli:    m,
			tickTimer: 1,
			cache:     cache,
		}
		m.EXPECT().Movies().Return([]ghibli.Movie{ghibli.Movie{ID: "m1"}}, nil).AnyTimes()
		m.EXPECT().Peoples().Return([]ghibli.People{ghibli.People{ID: "p1"}}, nil).AnyTimes()
		go p.Start()
		defer p.Close()
		time.Sleep(10 * time.Millisecond)
		cache.Lock()
		defer cache.Unlock()
		if assert.NotNil(t, cache.movies) {
			assert.EqualValues(t, []ghibli.Movie{ghibli.Movie{ID: "m1"}}, *cache.movies)
		}
		if assert.NotNil(t, cache.peoples) {
			assert.EqualValues(t, []ghibli.People{ghibli.People{ID: "p1"}}, *cache.peoples)
		}
	})
}
