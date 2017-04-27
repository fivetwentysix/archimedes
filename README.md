# archimedes
Slack bot that is designed to capture Slack messages and store them in a backend (Currently Confluence)

Current planning docs at https://newcontext.atlassian.net/wiki/display/NCIS/Project+Archimedes

## Contributing

This project uses [godep][godep] for dependencies, following their recommendation and source controlling `Godeps/Godeps.json` and the resulting `vendor/`.

After cloning and making code changes, you can just run `go test` to see if it still works. Well, as long as we've added tests by then.

To run locally, run `go install` and if your `$GOPATH/bin` is on your path, run:

```
SLACK_TOKEN=secret WIKI_USER=secret_user WIKI_PASS=secret_pass archimedes
```

[godep]: https://github.com/tools/godep
