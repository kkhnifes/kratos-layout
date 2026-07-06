package data

import (
	"github.com/kkhnifes/kratos-layout/internal/conf"

	"github.com/go-kratos/kratos/v3/log"
)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data) (*Data, func(), error) {
	cleanup := func() {
		log.Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
