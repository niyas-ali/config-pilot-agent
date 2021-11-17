# config-pilot-job

## Usage

To run the binary, please download the executable file along with the following configuration files from the binary folder

[*Go to*](https://github.com/niyas-ali/config-pilot-agent/tree/master/binary)
```
1. config-pilot-job.exe
```

```
2. patch_configuration.json
```

```
3. repository.json
```

## Build manually

Clone the repo and then run `run.sh` file

## Configurations

### repository configuration - *repository.json*

```JSON
{
    "checkout_branch":"features/__package-upgrade",
    "pr_title":"Automation - Package dependency upgrade",
    "pr_message":"",
    "azure_devops":{
        "organization":"",
        "project_name":"",
        "repository":[
            {
                "name":"",
                "url":"",
                "merge_branch":""
            },
            ...
        ]
    },
    "github":{
        "organization":"",
        "repository":[
            {
                "name":"",
                "url":"",
                "merge_branch":""
            },
            ...
        ]
    }
}
```
### patch configuration - *patch_configuration.json*

```JSON
[
    {
        "packageName":"",
        "minVersion":"^1.0.0",
        "forceUpgrade": true
    },
    ...
]
```
### Configure tokens

In order to clone/raise pr for private repositories, token needs to be configured in the system environment variable for azure DevOps and Github
For Azure

| key | value |
|--|--|
| AZ_TOKEN | secret |
| GITHUB_TOKEN| secret |
