package goshred

import (
	"os"
	"math/rand"
)



// writeRandomData overwrites the file's content with random bytes.
// If path is a directory or does not exist, returns the corresponding error.
// If the user does not own the file, returns an error.
func writeRandomData(path string) error {

	finfo, err := os.Stat(path)

	if err != nil {
		return err
	}

	size := finfo.Size()

	newdata := make([]byte, size)
	rand.Read(newdata)

	err = os.Chmod(path, 0600)

	if err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_WRONLY, 0700)

	if err != nil {
		return err
	}

	_, err = f.WriteAt(newdata, 0)

	if err != nil {
		return err
	}

	err =f.Close()

	if err != nil {
		return err
	}

	return nil
}

// Shred overwrites a file's contents with random data three times, and deletes the file.
// If path does not exist, or is a directory, returns the corresponding error.
// If path is a link, it will overwrite the file's contents, and delete the link, but not the target.
// If the user does not own the file, returns an error.
func Shred(path string) error {

	var err error

	for i := 0; i < 3; i++ {

		err = writeRandomData(path)

		if err != nil {
			return err
		}
	}

	err = os.Remove(path)

	if err != nil {
		return err
	}

	return nil
}
