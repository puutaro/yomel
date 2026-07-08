<!-- <img width="512" height="474" alt="yomel4_1024" src="https://github.com/user-attachments/assets/880acff6-4e01-4869-8fd9-6a8582a74fd6" /> -->

<img width="512" height="474" alt="yomel4_siro_1024" src="https://github.com/user-attachments/assets/c90f8341-7ed6-4dde-a35a-1a64db71bf23" />



# yomel

`yomel` is a command-line utility designed to write multi-stage shell script pipelines using a structured, flat, and human-readable argument layout—inspired by the clear, nested visual style of YAML configuration files.

By breaking down complex, nested one-liners or lengthy shell scripts into highly visible declarative steps (`stage`), `yomel` simplifies shell automation without abandoning native command-line paradigms.

---

## 🚀 Key Features

- 🛠️ **YAML-Like Structure:** Compose sequentially grouped pipelines via continuous CLI arguments using explicit components (`stage`, `-cmd`, `-svc`, `-act`, etc.).
- 🔄 **Auto-Chaining Pipelines:** Standard output (`stdout`) from an earlier stage automatically streams directly into the standard input (`stdin`) of the next stage using concurrent `io.TeeReader` pipes.
- 📝 **Smart Logging & Isolation:** Enable automated logging (`--log`) per stage with standalone error capturing.
- 🧽 **Asynchronous Stream Filtering:** Apply global or stage-specific custom shell hooks (`--log-filter` or `--err-log-filter`) to process log fragments on-the-fly (e.g., streaming only specific info lines).
- 🔤 **Granular Quote Management:** Control parameter escaping instantly with semantic operators like `--val --s` (single quote wrapper) or `--val --n` (no quote wrapper).

---

## 🛠️ Complete Option Reference & Deep Dive

`yomel` parses arguments sequentially from left to right. Arguments are divided into global/stage telemetry controllers, structural elements, and value modifiers.

### 1. Telemetry and Filter Options
These options control debugging output and stream filtering. They do not alter the data passing through the core pipeline but manage what is written to `stderr`.

* **`--log`**
  * **Meaning:** Activates the internal logging system. When this flag is present, `yomel` prints detailed panel execution metrics, generated shell commands, and raw step statuses to `stderr`.
  * **Usage:** Place it at the very beginning of the command to apply globally, or within specific sections.

* **`--log-filter "<shell_command>"`**
  * **Meaning:** Attaches an asynchronous log interceptor for standard output (`stdout`). The log data captured from the stage is passed to this shell command (e.g., `grep`, `awk`, `sed`) via stdin before being printed.
  * **Usage:** `--log-filter "grep 'ERROR'"` will ensure only log lines containing "ERROR" are emitted to your console log view.

* **`--err-log-filter "<shell_command>"`**
  * **Meaning:** Attaches an asynchronous log interceptor for standard error (`stderr`). This functions exactly like `--log-filter` but processes error streams thrown by the executing binaries.
  * **Usage:** `--err-log-filter "awk '{print \"[ERR] \" \$0}'"` prefixes all error outputs with a custom tag.

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

### 3. Option and Argument Value Modifiers
Modifiers specify how parameters, options, and trailing arguments are constructed and quoted.

* **`--opt "<flag>"`**
  * **Meaning:** Generates a short-style option flag (prefixed with a single dash `-`).
  * **Usage:** `--opt "v"` generates `-v`.

* **`--lop "<flag>"`**
  * **Meaning:** Generates a long-style option flag (prefixed with double dashes `--`).
  * **Usage:** `--lop "region"` generates `--region`.

* **`--val`**
  * **Meaning:** Declares a value associated with the preceding option (`--opt` or `--lop`). It **must** be immediately followed by a quote control flag (`--s` or `--n`).
  * **Modifiers:**
    * `--val --s "<string>"`: Encloses the value in single quotes (`'value'`).
    * `--val --n "<string>"`: Emits the raw value without quotes (`value`), ideal for numbers or unquoted tokens.
  * **Usage:** `--lop "id" --val --s "123"` generates `--id '123'`. `--lop "count" --val --n "5"` generates `--count 5`.

* **`--arg`**
  * **Meaning:** Appends a standalone, positional argument to the tail end of the generated command string. It **must** be immediately followed by a quote control flag (`--s` or `--n`).
  * **Modifiers:**
    * `--arg --s "<string>"`: Appends a single-quoted positional argument.
    * `--arg --n "<string>"`: Appends an unquoted positional argument.
  * **Usage:** `--arg --s "/pattern/d"` appends `'/pattern/d'`.

---

## 💡 Practical Examples & Use Cases

### Example 1: Multi-Stage Cloud & Data Processing
This example showcases how a lengthy AWS log query can be piped directly into `grep` and `sed` dynamically, structured into readable, distinct stages.

```bash
yomel \
  --log \
  --log-filter "grep 'Exception'" \
  stage "fetch-cloud-logs" \
  -cmd "aws" \
  --lop "region" \
  --val --s "us-east-1" \
  -svc "logs" \
  -act "filter-log-events" \
  --lop "log-group-name" \
  --val --s "/aws/lambda/my-prod-service" \
  --lop "limit" \
  --val --n "100" \
  stage "mask-sensitive-data" \
  -cmd "sed" \
  --opt "e" \
  --arg --s "s/[0-9]\{4\}-[0-9]\{4\}/XXXX-XXXX/g"
