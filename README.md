# checkhost
Request a check-host.net TCP, UDP, HTTP or DNS check from the terminal

## Running
Download the latest release from [GitHub](https://github.com/azurejelly/checkhost/releases).
For a list of available options, run `checkhost --help`:
```
$ ./checkhost --help
🔧  Usage: checkhost [options] [target]

Options:
  -m, --mode string                 Type of check to request (http,
                                    tcp, udp, dns)
  -M, --max-nodes int               Amount of nodes to check from
  -n, --nodes string                Comma-separated list of nodes to use
  -a, --display-all-nodes boolean   Whether to display more than 5 nodes
                                    on the final output
  -y, --open-report boolean         Whether to automatically open the
                                    report in your web browser
  -h, --help boolean                Shows a list of available options
```

As an example, if you wanted to check if Google is online, you would do:
```
$ ./checkhost -m http https://www.google.com
✔ Successfully requested a HTTP check!
  • Report: https://check-host.net/check-report/21f5e6f4k8cf
  • Target: https://www.google.com
  • Nodes:
    - ru1.node.check-host.net @ Moscow, Russia (AS14576)
    - fr2.node.check-host.net @ Paris, France (AS12876)
    - id1.node.check-host.net @ Jakarta, Indonesia (AS21859)
    - ir1.node.check-host.net @ Tehran, Iran (AS49022)
    - ir6.node.check-host.net @ Karaj, Iran (AS208264)
    - ... and 18 more node(s)
      (use '--display-all-nodes' for a full list)

Do you want to open the report on your web browser? (y/N):
```
If you want to skip the prompt to open the report in your web browser, use `-y`.

### Docker
There's a Docker image available at [Docker Hub](https://hub.docker.com). Pull it with:
```
$ docker pull azurejelly/checkhost:latest
```

The usage is similar:
```
$ docker run --rm \
    azurejelly/checkhost \
    -m http \
    https://www.google.com
```
- `--rm` will remove the container once it's done executing
- `azurejelly/checkhost` is the Docker image
- `-m http https://www.google.com` will be passed to `checkhost`

## Building
Clone the repository:
```sh
$ git clone https://github.com/azurejelly/checkhost
$ cd ./checkhost/
```
Then run:
```sh
$ go build -o ./checkhost
```
