MATCH
  (u:User {id: $userID})-[:FRIENDS*0..1]->(creator:User)-[:POSTED]->(post:Post)
WITH
  u,
  creator,
  post
ORDER BY post.createTime DESC
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
  COUNT(DISTINCT r) AS like_count,
  u.id = creator.id AS is_creator
RETURN
  {
     post:          PROPERTIES(post),
     user:          PROPERTIES(creator),
     like_count:    like_count,
     has_liked:     has_liked,
     is_creator:    is_creator
   }
