
<img width="512" height="474" alt="yomel4_siro_1024" src="https://github.com/user-attachments/assets/c90f8341-7ed6-4dde-a35a-1a64db71bf23" />

<!-- CIステータスバッジ -->
[![CI](https://github.com/puutaro/yomel/actions/workflows/ci.yaml/badge.svg)](https://github.com/puutaro/yomel/actions/workflows/ci.yaml)
[![Release](https://img.shields.io/github/v/release/puutaro/yomel)](https://github.com/puutaro/yomel/releases)
[![License](https://img.shields.io/github/license/puutaro/yomel)](https://github.com/puutaro/yomel/blob/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/puutaro/yomel.svg)](https://pkg.go.dev/puutaro/yomel)
[![codecov](https://codecov.io/gh/puutaro/yomel/branch/master/graph/badge.svg)](https://codecov.io/gh/puutaro/yomel)

# yomel

`yomel` is a command-line utility designed to write multi-stage shell script pipelines using a structured, flat, and human-readable argument layout—inspired by the clear, nested visual style of YAML configuration files.

By breaking down complex, nested one-liners or lengthy shell scripts into highly visible declarative steps (`stage`), `yomel` simplifies shell automation without abandoning native command-line paradigms.

---

## 📦 Installation (Linux/Mac)

### General

curl https://raw.githubusercontent.com/puutaro/yomel/refs/heads/master/install.sh | sh

### go install

go install github.com/puutaro/yomel/cmd/yomel@latest

---

## 🚀 Key Features

- 🛠️ **YAML-Like Structure:** Compose sequentially grouped pipelines via continuous CLI arguments using explicit components (`stage`, `-cmd`, `-svc`, `-act`, etc.).
- 🔄 **Auto-Chaining Pipelines:** Standard output (`stdout`) from an earlier stage automatically streams directly into the standard input (`stdin`) of the next stage using concurrent `io.TeeReader` pipes.
- 📝 **Smart Logging & Isolation:** Enable automated logging (`--log`) per stage with standalone error capturing.
- 🧽 **Asynchronous Stream Filtering:** Apply global or stage-specific custom shell hooks (`--log-filter` or `--err-log-filter`) to process log fragments on-the-fly (e.g., streaming only specific info lines).
- 🔤 **Granular Quote & Description Management:** Control parameter escaping instantly with semantic operators (`--val --s`, `--val --n`, `--arg --s`, `--arg --n`) and append optional Alphanumeric PascalCase description suffixes (e.g., `--optS3Path`, `--lopRegion`, `--valDest`, `--argPattern`) to document argument intent clearly.

---

## 🛠️ Complete Option Reference & Deep Dive

`yomel` parses arguments sequentially from left to right. Arguments are divided into global/stage telemetry controllers, structural elements, and value/option modifiers.

### 1. Telemetry and Filter Options
These options control debugging output and stream filtering. They do not alter the data passing through the core pipeline but manage what is written to `stderr`.

* **`title "<pipeline_title>"`**

  * **Meaning:** Sets a title for the overall pipeline. When multiple stages are executed and a title is specified, it displays a distinct header banner (`YOMEL-LOG-TITLE:`) showing the title and the total generated pipeline command.

  * **Usage:** Place it at the beginning of the command (global control section).

* **`--no-live-stdout`**

  * **Meaning:** Suppresses real-time streaming of standard output (stdout) to the console while the pipeline commands run in the background.

  * **Usage:** Useful for muting noisy background stream outputs during execution.

* **`--no-live-stderr`**

  * **Meaning:** Suppresses real-time streaming of standard error (stderr) to the console during execution.

  * **Usage:** Useful for hiding intermediate progress or warning logs until the final error handling or reporting stage.

* **`--log`**
  * **Meaning:** Activates the internal logging system. When this flag is present, `yomel` prints detailed panel execution metrics, generated shell commands, and raw step statuses to `stderr`.
  * **Usage:** Place it at the very beginning of the command to apply globally, or within specific sections.

* **`--gen`**
  * **Meaning:** Outputs the total pipeline command. 
  * **Usage:** Useful for a `dry-run` to confirm the entire generated pipeline command before execution.

* **`--direct`**
  * **Meaning:** Executes the shell pipeline directly without internal logging decoration. 
  * **Usage:** Enables faster shell pipe execution and captures real-time `stderr` logs directly.

* **`--log-filter "<shell_command>"`**
  * **Meaning:** Attaches an asynchronous log interceptor for standard output (`stdout`). The log data captured from the stage is passed to this shell command (e.g., `grep`, `awk`, `sed`) via stdin before being printed.
  * **Usage:** `--log-filter "grep 'ERROR'"` ensures only log lines containing "ERROR" are emitted to your console log view.

* **`--err-log-filter "<shell_command>"`**
  * **Meaning:** Attaches an asynchronous log interceptor for standard error (`stderr`). This functions exactly like `--log-filter` but processes error streams thrown by the executing binaries.
  * **Usage:** `--err-log-filter "awk '{print "[ERR] " \$0}'"` prefixes all error outputs with a custom tag.

### 2. Structural Stage Elements
These keywords separate different processes and define command parts.

* **`stage "<stage_name>"`**
  * **Meaning:** Initializes a new execution boundary (a pipeline stage). All subsequent parameters (`-cmd`, `--opt`, etc.) are assigned to this stage until a new `stage` keyword appears.
  * **Usage:** `stage "fetch-data"` creates a clear logical separator for documentation and logging.

* **`-cmd "<binary>"`**
  * **Meaning:** Specifies the main executable or binary command to be run in the current stage.
  * **Usage:** `-cmd "aws"`, `-cmd "curl"`, or `-cmd "docker"`.

* **`-svc "<service_name>"`**
  * **Meaning:** Declares a sub-service or second-level command hierarchy. This is highly useful for modern cloud CLIs.
  * **Usage:** In `aws s3api`, `s3api` is the service. Example: `-svc "s3api"`.

* **`-act "<action_name>"`**
  * **Meaning:** Declares the operation, verb, or action to be performed under the specified command or service.
  * **Usage:** In `docker container run`, `run` is the action. Example: `-act "list-objects"`.

### 3. Option and Argument Value Modifiers with PascalCase Suffixes
Modifiers specify how parameters, options, and trailing arguments are constructed and quoted. You can append an optional Alphanumeric PascalCase description suffix to `--opt`, `--lop`, `--val`, and `--arg` to document the role of each argument clearly.

* **`--opt[PascalCase] "<flag>"`**
  * **Meaning:** Generates a short-style option flag (prefixed with a single dash `-`). An optional PascalCase suffix can be appended for documentation.
  * **Usage:** `--optVerbose "v"` generates `-v`.

* **`--lop[PascalCase] "<flag>"`**
  * **Meaning:** Generates a long-style option flag (prefixed with double dashes `--`). An optional PascalCase suffix can be appended for documentation.
  * **Usage:** `--lopRegion "region"` generates `--region`.

* **`--val[PascalCase]`**
  * **Meaning:** Declares a value associated with the preceding option (`--opt` or `--lop`). It **must** be immediately followed by a quote control flag (`--s` or `--n`), and can include an optional PascalCase description suffix.
  * **Modifiers:**
    * `--val[PascalCase] --s "<string>"`: Encloses the value in single quotes (`'value'`).
    * `--val[PascalCase] --n "<string>"`: Emits the raw value without quotes (`value`), ideal for numbers or unquoted tokens.
  * **Usage:** `--lopId --valId --s "123"` generates `--id '123'`. `--lopCount --valCount --n "5"` generates `--count 5`.

* **`--arg[PascalCase]`**
  * **Meaning:** Appends a standalone, positional argument to the tail end of the generated command string. It **must** be immediately followed by a quote control flag (`--s` or `--n`), and can include an optional PascalCase description suffix.
  * **Modifiers:**
    * `--arg[PascalCase] --s "<string>"`: Appends a single-quoted positional argument.
    * `--arg[PascalCase] --n "<string>"`: Appends an unquoted positional argument.
  * **Usage:** `--argPattern --s "/pattern/d"` appends `'/pattern/d'`.

---

## 💡 Practical Examples & Use Cases

### Example 1: Multi-Stage Cloud & Data Processing with PascalCase Suffixes
This example showcases how option and argument modifiers can be paired with Alphanumeric PascalCase description suffixes for high readability.

```sh.sh
yomel \
	title "list home dir con by smart"\
	--no-live-stderr \
	--no-live-stdout \
	--log \
	stage "list bellow home directory" \
	-cmd ls \
	--argTargetDir "${HOME}" \
	--log-filter "shuf | head -5 | sort" \
	--err-log-filter "shuf | head -5 | sort" \
	stage "newline to tab" \
	-cmd tr \
	--argRepSrc '\n' \
	--argRepDst '\t' \

```

- log
<img width="883" height="950" alt="image" src="https://github.com/user-attachments/assets/9bd8995a-e4e6-4508-9ae1-0d2ed81778dc" />


---


## 🧪 Development & Testing

Run unit tests across all internal packages to verify functionality:

go test ./... -v

---

## 📜 License

MIT License
