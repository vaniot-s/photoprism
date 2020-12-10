package meta

import (
	"testing"

	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/stretchr/testify/assert"
)

func TestExif(t *testing.T) {
	t.Run("photoshop.jpg", func(t *testing.T) {
		data, err := Exif("testdata/photoshop.jpg", fs.TypeJpeg)

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("all: %+v", data.All)

		assert.Equal(t, "Michael Mayer", data.Artist)
		assert.Equal(t, "2020-01-01T16:28:23Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2020-01-01T17:28:23Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "Example file for development", data.Description)
		assert.Equal(t, "This is a legal notice", data.Copyright)
		assert.Equal(t, 540, data.Height)
		assert.Equal(t, 720, data.Width)
		assert.Equal(t, float32(52.45969), data.Lat)
		assert.Equal(t, float32(13.321832), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, "1/50", data.Exposure)
		assert.Equal(t, "HUAWEI", data.CameraMake)
		assert.Equal(t, "ELE-L29", data.CameraModel)
		assert.Equal(t, "", data.CameraOwner)
		assert.Equal(t, "", data.CameraSerial)
		assert.Equal(t, 27, data.FocalLength)
		assert.Equal(t, 1, int(data.Orientation))

		// TODO: Values are empty - why?
		// assert.Equal(t, "HUAWEI P30 Rear Main Camera", data.LensModel)
	})

	t.Run("ladybug.jpg", func(t *testing.T) {
		data, err := Exif("testdata/ladybug.jpg", fs.TypeJpeg)

		if err != nil {
			t.Fatal(err)
		}

		//  t.Logf("all: %+v", data.All)

		assert.Equal(t, "Photographer: TMB", data.Artist)
		assert.Equal(t, "2011-07-10T17:34:28Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2011-07-10T19:34:28Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "", data.Title)    // Should be "Ladybug"
		assert.Equal(t, "", data.Keywords) // Should be "Ladybug"
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "", data.Copyright)
		assert.Equal(t, 540, data.Height)
		assert.Equal(t, 720, data.Width)
		assert.Equal(t, float32(51.254852), data.Lat)
		assert.Equal(t, float32(7.389468), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, "1/125", data.Exposure)
		assert.Equal(t, "Canon", data.CameraMake)
		assert.Equal(t, "Canon EOS 50D", data.CameraModel)
		assert.Equal(t, "Thomas Meyer-Boudnik", data.CameraOwner)
		assert.Equal(t, "2260716910", data.CameraSerial)
		assert.Equal(t, "", data.LensMake)
		assert.Equal(t, "EF100mm f/2.8 Macro USM", data.LensModel)
		assert.Equal(t, 100, data.FocalLength)
		assert.Equal(t, 1, int(data.Orientation))
	})

	t.Run("gopro_hd2.jpg", func(t *testing.T) {
		data, err := Exif("testdata/gopro_hd2.jpg", fs.TypeJpeg)

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("all: %+v", data.All)

		assert.Equal(t, "", data.Artist)
		assert.Equal(t, "2017-12-21T05:17:28Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2017-12-21T05:17:28Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "", data.Title)
		assert.Equal(t, "", data.Keywords)
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "", data.Copyright)
		assert.Equal(t, 180, data.Height)
		assert.Equal(t, 240, data.Width)
		assert.Equal(t, float32(0), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, "1/2462", data.Exposure)
		assert.Equal(t, "GoPro", data.CameraMake)
		assert.Equal(t, "HD2", data.CameraModel)
		assert.Equal(t, "", data.CameraOwner)
		assert.Equal(t, "", data.CameraSerial)
		assert.Equal(t, 16, data.FocalLength)
		assert.Equal(t, 1, int(data.Orientation))
	})

	t.Run("tweethog.png", func(t *testing.T) {
		_, err := Exif("testdata/tweethog.png", fs.TypePng)

		if err == nil {
			t.Fatal("err should NOT be nil")
		}

		assert.Equal(t, "metadata: no exif header in tweethog.png (parse png)", err.Error())
	})

	t.Run("iphone_7.heic", func(t *testing.T) {
		data, err := Exif("testdata/iphone_7.heic", fs.TypeHEIF)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "2018-09-10T03:16:13Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2018-09-10T12:16:13Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, float32(34.79745), data.Lat)
		assert.Equal(t, float32(134.76463), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, "1/4000", data.Exposure)
		assert.Equal(t, "Apple", data.CameraMake)
		assert.Equal(t, "iPhone 7", data.CameraModel)
		assert.Equal(t, 74, data.FocalLength)
		assert.Equal(t, 6, data.Orientation)
		assert.Equal(t, "Apple", data.LensMake)
		assert.Equal(t, "iPhone 7 back camera 3.99mm f/1.8", data.LensModel)

	})

	t.Run("gps-2000.jpg", func(t *testing.T) {
		data, err := Exif("testdata/gps-2000.jpg", fs.TypeJpeg)

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("GPS 2000: %+v", data.All)

		assert.Equal(t, "", data.Artist)
		assert.True(t, data.TakenAt.IsZero())
		assert.True(t, data.TakenAtLocal.IsZero())
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "", data.Copyright)
		assert.Equal(t, 0, data.Height)
		assert.Equal(t, 0, data.Width)
		assert.Equal(t, float32(-38.405193), data.Lat)
		assert.Equal(t, float32(144.18896), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, "", data.Exposure)
		assert.Equal(t, "", data.CameraMake)
		assert.Equal(t, "", data.CameraModel)
		assert.Equal(t, "", data.CameraOwner)
		assert.Equal(t, "", data.CameraSerial)
		assert.Equal(t, 0, data.FocalLength)
		assert.Equal(t, 1, int(data.Orientation))
	})

	t.Run("image-2011.jpg", func(t *testing.T) {
		data, err := Exif("testdata/image-2011.jpg", fs.TypeJpeg)

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("ALL: %+v", data.All)

		/*
		  Exiftool date information:

		  File Modification Date/Time     : 2020:05:15 08:25:46+00:00
		  File Access Date/Time           : 2020:05:15 08:25:47+00:00
		  File Inode Change Date/Time     : 2020:05:15 08:25:46+00:00
		  Modify Date                     : 2020:05:15 10:25:45
		  Create Date                     : 2011:07:19 11:36:38
		  Metadata Date                   : 2020:05:15 10:25:45+02:00

		*/

		// assert.Equal(t, "2011-07-19T11:36:38Z", data.TakenAt.Format("2006-01-02T15:04:05Z")) // TODO
		// assert.Equal(t, "2011-07-19T11:36:38Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))  // TODO
		assert.Equal(t, float32(0), data.Lat)
		assert.Equal(t, float32(0), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, "1/1100", data.Exposure)
		assert.Equal(t, "SAMSUNG", data.CameraMake)
		assert.Equal(t, "GT-I9000", data.CameraModel)
		assert.Equal(t, 3, data.FocalLength)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, "", data.LensMake)
		assert.Equal(t, "", data.LensModel)
	})

	t.Run("ship.jpg", func(t *testing.T) {
		data, err := Exif("testdata/ship.jpg", fs.TypeJpeg)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "2019-05-12T15:13:53Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2019-05-12T17:13:53Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, float32(53.12349), data.Lat)
		assert.Equal(t, float32(18.00152), data.Lng)
		assert.Equal(t, 63, data.Altitude)
		assert.Equal(t, "1/100", data.Exposure)
		assert.Equal(t, "Xiaomi", data.CameraMake)
		assert.Equal(t, "Mi A1", data.CameraModel)
		assert.Equal(t, 52, data.FocalLength)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, "", data.LensMake)
		assert.Equal(t, "", data.LensModel)
	})

	t.Run("no-exif-data.jpg", func(t *testing.T) {
		_, err := Exif("testdata/no-exif-data.jpg", fs.TypeJpeg)

		if err == nil {
			t.Fatal("err should NOT be nil")
		}

		assert.Equal(t, "metadata: no exif header in no-exif-data.jpg (search and extract)", err.Error())
	})

	t.Run("screenshot.png", func(t *testing.T) {
		data, err := Exif("testdata/screenshot.png", fs.TypePng)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "721", data.All["PixelXDimension"])
		assert.Equal(t, "332", data.All["PixelYDimension"])
	})

	t.Run("orientation.jpg", func(t *testing.T) {
		data, err := Exif("testdata/orientation.jpg", fs.TypeJpeg)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "3264", data.All["PixelXDimension"])
		assert.Equal(t, "1836", data.All["PixelYDimension"])
		assert.Equal(t, 3264, data.Width)
		assert.Equal(t, 1836, data.Height)
		assert.Equal(t, 6, data.Orientation) // TODO: Should be 1

		if err := data.JSON("testdata/orientation.json", "orientation.jpg"); err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 3264, data.Width)
		assert.Equal(t, 1836, data.Height)
		assert.Equal(t, 1, data.Orientation)

		if err := data.JSON("testdata/orientation.json", "foo.jpg"); err != nil {
			assert.EqualError(t, err, "metadata: original name foo.jpg does not match orientation.jpg (exiftool)")
		} else {
			t.Error("error expected when providing wrong original name")
		}
	})

	t.Run("gopher-preview.jpg", func(t *testing.T) {
		_, err := Exif("testdata/gopher-preview.jpg", fs.TypeJpeg)

		assert.EqualError(t, err, "metadata: no exif header in gopher-preview.jpg (search and extract)")
	})

	t.Run("huawei-gps-error.jpg", func(t *testing.T) {
		data, err := Exif("testdata/huawei-gps-error.jpg", fs.TypeJpeg)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "2020-06-16T16:52:46Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2020-06-16T18:52:46Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, float32(48.302776), data.Lat)
		assert.Equal(t, float32(8.9275), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, "1/110", data.Exposure)
		assert.Equal(t, "HUAWEI", data.CameraMake)
		assert.Equal(t, "ELE-L29", data.CameraModel)
		assert.Equal(t, 27, data.FocalLength)
		assert.Equal(t, 0, data.Orientation)
		assert.Equal(t, "", data.LensMake)
		assert.Equal(t, "", data.LensModel)
	})

	t.Run("panorama360.jpg", func(t *testing.T) {
		data, err := Exif("testdata/panorama360.jpg", fs.TypeJpeg)

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("all: %+v", data.All)

		assert.Equal(t, "", data.Artist)
		assert.Equal(t, "2020-05-24T08:55:21Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2020-05-24T11:55:21Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "", data.Title)
		assert.Equal(t, "", data.Keywords)
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "", data.Copyright)
		assert.Equal(t, 3600, data.Height)
		assert.Equal(t, 7200, data.Width)
		assert.Equal(t, float32(59.84083), data.Lat)
		assert.Equal(t, float32(30.51), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, "1/1250", data.Exposure)
		assert.Equal(t, "SAMSUNG", data.CameraMake)
		assert.Equal(t, "SM-C200", data.CameraModel)
		assert.Equal(t, "", data.CameraOwner)
		assert.Equal(t, "", data.CameraSerial)
		assert.Equal(t, 6, data.FocalLength)
		assert.Equal(t, 0, data.Orientation)
		assert.Equal(t, "", data.Projection)
	})

	t.Run("exif-example.tiff", func(t *testing.T) {
		data, err := Exif("testdata/exif-example.tiff", fs.TypeTiff)

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("all: %+v", data.All)

		assert.Equal(t, "", data.Artist)
		assert.Equal(t, "0001-01-01T00:00:00Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "0001-01-01T00:00:00Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "", data.Title)
		assert.Equal(t, "", data.Keywords)
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "", data.Copyright)
		assert.Equal(t, 43, data.Height)
		assert.Equal(t, 65, data.Width)
		assert.Equal(t, float32(0), data.Lat)
		assert.Equal(t, float32(0), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, "", data.Exposure)
		assert.Equal(t, "", data.CameraMake)
		assert.Equal(t, "", data.CameraModel)
		assert.Equal(t, "", data.CameraOwner)
		assert.Equal(t, "", data.CameraSerial)
		assert.Equal(t, 0, data.FocalLength)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, "", data.Projection)
	})

	t.Run("out-of-range-500.jpg", func(t *testing.T) {
		data, err := Exif("testdata/out-of-range-500.jpg", fs.TypeJpeg)

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("all: %+v", data.All)

		assert.Equal(t, "", data.Artist)
		assert.Equal(t, "2017-04-09T18:33:44Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2017-04-09T18:33:44Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "", data.Title)
		assert.Equal(t, "", data.Keywords)
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "", data.Copyright)
		assert.Equal(t, 2448, data.Height)
		assert.Equal(t, 3264, data.Width)
		assert.Equal(t, float32(0), data.Lat)
		assert.Equal(t, float32(0), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, "1/387", data.Exposure)
		assert.Equal(t, "Apple", data.CameraMake)
		assert.Equal(t, "iPhone 5s", data.CameraModel)
		assert.Equal(t, "", data.CameraOwner)
		assert.Equal(t, "", data.CameraSerial)
		assert.Equal(t, 29, data.FocalLength)
		assert.Equal(t, 3, data.Orientation)
		assert.Equal(t, "", data.Projection)
	})

	t.Run("digikam.jpg", func(t *testing.T) {
		data, err := Exif("testdata/digikam.jpg", fs.TypeJpeg)

		if err != nil {
			t.Fatal(err)
		}

		//  t.Logf("all: %+v", data.All)

		assert.Equal(t, "", data.Codec)
		assert.Equal(t, "", data.Artist)
		assert.Equal(t, "2020-10-17T15:48:24Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2020-10-17T17:48:24Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "Europe/Berlin", data.TimeZone)
		assert.Equal(t, "", data.Title)
		assert.Equal(t, "", data.Keywords)
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "", data.Copyright)
		assert.Equal(t, 2736, data.Height)
		assert.Equal(t, 3648, data.Width)
		assert.Equal(t, float32(52.46052), data.Lat)
		assert.Equal(t, float32(13.331402), data.Lng)
		assert.Equal(t, 84, data.Altitude)
		assert.Equal(t, "1/50", data.Exposure)
		assert.Equal(t, "HUAWEI", data.CameraMake)
		assert.Equal(t, "ELE-L29", data.CameraModel)
		assert.Equal(t, "", data.CameraOwner)
		assert.Equal(t, "", data.CameraSerial)
		assert.Equal(t, "", data.LensMake)
		assert.Equal(t, "", data.LensModel)
		assert.Equal(t, 27, data.FocalLength)
		assert.Equal(t, 0, int(data.Orientation))
	})

	t.Run("notebook.jpg", func(t *testing.T) {
		data, err := Exif("testdata/notebook.jpg", fs.TypeJpeg)

		if err != nil {
			t.Fatal(err)
		}

		//  t.Logf("all: %+v", data.All)

		assert.Equal(t, 3000, data.Height)
		assert.Equal(t, 4000, data.Width)
		assert.Equal(t, float32(0), data.Lat)
		assert.Equal(t, float32(0), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, "1/24", data.Exposure)
		assert.Equal(t, "HMD Global", data.CameraMake)
		assert.Equal(t, "Nokia X71", data.CameraModel)
		assert.Equal(t, 26, data.FocalLength)
		assert.Equal(t, 6, int(data.Orientation))
	})
}
