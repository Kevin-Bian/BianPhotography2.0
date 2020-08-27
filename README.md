# Collageify
2.0 my friends

## Postgres (Currently Running on AWS)
```
docker run -it -p 5432:5432 -d postgres
```

### Create TABLE
```
CREATE TABLE photos (
    photoid SERIAL PRIMARY KEY,
    collageid TEXT,
    name TEXT,
    link TEXT,
    description TEXT
);
```
