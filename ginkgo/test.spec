Given a shopping cart
  initially
    it has 0 items
    it has 0 units
    the total amount is 0.00

  when a new item is added
    the shopping cart has 1 more unique item than it had earlier
    the shopping cart has 1 more unit than it had earlier
    the total amount increases by item price

  when an existing item is added
    the shopping cart has the same number of unique items as earlier
    the shopping cart has 1 more unit than it had earlier
    the total amount increases by item price

  that has 0 unit of item A
    removing item A
      should not change the number of items
      should not change the number of units
      should not change the amount

  that has 1 unit of item A
    removing 1 unit item A
      should reduce the number of items by 1
      should reduce the number of units by 1
      should reduce the amount by the item price

  that has 2 units of item A
    removing 1 unit of item A
      should not reduce the number of items
      should reduce the number of units by 1
      should reduce the amount by the item price

    removing 2 units of item A
      should reduce the number of items by 1
      should reduce the number of units by 2
      should reduce the amount by twice the item price
