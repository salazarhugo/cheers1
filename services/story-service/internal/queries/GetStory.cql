MATCH
  (u:User {id: $userID})
MATCH
  (creator:User)-[:POSTED]->(story:Story {id: $storyID})
OPTIONAL MATCH
  (:User)-[likes:LIKED]->(story)
OPTIONAL MATCH
  (:User)-[viewers:VIEWED]->(story)
WITH
  story,
  creator,
  EXISTS((u)-[:LIKED]->(story)) AS has_liked,
  COUNT(DISTINCT likes) AS like_count,
  COUNT(DISTINCT viewers) AS viewer_count
RETURN
  {
     story:         PROPERTIES(story),
     user:          PROPERTIES(creator),
     like_count:    like_count,
     viewer_count:  viewer_count,
     has_liked:     has_liked
  }
ORDER BY story.createTime DESC