MATCH
  (u:User {id: $userID})
MATCH
  (:User)-[:POSTED]->(story:Story {id: $storyID})
WITH
  story,
  EXISTS((u)-[:VIEWED]->(story)) AS has_viewed,
  EXISTS((u)-[:LIKED]->(story)) AS has_liked
RETURN
  {
    story:         PROPERTIES(story),
    has_liked:     has_liked,
    has_viewed:    has_viewed
  }
ORDER BY story.createTime DESC