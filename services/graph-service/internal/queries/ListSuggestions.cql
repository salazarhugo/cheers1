MATCH
  (user:User)-[:FRIENDS]->(friend:User)-[:FRIENDS]->(potentialFriend:User)
WHERE user.id = $userID
AND NOT (user)-[:FRIENDS]->(potentialFriend)
AND NOT (user)-[:REQUESTED]->(potentialFriend)
WITH
  potentialFriend,
  count(friend) as mutualFriendCount
ORDER BY mutualFriendCount DESC
LIMIT $pageSize
RETURN
  potentialFriend.id
UNION
MATCH
  (user: User {id: $userID})
MATCH
  (potentialFriend:User)
WHERE NOT (user)-[:FRIENDS]->(potentialFriend)
AND NOT (user)-[:REQUESTED]->(potentialFriend)
WITH
  potentialFriend,
  count(*) as friendCount
ORDER BY friendCount DESC
LIMIT $pageSize
RETURN potentialFriend.id
LIMIT $pageSize