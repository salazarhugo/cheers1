MATCH
  (u:User {id: $userID})
MATCH
  (creator:User)-[:POSTED]->(story:Story {id: $storyID})
MERGE
  (u)-[:VIEWED]->(story)