MATCH
  (creator:User {id: $userID})
CREATE
  (story:Story $story)
CREATE
  (creator)-[:POSTED]->(story)