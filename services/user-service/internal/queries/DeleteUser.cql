MATCH
  (user:User {id: $userID})
DETACH DELETE user
