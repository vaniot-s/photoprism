package entity

import (
	"testing"
	"time"

	"github.com/photoprism/photoprism/internal/form"

	"github.com/gosimple/slug"
	"github.com/photoprism/photoprism/pkg/txt"
	"github.com/stretchr/testify/assert"
)

func TestNewAlbum(t *testing.T) {
	t.Run("name Christmas 2018", func(t *testing.T) {
		album := NewAlbum("Christmas 2018", AlbumDefault)
		assert.Equal(t, "Christmas 2018", album.AlbumTitle)
		assert.Equal(t, "christmas-2018", album.AlbumSlug)
	})
	t.Run("name empty", func(t *testing.T) {
		album := NewAlbum("", AlbumDefault)

		defaultName := time.Now().Format("January 2006")
		defaultSlug := slug.Make(defaultName)

		assert.Equal(t, defaultName, album.AlbumTitle)
		assert.Equal(t, defaultSlug, album.AlbumSlug)
	})
	t.Run("type empty", func(t *testing.T) {
		album := NewAlbum("Christmas 2018", "")
		assert.Equal(t, "Christmas 2018", album.AlbumTitle)
		assert.Equal(t, "christmas-2018", album.AlbumSlug)
	})
}

func TestAlbum_SetName(t *testing.T) {
	t.Run("valid name", func(t *testing.T) {
		album := NewAlbum("initial name", AlbumDefault)
		assert.Equal(t, "initial name", album.AlbumTitle)
		assert.Equal(t, "initial-name", album.AlbumSlug)
		album.SetTitle("New Album Name")
		assert.Equal(t, "New Album Name", album.AlbumTitle)
		assert.Equal(t, "new-album-name", album.AlbumSlug)
	})
	t.Run("empty name", func(t *testing.T) {
		album := NewAlbum("initial name", AlbumDefault)
		assert.Equal(t, "initial name", album.AlbumTitle)
		assert.Equal(t, "initial-name", album.AlbumSlug)

		album.SetTitle("")
		expected := album.CreatedAt.Format("January 2006")
		assert.Equal(t, expected, album.AlbumTitle)
		assert.Equal(t, slug.Make(expected), album.AlbumSlug)
	})
	t.Run("long name", func(t *testing.T) {
		longName := `A value in decimal degrees to a precision of 4 decimal places is precise to 11.132 meters at the 
equator. A value in decimal degrees to 5 decimal places is precise to 1.1132 meter at the equator. Elevation also 
introduces a small error. At 6,378 m elevation, the radius and surface distance is increased by 0.001 or 0.1%. 
Because the earth is not flat, the precision of the longitude part of the coordinates increases 
the further from the equator you get. The precision of the latitude part does not increase so much, 
more strictly however, a meridian arc length per 1 second depends on the latitude at the point in question. 
The discrepancy of 1 second meridian arc length between equator and pole is about 0.3 metres because the earth 
is an oblate spheroid.`
		expected := txt.Clip(longName, txt.ClipDefault)
		slugExpected := txt.Clip(longName, txt.ClipSlug)
		album := NewAlbum(longName, AlbumDefault)
		assert.Equal(t, expected, album.AlbumTitle)
		assert.Contains(t, album.AlbumSlug, slug.Make(slugExpected))
	})
}

func TestAlbum_SaveForm(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		album := NewAlbum("Old Name", AlbumDefault)

		assert.Equal(t, "Old Name", album.AlbumTitle)
		assert.Equal(t, "old-name", album.AlbumSlug)

		album2 := Album{ID: 123, AlbumTitle: "New name", AlbumDescription: "new description", AlbumCategory: "family"}

		albumForm, err := form.NewAlbum(album2)

		if err != nil {
			t.Fatal(err)
		}

		err = album.SaveForm(albumForm)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "New name", album.AlbumTitle)
		assert.Equal(t, "new description", album.AlbumDescription)
		assert.Equal(t, "Family", album.AlbumCategory)

	})
}

func TestAddPhotoToAlbums(t *testing.T) {
	t.Run("success one album", func(t *testing.T) {
		err := AddPhotoToAlbums("pt9jtxrexxvl0yh0", []string{"at6axuzitogaaiax"})

		if err != nil {
			t.Fatal(err)
		}

		a := Album{AlbumUID: "at6axuzitogaaiax"}

		if err := a.Find(); err != nil {
			t.Fatal(err)
		}

		var entries []PhotoAlbum

		if err := Db().Where("album_uid = ? AND photo_uid = ?", "at6axuzitogaaiax", "pt9jtxrexxvl0yh0").Find(&entries).Error; err != nil {
			t.Fatal(err)
		}

		if len(entries) < 1 {
			t.Error("at least one album entry expected")
		}

		// t.Logf("photo album entries: %+v", entries)
	})

	t.Run("empty photo", func(t *testing.T) {
		err := AddPhotoToAlbums("", []string{"at6axuzitogaaiax"})

		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("invalid photo uid", func(t *testing.T) {
		assert.Error(t, AddPhotoToAlbums("xxx", []string{"at6axuzitogaaiax"}))
	})

	t.Run("success two album", func(t *testing.T) {
		err := AddPhotoToAlbums("pt9jtxrexxvl0yh0", []string{"at6axuzitogaaiax", ""})

		if err != nil {
			t.Fatal(err)
		}

		a := Album{AlbumUID: "at6axuzitogaaiax"}

		if err := a.Find(); err != nil {
			t.Fatal(err)
		}

		var entries []PhotoAlbum

		if err := Db().Where("album_uid = ? AND photo_uid = ?", "at6axuzitogaaiax", "pt9jtxrexxvl0yh0").Find(&entries).Error; err != nil {
			t.Fatal(err)
		}

		if len(entries) < 1 {
			t.Error("at least one album entry expected")
		}

		// t.Logf("photo album entries: %+v", entries)
	})
}

func TestNewFolderAlbum(t *testing.T) {
	t.Run("name Christmas 2018", func(t *testing.T) {
		album := NewFolderAlbum("Dogs", "dogs", "label:dog")
		assert.Equal(t, "Dogs", album.AlbumTitle)
		assert.Equal(t, "dogs", album.AlbumSlug)
		assert.Equal(t, AlbumFolder, album.AlbumType)
		assert.Equal(t, SortOrderAdded, album.AlbumOrder)
		assert.Equal(t, "label:dog", album.AlbumFilter)
	})
	t.Run("title empty", func(t *testing.T) {
		album := NewFolderAlbum("", "dogs", "label:dog")
		assert.Nil(t, album)
	})
}

func TestNewMomentsAlbum(t *testing.T) {
	t.Run("name Christmas 2018", func(t *testing.T) {
		album := NewMomentsAlbum("Dogs", "dogs", "label:dog")
		assert.Equal(t, "Dogs", album.AlbumTitle)
		assert.Equal(t, "dogs", album.AlbumSlug)
		assert.Equal(t, AlbumMoment, album.AlbumType)
		assert.Equal(t, SortOrderOldest, album.AlbumOrder)
		assert.Equal(t, "label:dog", album.AlbumFilter)
	})
	t.Run("title empty", func(t *testing.T) {
		album := NewMomentsAlbum("", "dogs", "label:dog")
		assert.Nil(t, album)
	})
}

func TestNewStateAlbum(t *testing.T) {
	t.Run("name Christmas 2018", func(t *testing.T) {
		album := NewStateAlbum("Dogs", "dogs", "label:dog")
		assert.Equal(t, "Dogs", album.AlbumTitle)
		assert.Equal(t, "dogs", album.AlbumSlug)
		assert.Equal(t, AlbumState, album.AlbumType)
		assert.Equal(t, SortOrderNewest, album.AlbumOrder)
		assert.Equal(t, "label:dog", album.AlbumFilter)
	})
	t.Run("title empty", func(t *testing.T) {
		album := NewStateAlbum("", "dogs", "label:dog")
		assert.Nil(t, album)
	})
}

func TestNewMonthAlbum(t *testing.T) {
	t.Run("name Christmas 2018", func(t *testing.T) {
		album := NewMonthAlbum("Dogs", "dogs", 2020, 7)
		assert.Equal(t, "Dogs", album.AlbumTitle)
		assert.Equal(t, "dogs", album.AlbumSlug)
		assert.Equal(t, AlbumMonth, album.AlbumType)
		assert.Equal(t, SortOrderOldest, album.AlbumOrder)
		assert.Equal(t, "public:true year:2020 month:7", album.AlbumFilter)
		assert.Equal(t, 7, album.AlbumMonth)
		assert.Equal(t, 2020, album.AlbumYear)
	})
	t.Run("title empty", func(t *testing.T) {
		album := NewMonthAlbum("", "dogs", 2020, 8)
		assert.Nil(t, album)
	})
}

func TestFindAlbumBySlug(t *testing.T) {
	t.Run("1 result", func(t *testing.T) {
		album := FindAlbumBySlug("holiday-2030", AlbumDefault)

		if album == nil {
			t.Fatal("expected to find an album")
		}

		assert.Equal(t, "Holiday2030", album.AlbumTitle)
		assert.Equal(t, "holiday-2030", album.AlbumSlug)
	})
	t.Run("no result", func(t *testing.T) {
		album := FindAlbumBySlug("holiday-2030", AlbumMonth)

		if album != nil {
			t.Fatal("album should be nil")
		}
	})
}

func TestAlbum_String(t *testing.T) {
	t.Run("return slug", func(t *testing.T) {
		album := Album{
			AlbumUID:   "abc123",
			AlbumSlug:  "test-slug",
			AlbumType:  AlbumDefault,
			AlbumTitle: "Test Title",
		}
		assert.Equal(t, "test-slug", album.String())
	})
	t.Run("return title", func(t *testing.T) {
		album := Album{
			AlbumUID:   "abc123",
			AlbumSlug:  "",
			AlbumType:  AlbumDefault,
			AlbumTitle: "Test Title",
		}
		assert.Contains(t, album.String(), "Test Title")
	})
	t.Run("return uid", func(t *testing.T) {
		album := Album{
			AlbumUID:   "abc123",
			AlbumSlug:  "",
			AlbumType:  AlbumDefault,
			AlbumTitle: "",
		}
		assert.Equal(t, "abc123", album.String())
	})
	t.Run("return unknown", func(t *testing.T) {
		album := Album{
			AlbumUID:   "",
			AlbumSlug:  "",
			AlbumType:  AlbumDefault,
			AlbumTitle: "",
		}
		assert.Equal(t, "[unknown album]", album.String())
	})
}

func TestAlbum_IsMoment(t *testing.T) {
	t.Run("false", func(t *testing.T) {
		album := Album{
			AlbumUID:   "abc123",
			AlbumSlug:  "test-slug",
			AlbumType:  AlbumDefault,
			AlbumTitle: "Test Title",
		}
		assert.False(t, album.IsMoment())
	})
	t.Run("true", func(t *testing.T) {
		album := Album{
			AlbumUID:   "abc123",
			AlbumSlug:  "test-slug",
			AlbumType:  AlbumMoment,
			AlbumTitle: "Test Title",
		}
		assert.True(t, album.IsMoment())
	})
}

func TestAlbum_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		album := Album{
			AlbumUID:   "abc123",
			AlbumSlug:  "test-slug",
			AlbumType:  AlbumDefault,
			AlbumTitle: "Test Title",
		}
		assert.Equal(t, "test-slug", album.AlbumSlug)

		err := album.Update("AlbumSlug", "new-slug")

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "new-slug", album.AlbumSlug)
	})
}

func TestAlbum_Save(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		album := NewStateAlbum("Dogs", "dogs", "label:dog")

		initialDate := album.UpdatedAt

		err := album.Save()

		if err != nil {
			t.Fatal(err)
		}
		afterDate := album.UpdatedAt
		t.Log(initialDate)
		t.Log(afterDate)
		//TODO Why does it fail?
		//assert.True(t, afterDate.After(initialDate))
	})
}

func TestAlbum_Create(t *testing.T) {
	t.Run("album", func(t *testing.T) {
		album := Album{
			AlbumType: AlbumDefault,
		}

		err := album.Create()

		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("moment", func(t *testing.T) {
		album := Album{
			AlbumType: AlbumMoment,
		}

		err := album.Create()

		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("month", func(t *testing.T) {
		album := Album{
			AlbumType: AlbumMonth,
		}

		err := album.Create()

		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("folder", func(t *testing.T) {
		album := Album{
			AlbumType: AlbumFolder,
		}

		err := album.Create()

		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestAlbum_Title(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		album := Album{
			AlbumUID:   "abc123",
			AlbumSlug:  "test-slug",
			AlbumType:  AlbumDefault,
			AlbumTitle: "Test Title",
		}
		assert.Equal(t, "Test Title", album.Title())
	})
}

func TestAlbum_Links(t *testing.T) {
	t.Run("1 result", func(t *testing.T) {
		album := AlbumFixtures.Get("christmas2030")
		links := album.Links()
		assert.Equal(t, "4jxf3jfn2k", links[0].LinkToken)
	})
}

func TestAlbum_AddPhotos(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		album := Album{
			AlbumUID:   "abc123",
			AlbumSlug:  "test-slug",
			AlbumType:  AlbumDefault,
			AlbumTitle: "Test Title",
		}
		added := album.AddPhotos([]string{"ab", "cd"})
		assert.Equal(t, 2, len(added))
	})
}

func TestAlbum_RemovePhotos(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		album := Album{
			AlbumUID:   "abc123",
			AlbumSlug:  "test-slug",
			AlbumType:  AlbumDefault,
			AlbumTitle: "Test Title",
		}
		removed := album.RemovePhotos([]string{"ab", "cd"})
		assert.Equal(t, 2, len(removed))
	})
}
