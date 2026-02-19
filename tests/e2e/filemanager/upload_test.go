/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package filemanager_test

import (
	"bytes"
	"crypto/rand"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/protobuf/encoding/protojson"

	fmv1 "origadmin/application/admin/api/v1/services/filemanager"
	"origadmin/application/admin/tests/e2e"
	"origadmin/application/admin/tests/tools"
)

func TestSimpleUploadAndDownload(t *testing.T) {
	// 1. Login to get a token
	token := e2e.LoginAndGetToken(t)
	require.NotEmpty(t, token, "Failed to get token")

	// 2. Prepare file content
	fileSize := 1024 // 1KB
	fileContent := make([]byte, fileSize)
	_, err := rand.Read(fileContent)
	require.NoError(t, err)

	fileName := "test-upload-" + strconv.FormatInt(time.Now().Unix(), 10) + ".txt"

	// 3. Upload the file
	uploadReq := &fmv1.UploadFileRequest{
		Data:        fileContent,
		Name:        fileName,
		ContentType: "text/plain",
		Visibility:  "private",
	}

	client := tools.NewSystemTestClient("http://localhost:8000")
	client.SetPrefix("/api/v1") // Ensure prefix is correct

	// The prefix is now handled automatically by the client, so we pass the relative path.
	resp := client.Post(t, "/fm/files", uploadReq, token)
	defer resp.Body.Close()

	tools.AssertHTTPStatusCode(t, resp, http.StatusOK)

	var uploadResp fmv1.UploadFileResponse
	bodyBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	err = protojson.Unmarshal(bodyBytes, &uploadResp)
	require.NoError(t, err)

	require.NotNil(t, uploadResp.FileMetadata, "FileMetadata should not be nil")
	require.NotZero(t, uploadResp.FileMetadata.Id, "File ID should not be zero")
	require.Equal(t, fileName, uploadResp.FileMetadata.Name)
	require.Equal(t, int64(fileSize), uploadResp.FileMetadata.Size)
	fileID := uploadResp.FileMetadata.Id

	// 4. Get file metadata and download URL
	getResp := client.Get(t, "/fm/files/"+strconv.FormatInt(fileID, 10), token)
	defer getResp.Body.Close()

	tools.AssertHTTPStatusCode(t, getResp, http.StatusOK)

	var getFileResp fmv1.GetFileResponse
	bodyBytes, err = io.ReadAll(getResp.Body)
	require.NoError(t, err)
	err = protojson.Unmarshal(bodyBytes, &getFileResp)
	require.NoError(t, err)

	require.NotNil(t, getFileResp.FileMetadata, "FileMetadata should not be nil")
	require.Equal(t, fileID, getFileResp.FileMetadata.Id)
	require.NotEmpty(t, getFileResp.DownloadUrl, "Download URL should not be empty")

	// 5. Download the file content
	downloadURL := getFileResp.DownloadUrl
	// If the URL is relative, prepend the gateway base URL
	if strings.HasPrefix(downloadURL, "/") {
		downloadURL = client.GetBaseURL() + downloadURL
	}

	// Use a client with a longer timeout for downloads if necessary
	downloadClient := http.Client{Timeout: 30 * time.Second}
	t.Logf("downloadURL: %s", downloadURL)
	// Create request manually to add Authorization header
	req, err := http.NewRequest("GET", downloadURL, nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+token)

	downloadHTTPResp, err := downloadClient.Do(req)
	require.NoError(t, err)
	defer downloadHTTPResp.Body.Close()

	tools.AssertHTTPStatusCode(t, downloadHTTPResp, http.StatusOK)

	// Read the raw response body
	downloadedRawBody, err := io.ReadAll(downloadHTTPResp.Body)
	require.NoError(t, err)

	// Unmarshal the raw response body into google.api.HttpBody
	var httpBody httpbody.HttpBody
	err = protojson.Unmarshal(downloadedRawBody, &httpBody)
	require.NoError(t, err, "Failed to unmarshal downloaded content into HttpBody")

	// The actual file content is in the Data field of HttpBody
	downloadedContent := httpBody.Data

	// 6. Verify content
	if !bytes.Equal(fileContent, downloadedContent) {
		t.Logf("Uploaded size: %d, Downloaded size: %d", len(fileContent), len(downloadedContent))
		if len(fileContent) > 0 && len(downloadedContent) > 0 {
			t.Logf("Uploaded first 10 bytes: %x", fileContent[:10])
			t.Logf("Downloaded first 10 bytes: %x", downloadedContent[:10])
		}
	}
	require.True(t, bytes.Equal(fileContent, downloadedContent), "Downloaded content does not match uploaded content")

	// 7. Delete the file
	deleteResp := client.Delete(t, "/fm/files/"+strconv.FormatInt(fileID, 10), token)
	defer deleteResp.Body.Close()
	tools.AssertHTTPStatusCode(t, deleteResp, http.StatusOK)

	// 8. Verify deletion (with retry to handle potential eventual consistency)
	require.Eventually(t, func() bool {
		verifyResp := client.Get(t, "/fm/files/"+strconv.FormatInt(fileID, 10), token)
		defer verifyResp.Body.Close()
		// Keep trying until we get a 404
		return verifyResp.StatusCode == http.StatusNotFound
	}, 5*time.Second, 500*time.Millisecond, "File was not deleted within the time limit. Last status was not 404.")
}
