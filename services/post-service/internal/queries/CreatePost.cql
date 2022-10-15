MATCH
  (creator:User {id: $userID})
CREATE
  (post:Post $post)
CREATE
  (creator)-[:POSTED]->(post)