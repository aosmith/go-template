package main



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


// There are going to be a lot of generic functions in here,
// we can write a really simple genmrator for these that
// creates the stub.go with interpolation.

func (user Example) save() (Example, err) {


}

func (user Example) find() ([]Example) {


}

func (user Example) find_or_create_by(keys []string, key string) (Example) {

}
