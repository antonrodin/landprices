# landprices
Api to access land prices inside UK, using kaggle dataset.


## Import all transactions

1. Download the source: https://www.kaggle.com/datasets/willianoliveiragibin/uk-property-price-data-1995-2023-04?resource=download

```
sqlite> CREATE TABLE transactions (
    Transaction_unique_identifier TEXT PRIMARY KEY,
    price REAL,
    Date_of_Transfer TEXT,
    postcode TEXT,
    Property_Type TEXT,
    Old_New TEXT,
    Duration TEXT,
    PAON TEXT,
    SAON TEXT,
    Street TEXT,
    Locality TEXT,
    Town_City TEXT,
    District TEXT,
    County TEXT,
    PPDCategory_Type TEXT,
    Record_Status_monthly_file_only TEXT
);

sqlite> .mode csv
sqlite> .import ./202304.csv transactions
```

## I have uploaded the DB to my bucket

https://s3.amazonaws.com/azr-es/prices.db

## Compile binary with CGO enabled

```
# CGO_ENABLED if not works
sudo apt-get install build-essential

# Compile
env GOOS=linux CGO_ENABLED=1 go build -o landApp ./cmd/api

# For MacOs
env GOOS=darwin GOARCH=amd64 go build -o landAppMac ./cmd/api/*
```

## Supervisor Conf

````
cd /etc/supervisor/conf.d/

# Create supervisor file
sudo vim landtitles.conf
````

````
anton@gosha:/etc/supervisor/conf.d$ cat landtitles.conf
[program:landtitles]
command=/home/anton/www/landprices/landApp
directory=/home/anton/www/landprices
autorestart=true
autostart=true
stdout_logfile=/home/anton/www/landprices/supervisor.log
````

Then

````
sudo supervisorctl
> status
> update
> status
````

## Nginx Configuration

````
server {
        listen 80;

        server_name landtitles.azr.es;

        location / {
                proxy_pass http://localhost:3334;
        }
}
````

## Fields explanations from here

- primaryAddress TEXT: Typically the house number or name.
- secondaryAddress TEXT: Additional information if the building is divided into flats or sub-buildings.
- propertyType TEXT: D, T, S, A <- no idea.
- categoryType TEXT: F ? FreeHoland and LeaseHold???