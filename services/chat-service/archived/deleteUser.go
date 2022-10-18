package archived

//func DeleteUser(userId string) {
//	rooms := roomCache.GetRooms(userId)
//
//	var wg sync.WaitGroup
//
//	for _, room := range rooms {
//		wg.Add(1)
//		room := room
//		go func() {
//			defer wg.Done()
//			roomCache.LeaveRoom(userId, room.Id)
//		}()
//	}
//
//	roomCache.DeleteTokens(userId)
//	wg.Wait()
//}
