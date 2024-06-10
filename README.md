# GPTScript Generic Credential Tool

This is a generic credential tool that prompts the user to enter a token, which will then be set to an
environment variable and saved in the credential store.

## Arguments:

- `message`: The message to present to the user when asking for the token
- `field`: The name of the field that the user will fill in (i.e. "token")
- `env`: The name of the environment variable to set to the user's input
- `sensitive`: (optional, default true) If true, asterisks will be shown as the user types, rather than plain text

## Example:

This script gets a GitHub API token from the user and uses it to make a request to the GitHub API.

```yaml
credential: github.com/gptscript-ai/credential as githubToken with "Please enter your GitHub token" as message and token as field and "GITHUB_TOKEN" as env

#!/usr/bin/env bash

curl -H "Authorization: token $GITHUB_TOKEN" \
    "https://api.github.com/search/issues?q=is:pr+repo:gptscript-ai/gptscript"
```

The credential will be saved as `githubToken`, and can be listed with `gptscript credential` and deleted with
`gptscript credential delete githubToken`. Once the credential is used, it will be automatically retrieved from
the credential store on subsequent runs.

Check out [the docs](https://docs.gptscript.ai/tools/credentials) to learn more about credential tools.

## License

Copyright (c) 2024 Acorn Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.
