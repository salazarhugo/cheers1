MATCH
  (u:User {id: $userID})
MATCH
  (creator:User)-[:POSTED]->(post:Post {id: $postID})
OPTIONAL MATCH
  (:User)-[likes:LIKED]->(post)
WITH
  post,
  creator,
  EXISTS((u)-[:LIKED]->(post)) AS has_liked,
  COUNT(DISTINCT likes) AS like_count,
  u.id = creator.id AS is_creator
RETURN
  {
     post:          PROPERTIES(post),
     user:          PROPERTIES(creator),
     like_count:    like_count,
     has_liked:     has_liked,
     is_creator:    is_creator
   }
ORDER BY post.createTime DESC