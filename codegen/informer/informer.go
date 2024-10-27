package informer

import (
	"k8s.io/code-generator/cmd/informer-gen/args"
	"k8s.io/code-generator/cmd/informer-gen/generators"
	"k8s.io/gengo/v2"
	"k8s.io/gengo/v2/generator"
)

func Generate(a *args.Args, packagesToGenerateFor ...string) (err error) {

	if err = a.Validate(); err != nil {
		return err
	}

	var targets = func(context *generator.Context) []generator.Target {
		var ts = generators.GetTargets(context, a)
		return ts
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
