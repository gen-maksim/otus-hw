package memorystorage

import (
	"github.com/gen-maksim/otus-hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestStorage(t *testing.T) {
	t.Run("check CRUD", func(t *testing.T) {
		s := New()
		require.Equal(t, 0, len(s.collection))

		e := storage.Event{
			Title:     "Hello World",
			StartTime: time.Now(),
			EndTime:   time.Now().Add(1 * time.Hour),
		}
		e = s.Store(e)
		require.Equal(t, 1, len(s.collection))

		e.Title = "Check"
		s.Update(e.ID, e)

		es := s.ListDay(time.Now())
		require.Equal(t, e.Title, es[0].Title)

		s.Delete(e.ID)
		require.Equal(t, 0, len(s.collection))
	})

	t.Run("check month list", func(t *testing.T) {
		s := New()
		e := storage.Event{
			StartTime: time.Now(),
			EndTime:   time.Now().Add(1 * time.Hour),
		}

		e = s.Store(e)

		es := s.ListMonth(time.Now())
		require.Contains(t, es, e)

		e = storage.Event{
			StartTime: time.Now().AddDate(1, 0, 0),
			EndTime:   time.Now().AddDate(1, 1, 0),
		}
		e = s.Store(e)
		es = s.ListMonth(time.Now())
		require.NotContains(t, es, e)

		monthStart := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Now().Location())
		e = storage.Event{
			StartTime: monthStart.AddDate(0, 0, 10),
			EndTime:   monthStart.AddDate(0, 0, 15),
		}
		e = s.Store(e)
		es = s.ListMonth(time.Now())
		require.Contains(t, es, e)
		require.Equal(t, 3, len(s.collection))
		require.Equal(t, 2, len(es))

	})
}
