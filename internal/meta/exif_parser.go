package meta

import (
	"fmt"
	"path/filepath"
	"runtime/debug"
	"strings"

	"github.com/dsoprea/go-exif/v3"
	heicexif "github.com/dsoprea/go-heic-exif-extractor"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure"
	pngstructure "github.com/dsoprea/go-png-image-structure"
	tiffstructure "github.com/dsoprea/go-tiff-image-structure"
	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/photoprism/photoprism/pkg/txt"
)

func RawExif(fileName string, fileType fs.FileType) (rawExif []byte, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("metadata: %s in %s (raw exif panic)\nstack: %s", e, txt.Quote(filepath.Base(fileName)), debug.Stack())
		}
	}()

	// Extract raw EXIF block.
	var parsed bool

	logName := txt.Quote(filepath.Base(fileName))

	if fileType == fs.TypeJpeg {
		jpegMp := jpegstructure.NewJpegMediaParser()

		sl, err := jpegMp.ParseFile(fileName)

		if err != nil {
			log.Warnf("metadata: %s in %s (parse jpeg)", err, logName)
		} else {
			_, rawExif, err = sl.Exif()

			if err != nil {
				if strings.HasPrefix(err.Error(), "no exif header") {
					return rawExif, fmt.Errorf("metadata: no exif header in %s (parse jpeg)", logName)
				} else if strings.HasPrefix(err.Error(), "no exif data") {
					log.Debugf("metadata: failed parsing %s, starting brute-force search (parse jpeg)", logName)
				} else {
					log.Warnf("metadata: %s in %s, starting brute-force search (parse jpeg)", err, logName)
				}
			} else {
				parsed = true
			}
		}
	} else if fileType == fs.TypePng {
		pngMp := pngstructure.NewPngMediaParser()

		cs, err := pngMp.ParseFile(fileName)

		if err != nil {
			return rawExif, fmt.Errorf("metadata: %s in %s (parse png)", err, logName)
		}

		_, rawExif, err = cs.Exif()

		if err != nil {
			if err.Error() == "file does not have EXIF" {
				return rawExif, fmt.Errorf("metadata: no exif header in %s (parse png)", logName)
			} else {
				log.Warnf("metadata: %s in %s (parse png)", err, logName)
			}
		} else {
			parsed = true
		}
	} else if fileType == fs.TypeHEIF {
		heicMp := heicexif.NewHeicExifMediaParser()

		cs, err := heicMp.ParseFile(fileName)

		if err != nil {
			return rawExif, fmt.Errorf("metadata: %s in %s (parse heic)", err, logName)
		}

		_, rawExif, err = cs.Exif()

		if err != nil {
			if err.Error() == "file does not have EXIF" {
				return rawExif, fmt.Errorf("metadata: no exif header in %s (parse heic)", logName)
			} else {
				log.Warnf("metadata: %s in %s (parse heic)", err, logName)
			}
		} else {
			parsed = true
		}
	} else if fileType == fs.TypeTiff {
		tiffMp := tiffstructure.NewTiffMediaParser()

		cs, err := tiffMp.ParseFile(fileName)

		if err != nil {
			return rawExif, fmt.Errorf("metadata: %s in %s (parse tiff)", err, logName)
		}

		_, rawExif, err = cs.Exif()

		if err != nil {
			if err.Error() == "file does not have EXIF" {
				return rawExif, fmt.Errorf("metadata: no exif header in %s (parse tiff)", logName)
			} else {
				log.Warnf("metadata: %s in %s (parse tiff)", err, logName)
			}
		} else {
			parsed = true
		}
	}

	if !parsed {
		rawExif, err = exif.SearchFileAndExtractExif(fileName)

		if err != nil {
			return rawExif, fmt.Errorf("metadata: no exif header in %s (search and extract)", logName)
		}
	}

	return rawExif, nil
}
