# Promviz [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat)](http://makeapullrequest.com) [![MIT Licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/nghialv/promviz/blob/master/LICENSE)


Promviz is an application that helps you visualize the traffic of your cluster from Prometheus data.

It has 2 components:

- Promviz: retrieves data from Prometheus servers, aggregates them and provides an API to get the graph data.

- [Promviz-front](https://github.com/mjhd-devlion/promviz-front): based on Netflix's [vizceral](https://github.com/Netflix/vizceral) to render traffic graph.

#### Features:
- Generates and renders traffic graph in realtime
- Able to replay from any time in the past
- Able to generate notices on node and connection from prom query
- Provides a sidecar application for k8s that watches config changes and reload Promviz server in runtime
- Fits with [Istio](https://istio.io)'s metrics

![](https://github.com/nghialv/promviz/blob/master/documentation/sample_filter.png)

## Development
Start a promviz server at 9091 port with the configuration of promviz.yaml
```bash
go run cmd/promviz/main.go  --config.file="./promviz.yaml" --storage.path=/tmp/promvizdata
```

You can access `http://localhost:9091/graph` to check the metrics.

## Architecture

![](https://github.com/nghialv/promviz/blob/master/documentation/architecture.png)

## Configuration

See [configuration.md](https://github.com/nghialv/promviz/blob/master/documentation/configuration.md) in documentation directory.

## Contributing

Please feel free to create an issue or pull request.

## LICENSE

Promviz is released under the MIT license. See LICENSE file for details.
