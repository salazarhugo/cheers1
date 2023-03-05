MATCH (u:User) WHERE u.id = $userID OR u.username = $userID
RETURN PROPERTIES(u)