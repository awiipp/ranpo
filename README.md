# ranpo

> Terminal-based API client. Lighter than Postman, more honest than curl.

A TUI + CLI tool for testing APIs without leaving your terminal. Drop into the interactive UI to browse saved requests and manage environments, or fire off one-liners from the command line.

## Install

Requires Go 1.24+.

```bash
go install github.com/awiipp/ranpo@latest
```

Binary installs to `$GOPATH/bin` — make sure it's in your `PATH`.

**Build from source**

```bash
git clone https://github.com/awiipp/ranpo
cd ranpo
go build -o ranpo .
```

## Usage

### TUI

```bash
ranpo
```

Launches the interactive UI. Navigate with arrow keys or `j`/`k`, select with `enter`.

| Key | Action |
|-----|--------|
| `g` `p` `u` `P` `d` | Quick-select HTTP method |
| `c` | Collections browser |
| `e` | Environment manager |
| `ctrl+r` | Send request |
| `tab` / `shift+tab` | Cycle form fields |
| `1` / `2` | Body / Headers tab |
| `esc` | Back to home |
| `q` | Quit |

> `ctrl+c` only quits from the home screen — intentional, so you don't lose a half-typed request.

### CLI

```bash
ranpo get <url>
ranpo post <url> -b '{"key":"value"}'
ranpo put <url> -b '{"key":"value"}'
ranpo patch <url> -b '{"key":"value"}'
ranpo delete <url>
```

Run a saved request:

```bash
ranpo run <name>
ranpo run <collection>/<name>
```

Open TUI with a pre-filled form:

```bash
ranpo post https://api.example.com/data -i
```

**Flags**

| Flag | Description |
|------|-------------|
| `-t` `--token` | Bearer token |
| `-H` `--header` | Header in `Key:Value` format (repeatable) |
| `-b` `--body` | JSON request body |
| `-s` `--save` | Save request with this name |
| `-c` `--collection` | Collection to save into (default: `default`) |
| `-i` `--interactive` | Open TUI form instead of sending immediately |

**Example**

```bash
ranpo post https://api.example.com/users \
  -t "my-token" \
  -H "X-Trace-Id: abc-123" \
  -b '{"name":"Alice"}' \
  -s create-alice \
  -c users
```

## Environments

Switch between dev, staging, and prod without editing requests.

```bash
ranpo env set staging BASE_URL https://staging.api.example.com
ranpo env set staging TOKEN staging-token-abc
ranpo env use staging
ranpo get {{BASE_URL}}/health
```

`{{KEY}}` placeholders resolve from the active environment — works in URLs, headers, body, and token fields. Unknown keys are left as-is.

Other commands: `env list`, `env show`, `env delete`.

## Auth

Resolved in order: inline `--token` → env `TOKEN` → config default.

## Saved Requests

Saved via `-s` flag or the TUI "Save as" field. Existing names are overwritten.

```
~/.ranpo/collections/default.json
~/.ranpo/collections/<name>.json
```

## Data

Everything lives in `~/.ranpo/`:

| Path | Description |
|------|-------------|
| `config.yaml` | Active env, default auth |
| `collections/<name>.json` | Saved requests |
| `environments/<name>.json` | Variables |

Default config:

```yaml
active_env: local
default_auth:
  type: bearer
  token: ""
```

## License

Copyright (c) 2026 awiipp. Released under the [MIT License](https://opensource.org/licenses/MIT).
