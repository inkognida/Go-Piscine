myXargs tool that builds a command by appending all lines that are fed to program's stdin as this command's arguments, then execute it.

Usage: go build && ./echo -e "dirs_for os.stdin" | ./myXargs ls -la
