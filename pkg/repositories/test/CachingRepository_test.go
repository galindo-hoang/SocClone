package repositories_test

import (
	_ "embed"
	"fmt"
	"github.com/AuthService/pkg/database"
	"github.com/AuthService/pkg/repositories"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var (
	//go:embed env.md
	environment string

	caching *repositories.CachingRepository = nil
)

func TestAddCache(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		key    string
		value  string
		expect string
	}{
		{
			key:    "k1",
			value:  "v1",
			expect: "v1",
		},
		{
			key:    "k2",
			value:  "v2",
			expect: "v2",
		},
		{
			key:    "k1",
			value:  "v3",
			expect: "v3",
		},
	}
	t.Cleanup(func() {
		for _, testCase := range testCases {
			err := caching.DeleteCache("path", testCase.key)
			if err != nil {
				t.Error(err)
			}
		}
	})

	for _, testCase := range testCases {
		if err := caching.AddCache("path", testCase.key, testCase.value, 10); err != nil {
			t.Error(err)
		}
		var item, ok = caching.GetCache("path", testCase.key)
		if !ok {
			t.Fail()
		}
		assert.Equal(t, testCase.expect, string(item), fmt.Sprintf("Expect: %v || actual: %v", testCase.expect, string(item)))
	}
}

func TestMain(m *testing.M) {
	if err := readEnvironment(); err != nil {
		os.Exit(0)
	}

	database.InitCache()
	caching = &repositories.CachingRepository{}
	os.Exit(m.Run())
}

func readEnvironment() error {
	var fields = strings.Split(environment, "\n")
	for _, field := range fields {
		if strings.Contains(field, "=") {
			var lines = strings.Split(field, "=")
			err := os.Setenv(lines[0], lines[1])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
