package querytesting

import (
	"context"
	"testing"

	db "github.com/aniket0951/video_status/sqlc_lib"
	"github.com/aniket0951/video_status/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func uploadRandomVideo(t *testing.T) {
	uploaded_by, err := uuid.Parse("03beec37-4362-49ea-a0e4-1b25279e321c")
	require.NoError(t, err)

	args := db.UploadVideoByAdminParams{
		Title:       utils.RandomString(8),
		FileAddress: utils.RandomString(4),
		UploadedBy:  uploaded_by,
		Status:      "verification",
	}

	video, err := testQueries.UploadVideoByAdmin(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, video)

	require.Equal(t, video.Title, args.Title)
	require.Equal(t, video.FileAddress, args.FileAddress)
	require.Equal(t, video.UploadedBy, args.UploadedBy)
	require.Equal(t, video.Status, args.Status)

	require.NotZero(t, video.ID)

}

func TestUploadVideoByAdmin(t *testing.T) {
	uploadRandomVideo(t)
}

func TestUpdateVideoStat(t *testing.T) {
	video_id, err := uuid.Parse("08da9c3e-6b37-4f54-893b-a42f74503e76")
	require.NoError(t, err)
	args := db.UpdateVideoStatusParams{
		ID:     video_id,
		Status: "VIDEO_INIT",
	}

	result, err := testQueries.UpdateVideoStatus(context.Background(), args)

	require.NoError(t, err)
	rows_affected, _ := result.RowsAffected()
	require.NotZero(t, rows_affected)
}

func TestFetchVideoByAdminFullDetails(t *testing.T) {
	video_id, err := uuid.Parse("7e3aa423-b265-4861-9041-fe308dfb5069")

	require.NoError(t, err)

	result, err := testQueries.GetVideoByAdminFullDetail(context.Background(), video_id)

	require.NoError(t, err)
	require.NotEmpty(t, result)
}
