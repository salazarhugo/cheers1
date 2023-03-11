MATCH
  (u:User {id: $userID})
MATCH
  (creator:User)-[:POSTED]->(post:Post {id: $postID})
MERGE
  (u)-[:COMMENTED]->(post)