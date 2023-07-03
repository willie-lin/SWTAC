package main

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestUpload(t *testing.T) {

	// テストサーバ起動
	e := echo.New()
	e.POST("/upload", upload)
	testServer := httptest.NewServer(e)
	defer testServer.Close()

	// クライアントの画像ファイルアップロード処理
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	fileWriter, err := writer.CreateFormFile("file", "test.png")
	if err != nil {
		t.Fatalf("Failed to create file writer. %s", err)
	}

	readFile, err := os.Open("test.png")
	if err != nil {
		t.Fatalf("Failed to open file. %s", err)
	}
	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	writer.Close()

	res, err := http.Post(testServer.URL+"/upload", writer.FormDataContentType(), &buffer)
	if err != nil {
		t.Fatalf("Failed to POST request. %s", err)
	}

	// API レスポンス検証
	message, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatalf("Failed to read HTTP response body. %s", err)
	}
	assert.Equal(t, "Image has been saved.", string(message))

	// アップロードを行った画像ファイルとサーバ側が保存した画像の内容が一致しているかどうか検証
	outFile, err := os.Open("test_out.png")
	if err != nil {
		t.Fatalf("Failed to open file. %s", err)
	}
	defer outFile.Close()
	var expectedBytes []byte
	readFile.Write(expectedBytes)
	var actualBytes []byte
	outFile.Write(actualBytes)
	assert.Equal(t, expectedBytes, actualBytes)
}
