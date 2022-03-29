package register

import (
	"github.com/deissh/rf-cli/internal/factory"
	"github.com/deissh/rf-cli/internal/utils"
	"github.com/deissh/rf-cli/pkg/log"
	"github.com/deissh/rf-cli/pkg/manifest"
	"github.com/deissh/rf-cli/pkg/rf"
	"github.com/deissh/rf-cli/pkg/view"
	"github.com/spf13/cobra"
	"os"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register",
		Short: "",
		Run:   run,
	}

	cmd.Flags().StringP("file", "f", "manifest.yaml", "file <path>")
	cmd.Flags().BoolP("yes", "y", false, "Automatically answer \"yes\" to any prompts")

	return cmd
}

func run(cmd *cobra.Command, _ []string) {
	client := factory.ClientInstance

	forceYes, err := cmd.Flags().GetBool("yes")
	utils.ExitIfError(err)

	path, err := cmd.Flags().GetString("file")
	utils.ExitIfError(err)

	// todo: more complex errors
	info, err := loadManifest(path)
	if err != nil {
		log.Warn("manifest not loaded, %s", err)
		log.Warn("Run command to create manifest.")
		log.Warn(" $ rf-cli extension init --help")
		utils.Exit("")
	}

	err = view.NewManifest(info).Render()
	utils.ExitIfError(err)

	if err = info.Validate(); err != nil {
		utils.ExitIfError(err)
	}

	log.Info("Manifest validated")

	if ok := utils.Confirm(forceYes); !ok {
		utils.Exit("aborted")
	}

	ext := info.ToExtension()

	data, err := func() (*rf.Extension, error) {
		s := utils.PrintSpinner("Creating extension from manifest...")
		defer s.Stop()

		return client.Ext.Create(ext)
	}()
	utils.ExitIfError(err)

	log.Info("Extension created, updating manifest")

	err = writeManifest(path, manifest.FromExtension(data))
	utils.ExitIfError(err)
}

func loadManifest(path string) (*manifest.Manifest, error) {
	if !utils.FileExists(path) {
		return nil, os.ErrNotExist
	}

	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return manifest.Read(f)
}

func writeManifest(path string, info *manifest.Manifest) error {
	if !utils.FileExists(path) {
		return os.ErrNotExist
	}

	f, err := os.OpenFile(path, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return manifest.Write(f, info)
}
