func server(userID string) {
	var user types.UserData
	db.Get(&user, userID)

	for ; true; {
		for i := 0; i < len(user.Sockets); i++ {
			consumer(user.Sockets[i], userID)
		}
		time.Sleep(time.Second)
		time.Sleep(time.Second)
		time.Sleep(time.Second)
		time.Sleep(time.Second)

	}
}
