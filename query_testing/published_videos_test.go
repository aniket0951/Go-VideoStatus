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

	// update video status in verify_video
	id, _ := uuid.Parse("ec76bd81-2264-4b13-a153-c1b40add7855")

	update_args := db.UpdateVerifyVideoStatusParams{
		VideoID: id,
		Status:  helper.VIDEO_PUBLISHED,
	}

	verify_video, err := testQueries.UpdateVerifyVideoStatus(context.Background(), update_args)

	require.NoError(t, err)
	require.NotEmpty(t, verify_video)

	require.Equal(t, verify_video.Status, update_args.Status)

}

func TestCreatePublishedVideo(t *testing.T) {
	createRandomPublishedVideo(t)
}

func TestFetchPublishedVideos(t *testing.T) {
	args := db.FetchAllPublishedVideosParams{
		Limit:  10,
		Offset: 0,
	}

	videos, err := testQueries.FetchAllPublishedVideos(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, videos)
}

func TestUnPublishedVideo(t *testing.T) {
	video_id, err := uuid.Parse("73281f1b-1033-4636-a8d7-f4963eb33915")

	require.NoError(t, err)

	args := db.UpdatePublishedVideoStatusParams{
		VideoID: video_id,
		Status:  helper.VIDEO_UNPUBLISHED,
	}

	result, err := testQueries.UpdatePublishedVideoStatus(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, result.Status, args.Status)
}
