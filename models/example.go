



// This is an example model.  

type Example struct {
  Name      string
  Email     string
  Password  string
  Salt      string
}

// Let's talk about go funcation signatures for a second.
// Firstly, they're awesome.  You can return multiple
// objects, complete with type checking.

// In this example we are adding an interface to the
// Example type.  The subject example will be called
// user.  Next comes the name, followed by the return
// type.

func (user Example) send_welcome_email() bool {

}
