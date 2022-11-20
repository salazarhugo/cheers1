MATCH
  (u:User {id: $userID})
MATCH
  (creator:User {username: $otherUserID})-[:POSTED]->(post:Post)
WITH
  u,
  creator,
  post
SKIP
  $skip
LIMIT
  $pageSize
OPTIONAL MATCH
  (:User)-[r:LIKED]->(post)
WITH
  post,
  creator,
  EXISTS((u)-[:LIKED]->(post)) AS has_liked,
  COUNT(DISTINCT r) AS like_count
RETURN
  {
     post:          PROPERTIES(post),
     user:          PROPERTIES(creator),
     like_count:    like_count,
     comment_count: 0,
     has_liked:     has_liked
   }
ORDER BY post.createTime DESC