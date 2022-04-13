package update

import (
	"github.com/deissh/rf-cli/internal/factory"
	"github.com/deissh/rf-cli/internal/utils"
	"github.com/deissh/rf-cli/pkg/log"
	"github.com/deissh/rf-cli/pkg/manifest"
	"github.com/deissh/rf-cli/pkg/rf"
	"github.com/deissh/rf-cli/pkg/view"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [extId] [-f format]",
		Short: "Update information about extension in registry",
		Run:   run,
	}

	cmd.Flags().String("format", "pretty", "output format (json, pretty-json, yaml, pretty)")
	cmd.Flags().StringP("file", "f", "manifest.yaml", "file <path>")

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	client := factory.ClientInstance

	path, err := cmd.Flags().GetString("file")
	utils.ExitIfError(err)

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

	utils.Confirm(false)

	ext := info.ToExtension()

	// override extId by args
	if len(args) > 0 {
		ext.ID = args[0]
	}

	data, err := func() (*rf.Extension, error) {
		s := utils.PrintSpinner("Updating extension from manifest...")
		defer s.Stop()

		return client.Ext.Update(ext)
	}()
	utils.ExitIfError(err)

	log.Info("Extension created, updating manifest")

	err = writeManifest(path, manifest.FromExtension(data))
	utils.ExitIfError(err)
}

func loadManifest(path string) (*manifest.Manifest, error) {
	info := &manifest.Manifest{}

	if !utils.FileExists(path) {
		return info, nil
	}

	return manifest.ReadByPath(path)
}

func writeManifest(path string, info *manifest.Manifest) error {
	err := utils.CreateFileAndBackup(path)
	if err != nil {
		return err
	}

	return manifest.WriteByPath(path, info)
}
