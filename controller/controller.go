package controller

import (
	"reflect"
	"sync"

	"shop-test/cmd/config"
	"shop-test/errRes"
	"shop-test/model"
	"shop-test/pkg/log"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type (
	Controller struct {
		config    *config.Config
		logger    log.ILogger
		debug     bool
		validator *validator.Validate
		sync.Mutex
	}
)


func NewController(config *config.Config, logger log.ILogger, debug bool) *Controller {
	validator := validator.New()

	validator.RegisterValidation("ids", ValidateIds)
	ctrl := Controller{
		config:    config,
		logger:    logger,
		debug:     debug,
		validator: validator,
	}
	return &ctrl
}

func (c *Controller) Load() error {
	c.Lock()
	defer c.Unlock()

	// do the magic during loads

	return nil
}

// Start is non-blocking
func (c *Controller) Start() {
	c.Lock()
	defer c.Unlock()

	c.start()
}

func (c *Controller) start() {
	// call any additionl handler -> must be non-blocking
}

// Stop is non-blocking
func (c *Controller) Stop() {
	c.Lock()
	defer c.Unlock()

	c.stop()
}

func (c *Controller) stop() {
	// shutdown all additional handler -> must be non-blocking
}

func (c *Controller) Config() *config.Config {
	return c.config
}

func (c *Controller) Logger() log.ILogger {
	return c.logger
}

func (c *Controller) NewPagingRequest(page, pageSize int) *model.PagingRequest {
	return &model.PagingRequest{
		Page:     page,
		PageSize: pageSize,
	}
}

func (c *Controller) StructToMap(obj interface{}) map[string]float32 {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() != reflect.Struct {
		panic("Input is not a struct")
	}

	result := make(map[string]float32)

	for i := 0; i < objValue.NumField(); i++ {
		field := objValue.Field(i)
		fieldName := objValue.Type().Field(i).Tag.Get("bson")
		if fieldName == "" {
			fieldName = objValue.Type().Field(i).Tag.Get("json")
		}
		if fieldName == "" {
			fieldName = objValue.Type().Field(i).Name
		}
		if field.Kind() == reflect.Float32 {
			result[fieldName] = float32(field.Float())
		} else {
			panic("Field is not of type float32")
		}
	}

	return result
}

func (c *Controller) Validate(i interface{}) error {
	if err := c.validator.Struct(i); err != nil {
		return errRes.ErrValidatingData
	}
	return nil
}
func ValidateIds(fl validator.FieldLevel) bool {
	for _, value := range fl.Field().Interface().([]string) {
		_, err := uuid.Parse(value)
		if err != nil {
			return false
		}
	}
	return true
}
