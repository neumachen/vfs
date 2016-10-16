package vfs

import "testing"

func TestLoad(t *testing.T) {
	tests := []struct {
		desc    string
		file    string
		matcher func(e error) bool
	}{
		{
			desc:    "directory does not exist",
			file:    "/manbearpig/",
			matcher: func(e error) bool { return e != nil },
		},
		{
			desc:    "file not found",
			file:    "nomanbearpig.json",
			matcher: func(e error) bool { return e != nil },
		},
		{
			desc:    "invalid file type",
			file:    "../../tests/fixtures/filestore/invalid_files/manbearpig.mbp",
			matcher: func(e error) bool { return e != nil },
		},
		{
			desc:    "directory exists with valid file types",
			file:    "../../tests/fixtures/filestore/valid_files/",
			matcher: func(e error) bool { return e == nil },
		},
		{
			desc:    "directory exists with invalid file types",
			file:    "../../tests/fixtures/filestore/invalid_files",
			matcher: func(e error) bool { return e != nil },
		},
		{
			desc:    "file exists and is a valid file type",
			file:    "../../tests/fixtures/filestore/valid_files/test.json",
			matcher: func(e error) bool { return e == nil },
		},
	}

	for _, test := range tests {
		_, err := LoadFiles(test.file)
		if !test.matcher(err) {
			t.Errorf("Test failed, desc: %v, returned %v. matched %t.", test.desc, err, test.matcher(err))
		}
	}
}

func TestLookupJSONSchema(t *testing.T) {
	files := map[string]string{"manbearpig-schema": "manbearpig"}

	s := &store{files}
	tests := []struct {
		desc    string
		schema  string
		matcher func(e error) bool
	}{
		{
			desc:    "schema not found",
			schema:  "manbearpiggly-schema",
			matcher: func(e error) bool { return e != nil },
		},
		{
			desc:    "schema found",
			schema:  "manbearpig-schema",
			matcher: func(e error) bool { return e == nil },
		},
	}

	for _, test := range tests {
		_, err := s.Get(test.schema)

		if !test.matcher(err) {
			t.Errorf("Test failed, desc: %v, returned %v. matched %t.", test.desc, err, test.matcher(err))
		}
	}

}
