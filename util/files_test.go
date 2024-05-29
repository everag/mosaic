package util

import "testing"

func TestGetNewImageFilename(t *testing.T) {
	// Empty filename
	_, err := GetNewImageFilename("", "token")
	if err == nil {
		t.Errorf("error expected for empty or spaces-only srcFilename")
	}

	// Filename with spaces
	_, err = GetNewImageFilename("   ", "token")
	if err == nil {
		t.Errorf("error expected for empty or spaces-only srcFilename")
	}

	// Unix format + Missing suffix
	actual, _ := GetNewImageFilename("/tmp/test.png", "")
	expected := "/tmp/test_new.png"
	if actual != expected {
		t.Errorf("expected: %s, actual: %s, for empty suffix", expected, actual)
	}

	// Windows file format
	actual, _ = GetNewImageFilename(`C:\Temp\test.png`, "")
	expected = `C:\Temp\test_new.png`
	if actual != expected {
		t.Errorf("expected: %s, actual: %s, for windows file name", expected, actual)
	}

	// Some token
	actual, _ = GetNewImageFilename("/tmp/test.png", "100x100")
	expected = "/tmp/test_100x100.png"
	if actual != expected {
		t.Errorf("expected: %s, actual: %s, for %% suffix", expected, actual)
	}
}
