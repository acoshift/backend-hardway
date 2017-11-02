package main

func main() {
	// create tcp listener at :3333
	// don't forget to close tcp listener when done

	for {
		// accept connection from listener

		go func() {
			// don't forget to close connection when done

			// create bufio.NewReader from connection
			for {
				// read line from reader

				// print data out to console

				if true { // check is data is empty string
					// response HTTP to connection
				}
			}
		}()
	}
}
