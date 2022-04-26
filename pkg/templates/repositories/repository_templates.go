package repositories

var MicroservicesTemplates = map[string]string{
	"go-fiber":          "https://github.com/veda-p/go-fiber.git",
	"python-fastapi":    "https://github.com/veda-p/python-fastapi-sql.git",
	"node-express":      "https://github.com/The-Grand-Stew/nodejs-express-skeleton.git",
	"go-fiber-postgres": "https://gist.github.com/92e6dbf7502187d76f23e73b7133f213.git",
	"aws-fargate":       "https://github.com/The-Grand-Stew/container-based-arch-tf.git",
	"gcp-cloudrun":      "",
}

var CloudInfraTemplates = map[string]string{
	"aws-setup":       "https://github.com/The-Grand-Stew/tf-aws-state.git",
	"aws-ecs-fargate": "https://github.com/The-Grand-Stew/container-based-arch-tf.git",
	// YET TO ARRIVE
	"gcp-setup":    "",
	"gcp-cloudrun": "",
}
