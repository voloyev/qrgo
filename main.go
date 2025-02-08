package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func generateQrCodeActio(ctx context.Context, cmd *cli.Command) error {
	text := cmd.Args().Get(0)
	if text == "" {
		return errors.New("text is required")
	}

	qrc, err := qrcode.New(text)
	if err != nil {
		return fmt.Errorf("could not generate QRCode: %v", err)
	}

	w, err := standard.New("repo-qrcode.jpeg")
	if err != nil {
		return fmt.Errorf("standard.New failed: %v", err)
	}

	if err = qrc.Save(w); err != nil {
		return fmt.Errorf("could not save image: %v", err)
	}

	return nil
}

func main() {
	cmd := &cli.Command{
		Name:   "Generate QR code",
		Usage:  "fight the loneliness!",
		Action: generateQrCodeActio,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
