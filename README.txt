install w/:
goinstall github.com/JayTeeSF/MyFibonacci

from your code you can:
  include "MyFibonacci"

which gives you the FibCommander (Struct) returning method:
   MyFibonacci.Fibber()
   with which you can call previous, current, or next in order to cycle through fibonacci sequence

ls -al $GOROOT/pkg/*/github.com/JayTeeSF/

# if you download the repo from github, then test it by RUNNING:
make test
