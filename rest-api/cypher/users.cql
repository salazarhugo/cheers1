MATCH (u:User {id: $userId})<-[:FOLLOWS]-(follower:User)
WITH follower, exists((me)-[:FOLLOWS]->(follower)) as followBack
RETURN follower {.id, .username, .verified, .picture, .name, followBack:followBack}
