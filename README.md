# RF CLI

## Setup

```bash
$ export PATH=$GOPATH/bin:$PATH
$ go install github.com/deissh/rf-cli/cmd/rf@latest
```

```bash
$ rf version && rf help
```

```bash
$ rf config edit
```


## Usage
### Extension
```bash
$ rf extension --help
```
#### List public extensions

```bash
$ rf extension list
ID                                    NAME                     AUTHOR                           BASE URL
fed325b0-c3e0-4f9f-8c3b-7f958f8a28a5  Timeline                 tech@redforester.com             https://rf-b2b-timeline-frontend.extensions.redforester.com
9e07269a-8e9b-4114-a4d7-1940a67d120c  Exporter                 tech@redforester.com             https://satek-exporter.extensions.redforester.com
e9673b54-d6e7-4ca1-bd6a-341766fc4fb5  Aggregation              tech@redforester.com             https://satek-aggregation.extensions.redforester.com
bde7148b-5ac9-45fd-8fff-e1a729ce0ba3  Google Calendar          tech@redforester.com             https://satek-remind-me-ext.extensions.redforester.com:443
c8d4068e-7ac6-4f4a-b787-8a3db50e590c  Time Tracker             tech@redforester.com             https://rf-time-tracker.extensions.redforester.com:443
...
```

#### View extension info

```bash
$ rf extension info 12345678-7ac6-4f4a-b787-8a3db50e590c
ID: 12345678-7ac6-4f4a-b787-8a3db50e590c
Name: Example extension
Description: some description
...

$ rf extension info -f manifest.yaml
Name: Example manifest
Description: lorem ipsu
...
```

#### Update extension info

```bash
$ rf extension update -f manifest.yaml
ID: 12345678-7ac6-4f4a-b787-8a3db50e590c
Name: Example extension
Description: some description
...

# override extension id from manifest.yaml
$ rf extension update
ID: 0000000-1234-4f4a-b787-8a3db50e590c
Name: Example extension
Description: some description
...
```

#### Init extension manifest

```bash
$ rf extension init

? Extensions name: New extension
? Description: -
? Author email: user@example.com
? Extension base url: https://example.com
? Extension commands Add new command
? Command name: Test command
? Description:
? Extension commands Add new command
? Command name: Second command
? Description: lorem ipsu
? Extension commands  [Use arrows to move, type to filter]
  Test Command
  Second Command
  Add new command
> Exit

```
