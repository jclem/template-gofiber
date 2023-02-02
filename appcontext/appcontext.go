package appcontext

import (
	"context"
	"time"

	"github.com/jclem/template-gofiber/config"
	"go.uber.org/zap"
)

// Context provides access to application-level config, context, logging.
type Context struct {
	Config  *config.Config
	Context context.Context
	Logger  *zap.SugaredLogger
}

// Deadline implements the context.Context interface.
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.Context.Deadline()
}

// Done implements the context.Context interface.
func (c *Context) Done() <-chan struct{} {
	return c.Context.Done()
}

// Err implements the context.Context interface.
func (c *Context) Err() error {
	return c.Context.Err()
}

// Value implements the context.Context interface.
func (c *Context) Value(key any) any {
	return c.Context.Value(key)
}

// Opt is a function that configures a Context.
type Opt func(*Context) error

// WithOptions returns a new Context with the provided options merged.
func (c *Context) WithOptions(opts ...Opt) (*Context, error) {
	newc := &Context{
		Config:  c.Config,
		Context: c.Context,
		Logger:  c.Logger,
	}

	for _, opt := range opts {
		if err := opt(newc); err != nil {
			return nil, err
		}
	}

	return newc, nil
}

// New returns a new Context with the provided options applied.
func NewWithOptions(opts ...Opt) (*Context, error) {
	c := &Context{}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// WithConfig returns an Opt that sets the Config.
func WithConfig(cfg *config.Config) Opt {
	return func(c *Context) error {
		c.Config = cfg
		return nil
	}
}

// WithContext returns an Opt that sets the Context.
func WithContext(ctx context.Context) Opt {
	return func(c *Context) error {
		c.Context = ctx
		return nil
	}
}

// WithLogger returns an Opt that sets the Logger.
func WithLogger(logger *zap.SugaredLogger) Opt {
	return func(c *Context) error {
		c.Logger = logger
		return nil
	}
}
