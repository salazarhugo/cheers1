MATCH
  (u:User) WHERE u.id IN $userIDs
RETURN
  properties(u)