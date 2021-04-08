// +build mage

package main

import (
	// mage:import

	"os"

	"get.porter.sh/porter/mage/releases"
)

// We are migrating to mage, but for now keep using make as the main build script interface.

// Publish the cross-compiled binaries.
func Publish(mixin string, version string, permalink string) {
	//The following demonstrates how to publish a mixin. As an example, we show how to publish our own mixins.
	// The porter mixins feed generate command is used to build an ATOM feed for sharing mixins once published

	// Set the repository name, e.g. github.com/YOURNAME/YOUR_MIXIN
	os.Setenv("PORTER_RELEASE_REPOSITORY", "TODO")

	releases.PrepareMixinForPublish(mixin, version, permalink)
	releases.PublishMixin(mixin, version, permalink)
}
