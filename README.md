# RF CLI

## Setup

```bash
$ export PATH=$GOPATH/bin:$PATH
$ go install github.com/deissh/rf-cli@latest
```

```bash
$ rf-cli version && rf-cli help
```

```bash
$ rf-cli config
```


## Usage
### Extension
```bash
$ rf-cli extension --help
```
#### List public extensions

```bash
$ rf-cli extension
ID                                    NAME             EMAIL                    BASE URL                                                           DESCRIPTION
62562972-ffcc-4225-a779-d63d9d944af3  Kanban EN        support@redforester.com                                                                     Open kanban for sprints, backlog, milestones and projects
bde7148b-5ac9-45fd-8fff-e1a729ce0ba3  Google Calendar  support@redforester.com  https://satek-remind-me-ext.extensions.redforester.com:443         Create event in your Google Calendar
286a47a0-504d-457c-8607-fec896cff599  Timeline RU      support@redforester.com  https://satek-timeline-frontend.extensions.redforester.com:443     Отображает проекты, этапы и спринты в виде временной диаграммы
...
```
