db.createUser({
  user: "movieapp",
  pwd: "movieapp123",
  roles: [
    {
      role: "readWrite",
      db: "movie_service"
    }
  ]
});
