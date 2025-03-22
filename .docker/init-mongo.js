db.createUser({
  user: "movie",
  pwd: "pass123",
  roles: [
    {
      role: "readWrite",
      db: "movie_service"
    }
  ]
});

db.createUser({
  user: "series",
  pwd: "pass123",
  roles: [
    {
      role: "readWrite",
      db: "series_service"
    }
  ]
});
