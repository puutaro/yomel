# `yomel ` deep dive

`yomel` parses arguments sequentially from left to right. Arguments are divided into global/stage telemetry controllers, structural elements, and value/option modifiers.

## Table 
- [`yomel ` deep dive](#yomel--deep-dive)
  - [1. Telemetry and Filter Options](#1-telemetry-and-filter-options)
    - [`title "<pipeline_title>"`](#title-pipeline_title)
    - [`--no-live-stdout`](#--no-live-stdout)
    - [`--no-live-stderr`](#--no-live-stderr)
    - [`--log`](#--log)
    - [`--gen`](#--gen)
    - [`--direct`](#--direct)
    - [`--log-filter "<shell_command>"`](#--log-filter-shell_command)
    - [`--err-log-filter "<shell_command>"`](#--err-log-filter-shell_command)
  - [2. Structural Stage Elements](#2-structural-stage-elements)
    - [`stage "<stage_name>"`](#stage-stage_name)
    - [`-cmd "<binary>"`](#-cmd-binary)
    - [`-svc "<service_name>"`](#-svc-service_name)
    - [`-act "<action_name>"`](#-act-action_name)
  - [3. Option and Argument Value Modifiers with PascalCase Suffixes](#3-option-and-argument-value-modifiers-with-pascalcase-suffixes)
    - [`--opt[PascalCase] "<flag>"`](#--optpascalcase-flag)
    - [`--lop[PascalCase] "<flag>"`](#--loppascalcase-flag)
    - [`--val[PascalCase]`](#--valpascalcase)
    - [`--arg[PascalCase]`](#--argpascalcase)


## 1. Telemetry and Filter Options
These options control debugging output and stream filtering. They do not alter the data passing through the core pipeline but manage what is written to `stderr`.

### `title "<pipeline_title>"`

  * **Meaning:** Sets a title for the overall pipeline. When multiple stages are executed and a title is specified, it displays a distinct header banner (`YOMEL-LOG-TITLE:`) showing the title and the total generated pipeline command.

  * **Usage:** Place it at the beginning of the command (global control section).

### `--no-live-stdout`

  * **Meaning:** Suppresses real-time streaming of standard output (stdout) to the console while the pipeline commands run in the background.

  * **Usage:** Useful for muting noisy background stream outputs during execution.

### `--no-live-stderr`

  * **Meaning:** Suppresses real-time streaming of standard error (stderr) to the console during execution.

  * **Usage:** Useful for hiding intermediate progress or warning logs until the final error handling or reporting stage.

### `--log`
  * **Meaning:** Activates the internal logging system. When this flag is present, `yomel` prints detailed panel execution metrics, generated shell commands, and raw step statuses to `stderr`.
  * **Usage:** Place it at the very beginning of the command to apply globally, or within specific sections.

### `--gen`
  * **Meaning:** Outputs the total pipeline command. 
  * **Usage:** Useful for a `dry-run` to confirm the entire generated pipeline command before execution.

### `--direct`
  * **Meaning:** Executes the shell pipeline directly without internal logging decoration. 
  * **Usage:** Enables faster shell pipe execution and captures real-time `stderr` logs directly.

### `--log-filter "<shell_command>"`
  * **Meaning:** Attaches an asynchronous log interceptor for standard output (`stdout`). The log data captured from the stage is passed to this shell command (e.g., `grep`, `awk`, `sed`) via stdin before being printed.
  * **Usage:** `--log-filter "grep 'ERROR'"` ensures only log lines containing "ERROR" are emitted to your console log view.

### `--err-log-filter "<shell_command>"`
  * **Meaning:** Attaches an asynchronous log interceptor for standard error (`stderr`). This functions exactly like `--log-filter` but processes error streams thrown by the executing binaries.
  * **Usage:** `--err-log-filter "awk '{print "[ERR] " \$0}'"` prefixes all error outputs with a custom tag.

## 2. Structural Stage Elements
These keywords separate different processes and define command parts.

### `stage "<stage_name>"`
  * **Meaning:** Initializes a new execution boundary (a pipeline stage). All subsequent parameters (`-cmd`, `--opt`, etc.) are assigned to this stage until a new `stage` keyword appears.
  * **Usage:** `stage "fetch-data"` creates a clear logical separator for documentation and logging.

### `-cmd "<binary>"`
  * **Meaning:** Specifies the main executable or binary command to be run in the current stage.
  * **Usage:** `-cmd "aws"`, `-cmd "curl"`, or `-cmd "docker"`.

### `-svc "<service_name>"`
  * **Meaning:** Declares a sub-service or second-level command hierarchy. This is highly useful for modern cloud CLIs.
  * **Usage:** In `aws s3api`, `s3api` is the service. Example: `-svc "s3api"`.

### `-act "<action_name>"`
  * **Meaning:** Declares the operation, verb, or action to be performed under the specified command or service.
  * **Usage:** In `docker container run`, `run` is the action. Example: `-act "list-objects"`.

## 3. Option and Argument Value Modifiers with PascalCase Suffixes
Modifiers specify how parameters, options, and trailing arguments are constructed and quoted. You can append an optional Alphanumeric PascalCase description suffix to `--opt`, `--lop`, `--val`, and `--arg` to document the role of each argument clearly.

### `--opt[PascalCase] "<flag>"`
  * **Meaning:** Generates a short-style option flag (prefixed with a single dash `-`). An optional PascalCase suffix can be appended for documentation.
  * **Usage:** `--optVerbose "v"` generates `-v`.

### `--lop[PascalCase] "<flag>"`
  * **Meaning:** Generates a long-style option flag (prefixed with double dashes `--`). An optional PascalCase suffix can be appended for documentation.
  * **Usage:** `--lopRegion "region"` generates `--region`.

### `--val[PascalCase]`
  * **Meaning:** Declares a value associated with the preceding option (`--opt` or `--lop`). It **must** be immediately followed by a quote control flag (`--s` or `--n`), and can include an optional PascalCase description suffix.
  * **Modifiers:**
    * `--val[PascalCase] --s "<string>"`: Encloses the value in single quotes (`'value'`).
    * `--val[PascalCase] --n "<string>"`: Emits the raw value without quotes (`value`), ideal for numbers or unquoted tokens.
  * **Usage:** `--lopId --valId --s "123"` generates `--id '123'`. `--lopCount --valCount --n "5"` generates `--count 5`.

### `--arg[PascalCase]`
  * **Meaning:** Appends a standalone, positional argument to the tail end of the generated command string. It **must** be immediately followed by a quote control flag (`--s` or `--n`), and can include an optional PascalCase description suffix.
  * **Modifiers:**
    * `--arg[PascalCase] --s "<string>"`: Appends a single-quoted positional argument.
    * `--arg[PascalCase] --n "<string>"`: Appends an unquoted positional argument.
  * **Usage:** `--argPattern --s "/pattern/d"` appends `'/pattern/d'`.



