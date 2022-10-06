# stew
## Install

### Brew Installlation

- Install using brew
``` 
brew tap The-Grand-Stew/stew
brew install stew
```

## Usage
- Export your access github token  
`export HOMEBREW_GITHUB_ACCESS_TOKEN="your-token"`
- Run `stew --help` for all command options
- Run a command and follow the steps
- To deploy on AWS, make sure the credentials are exported as env variables (Stew checks for only env vars now, profiles and loading credentials is in development)
```export AWS_ACCESS_KEY_ID=""
    export AWS_SECRET_ACCESS_KEY=""
```
- To get an CLI "app" tour:
`stew play`

## Supported Languages and Frameworks:
### Go
- Fiber
### Node
- Express
### Python
- FastAPI (coming soon)

## Support Cloud Architectures:
### AWS
- ECS Fargate
### GCP
- Cloudrun (Coming soon)

## Contacts:
- Vedashree Patil   
- Aashrit Shankar   
- Arun Kutty        
