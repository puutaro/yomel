#!/bin/bash
# gen_version.sh

readonly OUTPUT_FILE="internal/apps/yomel/pkg/info/z_yomel_toml.go"

# 自動生成される Go ファイルの構造をヒアドキュメントで書き出す
cat << EOF > "$OUTPUT_FILE"
package info

const YomelInfoRaw = \`$(cat yomel.toml)\`
EOF