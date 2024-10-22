.PHONY: install-hooks
install-hooks: .git/hooks/some-hook .git/hooks/another-hook

.git/hooks/%:
	ln -s hooks/$* $@
