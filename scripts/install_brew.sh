#!/bin/bash

set -euo pipefail

# Auto install brew only if missing
export NONINTERACTIVE=1
# Avoid Homebrew auto-update (keeps script deterministic)
export HOMEBREW_NO_AUTO_UPDATE=1

had_preexisting_brew=0
if command -v brew >/dev/null 2>&1; then
	had_preexisting_brew=1
else
	/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
fi

find_brew_bin() {
	if command -v brew >/dev/null 2>&1; then
		command -v brew
		return 0
	fi

	local candidates=(
		"/opt/homebrew/bin/brew" # macOS Apple Silicon
		"/usr/local/bin/brew"    # macOS Intel
		"/home/linuxbrew/.linuxbrew/bin/brew" # Linuxbrew default
	)

	for path in "${candidates[@]}"; do
		if [[ -x "$path" ]]; then
			echo "$path"
			return 0
		fi
	done

	return 1
}

if ! brew_bin="$(find_brew_bin)"; then
	echo "Homebrew installation not detected after install script" >&2
	exit 1
fi

shellenv_cmd="eval \"\$($brew_bin shellenv)\""

# Detect the current shell so we update the matching rc file; fall back to .profile.
if [ -n "${SHELL:-}" ]; then
	shell_name="$(basename "$SHELL")"
else
	shell_name="$(ps -p "$$" -o comm= 2>/dev/null | tail -n1)"
	shell_name="${shell_name##*/}"
	shell_name="${shell_name#-}"
	shell_name="${shell_name%% *}"
	shell_name="${shell_name:-sh}"
fi

case "$shell_name" in
	bash) rc_file="${HOME}/.bashrc" ;;
	zsh) rc_file="${HOME}/.zshrc" ;;
	*) rc_file="${HOME}/.profile" ;;
esac

if [[ "$had_preexisting_brew" -eq 0 ]]; then
	echo "Configuring Homebrew environment for $shell_name via $rc_file"

	touch "$rc_file"
	if ! grep -Fqx "$shellenv_cmd" "$rc_file"; then
		printf '\n%s\n' "$shellenv_cmd" >> "$rc_file"
	fi
fi

eval "$($brew_bin shellenv)"
