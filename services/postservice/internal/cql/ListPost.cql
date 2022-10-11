MATCH (u:User {id: $userId})-[:FOLLOWS*0..1]->(author:User)-[:POSTED]->(post:Post)
WITH
    u,
    author,
    post
SKIP
    $skip
LIMIT
    $pageSize
OPTIONAL MATCH (:User)-[r:LIKED]->(post)
WITH
    p,
    author,
    exists((u)-[:LIKED]->(post)) as has_liked,
    count(DISTINCT r) as like_count
RETURN {
    post: p,
    user: author,
    like_count: like_count,
    comment_count : 1,
    has_liked: has_liked
}
ORDER BY p.created DESC
