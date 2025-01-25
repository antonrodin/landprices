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