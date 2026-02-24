/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package filemanager_test

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"origadmin/application/admin/tests/e2e"
	"origadmin/application/admin/tests/tools"
)

func TestSimpleUploadAndDownload(t *testing.T) {
	// 1. Initialize Test HTTP Client
	client := tools.NewTestHTTPClientWithTimeout("http://localhost:8000", 30*time.Second)
	client.SetPrefix("/api/v1")

	// 2. Login to get a token using the default helper
	token := e2e.LoginAndGetToken(t)
	require.NotEmpty(t, token, "Failed to get token")

	// 3. Prepare file content
	fileSize := 1024 // 1KB
	fileContent := make([]byte, fileSize)
	_, err := rand.Read(fileContent)
	require.NoError(t, err)

	fileName := "test-upload-" + strconv.FormatInt(time.Now().Unix(), 10) + ".txt"

	// 4. Upload the file using the new Upload helper to the proxy endpoint
	formFields := map[string]string{"visibility": "private"}
	uploadResp := client.Upload(t, "/storage/upload", fileContent, fileName, formFields, token)
	defer uploadResp.Body.Close()

	tools.AssertHTTPStatusCode(t, uploadResp, http.StatusOK)

	// The proxy returns a custom JSON structure, not the proto UploadFileResponse
	var uploadResult struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Size int64  `json:"size"`
		URL  string `json:"url"`
	}
	err = json.NewDecoder(uploadResp.Body).Decode(&uploadResult)
	require.NoError(t, err)

	require.NotZero(t, uploadResult.ID, "File ID should not be zero")
	require.Equal(t, fileName, uploadResult.Name)
	require.Equal(t, int64(fileSize), uploadResult.Size)
	fileID := uploadResult.ID

	// 5. Get file metadata
	getResp := client.Get(t, "/storage/download/"+strconv.FormatInt(fileID, 10), token)
	defer getResp.Body.Close()

	tools.AssertHTTPStatusCode(t, getResp, http.StatusOK)

	//var getFileRespBody fmv1.GetFileResponse
	//bodyBytes, err := io.ReadAll(getResp.Body)
	//require.NoError(t, err)
	//err = protojson.Unmarshal(bodyBytes, &getFileRespBody)
	//require.NoError(t, err)
	//
	//require.NotNil(t, getFileRespBody.FileMetadata, "FileMetadata should not be nil")
	//require.Equal(t, fileID, getFileRespBody.FileMetadata.Id)
	//require.NotEmpty(t, getFileRespBody.DownloadUrl, "Download URL should not be empty")
	//
	//// 6. Download the file content
	//// The download URL might be relative or absolute (presigned).
	//downloadURL := getFileRespBody.DownloadUrl
	//t.Logf("Downloading from: %s", downloadURL)
	//
	//var downloadHTTPResp *http.Response
	//if strings.HasPrefix(downloadURL, "http") {
	//	// It's an absolute URL (likely a presigned S3/MinIO link)
	//	downloadClient := http.Client{Timeout: 30 * time.Second}
	//	var err error
	//	downloadHTTPResp, err = downloadClient.Get(downloadURL)
	//	require.NoError(t, err)
	//} else {
	//	// It's a relative URL pointing back to our gateway
	//	// Use our test client which already has the baseURL and Prefix
	//	// We need to strip the prefix if the URL already contains it
	//	urlToGet := downloadURL
	//	apiPrefix := "/api/v1"
	//	if strings.HasPrefix(urlToGet, apiPrefix) {
	//		urlToGet = strings.TrimPrefix(urlToGet, apiPrefix)
	//	}
	//	downloadHTTPResp = client.Get(t, urlToGet, token)
	//}
	//defer downloadHTTPResp.Body.Close()

	//tools.AssertHTTPStatusCode(t, downloadHTTPResp, http.StatusOK)

	downloadedContent, err := io.ReadAll(getResp.Body)
	require.NoError(t, err)

	// 7. Verify content
	if !bytes.Equal(fileContent, downloadedContent) {
		t.Errorf("Downloaded content does not match uploaded content. Sizes: uploaded=%d, downloaded=%d", len(fileContent), len(downloadedContent))
		previewLen := 32
		if len(fileContent) < previewLen {
			previewLen = len(fileContent)
		}
		if len(downloadedContent) < previewLen {
			previewLen = len(downloadedContent)
		}
		t.Errorf("Uploaded preview: %x", fileContent[:previewLen])
		t.Errorf("Downloaded preview: %x", downloadedContent[:previewLen])
		t.FailNow()
	}

	// 8. Delete the file
	deleteResp := client.Delete(t, "/storage/delete/"+strconv.FormatInt(fileID, 10), token)
	defer deleteResp.Body.Close()
	tools.AssertHTTPStatusCode(t, deleteResp, http.StatusOK)

	// 9. Verify deletion
	require.Eventually(t, func() bool {
		verifyResp := client.Get(t, "/storage/download/"+strconv.FormatInt(fileID, 10), token)
		defer verifyResp.Body.Close()
		return verifyResp.StatusCode == http.StatusNotFound
	}, 5*time.Second, 500*time.Millisecond, "File was not deleted; expected 404")
}
