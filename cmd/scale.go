/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"image"
	"regexp"
	"strconv"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/spf13/cobra"
)

var scaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "Upscales or downscales an image",
	Long:  "Upscales or downscales an image, either as a copy or by changing the original image file",
	Run:   scaleRun,
}

func scaleRun(cmd *cobra.Command, args []string) {
	imgPath := args[0]
	fmt.Printf("mosaic: image path: %s\n", imgPath)

	img, err := imgio.Open(imgPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	newX, newY, err := calcNewImageDimensions(args[1], img)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("mosaic: new image size: %dx%d\n", newX, newY)
	newImg := transform.Resize(img, newX, newY, transform.Linear)

	if err := imgio.Save(`C:\Users\evert\new.png`, newImg, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
	}
}

func calcNewImageDimensions(value string, img image.Image) (int, int, error) {
	fmt.Printf("mosaic: value: %s\n", value)
	fmt.Printf("mosaic: current image size: %dx%d\n", img.Bounds().Size().X, img.Bounds().Size().Y)

	regexpPerc := regexp.MustCompile("^([0-9]+)%$")
	regexpDim := regexp.MustCompile("^([0-9]+)x([0-9]+)$")

	if regexpPerc.MatchString(value) {
		fmt.Printf("mosaic: %s is a percentage value\n", value)
		perc, err := strconv.Atoi(regexpPerc.FindStringSubmatch(value)[1])
		if err != nil {
			return 0, 0, err
		}
		fmt.Printf("mosaic: %d percent\n", perc)
		x, y := calcNewImageDimensionsForPercentage(img, float32(perc)/100.0)
		return x, y, nil
	} else if regexpDim.MatchString(value) {
		x, err := strconv.Atoi(regexpDim.FindStringSubmatch(value)[1])
		if err != nil {
			fmt.Println(err)
			return 0, 0, err
		}
		y, err := strconv.Atoi(regexpDim.FindStringSubmatch(value)[2])
		if err != nil {
			fmt.Println(err)
			return 0, 0, err
		}
		return x, y, nil
	} else {
		return 0, 0, fmt.Errorf("mosaic: %s is not a recognized percentage or dimension value", value)
	}
}

func calcNewImageDimensionsForPercentage(img image.Image, perc float32) (int, int) {
	newX := int(float32(img.Bounds().Size().X) * perc)
	newY := int(float32(img.Bounds().Size().Y) * perc)
	return newX, newY
}

func init() {
	rootCmd.AddCommand(scaleCmd)

	scaleCmd.Flags().String("value", "", "The scale value for the new image based on the source image. Accepts both % and px based values")
}
