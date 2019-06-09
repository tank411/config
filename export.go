package config

import (
	"errors"
	"fmt"
	"io"

	"github.com/imdario/mergo"
)

// MapStruct alias method of the 'Structure'
func MapStruct(key string, v interface{}) error { return dc.Structure(key, v) }

// MapStruct alias method of the 'Structure'
func (c *Config) MapStruct(key string, v interface{}) error {
	return c.Structure(key, v)
}

// MapTo alias method of the 'Structure'
func MapTo(key string, v interface{}) error { return dc.Structure(key, v) }

// MapTo alias method of the 'Structure'
func (c *Config) MapTo(key string, v interface{}) error {
	return c.Structure(key, v)
}

// ToStruct alias method of the 'Structure'
func ToStruct(key string, v interface{}) error { return dc.Structure(key, v) }

// ToStruct alias method of the 'Structure'
func (c *Config) ToStruct(key string, v interface{}) error {
	return c.Structure(key, v)
}

// Structure get config data and map to a structure.
// Usage:
// 	dbInfo := Db{}
// 	config.Structure("db", &dbInfo)
func (c *Config) Structure(key string, dst interface{}, driverName string) (err error) {
	var ok bool
	var data interface{}

	// map all data
	if key == "" {
		ok = true
		data = c.data
	} else {
		data, ok = c.GetValue(key)
	}

	if !ok {
		return
	}

	decoder := c.getDecoderByFormat(driverName)
	if decoder == nil {
		return mergo.Map(dst, data)
	}

	return
}

// WriteTo a writer
func WriteTo(out io.Writer) (int64, error) { return dc.WriteTo(out) }

// WriteTo Write out config data representing the current state to a writer.
func (c *Config) WriteTo(out io.Writer) (n int64, err error) {
	return c.DumpTo(out, c.opts.DumpFormat)
}

// DumpTo a writer and use format
func DumpTo(out io.Writer, format string) (int64, error) { return dc.DumpTo(out, format) }

// DumpTo use the format(json,yaml,toml) dump config data to a writer
func (c *Config) DumpTo(out io.Writer, format string) (n int64, err error) {
	var ok bool
	var encoder Encoder

	format = fixFormat(format)
	if encoder, ok = c.encoders[format]; !ok {
		err = errors.New("no exists or no register encoder for the format: " + format)
		return
	}

	// is empty
	if len(c.data) == 0 {
		return
	}

	// encode data to string
	encoded, err := encoder(&c.data)
	if err != nil {
		return
	}

	// write content to out
	num, err := fmt.Fprintln(out, string(encoded))
	if err != nil {
		return
	}

	return int64(num), nil
}
