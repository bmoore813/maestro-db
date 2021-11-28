# Design Spec for What Comprises a Log

- `Record`- The data that we store in a log
- `Store` - The file we store the records in
- `Index` - The file we store index entries in
- `Segment` - The abstraction that ties a `Store` and an `Index` together
- `Log` - The abstraction that ties all the segments together