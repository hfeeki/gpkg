// gpkg is the package manager for gvm
//
package main

import "os"

func main() {
	app := NewApp(os.Args)
	if app == nil {
		os.Exit(1)
	}

	switch app.command {
		case "install":
			app.install()
			break
		case "debug":
			//logger.Info(gpkg.gvm.FindPackage("manhattan"))
			return
			break
		case "uninstall":
			app.uninstall()
			break
		case "empty":
			err := app.EmptyPackages()
			if err != nil {
				app.Fatal("Failed to delete packages\n", err)
			} else {
				app.Message("Packages emptied")
			}
			break
		case "build":
			app.build()
			break
		case "list":
			app.list()
			break
		case "sources":
			app.sources()
			break
		case "version":
			app.Info(VERSION)
			break
		default:
			app.Info("Usage: gpkg [command]")
			app.Info()
			app.Info("Commands:")
			app.printUsage()
			if app.command != "" {
				app.Error("\nInvalid command (" + app.command + ")")
			}
			app.Info()
			break
	}
}

