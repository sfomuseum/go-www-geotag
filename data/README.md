This folder is used for storing user-specific data files that you may want to copy in to a Docker container. These are included in a container instance with the following commands in the `Dockerfile`:

```
COPY data/*.db /usr/local/data
COPY data/*.pmtiles /usr/local/data
```

Importantly any files included in this folder ending `.db`, `.sqlite`, `.sqlite3` or `.pmtiles` are explicitly _excluded from being included in any Git commits._