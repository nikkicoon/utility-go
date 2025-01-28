package pkg

import (
	"compress/gzip"
	"fmt"
	"github.com/phuslu/log"
	"io"
	"os"
	"path/filepath"
)

// GzipFile compresses the file `in` to `in`.gz
func GzipFile(logger *log.Logger, in string) {
	// Open the original file
	originalFile, err := os.Open(in)
	if err != nil {
		logger.Error().Err(err)
		return
	}
	defer originalFile.Close()

	// Create a new gzipped file
	gzippedFile, err := os.Create(in + ".gz")
	if err != nil {
		logger.Error().Err(err)
		return
	}
	defer gzippedFile.Close()

	// Create a new gzip writer
	gzipWriter := gzip.NewWriter(gzippedFile)
	defer gzipWriter.Close()

	// Copy the contents of the original file to the gzip writer
	_, err = io.Copy(gzipWriter, originalFile)
	if err != nil {
		logger.Error().Err(err)
		return
	}

	// Flush the gzip writer to ensure all data is written
	gzipWriter.Flush()

	// Delete original file
	err = os.Remove(in)
	if err != nil {
		logger.Error().Err(err)
		return
	}
}

// GzipAll compresses all the files found in `directory`.
func GzipAll(logger *log.Logger, directory string) {
	files, err := os.ReadDir(directory)
	path, _ := filepath.Abs(directory)
	if err != nil {
		switch logger {
		case nil:
			_, _ = fmt.Fprintf(os.Stderr, "%q\n", err)
		default:
			logger.Error().Err(err)
		}
	}
	for _, f := range files {
		GzipFile(logger, filepath.Join(path, f.Name()))
	}
}
