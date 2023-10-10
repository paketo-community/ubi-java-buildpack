package main

import (
	ubi8javabuildpack "github.com/paketo-community/ubi-java-buildpack/v1"

	libpak "github.com/paketo-buildpacks/libpak/v2"
)

func main() {
	libpak.BuildpackMain(
		ubi8javabuildpack.Detect,
		ubi8javabuildpack.Build,
	)
}
