MATCH
  (u:User {id: $userID})-[:FOLLOWS*0..1]->(creator:User)-[:POSTED]->(story:Story)
WITH
  u,
  creator,
  story
WITH
  story,
  creator,
  EXISTS((u)-[:LIKED]->(story)) AS has_liked,
  EXISTS((u)-[:VIEWED]->(story)) AS has_viewed
RETURN
  {
    user:    PROPERTIES(creator),
    stories: COLLECT(DISTINCT {
      story: PROPERTIES(story),
      has_liked:     has_liked,
      has_viewed:    has_viewed
    })
  }
SKIP
  $skip
LIMIT
  $pageSize