OPTIONAL MATCH
  (u:User {username: $username})
RETURN
  u IS NOT NULL