
<img width="512" height="474" alt="yomel4_siro_1024" src="https://github.com/user-attachments/assets/c90f8341-7ed6-4dde-a35a-1a64db71bf23" />

<!-- CIステータスバッジ -->
[![CI](https://github.com/puutaro/yomel/actions/workflows/ci.yaml/badge.svg)](https://github.com/puutaro/yomel/actions/workflows/ci.yaml)
[![Release](https://img.shields.io/github/v/release/puutaro/yomel)](https://github.com/puutaro/yomel/releases)
[![License](https://img.shields.io/github/license/puutaro/yomel)](https://github.com/puutaro/yomel/blob/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/puutaro/yomel.svg)](https://pkg.go.dev/puutaro/yomel)
[![codecov](https://codecov.io/gh/puutaro/yomel/branch/master/graph/badge.svg)](https://codecov.io/gh/puutaro/yomel)

![Linux](https://img.shields.io/badge/Linux-supported-success?logo=linux&logoColor=white)
![macOS](https://img.shields.io/badge/macOS-supported-success?logo=apple&logoColor=white)

# yomel

Super-readable and debuggable shell pipeline command tool, styled like `yaml`.  

## Innovative point about `yomel`

- Existing shellscript is not readable, but `yomel` give us super readable code like `yaml`!!
- Existing shellscript is not debugable, but `yomel` give us super structured logs!!

### Existing cmd drawback example

Bellow existing cmd has two big hard point.  

#### It's super hard for us to read.    
#### It's super hard for us to debug.  

- existing cmd

```sh.sh
step_num=$(\
	find \
	 "/home/xbabu/Desktop/share/temp/exp_py_for_yomel" \
		 -name "*.py" \
		 -type "f" \
	| xargs wc -l \
	| sort -nr \
	| head -1 \
	| sed 's/[^0-9]//g' \
);\
echo "${step_num}"
```

- existing log  

what?

- stdout 

<img width="939" height="31" alt="image" src="https://github.com/user-attachments/assets/bcdde97a-4553-4624-a710-7205eadfde1b" />

Generally, nobody wants to code in shell script if they can avoid it.  
We only do it because there isn't a simpler way, but it will absolutely put you through hell—whether you're writing it, reading it back later, or trying to improve it.  
Because the code is extremely difficult to read as prose/text, and  debug logs are not found.  



### Super advantages example of `yomel`

`yomel` give us two advantage like bellow.  


#### **Readable shellscript code line `yaml`**
#### **Structure log in shellscript pipeline**


- cmd

```sh.sh
step_num=$(yomel \
	title "agregate py step count" \
	stage "find py file" \
	-cmd find  \
	--argFindDir "/home/xbabu/Desktop/share/temp/exp_py_for_yomel" \
	--optFilterFileName name \
	--valOnlyPyExtend "*.py" \
	--optFilterType type \
	--valFile f \
	stage "count step num by each py file" \
	-cmd xargs \
	--argCountCmd  --n "wc -l" \
	stage "sort numerically in descending order" \
	--log \
	-cmd sort \
	--optNumrically n \
	--optDescendingOrder r \
	stage "get only first total line" \
	--log \
	-cmd head \
	--optOnlyFirstLine 1 \
	stage "get only step num" \
	--log \
	-cmd sed \
	--argSubstitute --s 's/[^0-9]//g' \
);\
echo "total ${step_num}"

```

The above code is very long.  
But we can easily recognize the pipeline's purpose from the `title`, `stage`, and `--opt` , `--val`, `--arg` suffix description.  

For a long time, I have been considering what's make a readable shell pipline.  
Eventually, I realized that the key is being rich in notes. By heavily annotating it, we can reach a readable shellpipeline.  
Although this approach is very simple, I am sure of its strength.  

Futhermore, look at the `yomel` log bellow.   
`yomel` log is a messive advantage.  
When I first saw this log, All the hassle associated with shell pipelines is disapeared.  

- log

<img width="942" height="872" alt="image" src="https://github.com/user-attachments/assets/fed95be1-97b0-4faf-8282-1e6ccfd1816b" />

<img width="942" height="960" alt="image" src="https://github.com/user-attachments/assets/1786ad03-2664-4386-897c-c2ade3ef4652" />

<img width="942" height="179" alt="image" src="https://github.com/user-attachments/assets/cc3bf3bf-7607-43b6-9ff4-8088587eb3a1" />

- stdout 

<img width="939" height="31" alt="image" src="https://github.com/user-attachments/assets/bcdde97a-4553-4624-a710-7205eadfde1b" />


I must point out that this log flows to `stderr`.  
So, it has no effect on `stdout` at all.In other words, we are free to put debug commands like `echo` and `tee` in the middle of the shell pipeline.  


Thanks to `yomel`, we can create a super readable and debuggable environment in our shell scripts.



## Demo

![yomel_demo2](https://github.com/user-attachments/assets/b8b5d0c6-ad0b-44bb-b64d-fe3cfcd3b20a)



## Installation (Linux/Mac)

### General

```sh.sh
curl https://raw.githubusercontent.com/puutaro/yomel/refs/heads/master/install.sh | sh
```

### go install

```sh.sh
go install github.com/puutaro/yomel/cmd/yomel@latest
```


## Complete Option Reference & Deep Dive

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




