package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_MapStructure(t *testing.T) {
	st := assert.New(t)

	cfg := Default()
	cfg.ClearAll()

	err := cfg.LoadStrings(JSON, `{
"age": 28,
"name": "inhere",
"sports": ["pingPong", "跑步"]
}`)

	st.Nil(err)

	user := &struct {
		Age    float64 // always float64 from JSON
		Name   string
		Sports []string
	}{}
	// map all
	err = MapStruct("", user)
	st.Nil(err)

	st.Equal(28, user.Age)
	st.Equal("inhere", user.Name)
	st.Equal("pingPong", user.Sports[0])

	// map some
	err = cfg.LoadStrings(JSON, `{
"sec": {
	"key": "val",
	"age": 120,
	"tags": [12, 34]
}
}`)
	st.Nil(err)

	some := struct {
		Age  int
		Kye  string
		Tags []int
	}{}
	err = cfg.ToStruct("sec", &some)
	st.Nil(err)
	st.Equal(120, some.Age)
	st.Equal(12, some.Tags[0])

	cfg.ClearAll()

	// custom data
	cfg = New("test")
	err = cfg.LoadData(map[interface{}]interface{}{
		"key": "val",
		"age": 120,
		"tags": []int{12, 34},
	})
	st.NoError(err)

	s1 := struct {
		Age  int
		Kye  string
		Tags []int
	}{}
	err = cfg.MapTo("", &s1)
	st.Nil(err)
	st.Equal(120, s1.Age)
	st.Equal(12, s1.Tags[0])

	cfg.ClearAll()
}
