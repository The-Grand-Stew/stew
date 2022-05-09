package serverless

const ServerlessFunctionConfigYq string = `{ 
	"{{ .FunctionName }}":{
	"name":"{{ .FunctionName }}-{{ .Environment }}",
	"memorySize":128,
	"environment":{},
	"handler":"./handlers/{{ .FunctionName }}/{{ .HandlerName }}",
	"events":[
		{
			"httpApi": {
				"path":"/{{ .PathPart }}",
				"method":"{{ .HttpMethod }}"
			}
		}
	]
	}
}`
