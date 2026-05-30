# slskd-exporter

## Install

Clone the repo and rename the .env file:
```
git clone https://github.com/Deb84/slskd-exporter
cd slskd-exporter
mv .env-example .env
```
Add your slskd informations to the .env file  

Start the docker container:
```
docker compose up
```

## Grafana
- add a postgres data source (login informations are in the .env file
- create a dashboard (pre-built dashboard soon)
- add a panel
- select the postgres data source and write your query (see the postgres tables & columns below)

## Postgres

### Tables:
- transfers
- files
- batches

Batches are a set of transfers downloaded together

### Columns
**transfers**
- id (pk)
- file_id (files fk)
- batch_id (batches fk)
- transfer_uuid (slskd transfers internal id, unique)
- username
- direction (upload/download)
- size (bytes (SI))
- state
- requested_at
- started_at
- ended_at
- bytes_transferred
- average_speed (bytes/sec (SI))
- exception
- attempts
- bytes_remaining
- elapsed_time
- percent_complete
- remaining_time

**files**
- id (pk)
- path (unique)
- file_name (unused, comming soon)
- artist_name (unused, comming soon)
- album_name (unused, comming soon)
- year (unused, comming soon)

**batches**
- id (pk)
