.PHONY: dev-mode
dev-mode:
	gh extension remove KazukiHayase/gh-requested-prs
	gh extension install .

.PHONY: prod-mode
prod-mode:
	gh extension remove ./gh-requested-prs
	gh extension install KazukiHayase/gh-requested-prs
