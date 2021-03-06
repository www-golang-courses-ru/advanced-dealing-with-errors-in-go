package errs

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPretty(t *testing.T) {
	firstErr := sql.ErrNoRows

	err := fmt.Errorf("cannot get user schedule: %w",
		fmt.Errorf("cannot build data for event: %w",
			fmt.Errorf("cannot get image: %w",
				fmt.Errorf("cannot get image: %w",
					fmt.Errorf("cannot get image: %w",
						fmt.Errorf("cannot get image for event: %w", firstErr))))))
	t.Log(err)

	pErr := Pretty(err)
	require.Error(t, pErr)
	t.Log(pErr)

	assert.ErrorIs(t, pErr, firstErr)
	assert.Equal(t, `cannot get user schedule:
cannot build data for event:
cannot get image:
cannot get image:
cannot get image:
cannot get image for event:
`+firstErr.Error(), pErr.Error())
}

func TestPretty_MoreColons(t *testing.T) {
	firstErr := sql.ErrNoRows

	err := fmt.Errorf("user service: cannot get user schedule: %w",
		fmt.Errorf("event service: cannot build data for event: %w",
			fmt.Errorf("storage: cannot get image: %w", firstErr)))
	t.Log(err)

	pErr := Pretty(err)
	require.Error(t, pErr)
	t.Log(pErr)

	assert.ErrorIs(t, pErr, firstErr)
	assert.Equal(t, `user service: cannot get user schedule:
event service: cannot build data for event:
storage: cannot get image:
`+firstErr.Error(), pErr.Error())
}

func TestPretty_NoError(t *testing.T) {
	err := Pretty(nil)
	require.NoError(t, err)
}
