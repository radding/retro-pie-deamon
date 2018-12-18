# Retropie Manager Daemon

This manager application runs a deamon that manages a multicast server for service discovery on the same network as well as an API to easily update the retropie and upload ROMs.

## Dependencies

Requires `avahi` to be installed (For multicasting service discovery).

## Usage

Run this application and this will use `avahi` to advertise its service.