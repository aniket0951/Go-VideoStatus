package querytesting

import (
	"context"
	"testing"

	"github.com/aniket0951/video_status/apis/helper"
	db "github.com/aniket0951/video_status/sqlc_lib"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomVerifyVideo(t *testing.T) {
	video_id, _ := uuid.Parse("08da9c3e-6b37-4f54-893b-a42f74503e76")
	// verify_by, _ := uuid.Parse("03beec37-4362-49ea-a0e4-1b25279e321c")
	verify_by := uuid.New()

	args := db.CreateVerifyVideoParams{
		VideoID:  video_id,
		VerifyBy: verify_by,
		Status:   helper.VIDEO_VERIFY,
	}

	verify_video, err := testQueries.CreateVerifyVideo(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, verify_video)

	require.Equal(t, verify_video.VideoID, args.VideoID)
	require.Equal(t, verify_video.VerifyBy, args.VerifyBy)
	require.Equal(t, verify_video.Status, args.Status)

	require.NotZero(t, verify_video.ID)
}

func TestCreateVerifyVideo(t *testing.T) {
	createRandomVerifyVideo(t)
}

func TestUpdateVerifyVideoStatus(t *testing.T) {
	id, _ := uuid.Parse("52b54440-2625-44ce-b555-533fd9f8b141")

	args := db.UpdateVerifyVideoStatusParams{
		VideoID: id,
		Status:  helper.VIDEO_PUBLISHED,
	}

	verify_video, err := testQueries.UpdateVerifyVideoStatus(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, verify_video)

	require.Equal(t, verify_video.Status, args.Status)
}

func TestFetchAllVerifyVideos(t *testing.T) {
	args := db.GetAllVerifyVideosParams{
		Limit:  10,
		Offset: 0,
	}

	videos, err := testQueries.GetAllVerifyVideos(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, videos)

}

// failed the video verification
func TestCreateVerificationFailed(t *testing.T) {
	video_id, _ := uuid.Parse("08da9c3e-6b37-4f54-893b-a42f74503e76")
	verified_by, _ := uuid.Parse("03beec37-4362-49ea-a0e4-1b25279e321c")

	args := db.CreateVerificationFailedParams{
		VideoID: video_id,
		VerificationFailedBy: uuid.NullUUID{
			UUID:  verified_by,
			Valid: true,
		},
		Status: helper.VIDEO_VIRIFICATION_FAILED,
		Reason: "content miss matched",
	}

	result, err := testQueries.CreateVerificationFailed(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, result.VideoID, args.VideoID)
	require.Equal(t, result.VerificationFailedBy, args.VerificationFailedBy)
	require.Equal(t, result.Status, args.Status)
	require.Equal(t, result.Reason, args.Reason)
}

func TestCreateUnPublieshVideo(t *testing.T) {
	video_id, _ := uuid.Parse("08da9c3e-6b37-4f54-893b-a42f74503e76")
	unpublished_by, _ := uuid.Parse("03beec37-4362-49ea-a0e4-1b25279e321c")

	args := db.CreateUnPublishedVideoParams{
		VideoID: video_id,
		UnpublishedBy: uuid.NullUUID{
			UUID:  unpublished_by,
			Valid: true,
		},
		Status: helper.VIDEO_UNPUBLISHED,
		Reason: "content miss matched",
	}

	result, err := testQueries.CreateUnPublishedVideo(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, result.VideoID, args.VideoID)
	require.Equal(t, result.UnpublishedBy, args.UnpublishedBy)
	require.Equal(t, result.Status, args.Status)
	require.Equal(t, result.Reason, args.Reason)
}

func TestDeleteVerificationFailed(t *testing.T) {
	video_id := uuid.New()

	result, err := testQueries.DeleteVerificationFailed(context.Background(), video_id)

	require.NoError(t, err)
	require.NotEmpty(t, result)

}
