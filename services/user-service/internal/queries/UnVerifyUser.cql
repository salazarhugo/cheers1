MATCH
  (u:User {id: $userID})
SET
  u.verified = false
