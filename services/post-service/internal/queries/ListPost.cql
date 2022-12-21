MATCH
  (u:User {id: $userID})
MATCH
  (creator:User)-[:POSTED]->(post:Post)
WHERE
  creator.id = $userIdOrUsername OR creator.username = $userIdOrUsername
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
  exists((u)-[:LIKED]->(post)) AS has_liked,
  count(DISTINCT r) AS like_count,
  u.id = creator.id AS is_creator
RETURN
  {
    post:          properties(post),
    user:          properties(creator),
    like_count:    like_count,
    comment_count: 0,
    has_liked:     has_liked,
    is_creator:    is_creator
  }
  ORDER BY post.createTime DESC