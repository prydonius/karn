# karn: Manage multiple Git identities with ease

Setup your Git repositories to always use a specific identity based on the directory tree.

With karn, you never have to manually change the local repository configuration to a different
identity from your global.

karn will change your repository's local user.name and user.email configuration if necessary, but will never modify your global configuration.

## Install

### Pre-built binary
Head to the [releases](https://github.com/prydonius/karn/releases) page to download pre-built binaries for OS X. Note that karn is currently only tested on OS X.

### Go
You can install karn using Go with the following command:
```
go get github.com/prydonius/karn/cmd/karn
```

## Usage
karn can be used in two ways!

### Automatically check for identity updates before running a Git command
**Note: this method overrides the `git` command with a function that runs `karn update` before executing the original Git command.**
*Run `karn init` to see exactly what the `git` command is overriden with*

If you're okay with the scary warnings above, add the following line to your shell startup script (e.g. .bash_profile, .zshrc)
```
if which karn > /dev/null; then eval "$(karn init)"; fi
```

### Run manually when you want to update
Alternatively, you can run `karn update` manually in a Git repository whenever you need to update your identity for that repository.

## Configuring Identities
karn looks for a YAML configuration file in your home directory, `~/.karn.yml`.

A sample configuration looks like this:
```
---
/Fun:
  name: Adnan Abdulhussein
  email: adnan@prydoni.us
/Projects:
  name: Adnan A
  email: adnan@bitnami.com
```
