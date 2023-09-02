package querytesting

import (
	"context"
	"testing"

	"github.com/aniket0951/video_status/apis/helper"
	db "github.com/aniket0951/video_status/sqlc_lib"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomPublishedVideo(t *testing.T) {
	video_id, _ := uuid.Parse("08da9c3e-6b37-4f54-893b-a42f74503e76")
	published_by, _ := uuid.Parse("03beec37-4362-49ea-a0e4-1b25279e321c")

	args := db.CreatePublishedVideoParams{
		VideoID:     video_id,
		PublishedBy: published_by,
		Status:      helper.VIDEO_PUBLISHED,
	}

	publishe_video, err := testQueries.CreatePublishedVideo(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, publishe_video)

	require.Equal(t, publishe_video.VideoID, args.VideoID)
	require.Equal(t, publishe_video.PublishedBy, args.PublishedBy)
	require.Equal(t, publishe_video.Status, args.Status)
	require.NotEqual(t, publishe_video.VideoID, publishe_video.PublishedBy)

	require.NotZero(t, publishe_video.ID)
	require.NotZero(t, publishe_video.CreatedAt)
}

func TestCreatePublishedVideo(t *testing.T) {
	createRandomPublishedVideo(t)
}
