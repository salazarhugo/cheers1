MATCH
  (u:User {id: $userID})-[:FOLLOWS*0..1]->(creator:User)-[:POSTED]->(story:Story)
WITH
  u,
  creator,
  story
SKIP
  $skip
LIMIT
  $pageSize
OPTIONAL MATCH
  (:User)-[r:LIKED]->(story)
WITH
  story,
  creator,
  COUNT(DISTINCT r) AS like_count,
  EXISTS((u)-[:LIKED]->(story)) AS has_liked
RETURN
  {
     story:         PROPERTIES(story),
     user:          PROPERTIES(creator),
     like_count:    like_count,
     has_liked:     has_liked
   }
ORDER BY story.createTime DESC