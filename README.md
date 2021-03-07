# Shelly exporter for prometheus and grafana

This exporter exports some temperature and humdidity values from Shelly H&T to prometheus.

![ShellyPIC](https://shelly.cloud/wp-content/uploads/2020/06/Shelly_HT.png)

## Shelly H&T

Shelly API: https://shelly-api-docs.shelly.cloud/#shelly-h-amp-t

Shelly cloud API access: https://shelly.cloud/documents/developers/shelly_cloud_api_access.pdf

## Building

    go get github.com/aexel90/shelly_exporter/
    cd $GOPATH/src/github.com/aexel90/shelly_exporter
    go install

## shelly-metrics.json

Determine "auth_key", "id" and "url" of your device via Shelly cloud API access and update in shelly-metrics.json.
"shelly_name" and "name" can be determined on your own.

    "account": [
                {
                    "auth_key": "NGMxZTd1aWQ816EA26E4CD0F0A0DB602B9A77C3D195CD169EB86703ED0...",
                    "url": "shelly-20-eu.shelly.cloud"
                }
            ]

    "devices": [
                {
                    "id": "956...",                  
                    "shelly_name": "S1",
                    "name": "Outdoor"
                },
                {
                    "id": "d50...",                   
                    "shelly_name": "S2",
                    "name": "Indoor"
                }
            ]

## Running

Usage:

    $GOPATH/bin/shelly_exporter -h

    Usage of ./shelly_exporter:
        -listen-address string
                The address to listen on for HTTP requests. (default "127.0.0.1:9784")
        -metrics-file string
                The JSON file with the metric definitions. (default "shelly-metrics.json")
        -test
                Test configured metrics

## Example execution

### Running within prometheus:

    TODO

### Test exporter:

    TODO

## Grafana Dashboard

### weather widget

https://weatherwidget.io/

    <!doctype html> <html lang="de">
    <head>
    </head>
    <body>
    <a class="weatherwidget-io" href="https://forecast7.com/de/51d0513d74/dresden/" data-label_1="DRESDEN" data-theme="original" data-highcolor="#88d976" >DRESDEN</a>
    <script>
    !function(d,s,id){
        var js,fjs=d.getElementsByTagName(s)[0];
        if(!d.getElementById(id)){
        js=d.createElement(s);
        js.id=id;
        js.src='https://weatherwidget.io/js/widget.min.js';
        fjs.parentNode.insertBefore(js,fjs);
        setInterval('__weatherwidget_init()', 1800000) <!-- refresh widget every 30 minutes (1800000 milliseconds): -->
        }
    }(document,'script','weatherwidget-io-js');
    </script>
    </body>
    </html>

### dashboard

Grafana-ID: 13739

https://grafana.com/grafana/dashboards/13739

![Grafana](https://raw.githubusercontent.com/aexel90/shelly_exporter/main/grafana/screenshot.jpg)