package basename

import "testing"

func TestBasename(t *testing.T) {
	expected := []struct {
		Path     string
		Basename string
	}{
		{"/home/justin/.config/init.vim", "init"},
		{"/home/justin/.config/.init.vim.swp", ".init.vim"},
	}

	for i, test := range expected {
		if act := basename(test.Path); act != test.Basename {
			t.Errorf("Test case #%d failed due to expected: basename(%s) == %s, actual: %s",
				i+1, test.Path, test.Basename, act)
		}
	}

}
