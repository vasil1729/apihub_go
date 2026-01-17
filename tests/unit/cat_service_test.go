package public_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/ultimatum/apihub_go/internal/service/public"
)

func TestCatService_GetAll(t *testing.T) {
	service, err := public.NewCatService("../../data")
	assert.NoError(t, err)

	cats, total, err := service.GetAll(1, 10)
	assert.NoError(t, err)
	assert.Greater(t, total, 0)
	assert.LessOrEqual(t, len(cats), 10)
}

func TestCatService_GetByID(t *testing.T) {
	service, err := public.NewCatService("../../data")
	assert.NoError(t, err)

	cat, err := service.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, cat.ID)
}

func TestCatService_GetRandom(t *testing.T) {
	service, err := public.NewCatService("../../data")
	assert.NoError(t, err)

	cat, err := service.GetRandom()
	assert.NoError(t, err)
	assert.Greater(t, cat.ID, 0)
}
