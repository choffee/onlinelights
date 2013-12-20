onlinelights
============

Some simple code for running some RGB lights on a SparkCore

Attache a red, green and blue led to pins A4, A5 and A6. Then setup the server as below.

    go install github.com/choffee/onlinelights
    cd $GPOATH 
    cp lights.cfg.template lights.cfg
    # Then put in your device id and code that can be found at http://spark.com
    $EDITOR lights.cfg
    onlinelights

Now visit http://localhost:8080/ and enjoy the lights!
