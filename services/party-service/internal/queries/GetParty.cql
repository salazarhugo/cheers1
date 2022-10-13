MATCH (u:User {id: $userId})
MATCH (host:User)-[:POSTED]->(party: Party {id: $partyId})
OPTIONAL MATCH (:User)-[interest:INTERESTED]->(party)
OPTIONAL MATCH (:User)-[going:GOING]->(party)
WITH
  host,
  party,
  exists((u)-[:INTERESTED]->(party)) as interested,
  exists((u)-[:GOING]->(party)) as going,
  count(DISTINCT interest) as interestedCount,
  count(DISTINCT going) as goingCount,
  host.id = u.id as isHost
RETURN
  party {
  .*,
    isHost: isHost,
    hostId: host.id,
    hostName: host.name,
    interested: interested,
    interestedCount: interestedCount,
    going: going,
    goingCount: goingCount
  }`
