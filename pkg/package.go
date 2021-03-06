package pkg

import "path/filepath"

import . "gpkg/source"
import . "gpkg/container"
import . "gpkg/version"
import . "gpkg/tools"

type Package struct {
	Source
	Tool
	Name    string
	version *Version
}

func NewPackage(name string, version *Version, source Source) *Package {
	p := &Package{
		Source:  source,
		Name:    name,
		version: version,
	}
	return p
}

func (p *Package) Clone(c Container) error {
	err := p.Source.Clone(p.Name, p.version, c.SrcDir())
	if err != nil {
		return err
	}
	p.Tool, err = NewTool(filepath.Join(c.SrcDir(), p.Name))
	if err != nil {
		return err
	}
	return nil
}
