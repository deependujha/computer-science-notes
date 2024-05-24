# NoSQL (`not only SQL`)

- NoSQL (`not only sql`)

- mongodb acid compliant:

```
s.start_transaction()
		orders.insert_one(order, session=s)
		stock.update_one(item, stockUpdate, session=s)
	s.commit_transaction()
```