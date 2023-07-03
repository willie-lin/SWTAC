package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"os"
	"strings"
)

func upload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	fileModel := strings.Split(file.Filename, ".")
	fileName := fileModel[0]
	extension := fileModel[1]
	dst, err := os.Create(fmt.Sprintf("%s_out.%s", fileName, extension))
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(200, "Image has been saved")

}
