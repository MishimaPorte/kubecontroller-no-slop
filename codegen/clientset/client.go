package clientset

import (
	"k8s.io/code-generator/cmd/client-gen/args"
	"k8s.io/code-generator/cmd/client-gen/generators"
	"k8s.io/gengo/v2"
	"k8s.io/gengo/v2/generator"
)

func Generate(a *args.Args) (err error) {
	if err = a.Validate(); err != nil {
		return err
	}

	var targets = func(context *generator.Context) []generator.Target {
		var ts = generators.GetTargets(context, a)
		return ts
	}
	var packagesToGenerateFor = make([]string, 0, len(a.Groups))
	for i := range a.Groups {
		for x := range a.Groups[i].Versions {
			packagesToGenerateFor = append(packagesToGenerateFor, a.Groups[i].Versions[x].Package)
		}
	}

	var nameSystems = generators.NameSystems(map[string]string{})
	var dNameSystem = generators.DefaultNameSystem()
	return gengo.Execute(
		nameSystems,
		dNameSystem,
		targets,
		gengo.StdBuildTag,
		packagesToGenerateFor,
	)
}
