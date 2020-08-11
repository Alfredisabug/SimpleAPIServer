# About This

## Query String

```SELECT Data.Id, Data.DateAdded, Location.Lat, Location.Long FROM Data inner join Location On Data.Id=Location.Id;```

## Write String

### Table:Data

```INSERT INTO Data(DateAdded) VALUES(NOW());```

### Table:Location

```INSERT INTO Location() VALUES(LAST_INSERT_ID(), 100.0, 200.0);```
