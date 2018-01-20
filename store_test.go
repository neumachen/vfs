package vfs

import "testing"

func TestLoadFolders(t *testing.T) {
	tests := []struct {
		desc    string
		folders []string
		matcher func(e error) bool
	}{
		{
			desc:    "directory does not exist",
			folders: []string{"/manbearpig/", "/hoho/dodo"},
			matcher: func(e error) bool { return e != nil },
		},
		{
			desc: "directories exists",
			folders: []string{
				"tests/fixtures/folder1/",
				"tests/fixtures/folder2/",
			},
			matcher: func(e error) bool { return e == nil },
		},
	}

	for _, test := range tests {
		_, err := LoadFolders(test.folders...)
		if !test.matcher(err) {
			t.Errorf("Test failed, desc: %v, returned %v. matched %t.", test.desc, err, test.matcher(err))
		}
	}
}

func TestLoadFolder(t *testing.T) {
	tests := []struct {
		desc    string
		folder  string
		matcher func(e error) bool
	}{
		{
			desc:    "directory does not exist",
			folder:  "/manbearpig/",
			matcher: func(e error) bool { return e != nil },
		},
		{
			desc:    "folder exists",
			folder:  "tests/fixtures/folder1",
			matcher: func(e error) bool { return e == nil },
		},
	}

	for _, test := range tests {
		_, err := LoadFolder(test.folder)
		if !test.matcher(err) {
			t.Errorf("Test failed, desc: %v, returned %v. matched %t.", test.desc, err, test.matcher(err))
		}
	}
}
